package main

import "fmt"

func ch1() {
	ch := make(chan int)

	go func() {
		ch <- 42 // send
	}()

	val := <-ch    
	fmt.Println(val) // 42
}

func ch2() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() { ch1 <- "one" }()
    go func() { ch2 <- "two" }()

    select {
    case msg := <-ch1:
        fmt.Println(msg)
    case msg := <-ch2:
        fmt.Println(msg)
    }
}