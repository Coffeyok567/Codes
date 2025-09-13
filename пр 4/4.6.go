//Корчагин Евгений 363

package main

import (
	"fmt"
	"sync"
)

func processDatabase(dbName string, results chan<- string, wg *sync.WaitGroup, done <-chan struct{}) {
	defer wg.Done()

	select {
	case <-done:
		return
	case results <- "Result from " + dbName:
	}
}

func main() {
	var wg sync.WaitGroup
	results := make(chan string, 1)
	done := make(chan struct{})

	DataBases := []string{
		"DataBase1",
		"DataBase2",
		"DataBase3",
		"DataBase4",
		"DataBase5",
		"DataBase6",
	}

	for _, dbName := range DataBases {
		wg.Add(1)
		go processDatabase(dbName, results, &wg, done)
	}

	result := <-results
	fmt.Println(result)

	close(done)

	wg.Wait()
}
