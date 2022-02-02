package list

import (
	"io/ioutil"
	"net/http"
	"os"
)

func ListHandler(rw http.ResponseWriter, r *http.Request) {
	println(r.URL.Path)
	path := r.URL.Path[len("/list/"):]
	println(path)
	file, err := os.Open(path)
	if err != nil {
		println(1, err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		println(2, err)
		return
	}
	rw.Write(all)
}
