package channel

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func DemoSelect() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	var wg sync.WaitGroup
	//第一引数は親コンテキスト(最初はBackground）, 第二引数はtimeoutするまでの時間
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	// cancel()が呼ばれたらctx.Doneをクローズする
	defer cancel()
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		ch1 <- "A"
	}()
	go func() {
		defer wg.Done()
		time.Sleep(800 * time.Millisecond)
		ch2 <- "B"
	}()

loop:
	// ch1, ch2のどちらかがnilになるまで受信を繰り返す
	for ch1 != nil || ch2 != nil {
		select {
		case <-ctx.Done():
			// timeoutしたらループを抜ける
			// このときchannelがバッファなしの場合はloopを抜けてgo funcの書き込み処理が残るのでデッドロックになる
			// バッファありなら書き込み処理はできる
			fmt.Println("timeout")
			break loop
		// 受け取ったらnilにする
		case v := <-ch1:
			fmt.Println(v)
			ch1 = nil
		case v := <-ch2:
			fmt.Println(v)
			ch2 = nil
		}
	}
	// goroutineが終了するまで待つ (この関数が終了するとgoroutineも終了するため)
	wg.Wait()
	fmt.Println("finish")
}
