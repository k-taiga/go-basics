package pipeline

import (
	"context"
	"fmt"
)

func DemoPipeline() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nums := []int{1, 2, 3, 4, 5}
	var i int
	flag := true
	// pipelineの実行 (generator -> double -> offset -> double)
	// 一番外のdoubleをvで受け取りそれを出力する
	// doubleのchannelがクローズされるまでforで回す
	// 呼び出すたびにiが増える
	for v := range double(ctx, offset(ctx, double(ctx, generator(ctx, nums...)))) {
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

// 返り値は読み取り専用のchannel
func generator(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		// 関数が終了したらchannelを閉じる
		defer close(out)
		// 受け取ったsliceをrangeで回して、それぞれの値をchannelに送る
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			// 取り出した値をchannelに送る
			case out <- n:
			}
		}
	}()
	return out
}

func double(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n * 2:
			}
		}
	}()
	return out
}

func offset(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n + 2:
			}
		}
	}()
	return out
}
