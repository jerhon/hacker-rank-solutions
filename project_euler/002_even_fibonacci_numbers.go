// Project Euler #2  https://www.hackerrank.com/contests/projecteuler/challenges/euler002/problem
//
// Just generate a Fibonacci Sequence and check for even values and sum them together.
// The trick here is the execution time.
//
// I can't treat each test case as independent as re-generating the fibonacci sequence each time
// is costly.  So I reuse the same sequence, sort the test cases, and hold the results until they've
// all been calculated at that time I then output them all.
//
// By doing that I only ever generate one Fibonacci sequence at the cost of the extra memory
// storing the results.

package main

import (
	"fmt"
	"os"
	"sort"
)

// FibonacciSequence stores state for a Fibonacci Sequence generator
type FibonacciSequence struct {
	first int64
	second int64
}

// NewFibonacciSequence creates a new FibonacciSequence
func NewFibonacciSequence() *FibonacciSequence {
	return &FibonacciSequence{}
}

// next number in a Fibonacci Sequence
func (seq *FibonacciSequence) next() int64 {
	if seq.first == int64(0) {
		seq.first = 1
		return 1
	}
	if seq.second == int64(0) {
		seq.second = 2
		return 2
	}

	previousFirst := seq.first
	seq.first = seq.second
	seq.second += previousFirst

	return seq.second
}

// Calculates the results for the fibonacci sequence and stores them in a map
func getFibonacciResults(testCases []int) map[int]int64 {

	// By sorting the list, we can calculate a single sequence
	sortedTestCases := make([]int, len(testCases))
	results := make(map[int]int64)
	copy(sortedTestCases, testCases)
	sort.Ints(sortedTestCases)

	// Go through and keep adding to the sum until the max has been reached
	// then resume for the next test case.
	sequence := NewFibonacciSequence()
	seqVal := sequence.next()
	sum := int64(0)
	for _, testCaseMaximum := range sortedTestCases {
		for seqVal <= int64(testCaseMaximum) {
			if seqVal % 2 == 0 {
				sum += seqVal
			}
			seqVal = sequence.next()
		}
		results[testCaseMaximum] = sum
	}
	return results
}

// readTestCases reads the test cases from standard input
func readTestCases() []int {
	var lines int

	_, err := fmt.Fscanln(os.Stdin, &lines)
	if err != nil {
		panic("Expected the number of test cases")
	}

	testCases := make([]int, lines)
	for i := 0; i < lines; i++ {
		var testCase int
		_, err = fmt.Fscanln(os.Stdin, &testCase)
		if err != nil {
			panic("Missing test case")
		}
		testCases[i] = testCase
	}
	return testCases
}

// Run the sample program
func main() {

	testCases := readTestCases()
	results := getFibonacciResults(testCases)

	// Print the result for each test case
	for _, val := range testCases {
		fmt.Println(results[val])
	}
}