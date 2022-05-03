package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {

	//1-2.
	defer func() {
		if v := recover(); v != nil {

			fmt.Println("it was a a panic  situation: ", v)
			err := errors.New("My error")
			fmt.Println("it was a error situation: ", err)
			fmt.Println("Error time is ", time.Now())
		}

	}()
	dopanic()

	//3.
	makeFile()
}

func dopanic() {
	new_array := [3]int{1, 2, 3}

	var len_array = len(new_array) + 1

	fmt.Println(new_array[len_array+1])
}

func makeFile() {

	myFile, err := os.Create("New file")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		myFile.Close()
	}()
}
