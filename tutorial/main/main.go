package main

import (
	"fmt"
	"runtime"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	//path := os.Getenv("PATH")
	//fmt.Printf("Path is %s\n", path)
	var t = 7
	var a = &t
	*a = *a + 1
	b := a
	fmt.Println(*b, *a)
	fmt.Println(&a, a)

	slice := []int{1, 2, 3}
	slice2 := slice
	fmt.Println(&slice2[0], &slice[0])
}
