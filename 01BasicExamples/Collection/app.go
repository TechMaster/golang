package main

import (
	"fmt"
	"sort"
)

type KeyValue struct {
	Key   string
	Value int
}

func SortMapByValue(input map[string]int, isAscending bool) (result []KeyValue) {
	//Tạo slice của các key từ tham số input
	keys := make([]string, 0, len(input))
	for key := range input {
		keys = append(keys, key)
	}

	//Sắp sếp danh sách keys theo value thuộc map input
	if isAscending {
		sort.Slice(keys, func(i, j int) bool {
			return input[keys[i]] < input[keys[j]]
		})
	} else {
		sort.Slice(keys, func(i, j int) bool {
			return input[keys[i]] > input[keys[j]]
		})
	}

	result = make([]KeyValue, len(input))

	for index, key := range keys {
		result[index] = KeyValue{Key: key, Value: input[key]}
	}
	return
}

func main() {
	scores := map[string]int{"Alma": 23, "Cecilia": 12, "David": 37, "Berta": 78}
	fmt.Println(len(scores))
	fmt.Println(scores)
	fmt.Println()

	result := SortMapByValue(scores, false)
	for _, name := range result {
		//fmt.Printf("%-7v %v\n", name, scores[name])
		fmt.Println(name.Key, name.Value)
	}
	/*	for name, score := range scores {
			fmt.Printf("%-7v %v\n", name, score)
		}
		fmt.Println()

		names := make([]string, 0, len(scores))
		for name := range scores {
			names = append(names, name)
		}
		fmt.Println(names)
		fmt.Println()

		sort.Slice(names, func(i, j int) bool {
			return scores[names[i]] > scores[names[j]]
		})
		fmt.Println(names)
		fmt.Println()

		for _, name := range names {
			fmt.Printf("%-7v %v\n", name, scores[name])
		}*/
}
