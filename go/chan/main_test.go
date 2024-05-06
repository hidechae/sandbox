package main

import (
	"fmt"
	"testing"
	"time"
)

// chanの送信と受信
func Test_Basic(t *testing.T) {
	ch := make(chan int)

	// 送信
	go func() { ch <- 1 }()

	// 受信
	fmt.Println(<-ch)
}

// NG capを超えるとdeadlockする
func Test_Basic_NG(t *testing.T) {
	ch := make(chan int, 1)

	// 送信
	ch <- 1
	// capを超えているのでdeadlockする
	ch <- 2

	// 受信
	v := <-ch
	fmt.Println(v)
}

// 単一チャンネルからすべての値を受信したいときはrangeを使う
func Test_Range(t *testing.T) {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	// rangeはcloseされるまでブロックする
	for v := range ch {
		fmt.Println(v)
	}
}

// selectを使ってチャンネルの送受信を監視する
func Test_Select(t *testing.T) {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				// closeされたらnilを代入することで、次回のselectでcaseが無視される
				ch = nil
				break
			}
			fmt.Println(v)
		}

		// チャンネルがcloseされたら終了
		if ch == nil {
			break
		}
	}
}

// 複数のチャンネルを使う
// タイムアウトの処理をする
func Test_Select_Timeout(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		time.Sleep(1 * time.Second)
		close(ch1)
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 2
		close(ch2)
	}()

	for {
		select {
		case v, ok := <-ch1:
			if ok {
				fmt.Printf("ch1 received %d\n", v)
			} else {
				// closeされた後はブロックされずにokがfalseになりここに入ってくる
				fmt.Println("ch1 closed")
				// nilを代入すると、次回のselectでcaseが無視される
				ch1 = nil
			}
		case v, ok := <-ch2:
			if ok {
				fmt.Printf("ch2 received %d\n", v)
			} else {
				fmt.Println("ch2 closed")
				ch2 = nil
			}
		case <-time.After(500 * time.Millisecond):
			// タイムアウト時の処理
			fmt.Println("timeout")
		default:
			// default句があればブロックされずにここに入ってくる
		}

		// すべてのチャンネルがcloseされたら終了
		if ch1 == nil && ch2 == nil {
			fmt.Println("All channels are closed")
			break
		}
	}
}
