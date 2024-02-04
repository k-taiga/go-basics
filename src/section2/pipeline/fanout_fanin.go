package pipeline

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func DemoFanoutFanin() {
	// logical coresの数を取得(稼働している論理プロセッサの数)
	cores := runtime.NumCPU()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nums := []int{1, 2, 3, 4, 5}

	// string型のchannelのスライスを作成 サイズはcoresの数
	outChs := make([]<-chan string, cores)
	// generate関数でchannelに値を送る
	inData := generate(ctx, nums...)
	// for文でcoresの数だけfanOut関数を呼び出す
	for i := 0; i < cores; i++ {
		// fanOutの結果をoutChs(chのスライス)に格納
		outChs[i] = funOut(ctx, inData, i+1)
	}
	var i int
	flag := true
	// fanInにoutChsを展開して返す
	// fanInは複数のchannelから受け取った値を一つのchannelにまとめる
	for v := range fanIn(ctx, outChs...) {
		// 3回目のループでcancelを呼び出してcontextをキャンセルする
		if i == 3 {
			cancel()
			flag = false
		}
		if flag {
			fmt.Println(v)
		}
		i++
	}
	fmt.Println("finish")
}

func generate(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()
	return out
}

func funOut(ctx context.Context, in <-chan int, id int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		// 重い処理を行う関数
		heavyWork := func(i int, id int) string {
			time.Sleep(200 * time.Millisecond)
			return fmt.Sprintf("result: %v (id:%v)", i*i, id)
		}
		// 受け取った値を重い処理にかけてchannelに送る
		for v := range in {
			select {
			case <-ctx.Done():
				return
			case out <- heavyWork(v, id):
			}
		}
	}()
	return out
}

// 第２引数に可変長のchannelを受け取る
// fanInは複数のchannelから受け取った値を一つのchannelにまとめる
func fanIn(ctx context.Context, chs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	// 出力用のchannelを作成
	out := make(chan string)
	// multiplex関数を定義
	// 読み取り専用channelを受け取る
	multiplex := func(ch <-chan string) {
		defer wg.Done()
		// for文でchannelに書き込みがあるたびに値を受け取りoutに送る
		for text := range ch {
			select {
			case <-ctx.Done():
				return
			case out <- text:
			}
		}
	}
	wg.Add(len(chs))
	// 可変長のchannelの数だけgoroutineを立ち上げる
	for _, ch := range chs {
		go multiplex(ch)
	}
	// 立ち上げたgoroutineがwg.Doneするまで待つ
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
