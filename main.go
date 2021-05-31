package main

import (
	"fmt"
	"strings"
)

func main() {

	totalLength, upperName := lenAndUpper("koo")
	fmt.Println(totalLength, upperName)

}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)

}

func multiply(a, b int) int {
	return a * b
}
