package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(id int) chan string {
	out := make(chan string)

	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- fmt.Sprintf(" generator id: %d i %d", id, i)
			i++
		}
	}()

	return out
}

func createWorker(id int) chan string {
	worker := make(chan string)
	go func() {
		for n := range worker {
			// 消耗速度太慢
			time.Sleep(time.Second * 3)
			fmt.Printf("worker id: %d, %s\n", id, n)
		}
		close(worker)
	}()
	return worker
}

func main() {
	var c1, c2 chan string = generator(0), generator(1)
	w := createWorker(0)

	var values []string
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan string
		var activeValue string

		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			{
				values = append(values, n)
			}
		case n := <-c2:
			{
				values = append(values, n)
			}
		case activeWorker <- activeValue:
			{
				values = values[1:]
			}
		case <-time.After(800 * time.Microsecond):
			{
				println("超时")
			}
		case <-tick:
			{
				fmt.Printf("len values :%d\n", len(values))
			}
		default:
			{
				// fmt.Printf("Received Noting")
			}
		}
	}
}
