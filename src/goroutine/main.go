package main

import (
	"fmt"
	"sync"
)

// 单向通道
func send(ch chan <- int, value int) {
	ch <- value
}

// 单向通道
func receive(ch <- chan int)  {
	value := <- ch
	fmt.Println(value)
}

func main() {
	// case 1 测试单向通道
	//ch := make(chan int)
	//go send(ch, 1)
	//receive(ch)

	// case 2 测试带缓冲通道
	//ch := make(chan int, 5)
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		ch <- i
	//	}
	//	close(ch)
	//}()
	//for val := range ch {
	//	fmt.Println(val)
	//}

	//case 3 扇入
	ch := make(chan int, 5)
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	go func() {
		for i := 5; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 10; i < 15; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	for val := range merge(ch, ch1, ch2) {
		fmt.Printf("get number from out: %d\n", val)
	}
}

func merge(channels ...<- chan int) <- chan int {
	var wg sync.WaitGroup
	out := make(chan int, 15)
	output := func(ch <- chan int) {
		for val := range ch {
			out <- val
			fmt.Printf("put number on out: %d\n", val)
		}
		wg.Done()
	}
	wg.Add(len(channels))
	for _, ch := range channels {
		go output(ch)
	}
	wg.Wait()
	go func() {
		close(out)
	}()
	return out
}