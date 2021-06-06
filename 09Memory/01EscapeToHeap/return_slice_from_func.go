package main

func ReturnASlice(size int) (result []string) {
	result = make([]string, size) //make([]string, size) escapes to heap
	return
}
