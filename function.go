package main

import (
	"fmt"
	"strings"
)

func main() {
	// greet("Raditeo", 23)
	// greet("Raditeo", "test")

	// var names = []string{"Raditeo", "Warma"}
	// var printMsg = greet("Heii", names)

	// fmt.Println(printMsg)

	// var diameter float64 = 15

	// var area, circumference = calculate(diameter)

	// fmt.Println("Area:", area)
	// fmt.Println("Circumference:", circumference)

	// studentLists := print("Raditeo", "Warma", "test1", "test2")

	// fmt.Printf("%v", studentLists)

	// numbersLists := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// result := sum(numbersLists...)

	// fmt.Println("result:", result)

	profile("Raditeo", "mie ayam")
}

// func greet(name string, age int8) {
// 	fmt.Printf("Hello there! my name is %s and I'm %d years old", name, age)
// }

// func greet(name, address string) {
// 	fmt.Printf("Hello there! my name is %s and I live in %s", name, address)
// }

// func greet(msg string, names []string) string {
// 	var joinStr = strings.Join(names, " ")

// 	var result string = fmt.Sprintf("%s %s", msg, joinStr)

// 	return result
// }

// func calculate(d float64) (float64, float64) {
// 	var area float64 = math.Pi * math.Pow(d/2, 2)

// 	var circumference = math.Pi * d

// 	return area, circumference
// }

// func calculate(d float64) (area float64, circumference float64) {
// 	area = math.Pi * math.Pow(d/2, 2)

// 	circumference = math.Pi * d

// 	return
// }

// func print(names ...string) []map[string]string {
// 	var result []map[string]string
// 	for i, v := range names {
// 		key := fmt.Sprintf("Student%d", i+1)
// 		temp := map[string]string{
// 			key: v,
// 		}

// 		result = append(result, temp)
// 	}

// 	return result
// }

// func sum(numbers ...int) int {
// 	total := 0

// 	for _, v := range numbers {
// 		total += v
// 	}

// 	return total
// }

func profile(name string, favFoods ...string) {
	mergeFoods := strings.Join(favFoods, ",")

	fmt.Println("hello there! I'm", name)
	fmt.Println("I really love to eat", mergeFoods)
}
