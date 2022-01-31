package main

import (
	"fmt"
	"learn-golang/retriever/mock"
	"learn-golang/retriever/real"
	"time"
)

const url = "http://baidu.com"

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func inspect(r Retriever) {
	// use of .(type) outside type switchcompilerBadTypeKeyword
	// type asserstion
	switch r.(type) {
	case mock.Retriever:
		{
			print(1)
		}
	case *mock.Retriever:
		{
			print(2)
		}
	}
}

func main() {
	// 值接收者更加宽松
	mock := &mock.Retriever{"mock"}
	// mock2 := &mock.Retriever{"mock"}
	println(download(mock))
	real := &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	println(download(real))

	inspect(mock)

	fmt.Printf("mock: %v %T real: %v %T\n", mock, mock, real, real)
}
