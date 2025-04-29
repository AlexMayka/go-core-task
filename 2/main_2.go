package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func getRandSlice() []int {
	result := make([]int, 10)
	for i := 0; i < 10; i++ {
		result[i] = rand.Intn(100)
	}

	return result
}

func sliceExample(slice []int) []int {
	result := make([]int, 0, len(slice))
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}

	return result
}

func addElements(slice []int, num int) []int {
	return append(slice, num)
}

func copySlice(slice []int) []int {
	result := make([]int, len(slice))
	copy(result, slice)
	return result
}

func removeSlice(slice []int, index int) ([]int, error) {
	if index < 0 || index >= len(slice) {
		return nil, errors.New("index out of range")
	}
	return append(slice[:index], slice[index+1:]...), nil
}

func main() {
	originalSlice := getRandSlice()
	fmt.Println(originalSlice)

	sliceEven := sliceExample(originalSlice)
	fmt.Println(sliceEven)

	sliceWithElem := addElements(originalSlice, 2)
	fmt.Println(sliceWithElem)

	sliceCopy := copySlice(originalSlice)
	fmt.Println(sliceCopy)

	sliceRemove, err := removeSlice(originalSlice, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sliceRemove)

	var x = []int{}
	fmt.Println(removeSlice(x, 0))
}
