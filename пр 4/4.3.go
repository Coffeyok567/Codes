//Корчагин Евгений 363

package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(200 * time.Millisecond)
	requast := make(chan int, 15)
	for i := 1; i <= 15; i++ {
		requast <- i
	}
	close(requast)

	for req := range requast {
		<-tick
		fmt.Println("Обработка запроса:", req)
	}

}
