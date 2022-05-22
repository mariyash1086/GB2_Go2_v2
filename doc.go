package Go2

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

//It is a main function . Handles panic 1 and panic 2, prints the result to the console
func DoPanic() {

	defer func() {
		value := recover()
		err, ok := value.(error)
		if !ok {
			fmt.Println("Its panic")
		}

		if err != nil {
			log.Printf("recoveres %s", err)
		}
	}()
	p1 := CreateFirstPanic()
	p2 := CreateSecondPanic()
	fmt.Println(p1)
	fmt.Println(p2)
}

//Сreates a panic implicitly by getting a non-existent array element.
func CreateFirstPanic() int {

	arr := [3]int{0, 1, 2}

	var index = len(arr) + 1
	return arr[index]

}

//Сreates a panic implicitly by dividing by zero.
func CreateSecondPanic() int {

	x := 5 / 0
	return x
}

//Creates and saves a million files locally.
func DoFiles() error {
	for i := 1; i < 1000000; i++ {
		fileName := "file" + strconv.Itoa(i) + ".txt"

		file, err := os.Create(fileName)

		func() {
			defer file.Close()
		}()
		if err != nil {
			return err
		}

	}
	return nil
}
