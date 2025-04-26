package main

import "fmt"

func SliceUnification[T comparable](first []T, second []T) (bool, []T) {
	elements := make(map[T]bool)

	for _, value := range first {
		elements[value] = false
	}

	resultSlice := make([]T, 0, len(elements))

	for _, value := range second {
		if exists, ok := elements[value]; ok && !exists {
			elements[value] = true
			resultSlice = append(resultSlice, value)
		}
	}

	return len(resultSlice) > 0, resultSlice
}

func main() {
	a := []int{65, 3, 58, 678, 64, 2, 2}
	b := []int{64, 2, 3, 43, 2}
	fmt.Println(SliceUnification(a, b))
}
