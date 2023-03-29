package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Done: 1")
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Done: 2")
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Done: 3")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("done all tasks")

	ctx := context.Background()
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("Done: 1")
		return nil
	})
	eg.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("Done: 2")
		return nil
	})
	eg.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("Done: 3")
		return nil
	})
	err := eg.Wait()
	fmt.Println("Done all tasks", err)

	fa, pa := MakeFuturePromise()
	fb, pb := MakeFuturePromise()
	fc, pc := MakeFuturePromise()
	fd, pd := MakeFuturePromise()

	go func() {
		a := A()
		pa.Submit(a)
	}()
	go func() {
		b := B()
		pa.Submit(b)
	}()
	go func() {
		c := C(fa.Get(), fb.Get())
		pa.Submit(c)
	}()
	go func() {
		d := D(fa.Get(), fc.Get())
		pa.Submit(d)
	}()
	log.Printf("d: %d\n", fd.Get())

}

func A() int {
	time.Sleep(time.Second)
	return 10
}

func B() int {
	time.Sleep(time.Second * 2)
	return 5
}

func C(a, b int) int {
	time.Sleep(time.Second * 1)
	return a + b
}

func D(a, c int) int {
	time.Sleep(time.Second)
	return a + c
}

type Future struct {
	value int
	wait  chan struct{}
}

func (f *Future) IsDone() bool {
	select {
	case <-f.wait:
		return true
	default:
		return false
	}
}

func (f *Future) Get() int {
	<-f.wait
	return f.value
}

type Promise struct {
	f *Future
}

func (p *Promise) Submit(v int) {
	p.f.value = v
	close(p.f.wait)
}

func MakeFuturePromise() (*Future, *Promise) {
	f := &Future{
		wait: make(chan struct{}),
	}
	p := &Promise{
		f: f,
	}
	return f, p
}

var tokenContextKey = struct{}{}

func RegisterToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenContextKey, token)
}

func RetrieveToken(ctx context.Context) (string, error) {
	token, ok := ctx.Value(tokenContextKey).(string)
	if !ok {
		return "", errors.New("Token is not registered")
	}
	return token, nil
}
