package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var wg = sync.WaitGroup{}
var slice = []string{}
var mutex = sync.Mutex{}
var m = sync.RWMutex{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait() //wait for the count equals to zero
	fmt.Println("\nTotal execution time:", time.Since(t0))
	fmt.Println("the results are", slice)
}

func dbCall(i int) {
	//Simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	//fmt.Println("the result from the database is:", dbData[i])
	//mutex.Lock()
	//slice = append(slice, dbData[i])
	//mutex.Unlock()
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	slice = append(slice, result)
	m.Unlock()
}

func log() {
	//读锁可以让多个协程一起读，同时防止并发读+写导致panic
	m.RLock()
	fmt.Println("the current result is :", slice)
	m.RUnlock()
}
