package main

import (
	"fmt"
	"sort"
)

func main() {
	sampleMap := map[string]int{
		"banana": 3,
		"apple":  5,
		"cherry": 7,
	}

	for key, value := range sampleMap {
		fmt.Println(key, value)
	}

	keys := make([]string, 0, len(sampleMap))
	for key := range sampleMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	fmt.Println("Sorted keys:")
	for _, key := range keys {
		fmt.Printf("Key: %s, Value: %d\n", key, sampleMap[key])
	}
}
