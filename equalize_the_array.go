// Solution for the Hacker Rank Problem Equalize the Array
//
// https://www.hackerrank.com/challenges/equality-in-a-array/problem

// Reasoning is simple here
//
// It's really just what number occurs the most in the array
// How many numbers aren't equal to that

package main

func equalizeArray(arr []int32) int32 {
	// Write your code here

	counts := make(map[int32]int)
	for _, val := range arr {
		counts[val]++
	}

	maximum := 0
	for _, count := range counts {
		if count > maximum {
			maximum = count
		}
	}

	return int32(len(arr) - maximum)
}
