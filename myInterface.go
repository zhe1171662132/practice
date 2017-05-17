package main

import "fmt"

func main() {
	slice := make([]interface{}, 10)
	map1 := make(map[string]string)
	map2 := make(map[int]string)
	map1["name"] = "fangdi"
	map2[1] = "good"
	map3 := make(map[string]map[string]string)
	map3["my"] = map1
	slice[0] = map1
	slice[1] = map2
	slice[2] = map3
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
}
