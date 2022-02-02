package main

import (
	"learn-golang/error/web/list"
	"net/http"
)

// func errorWrapper(handler func(http.ResponseWriter, *http.Request)) error {

// }

func main() {
	http.HandleFunc("/list/", list.ListHandler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
