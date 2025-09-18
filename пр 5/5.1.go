package main

import (
	"fmt"
	"time"
)

func main(){
	result := make(chan struct{})
	go func(){
		time.Sleep(3 * time.Second)
		result <- struct{}{}
	}()

	select{
	case <- result: 
		fmt.Println("Ответ имеется")
	case <- time.After(2 * time.Second):
		fmt.Println("Время ожидания превышенно")
	}
}