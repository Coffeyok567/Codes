package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup, stop <-chan bool) {
    defer wg.Done()
    
    for {
        select {
        case <-stop:
            fmt.Printf("работник %d завершает работу\n", id)
            return
        default:
            fmt.Printf("работник %d выполняет задачу\n", id)
            time.Sleep(1 * time.Second)
        }
    }
}

func main() {
    var wg sync.WaitGroup
    stop := make(chan bool)

    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg, stop)
    }

    time.Sleep(2 * time.Second)
    close(stop) 

    wg.Wait()
    fmt.Println("Все завершенно")
}