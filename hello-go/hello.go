package main

import (
	"fmt"
	"time"
)

func main() {
	go printHello("Hello")
	go printHello("World!")
	fmt.Scanln()
}

func printHello(texto string) {
	for true {
		fmt.Println(texto)
		time.Sleep(1 * time.Second)
	}
}
