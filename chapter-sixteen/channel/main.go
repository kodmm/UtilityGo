package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	ic := make(chan int)
	go func() {
		ic <- 10
		ic <- 20
		close(ic) // クローズするとループ解除
	}()
	// チャネルに対してループ
	for v := range ic {
		fmt.Println(v)
	}
}

func recv(r <-chan string) {
	v := <-r
	// r <- "送信エラー" // error
}

func send(s chan<- string) {
	s <- "送信は可能"
	// v := <-s // エラー
}

type Account struct {
	balance int
	lock    sync.RWMutex
}

func (a *Account) GetBalance() int {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.balance
}

func (a *Account) Transfer(amount int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.balance += amount
}

type Account3 struct {
	balance int64
}

func (a Account3) GetBalance() int64 {
	return a.balance
}

func (a *Account3) Transfer(amount int64) {
	atomic.AddInt64(&a.balance, amount)
}
