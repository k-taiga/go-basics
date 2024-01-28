package channel

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func DemoMutex() {
	// var wg sync.WaitGroup
	// var mu sync.Mutex
	// var i int
	// wg.Add(2)
	// // この2つのgoroutineは同時に実行されるため、i++が競合するためmutexを使う
	// // go -raceで実行すると競合が発生していることがわかる
	// go func() {
	// 	defer wg.Done()
	// 	mu.Lock()
	// 	defer mu.Unlock()
	// 	i++
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	mu.Lock()
	// 	defer mu.Unlock()
	// 	i++
	// }()

	// wg.Wait()
	// fmt.Println(i)

	// var wg sync.WaitGroup
	// var rwMu sync.RWMutex
	// var c int
	// wg.Add(4)
	// go write(&rwMu, &wg, &c)
	// go read(&rwMu, &wg, &c)
	// go read(&rwMu, &wg, &c)
	// go read(&rwMu, &wg, &c)

	// wg.Wait()
	// fmt.Println("finish")

	var wg sync.WaitGroup
	var c int64

	// 5つのgoroutineで10回ずつcに1を足す = 50回
	// atomicはパッケージで提供されている関数を使うことで、複数のgoroutineから安全に変数を操作できる
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				atomic.AddInt64(&c, 1)
			}
		}()
	}
}

func read(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	time.Sleep(10 * time.Millisecond)
	// read lockはunlockしないでも複数のgoroutineがreadできる 書き込みはできない
	mu.RLock()
	// deferでUnlockすることで、関数が終了したときにUnlockされる
	defer mu.RUnlock()
	fmt.Println("read lock")
	fmt.Println(*c)
	time.Sleep(1 * time.Second)
	fmt.Println("read unlock")
}

func write(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	// write lockはunlockするまで同時に1つのgoroutineしか取得できない
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("write lock")
	*c += 1
	time.Sleep(1 * time.Second)
	fmt.Println("write unlock")
}
