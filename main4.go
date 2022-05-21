package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//1.
	var workers = make(chan int, 1000)

	for i := 0; i < 1000; {

		go dec(i, workers)
		i = <-workers
		fmt.Println(i)

	}
	//2.
	ourChan := make(chan os.Signal)
	signal.Notify(ourChan)
	select {
	case val := <-ourChan:
		if val == syscall.SIGTERM {
			os.Exit(1)

		}
	default:
		fmt.Println("Нет значения!")
	}

}

func dec(i int, work chan int) {

	work <- i + 1

}
