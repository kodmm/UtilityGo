package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Done: 1")
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Done: 2")
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Done: 3")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("done all tasks")

	ctx := context.Background()
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("Done: 1")
		return nil
	})
	eg.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("Done: 2")
		return nil
	})
	eg.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("Done: 3")
		return nil
	})
	err := eg.Wait()
	fmt.Println("Done all tasks", err)
}
