package tracer

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func DemoTracer() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	defer func() {
		// file closeした返り値をerrに代入しそれを条件式に使っている
		if err := f.Close(); err != nil {
			log.Fatalln("Error: ", err)
		}
	}()

	// traceの開始
	if err := trace.Start(f); err != nil {
		log.Fatalln("Error: ", err)
	}
	defer trace.Stop()
	ctx, t := trace.NewTask(context.Background(), "tracer")
	defer t.End()
	fmt.Println("The number of logical CPUs Cores ", runtime.NumCPU())
	// これは逐次処理 1sec x 3かかる
	// task(ctx, "task1")
	// task(ctx, "task2")
	// task(ctx, "task3")
	var wg sync.WaitGroup
	wg.Add(3)
	// これは並行処理 1secで終わる
	go cTask(ctx, &wg, "task1")
	go cTask(ctx, &wg, "task2")
	go cTask(ctx, &wg, "task3")
	wg.Wait()

	s := []int{1, 2, 3}
	for _, i := range s {
		wg.Add(1)
		// 無名関数をgo funcで並行処理
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()

	fmt.Println("main finished")
}

// func task(ctx context.Context, name string) {
// 	// deferはチェーンの最後のみ実行される この場合はEnd()のみが遅延して実行される
// 	// StartRegionはすぐに実行される
// 	defer trace.StartRegion(ctx, name).End()
// 	time.Sleep(time.Second)
// 	fmt.Println(name)
// }

func cTask(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer trace.StartRegion(ctx, name).End()
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println(name)
}
