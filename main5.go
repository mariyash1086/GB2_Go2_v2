package main

import (
	"fmt"
	"sync"
)

func main() {
	third()
}

//1.Напишите программу, которая запускает 𝑛 потоков и дожидается завершения их всех
func first() {
	var wg = sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {

			fmt.Println("Работает поток : ", i)
			wg.Done()

		}(i)

	}
	wg.Wait()

	fmt.Println("Дождались окончание 10 потоков.")

}

//2.Реализуйте функцию для разблокировки мьютекса с помощью defer
func second(ourMutex sync.Mutex, err error) string {
	ourMutex.Lock()
	defer ourMutex.Unlock()

	if err != nil {
		return "Возникла ошибка работы, разблокировали mutex"
	}
	return "Разблокировали mutex"
}

//3.тут я попыталась изобрать работу с каналом и значение
//и в зависимости от значения делать lock или rlock...не удалось.

func third() {
	var (
		rlock sync.RWMutex
		Lock  sync.Mutex
		wg    = sync.WaitGroup{}
		//	j  = 0
		ch chan int
	)

	ch <- 0
	for i := 0; i < 16; {
		wg.Add(1)
		fmt.Println(i)

		i = <-ch
		if i < 4 {

			go func() {
				Lock.Lock()
				i++
				ch <- i
				Lock.Unlock()
				wg.Done()
			}()

		} else {

			go func(i int) {
				rlock.RLock()
				i++
				ch <- i
				rlock.RLock()
				wg.Done()
			}(i)

		}
	}
	wg.Wait()
}
