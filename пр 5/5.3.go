package main

import (
	"fmt"
)

func main(){
	data := make(chan string , 1)
	data <- "а"
	select{
	case info := <- data:
		if info != ""{
			fmt.Println("Данные да ")
		}
	default:
		fmt.Println("Данные нет")
	}
}