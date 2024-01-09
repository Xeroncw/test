package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start")

	go func ()  {
		fmt.Println("goroutine start")
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("main end")
}
