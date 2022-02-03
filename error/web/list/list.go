package list

import (
	"io/ioutil"
	"net/http"
	"os"
)

type userError string

func (u userError) Error() string {
	return u.Message()
}

func (u userError) Message() string {
	return string(u)
}

func ListHandler(rw http.ResponseWriter, r *http.Request) error {
	// if !strings.HasPrefix(r.URL.Path, "/list/") {
	// 	return userError("path must start with /list/")
	// }
	path := r.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	rw.Write(all)
	return nil
}
