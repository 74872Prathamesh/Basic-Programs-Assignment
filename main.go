package main

import (
	"fmt"
	"math"
)

func main() {

	// Program to check whether a number is prime or not.
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	if IsPrime(number) {
		fmt.Println("The given number", number, "is prime")
	} else {
		fmt.Println("The given number", number, "is not prime")
	}

	//program to find sum of fibonaci series upto n.
	var number1 int
	fmt.Println("Enter number to find sum of fibo series upto number term")
	fmt.Scanln(&number1)

	FindSumOffiboSeries(number1)

	//program to find count of even, odd and zero
	var n int
	fmt.Println("Enter length of array elements to count:- ")
	fmt.Scanln(&n)
	var numbers = make([]int, n)
	fmt.Println("Enter the numbers to add in array :")

	for i := 0; i < len(numbers); i++ {
		fmt.Scanln(&numbers[i])
	}

	evenCount, oddCount, zeroCount := CountEvenOddZero(numbers)

	fmt.Println("Count of even numbers:  ", evenCount)
	fmt.Println("Count of odd numbers: ", oddCount)
	fmt.Println("Count of zeros: ", zeroCount)

}

func IsPrime(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func FindSumOffiboSeries(number int) {
	if number <= 0 {
		fmt.Println("The number of terms must be greater than 0")
		return
	}

	var sum int
	var first = 0
	var second = 1

	if number >= 1 {
		sum += first + second
	}

	for i := 2; i < number; i++ {
		third := first + second
		sum += third
		first = second
		second = third
	}

	fmt.Println("The sum of fibonacci series upto ", number, "th term is ", sum)
}

func CountEvenOddZero(numbers []int) (int, int, int) {
	var evenCount, oddCount, zeroCount int

	for _, num := range numbers {
		if num == 0 {
			zeroCount++
		} else if num%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	return evenCount, oddCount, zeroCount
}
