package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		fmt.Printf("Worker id: %d received %d\n", id, <-c)
	}
}

func channelDemo() {
	var channels [10]chan int

	// for i, c := range channels {
	// 	c = make(chan int)
	// 	go worker(i, c)
	// }

	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	// for i, c := range channels {
	// 	c <- i
	// }

	for i := 0; i < 10; i++ {
		channels[i] <- i
	}

	time.Sleep(time.Microsecond)
}

func createWorker(id int) chan int {
	c := make(chan int)

	go func() {
		for {
			fmt.Printf("Worker id: %d received %d\n", id, <-c)
		}
	}()

	return c
}

func channelDemo2() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- i
	}
	time.Sleep(time.Second)
}

// 告诉外面channel是如何用的
// chan<- 表示返回的是 send-only
// <-chan 表示返回的是 receive-only
func createWorker2(id int) chan<- int {
	c := make(chan int)

	go func() {
		for {
			fmt.Printf("Worker id: %d received %d\n", id, <-c)
		}
	}()

	return c
}

func channelDemo3() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker2(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- i
	}
	time.Sleep(time.Second)
}

func channelBuffer() {
	c := make(chan int, 3)

	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Second)
}

func worker2(id int, c chan int) {
	// for {
	// 	n, ok := <-c
	// 	if !ok {
	// 		println("channel 被 close 暂停接收")
	// 		break
	// 	}
	// 	fmt.Printf("Worker id: %d received %d\n", id, n)
	// }

	for n := range c {
		fmt.Printf("Worker id: %d received %d\n", id, n)
	}
}

func channelDemo4() {
	c := make(chan int, 3)

	go worker2(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4

	close(c)

	time.Sleep(time.Second)
}

func main() {
	// channelDemo()
	// channelDemo2()
	// channelDemo3()
	// channelBuffer()
	channelDemo4()
}
