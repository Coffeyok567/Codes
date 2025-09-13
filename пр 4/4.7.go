//Корчагин Евгений 363

package main

import (
	"fmt"
	"sync"
)

type State struct {
	value int
}

type Command struct {
	action   string
	value    int
	response chan int
}

func Manager(commands <-chan Command, wg *sync.WaitGroup) {
	defer wg.Done()

	state := State{value: 0}

	for cmd := range commands {
		switch cmd.action {
		case "get":
			cmd.response <- state.value
		case "set":
			state.value = cmd.value
			cmd.response <- state.value
		case "increment":
			state.value++
			cmd.response <- state.value
		case "decrement":
			state.value--
			cmd.response <- state.value
		case "add":
			state.value += cmd.value
			cmd.response <- state.value
		case "multiply":
			state.value *= cmd.value
			cmd.response <- state.value
		default:
			cmd.response <- state.value
		}
	}
}

func main() {
	var wg sync.WaitGroup
	commands := make(chan Command)

	wg.Add(1)
	go Manager(commands, &wg)

	operations := []struct {
		action string
		value  int
	}{
		{"set", 10},
		{"increment", 0},
		{"add", 5},
		{"multiply", 2},
		{"get", 0},
	}

	for _, op := range operations {
		response := make(chan int)
		commands <- Command{action: op.action, value: op.value, response: response}
		result := <-response
		fmt.Printf("После операции '%s': %d", op.action, result)
	}

	close(commands)
	wg.Wait()
	fmt.Println("Управление завершено")
}
