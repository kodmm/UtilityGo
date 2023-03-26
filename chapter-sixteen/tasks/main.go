package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync/atomic"
)

type Task string

type Result struct {
	Value int64
	Task  Task
	Err   error
}

func worker(id int, tasks <-chan Task, results chan<- Result) {
	for t := range tasks {
		fmt.Printf("workder: %d task: %s\n", id, t)
		s, err := os.Stat(string(t))
		if err == nil && s.IsDir() {
			err = fmt.Errorf("worker: %d err: %s is dir", id, string(t))
		}
		result := Result{
			Task: t,
		}
		if err != nil {
			result.Err = err
		} else {
			fmt.Printf("worker: %d path: %s size: %d\n", id, string(t), s.Size())
			result.Value = s.Size()
		}
		results <- result
	}
}

func TotalFileSize() int64 {
	tasks := make(chan Task)
	results := make(chan Result)

	// ワーカーを起動
	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, tasks, results)
	}
	// タスクを非同期でチャネルに投入
	inputDone := make(chan struct{})
	var remainedCount int64
	go func() {
		filepath.Walk(runtime.GOROOT(), func(path string, info os.FileInfo, err error) error {
			atomic.AddInt64(&remainedCount, 1)
			tasks <- Task(path)
			return nil
		})
		close(inputDone)
		close(tasks)
	}()

	var size int64
	for {
		select {
		case result := <-results:
			if result.Err != nil {
				fmt.Printf("err %v for %s\n", result.Err, result.Task)
			} else {
				atomic.AddInt64(&size, result.Value)
			}
			atomic.AddInt64(&remainedCount, -1)
		case <-inputDone:
			if remainedCount == 0 {
				return size
			}
		}
	}
}

func fixedTasks(taskSrcs []Task) int64 {
	// タスクの全量がわかっているなら、あらかじめ全部チャネルに入れてしまうのがシンプル
	tasks := make(chan Task, len(taskSrcs))
	results := make(chan Result)
	for _, src := range taskSrcs {
		tasks <- src
	}
	close(tasks)
	// コア数分ワーカーを起動
	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, tasks, results)
	}
	// 結果を受け取りつつ、全タスクの完了を待つ
	var count int
	var size int64
	for {
		result := <-results
		count += 1
		if result.Err != nil {
			fmt.Printf("err %v for %s\n", result.Err, result.Task)
		} else {
			size += result.Value
		}
		if count == len(taskSrcs) {
			break
		}
	}
	return size
}
