package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)

	return w
}

func doWorker(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker id: %d received %d\n", id, n)
		done <- true
	}
}

func channelDemo() {
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- i
	}

	for i := 0; i < 10; i++ {
		<-workers[i].done
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- i * 100
	}

	for i := 0; i < 10; i++ {
		<-workers[i].done
	}

	// time.Sleep(time.Second)
}

type worker2 struct {
	in chan int
	wg *sync.WaitGroup
}

func createWorker2(id int, wg *sync.WaitGroup) worker2 {
	w := worker2{
		in: make(chan int),
		wg: wg,
	}
	go doWorker2(id, w.in, w.wg)

	return w
}

func doWorker2(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("Worker id: %d received %d\n", id, n)
		wg.Done()
	}
}

func channelDemo2() {
	var wg sync.WaitGroup
	var workers [10]worker2
	for i := 0; i < 10; i++ {
		workers[i] = createWorker2(i, &wg)
	}
	wg.Add(20)

	for i, worker := range workers {
		worker.in <- i
	}

	for i, worker := range workers {
		worker.in <- i * 100
	}

	wg.Wait()
}

func main() {
	// channelDemo()
	channelDemo2()
}
