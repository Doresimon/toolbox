package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	listenAddr := "0.0.0.0:8088"
	// Hello world, the web server
	var infoMap = make(map[string]string)

	putHandler := func(w http.ResponseWriter, req *http.Request) {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		req.Body.Close()

		arrs := strings.Split(string(data), "@")

		infoMap[arrs[0]] = arrs[1] + ", " + time.Now().String()

		fmt.Printf("info received: %s\n", infoMap[arrs[0]])
		w.WriteHeader(200)
	}
	getHandler := func(w http.ResponseWriter, req *http.Request) {
		resData := ""
		for k, v := range infoMap {
			resData += (k + "@" + v + "\n")
		}
		io.WriteString(w, resData)
	}

	http.HandleFunc("/put", putHandler)
	http.HandleFunc("/get", getHandler)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
