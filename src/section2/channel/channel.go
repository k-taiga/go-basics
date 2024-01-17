package channel

import (
	"fmt"
	"runtime"
)

func DemoChannel() {
	// ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	// 送信は受信があるまでブロックされる
	// 	ch <- 10
	// 	time.Sleep(500 * time.Millisecond)
	// }()
	// fmt.Println(<-ch)
	// wg.Wait()

	ch1 := make(chan int)
	go func() {
		// 送信があるまでずっと待つ goroutine leak
		fmt.Println(<-ch1)
	}()
	// goroutine leakの回避
	ch1 <- 10
	fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())

	// buffer付きchannelは、bufferが埋まるまでブロックされない
	ch2 := make(chan int, 1)
	ch2 <- 2
	// これはbufferが埋まっているのでブロックされる deadlock!
	ch2 <- 3
	fmt.Println(<-ch2)

	// なにも書き込まれていないのに対して読み込もうとしている deadlock!
	// fmt.Println(<-ch2)
	// ch2 <- 2
}
