package main

import (
	"fmt"
	"time"
)

func main() {

	var c = make(chan int)
	go process(c)
	t0 := time.Now()
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
	fmt.Println(time.Since(t0))
}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 3; i++ {
		c <- i
		c <- i * 2
	}
	fmt.Println("process end")
}
