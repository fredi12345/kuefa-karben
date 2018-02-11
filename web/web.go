package web

import (
	"net/http"
	"path"
	"log"
	"io/ioutil"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	data, err := ioutil.ReadFile(path.Join("resources", "index.html"))
	if err != nil {
		log.Fatal(err)
	}

	w.Write(data)
}
