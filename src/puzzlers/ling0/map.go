package main

import (
	"fmt"
)

var (
	barVal = map[string]int{
		"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	invMap := make(map[int]string, len(barVal))
	copMap := make(map[string]int, len(barVal))

	for k, v := range barVal {
		invMap[v] = k
	}

	for k, v := range barVal {
		copMap[k] = v
	}

	fmt.Println("inverted:")
	// map 的键值对调
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v  ", k, v)
		fmt.Println()
	}

	fmt.Println("cop map:")
	// map 复制
	for k, v := range copMap {
		fmt.Printf("Key: %v, Value: %v", k, v)
		fmt.Println()
	}
}
