package main

import "fmt"

func main() {
	// var firstNumber *int
	// var secondNumber *int
	// _, _ = firstNumber, secondNumber

	// var firstNumber int = 4
	// var secondNumber *int = &firstNumber

	// fmt.Println("firstNumber value:", firstNumber)
	// fmt.Println("firstNumber memory address:", &firstNumber)

	// fmt.Println("secondNumber value:", *secondNumber)
	// fmt.Println("secondNumber memory address:", secondNumber)

	// var firstPerson string = "Raditeo"
	// var secondPerson *string = &firstPerson

	// fmt.Println("firstPerson value:", firstPerson)
	// fmt.Println("firstPerson memory address:", &firstPerson)

	// fmt.Println("secondPerson value:", *secondPerson)
	// fmt.Println("secondPerson memory address:", secondPerson)

	// fmt.Print("\n", strings.Repeat("#", 25), "\n")

	// *secondPerson = "Warma"

	// fmt.Println("firstPerson value:", firstPerson)
	// fmt.Println("firstPerson memory address:", &firstPerson)

	// fmt.Println("secondPerson value:", *secondPerson)
	// fmt.Println("secondPerson memory address:", secondPerson)

	var a int = 10

	fmt.Println("before:", a)

	changeValue(&a)

	fmt.Println("after:", a)
}

func changeValue(number *int) {
	*number = 123424
}
