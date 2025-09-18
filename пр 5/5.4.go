package main

import (
	"fmt"
	"time"
)

func main(){
	highpriority := make(chan string, 1)
	lowpriority := make(chan string, 1)

	go func(){
		time.Sleep(1 * time.Second)
		highpriority <- "Задача с высоким приоритетом"
	}()

	go func(){
		time.Sleep(2 * time.Second)
		lowpriority <- "Задача с низким приоритетом"
	}()

	for i := 0; i < 2; i++{
	select{
	case msg := <- highpriority:
		fmt.Println(msg)
	case msg := <- lowpriority:
		fmt.Println(msg)
		}
	}
}