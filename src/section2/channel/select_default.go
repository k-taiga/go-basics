package channel

import (
	"fmt"
	"sync"
	"time"
)

const bufSize1 = 3

func DemoSelectDefault() {
	var wg sync.WaitGroup
	// 3つのバッファを持つchannelを作成
	ch := make(chan string, bufSize1)
	// goroutineの終了を待つためwgに1を追加
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < bufSize1; i++ {
			time.Sleep(1000 * time.Millisecond)
			// バッファがいっぱいになるまで書き込み
			ch <- "hello"
		}
	}()

	for i := 0; i < bufSize1; i++ {
		select {
		// channelから読み込みができたら表示
		case m := <-ch:
			fmt.Println(m)
		// channelから読み込みができなかったらdefaultの処理を実行 1500msごとに実行されるので最初はdefaultが実行される
		default:
			fmt.Println("no msg arrived")
		}

		time.Sleep(1500 * time.Millisecond)
	}
}
