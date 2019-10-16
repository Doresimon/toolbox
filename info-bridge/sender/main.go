package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

func main() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	sendIP()
	for {
		select {
		case <-ticker.C:
			sendIP()
		}
	}
}

func sendIP() error {
	selfIP, err := getIP()
	if err != nil {
		return err
	}
	fmt.Printf("send ip %s\n", selfIP)
	host := "YOURIP"
	port := "8088"
	path := "put"
	url := fmt.Sprintf("http://%s:%s/%s", host, port, path)

	_, err = http.Post(url, "text/plain", strings.NewReader(selfIP))
	return err
}

func getIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, v := range addrs {
		if ipnet, ok := v.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", nil
}
