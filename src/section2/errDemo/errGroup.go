package errdemo

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func DemoErrGroup() {
	// timeout付きのcontextを作成
	ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel()

	// errgroupのcontextを作成
	eg, ctx := errgroup.WithContext(ctx)
	s := []string{"task1", "fake1", "task2", "task3", "task4"}

	// sliceをrangeで回して、それぞれのtaskをgoroutineで実行
	// 通常のgoroutineは返り値を受け取れないが、errgroupを使うと返り値を受け取れる
	for _, v := range s {
		task := v
		eg.Go(func() error {
			return doTask(ctx, task)
		})
	}
	// eg.Wait()で全てのgoroutineが終了するまで待つ
	if err := eg.Wait(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println("finish")
}

func doTask(ctx context.Context, task string) error {
	var t *time.Ticker
	switch task {
	case "task1":
		t = time.NewTicker(500 * time.Millisecond)
	case "task2":
		t = time.NewTicker(700 * time.Millisecond)
	default:
		t = time.NewTicker(1000 * time.Millisecond)
	}

	select {
	// errgroupのcontextはerrを検知したらcontextをDoneにする
	case <-ctx.Done():
		fmt.Printf("%v canceled : %v\n", task, ctx.Err())
		return ctx.Err()
	// tickerのchannelしたら実行
	case <-t.C:
		t.Stop()
		fmt.Printf("task %v completed\n", task)
	}
	return nil
}
