// Solution for problem: https://www.hackerrank.com/contests/projecteuler/challenges/euler001/problem
//
// Reasoning:
//
// Use the basic math equation to solve for some of a sequence.
//
// If we add the sequence of 3s plus sequence of multiples of 5, we'd have duplicate numbers that are multiple of 15s
// So, we remove that sequence
//
// To find sum of all numbers in a sequence where n is the maximum, equation = n * (n + 1) / 2

package main

import (
	"fmt"
	"os"
)

func getSumOfMultiples(max int64, multiple int64) int64 {
	n := (max / multiple)
	return multiple * n * (n + 1) / 2
}

func getMultiplesOf3And5(max int64) int64 {
	return getSumOfMultiples(max, 3) + getSumOfMultiples(max, 5) - getSumOfMultiples(max, 15)
}

func main() {
	var lines int

	_, err := fmt.Fscanln(os.Stdin, &lines)
	if err != nil {
		panic("Expected the number of test cases")
	}

	for i := 0; i < lines; i++ {
		var testCase int64
		_, err = fmt.Fscanln(os.Stdin, &testCase)
		if err != nil {
			panic("Missing test case")
		}
		fmt.Println(getMultiplesOf3And5(testCase - 1))
	}
}