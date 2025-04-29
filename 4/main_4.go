package main

import "fmt"

func SliceDifference[T comparable](first []T, second []T) []T {
	elements := make(map[T]struct{})

	for _, value := range first {
		elements[value] = struct{}{}
	}

	for _, value := range second {
		delete(elements, value)
	}

	result := make([]T, 0, len(elements))
	for key := range elements {
		result = append(result, key)
	}

	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	difference := SliceDifference(slice1, slice2)
	fmt.Println(difference)
}
