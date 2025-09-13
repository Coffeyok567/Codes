//Корчагин Евгений 363

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = "https://valorant/net/http"
var taskCount = 3
var wg sync.WaitGroup

func getURL(url string) {
	_, err := http.Get(url)
	fmt.Println(err)
	defer wg.Done()
}

func main() {
	for i := 1; i <= taskCount; i++ {
		wg.Add(1)
		go getURL(urls)
		fmt.Println(i)

	}
	wg.Wait()
	fmt.Print("Занято")
}
