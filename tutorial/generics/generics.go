package main

import "fmt"

func main() {
	var sliceInt = []int{1, 2, 3}
	var sliceFloat64 = []float64{1.1, 2.2, 3.3}
	var sliceFloat32 = []float32{1.1, 2.2, 3.3}
	fmt.Println(sliceInt, sliceFloat64, sliceFloat32)
	fmt.Println(sumSlice(sliceInt))
	fmt.Println(sumSlice(sliceFloat32))
	fmt.Println(sumSlice(sliceFloat64))
}

func sumSlice[T int | float32 | float64](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}
