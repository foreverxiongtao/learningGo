package main

import (
	"runtime"
	"fmt"
)

func main() {
	//go sayHello("aaaaaaa")
	//sayHello("hello")
	testChan()
}

func sayHello(String string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(String)
	}
}

func testChan() {
	c := make(chan int)
	var array []int
	array = append(array, 1, 2, 3, 4, 5)
	go sum(array, c)
	result := <-c
	fmt.Println("result is ", result)
}

func sum(arr []int, c chan int) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	c <- sum
}
