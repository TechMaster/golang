package main

import (
	"fmt"
	"math"
	"time"
)

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
	if bmi := weight / (height * height); bmi < 18.5 {
		return "Underweight"
	} else if bmi < 25 {
		return "Normal"
	} else {
		return "Overweight"
	}
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
func main() {
	fmt.Println(Quarter("May"))
	var x = "Cuong"
	var y = "Dung"
	x, y = y, x //Swap two numbers
	fmt.Println(x, y)
}
