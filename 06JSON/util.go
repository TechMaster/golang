package main

import (
	"fmt"
	"sort"
)

type KeyValue struct {
	Key   string
	Value int
}

type CityJobCount struct {
	City  string
	Job   string
	Count int
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

func PrintSliceKeyValue(input []KeyValue) {
	for index, item := range input {
		fmt.Printf("%d - %s - %d \n", index+1, item.Key, item.Value)
	}
}
