package channel

import (
	"fmt"
	"sync"
	"time"
)

func ChannelClose() {
	ch1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 送信は受信があるまでブロックされる
		fmt.Println(<-ch1)
	}()
	ch1 <- 10
	close(ch1)
	v, ok := <-ch1
	fmt.Printf("v: %d, ok: %v\n", v, ok)
	wg.Wait()

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	v, ok = <-ch2
	close(ch2)
	// 1, true
	fmt.Printf("v: %d, ok: %v\n", v, ok)
	v, ok = <-ch2
	// 2, true
	fmt.Printf("v: %d, ok: %v\n", v, ok)
	// 0, false バッファ付きチャネルの場合、close後もバッファに残っている値は取り出せる
	v, ok = <-ch2
	fmt.Printf("v: %d, ok: %v\n", v, ok)

	ch3 := generateCountStream()
	// チャネルがクローズされるまでrangeで受信を繰り返す
	for v := range ch3 {
		fmt.Println(v)
		time.Sleep(2 * time.Second)
	}

	// structのchanは0バイトなのでメモリを消費せず通知専用に適している
	nCh := make(chan struct{})
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("goroutine %v started\n", i)
			// チャネルがクローズされるまで受信を繰り返す
			<-nCh
			fmt.Println(i)
		}(i)
	}
	time.Sleep(2 * time.Second)
	close(nCh)
	wg.Wait()
	fmt.Println("finish")
}

// <-chanは受信専用チャネルを返す
func generateCountStream() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			// chに送信 読み取りがないとブロックされる
			ch <- i
			fmt.Println("write")
		}
	}()
	// chを受信専用チャネルとして返す
	return ch
}
