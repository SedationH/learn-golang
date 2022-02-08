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
			out <- fmt.Sprintf("id: %d i %d", id, i)
			i++
		}
	}()

	return out
}

func createWorker(id int) chan string {
	worker := make(chan string)
	go func() {
		for n := range worker {
			fmt.Printf("id: %d, Received %s\n", id, n)
		}
		close(worker)
	}()
	return worker
}

func main() {
	var c1, c2 chan string = generator(0), generator(1)
	w := createWorker(0)

	for {
		select {
		case n := <-c1:
			{
				w <- fmt.Sprintf("Received from c1 %s", n)
			}
		case n := <-c2:
			{
				w <- fmt.Sprintf("Received from c2 %s", n)
			}
		default:
			{
				// fmt.Printf("Received Noting")
			}
		}
	}
}
