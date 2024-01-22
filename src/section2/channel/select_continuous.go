package channel

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const bufSize2 = 5

func DemoSelectContinuous() {
	ch1 := make(chan int, bufSize2)
	ch2 := make(chan int, bufSize2)
	var wg sync.WaitGroup
	// cancel()が呼ばれたらctx.Doneをクローズする
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Millisecond)
	defer cancel()
	wg.Add(3)
	go countProducer(&wg, ch1, bufSize2, 50)
	go countProducer(&wg, ch2, bufSize2, 500)
	go countConsumer(ctx, &wg, ch1, ch2)
	wg.Wait()
	fmt.Println("finish")
}

// chan<-は書き込み専用 <-chanは読み込み専用のchannelを型指定する
func countProducer(wg *sync.WaitGroup, ch chan<- int, size int, sleep int) {
	defer wg.Done()
	defer close(ch)
	for i := 0; i < size; i++ {
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		ch <- i
	}
}

// ctxはキャンセル用 ch1, ch2は読み込み用のchannel
func countConsumer(ctx context.Context, wg *sync.WaitGroup, ch1 <-chan int, ch2 <-chan int) {
	defer wg.Done()

	// loop:
	// ch1, ch2のどちらかがnilになるまで受信を繰り返す
	for ch1 != nil || ch2 != nil {
		select {
		// ctx.Done()が呼ばれたらループを抜ける
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			// break loop
			for ch1 != nil || ch2 != nil {
				select {
				case v, ok := <-ch1:
					if !ok {
						ch1 = nil
						// caseを抜ける loopは抜けない
						break
					}
					fmt.Printf("ch1: %v\n", v)
				case v, ok := <-ch2:
					if !ok {
						ch2 = nil
						break
					}
					fmt.Printf("ch2: %v\n", v)
				}
			}
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				// caseを抜ける loopは抜けない
				break
			}
			fmt.Printf("ch1: %v\n", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			fmt.Printf("ch2: %v\n", v)
		}
	}
}
