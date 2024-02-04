package heartbeatwithwatchdogtimer

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func DemoHeartBeat() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	// MultiWriterでlogをファイルと標準出力に出力
	// log.LstdFlagsでログに日付を出力
	errorLogger := log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.LstdFlags)
	ctx, cancel := context.WithTimeout(context.Background(), 5100*time.Millisecond)
	defer cancel()
	const wdtTimeout = 800 * time.Millisecond
	const beatInterval = 500 * time.Millisecond
	heartbeat, v := task(ctx, beatInterval)

	// ウォッチドッグタイマーの設定
loop:
	// どれかのcaseを実行するとselectを抜ける
	// forなので次のselectを実行する
	// その際にtime.Afterが再設定される
	for {
		// selectは同じcaseが複数ある場合はランダムで選択される
		select {
		// heartbeatから受信した場合はbeat pulse ⚡️を出力
		case _, ok := <-heartbeat:
			// heartbeatがクローズされた場合はloopを抜ける
			if !ok {
				break loop
			}
			fmt.Println("beat pulse ⚡️")
		// vから受信した場合は値を出力
		case r, ok := <-v:
			// vがクローズされた場合はloopを抜ける
			if !ok {
				break loop
			}
			// m=でモノトニック時刻を取得(経過時間を取得するため)
			// 例: m=+0.500000000sなのでmがt[0]、+0.500000000sがt[1]になる
			t := strings.Split(r.String(), "m=")
			fmt.Printf("value: %v [s]\n", t[1])
		// 800msのウォッチドッグタイマーのタイムアウト
		case <-time.After(wdtTimeout):
			errorLogger.Println("doTask goroutine's heartbeat is stopped")
			break loop
		}
	}
}

func task(
	ctx context.Context,
	// ハートビートのインターバル
	beatInterval time.Duration,
	// データなしの読み取り専用のチャネル
	// タイム型の読み取り専用のチャネル
) (<-chan struct{}, <-chan time.Time) {
	heartbeat := make(chan struct{})
	out := make(chan time.Time)
	go func() {
		defer close(heartbeat)
		defer close(out)
		pulse := time.NewTicker(beatInterval)
		task := time.NewTicker(2 * beatInterval)
		sendPulse := func() {
			select {
			// heartbeatが受信ができる場合は空の構造体(struct{}{})をインスタンス化して送信
			// struct{}は宣言するときに使う struct{}{} という書き方はインスタンス化するときに使う
			// 空の構造体はメモリ効率が良いので通知するだけならこれを使う
			case heartbeat <- struct{}{}:
			// ハートビートの送信ができない場合は何もしないでselect文を抜ける
			default:
			}
		}
		sendValue := func(t time.Time) {
			for {
				select {
				// timeout付きのctxなのでその場合の処理
				case <-ctx.Done():
					return
				// pulseのチャネルから値を受信(<-pulse.C)できる場合はsendPulse()を実行しメインルーチンにハートビートを送信
				// case文を抜けるがforなので再度ループする
				case <-pulse.C:
					sendPulse()
				// outチャネルが受け取れる状態なら受け取った時刻を送信(out<-)し無名関数を抜ける
				case out <- t:
					return
				}
			}
		}
		var i int
		for {
			select {
			case <-ctx.Done():
				return
			case <-pulse.C:
				// pulseに受信するたびにiをインクリメントし3のときに1秒スリープすることで異常を発生させる
				if i == 3 {
					time.Sleep(1 * time.Second)
				}
				sendPulse()
				i++
			case t := <-task.C:
				sendValue(t)
			}
		}
	}()
	// ハートビートのチャネルと出力チャネルを返す
	return heartbeat, out
}
