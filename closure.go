package main

import "fmt"

func main() {
	// var eventNumbers = func(numbers ...int) []int {
	// 	var result []int

	// 	for _, v := range numbers {
	// 		if v%2 == 0 {
	// 			result = append(result, v)
	// 		}
	// 	}

	// 	return result
	// }

	// var numbers = []int{4, 93, 77, 10, 52, 22, 34}

	// fmt.Println(eventNumbers(numbers...))

	// var isPalindrome = func(str string) bool {
	// 	var temp string = ""

	// 	for i := len(str) - 1; i >= 0; i-- {
	// 		temp += string(byte(str[i]))
	// 	}

	// 	return temp == str
	// }("katak")

	// fmt.Println(isPalindrome)

	// var studentLists = []string{"Raditeo", "Warma"}

	// var find = findStudent(studentLists)

	// fmt.Println(find("Raditeo"))

	var numbers = []int{2, 3, 4, 5, 6, 7, 8, 9}

	var find = findOddNumbers(numbers, func(number int) bool { return number%2 != 0 })

	fmt.Println("total odd numbers:", find)
}

// func findStudent(students []string) func(string) string {
// 	return func(s string) string {
// 		var student string
// 		var position int

// 		for i, v := range students {
// 			if strings.ToLower(v) == strings.ToLower(s) {
// 				student = v
// 				position = i
// 				break
// 			}
// 		}

// 		if student == "" {
// 			return fmt.Sprintf("%s doesnt exist!", s)
// 		}

// 		return fmt.Sprintf("we found %s at position %d", s, position+1)
// 	}
// }

func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}

	return totalOddNumbers
}
