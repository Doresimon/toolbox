package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	listenAddr := "0.0.0.0:8088"
	// Hello world, the web server
	var savedData string

	putHandler := func(w http.ResponseWriter, req *http.Request) {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		req.Body.Close()
		savedData = string(data) + ", " + time.Now().String()

		fmt.Printf("info received: %s\n", savedData)
		w.WriteHeader(200)
	}
	getHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, savedData)
	}

	http.HandleFunc("/put", putHandler)
	http.HandleFunc("/get", getHandler)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
