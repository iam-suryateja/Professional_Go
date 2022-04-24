package main

import (
	"fmt"
)

func doCalculate(x int, y string) (result float64) {
	defaultValue := 150
	if y == "+" {
		result := defaultValue + x
		return float64(result)
	}

	if y == "-" {
		result := defaultValue + x
		return float64(result)
	}

	if y == "/" {
		result := defaultValue + x
		return float64(result)
	}

	if y == "*" {
		result := defaultValue + x
		return float64(result)
	}
	return

}

func main() {
	var result float64
	result = doCalculate(100, "+")
	fmt.Println(result)
	result = doCalculate(50, "-")
	fmt.Println(result)
	result = doCalculate(20, "/")
	fmt.Println(result)
	result = doCalculate(10, "*")
	fmt.Println(result)
}
