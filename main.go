package main

import "fmt"

func add() {
	fmt.Println("added")
}

func main() {
	var array = []int{1, 2, 3, 4}
	var age int = 66
	var str string = "hy"
	fmt.Println(array)
	fmt.Print(str)
	fmt.Print(age)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	if age > 55 {
		fmt.Print("hey \n")

	} else {
		fmt.Println("hi")
	}

	fmt.Println(array[0:4])
	add()
}
