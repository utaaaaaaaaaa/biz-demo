package main

import "fmt"

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])
	fmt.Println(arr[:])

	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(&slice[5])
	fmt.Printf("%v,%v\n", len(slice), cap(slice))
	slice = append(slice, 6, 7)
	fmt.Println(slice[3:])
	fmt.Println(slice)
	fmt.Println(&slice[4])
	fmt.Println(&slice[5])
	fmt.Println(&slice[6])
	fmt.Printf("%v,%v\n", len(slice), cap(slice))

	slice2 := []int{11, 12}
	fmt.Println(append(slice, slice2...))

	slice3 := make([]int, 3, 8)
	fmt.Println(slice3, " ", len(slice3), " ", cap(slice3))

	map1 := make(map[int]int)
	fmt.Println(map1)
	var map2 = map[string]uint8{"amy": 33, "lenerd": 35, "penny": 29}
	fmt.Println(map2["amy"])
	var age, ok = map2["penny"]
	if ok {
		fmt.Printf("age is : % v\n", age)
	} else {
		fmt.Println("not exist")
	}

	for key, value := range map2 {
		fmt.Printf("key:%v,value:%v\n", key, value)
	}
	delete(map2, "penny")
	for key, value := range map2 {
		fmt.Printf("key:%v,value:%v\n", key, value)
		//}
		//var i int = 0
		//for i < 10 {
		//	fmt.Println(i)
		//	i++
		//}
		//for j := 0; j < 10; j++ {
		//	fmt.Println("utaaa")
		//}
	}

	var str = []rune("resume")
	fmt.Println(str)
}
