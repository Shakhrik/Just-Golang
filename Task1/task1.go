package main

import (
	"fmt"
	"strings"
)

// Fibonachchi
func fibonachchi(number int) {
	first := 0
	second := 1
	for i := 0; i < number; i++ {
		if number-i == 1 {
			fmt.Println(first)
		} else {
			fmt.Print(first, "\t")
		}
		temp := first
		first = second
		second += temp
	}

}

// FizzBuzz
func fizzbuzz(upto int) {
	for i := 1; i < upto; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz", "\t")
		} else if i%3 == 0 {
			fmt.Print("Fizz", "\t")
		} else if i%5 == 0 {
			fmt.Print("Buzz", "\t")
		} else {
			fmt.Print(i, "\t")
		}
		if upto-i == 1 {
			fmt.Printf("\n")
		}
	}
}

// Palindrom
func palindrom(text string) (newText string) {
	for i := len(text) - 1; i >= 0; i-- {
		newText += string(text[i])
	}
	return strings.ToLower(newText)

}
func oddEvenSum(maxNumber int) (oddSum int, evenSum int) {
	oddSum = 0
	evenSum = 0
	for i := 0; i < maxNumber; i++ {
		if i%2 == 0 {
			evenSum += i
		} else {
			oddSum += i
		}
	}
	return oddSum, evenSum
}
func main() {
	// fibonachchi(10)
	// fizzbuzz(50)
	// println(palindrom("Shakhriyor Ismatov"))
	var (
		input int
		odd   int
		even  int
	)
	print("Enter number you want: ")
	fmt.Scan(&input)
	odd, even = oddEvenSum(input)
	println("Odd numbers sum until", input, "is: ", odd, "\nEven numbers sum until", input, "is: ", even)
}
