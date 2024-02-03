package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func DemoContextWithCancel() {
	var wg sync.WaitGroup
	// withCancelでcancel関数を返す
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(2)
	// goroutineでサブタスクを実行
	go func() {
		defer wg.Done()
		// main側で作成したctxを引数に渡す
		v, err := criticalTask(ctx)
		// errが発生していたらエラーメッセージとともにキャンセルを実行
		if err != nil {
			fmt.Printf("criticalTask canceled due to: %v\n", err)
			cancel()
			return
		}
		fmt.Println("success", v)
	}()
	go func() {
		defer wg.Done()
		v, err := normalTask(ctx)
		if err != nil {
			fmt.Printf("normalTask canceled due to: %v\n", err)
			cancel()
			return
		}
		fmt.Println("success", v)
	}()
	wg.Wait()
}

func criticalTask(ctx context.Context) (string, error) {
	// 親のコンテキストからタイムアウトを設定した新しいコンテキストを設定
	ctx, cancel := context.WithTimeout(ctx, 1200*time.Millisecond)
	defer cancel()
	// 1秒ごとのticker
	t := time.NewTicker(1000 * time.Millisecond)
	select {
	// ctxがtimeoutされたらエラーを返す
	case <-ctx.Done():
		return "", ctx.Err()
	// tickerからCに値を受信するまで待機
	case <-t.C:
		// 受け取ったら停止
		t.Stop()
	}
	return "A", nil
}

func normalTask(ctx context.Context) (string, error) {
	t := time.NewTicker(3000 * time.Millisecond)
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-t.C:
		t.Stop()
	}

	return "B", nil
}
