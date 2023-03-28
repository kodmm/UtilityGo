package main

import (
	"context"
	"os/exec"
)

func main() {
	ctx := context.Background()
	err := runJobs(ctx)
}

func runJobs(ctx context.Context) error {
	// cancel関数を作成。 deferで関数を抜けるときに自動で呼ばれるようにする
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// エラーと終了情報の受信用チャネル
	ec := make(chan error)
	done := make(chan struct{})

	// タスクを10回起動
	for i := 0; i < 10; i++ {
		go func() {
			cmd := exec.CommandContext(ctx, "sleep", "30")
			err := cmd.Run()
			if err != nil {
				ec <- err
			} else {
				done <- struct{}{}
			}
		}()
	}

	// 終了を待つ
	for i := 0; i < 10; i++ {
		select {
		case err := <-ec:
			// ここでエラーを返して関数を終了させること
			// defer cancel()が呼ばれ実行中の別のタスクが終了する
			return err
		case <-done:
		}
	}
	return nil
}
