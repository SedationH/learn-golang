package main

import (
	"learn-golang/infa"
	"learn-golang/testing"
)

func getRetriver() infa.Retriever {
	return infa.Retriever{}
}

func getTestingRetriver() testing.Retriever {
	return testing.Retriever{}
}

func main() {
	var retriever infa.Retriever = getRetriver()
	print(retriever.Get("http://baidu.com"))
	testingRetriever := getTestingRetriver()
	print(testingRetriever.Get("http://baidu.com"))
}
