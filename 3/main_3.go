package main

import (
	"fmt"
	"strings"
)

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{make(map[string]int)}
}

func (dict *StringIntMap) Add(key string, value int) {
	dict.data[key] = value
}

func (dict *StringIntMap) Remove(key string) {
	delete(dict.data, key)
}

func (dict *StringIntMap) Copy() map[string]int {
	result := make(map[string]int, len(dict.data))
	for key, value := range dict.data {
		result[key] = value
	}
	return result
}

func (dict *StringIntMap) Exists(key string) bool {
	_, ok := dict.data[key]
	return ok
}

func (dict *StringIntMap) Get(key string) (int, bool) {
	value, ok := dict.data[key]
	return value, ok
}

func (dict *StringIntMap) String() string {
	var builder strings.Builder
	for key, value := range dict.data {
		builder.WriteString(fmt.Sprintf("%-8s: %d\n", key, value))
	}
	return builder.String()
}

func main() {
	dict := NewStringIntMap()
	dict.Add("apple", 25)
	dict.Add("banana", 7)
	dict.Add("cherry", 3)

	fmt.Println(dict)

	dict.Remove("cherry")
	fmt.Println(dict)

	dictCopy := dict.Copy()
	fmt.Printf("%v\n\n", dictCopy)

	dictCopy["potato"] = 100
	fmt.Printf("Potato: %v\n", dict.Exists("potato"))
	fmt.Printf("Apple: %v\n", dict.Exists("apple"))

	value, ok := dict.Get("banana")
	if ok {
		fmt.Println("Banana", value)
	} else {
		fmt.Println("Banana not found")
	}
}
