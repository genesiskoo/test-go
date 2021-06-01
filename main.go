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

	// result := canIDrink(15)
	// fmt.Println(result)

	// POINTER
	// a := 2
	// b := &a
	// c := &a
	// fmt.Println(a, *b, c)

	//&포인터 지정, 메모리값
	//포인터 지정시 해당 메모리를 값으로 가짐

	// & address
	// * see though

	// ARRAY & SLICE
	// name := []string{"one", "two", "three"}
	// name = append(name, "four")

	// fmt.Println(name)

	// MAP
	// koo := map[string]string{"name": "koo", "age": "15"}
	// fmt.Println(koo)
	// for key, value := range koo {
	// 	fmt.Println(key, value)
	// }

	favFood := []string{"kim", "chi"}
	koo := person{name: "koo", age: 21, favFood: favFood}
	fmt.Println(koo)

}

type person struct {
	name    string
	age     int
	favFood []string
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
