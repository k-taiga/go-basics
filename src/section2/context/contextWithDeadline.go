package context

import (
	"context"
	"fmt"
	"time"
)

func DemoContextWithDeadline() {
	// 現在時刻から40ms後にキャンセルされるコンテキストを作成
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(40*time.Millisecond))
	// これだと20ms後にdeadlineが過ぎてしまうのでsubTaskの中身は実行されない
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(20*time.Millisecond))
	defer cancel()
	ch := deadlineSubTask(ctx)
	v, ok := <-ch
	// chから値を受信できたらokがtrueになる
	if ok {
		fmt.Println(v)
	}
	fmt.Println("finish")
}

func deadlineSubTask(ctx context.Context) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		// ctxのデッドラインの時刻を取得
		deadline, ok := ctx.Deadline()
		// 設定されていればtrue, されていなければfalse
		if ok {
			// subTaskは30msかかるので30ms過ぎたら実行できないので終了する
			// deadline - time.Now()で残り時間を取得 30msより小さければ実行不可能
			if deadline.Sub(time.Now().Add(30*time.Millisecond)) < 0 {
				fmt.Println("impossible to meet deadline")
				// returnしてdefer close(ch)を実行
				return
			}
		}
		time.Sleep(30 * time.Millisecond)
		// chに値を送信してdefer close(ch)を実行
		ch <- "hello"
	}()
	return ch
}
