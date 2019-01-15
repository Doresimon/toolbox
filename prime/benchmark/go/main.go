package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	var bit int = 256

	p, err := rand.Prime(rand.Reader, bit)

	fmt.Printf("p=%x \n", p)
	fmt.Printf("err=%s \n", err)
}
