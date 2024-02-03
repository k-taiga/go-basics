package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func DemoContextWithTimeout() {
	var wg sync.WaitGroup

	// context.Backgroundは親のコンテキストで空のコンテキストを返す
	// withTimeoutで親のコンテキストにタイムアウトを設定したコンテキストを返す
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	// ↓だとすべてのサブタスクがキャンセルされる
	// ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	// 600ms後にキャンセルされる
	defer cancel()
	wg.Add(3)
	// 3つのサブタスクを並列で実行
	go subTask(ctx, &wg, "a")
	go subTask(ctx, &wg, "b")
	go subTask(ctx, &wg, "c")
	wg.Wait()
}

func subTask(ctx context.Context, wg *sync.WaitGroup, id string) {
	defer wg.Done()
	// NewTickerは指定した間隔でチャネルに値を送信する 500ms間隔でt.Cに値が送信される
	t := time.NewTicker(500 * time.Millisecond)

	select {
	// ctx.Done()は親のコンテキストがキャンセルされたときに閉じられるチャネル
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return
	// 最初のCを受信するまで待機
	case <-t.C:
		t.Stop()
		fmt.Println(id)
		return
	}
}
