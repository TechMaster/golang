package main

import (
	"fmt"
	"math"
	"time"
)

var name = "random"

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("math: square root of negative number %g", f)
	}
	return math.Sqrt(f), nil
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func bmi(weight float32, height float32) string {
	var result = ""
	bmi := weight / (height * height)
	if bmi < 18.5 {
		result = "Underweight"
	} else if bmi < 25 {
		result = "Normal"
	} else {
		result = "Overweight"
	}
	fmt.Println(bmi)
	return result

}
func Greeting() {
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

func Quarter(month string) string {
	switch month {
	case "Jan", "Feb", "Mar":
		return "First Quarter"
	case "Apr", "May", "Jun":
		return "Second Quarter"
	case "Jul", "Aug", "Sep":
		return "Third Quarter"
	case "Oct", "Nov", "Dec":
		return "Forth Quarter"
	default:
		return "Unknown Quarter"
	}
}

func demoLoop() {
	cars := [3]string{"Toyota", "Mercedes", "BMW"}
	//fmt.Println(cars[0]) // Toyota
	for index, car := range cars {
		defer fmt.Println(index, car)
	}
}

func giaiThua(n int) int {
	if n == 1 {
		return 1
	}
	return n * giaiThua(n-1)
}

func demoArray2D() {
	langs := [][]string{{"C#", "C", "Python"},
		{"Java", "Scala", "Perl"},
		{"C++", "Go", "RUST", "Crystal", "OCAML"}}
	for _, v := range langs {
		for _, lang := range v {
			fmt.Print(lang, " ")
		}
		fmt.Println()
	}

}

func demoSlice() {
	letters := []string{"a", "b", "c", "d"}
	letters = append(letters, "e")
	length := len(letters)
	fmt.Println(length)
	fmt.Println(letters[:2])
}
func main() {
	demoSlice()
}
