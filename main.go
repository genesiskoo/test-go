package main

import (
	"fmt"
	"strings"
)

func main() {

	// totalLength, upperName := lenAndUpper("koo")
	// fmt.Println(totalLength, upperName)
	// repeatMe("koo", "hi", "many", "cool")

	// result := supperAdd(3, 2, 5, 2, 5, 6)
	// fmt.Println(result)

	result := canIDrink(15)
	fmt.Println(result)

}

func repeatMe(words ...string) {
	fmt.Println(words)
}

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)

// }

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func supperAdd(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
		fmt.Println(total)
	}
	return total

}

func lenAndUpper(name string) (lenth int, upperName string) {
	defer fmt.Println("lenAndUpper Fucn is Done")
	lenth = len(name)
	upperName = strings.ToUpper(name)
	return

}

func multiply(a, b int) int {
	return a * b
}
