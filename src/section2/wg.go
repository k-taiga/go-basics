package main

import (
	"fmt"
	"runtime"
	"sync"
)

func wg() {
	var wg sync.WaitGroup
	// wait for 1 goroutine
	wg.Add(1)
	go func() {
		// deferでgoroutineが終了したことを通知 wgのカウントを1減らす
		defer wg.Done()
		fmt.Println("goroutine invoked")
	}()
	// goroutineが終了するまで待機 wgのカウントが0になるまで待機
	wg.Wait()
	fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
	fmt.Println("main function finished")
}
