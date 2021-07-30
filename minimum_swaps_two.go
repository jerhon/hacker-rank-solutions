// Solution for hacker rank problem https://www.hackerrank.com/challenges/minimum-swaps-2/problem
//
// The trick is that there are essentially 3 operations that can occur on the array
//
// #1 two numbers could be out of order and swapping the two leads to them both being in the correct position
// #2 a number is one value out of position and another number close to it is further out of position, finding the value that should go in the other slot will allow to reduce the number of out of order elements to just 1 in 2 swaps
// #3 is just to swap an element to it's proper position
//
// The algorithm runs through #1 and #2 in a loop through the array
// After that, the algorithm again runs through the loop and performs #3
//

package main

import "fmt"


// ArrayElement represents a single element in a problem array.
type ArrayElement struct {
	// The index of the element
	index int
	// The actual value for the element at the index
	value int32
	// the expected value for the element at the index
	expected int32
}

// ProblemArray represents an array for the problem.
type ProblemArray struct {
	// The current values in order.
	values    []int32

	// Maps out of order items from their value to index.
	indexes   map[int32]int

	// The number of swaps to perform.
	swapCount int
}

// Creates a new array swapper which encapsolates operations to perform on the array.
func new(arr []int32) ProblemArray {
	positions := make(map[int32]int)
	for idx, v := range arr {
		value := v - 1
		arr[idx] = value
		if int32(idx) != value {
			positions[value] = idx
		}
	}
	return ProblemArray{values: arr, indexes: positions, swapCount: 0}
}

// get element for the position in the array.
func (arr *ProblemArray) get(idx int) ArrayElement {
	return ArrayElement{idx, arr.values[idx], int32(idx)}
}

// Checks if the array is large enought to support an index. */
func (arr *ProblemArray) hasIdx(idx int) bool {
	return idx >= 0 && idx < len(arr.values)
}

// Finds an array element by it's value.
func (arr *ProblemArray) findValue(value int32) ArrayElement {
	if idx, exists := arr.indexes[value]; exists {
		return ArrayElement{idx, value, int32(idx)}
	}
	return ArrayElement{int(value), value, value }
}

// Swap switch the position of two elements in the problem array.
func (arr *ProblemArray) swap(a int, b int) {

	fmt.Println(a, " <=> ", b)

	aVal := arr.values[a]
	bVal := arr.values[b]
	arr.values[b] = aVal
	arr.values[a] = bVal
	arr.indexes[bVal] = a
	arr.indexes[aVal] = b
	if arr.indexes[bVal] == int(bVal) {
		delete(arr.indexes, bVal)
	}
	if arr.indexes[aVal] == int(aVal) {
		delete(arr.indexes, aVal)
	}
	arr.swapCount++
}

// Can swap two elements in one operation. Returns true if it can and the indexes to swap.
func (arr *ProblemArray) canTwoForOneSwap(idx int) (bool, int, int) {
	current := arr.get(idx)
	if current.value == current.expected {
		return false, -1, -1
	}

	other := arr.findValue(current.expected)

	if current.expected == other.value && other.expected == current.value {
		return true, current.index, other.index
	}
	return false, -1, -1
}

// Can perform a 3 way swap. First is whether it matches a three way swap, the 3 integers are indexes to the elements that should be swapped in order.
func (arr *ProblemArray) canThreeWaySwap(idx int) (bool, int, int, int) {
	current := arr.get(idx)
	if current.value == current.expected {
		return false, -1, -1, -1
	}

	var adjacent ArrayElement
	if arr.hasIdx(idx + 1) && current.value == current.expected + 1 {
		adjacent = arr.get(idx + 1)
	} else if arr.hasIdx(idx - 1) && current.value == current.expected - 1 {
		adjacent = arr.get(idx - 1)
	} else {
		return false, -1, -1, -1
	}

	if adjacent.value == adjacent.expected {
		return false, -1, -1, -1
	}

	swap := arr.findValue(current.expected)
	return true, swap.index, adjacent.index, current.index
}

// Returns true if the array is out of order.
func (arr *ProblemArray) isOutOfOrder() bool {
	return len(arr.values) > 0
}

// Reorders the array.
func (arr *ProblemArray) reorderArray() {

	// First find any elements that can be switched with faster operations
	for idx := 0; idx < len(arr.values); idx++ {
		if match, a, b := arr.canTwoForOneSwap(idx); match {
			fmt.Println("Swap 2 elements")
			arr.swap(a, b)
		}
		if match, a, b, c := arr.canThreeWaySwap(idx); match {
			fmt.Println("Swap 3 elements")
			arr.swap(a, b)
			arr.swap(b, c)
		}
	}

	fmt.Println("Brute force the rest")
	// Find any elements that can be swapped one by on
	for idx := 0; idx < len(arr.values); idx++ {
		current := arr.get(idx)
		if current.value != current.expected {
			swap := arr.findValue(current.expected)
			arr.swap(current.index, swap.index)
		}
	}
}

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	swapper := new(arr)
	swapper.reorderArray()

	return int32(swapper.swapCount)
}

// Runs the program with a few samples
func main() {
	minimumSwaps([]int32{ 2, 3, 4, 1, 5 })
	fmt.Println("Next ---------------------------------------")
	minimumSwaps([]int32{ 1, 3, 5, 2, 4, 6, 7 })
}