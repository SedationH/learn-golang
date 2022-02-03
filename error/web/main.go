package main

import (
	"fmt"
	"learn-golang/error/web/list"
	"net/http"
	"os"
)

type handlerType func(http.ResponseWriter, *http.Request) error
type errorHandlerType func(http.ResponseWriter, *http.Request)

type userError interface {
	error
	Message() string
}

func errorWrapper(handler handlerType) errorHandlerType {
	return func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Panic: %v", r)
				http.Error(rw,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(rw, r)
		if err != nil {
			if userErr, ok := err.(userError); ok {
				http.Error(rw, userErr.Message(), http.StatusBadRequest)
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(rw, http.StatusText(code), code)
		}
	}
}

func main() {
	// http.HandleFunc("/list/", errorWrapper(list.ListHandler))
	http.HandleFunc("/", errorWrapper(list.ListHandler))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
