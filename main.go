package main

import (
	"HashMap/hashMap"
	"fmt"
)

func main() {
	hashMap := hashMap.NewHashMap(2)
	// hashMap.Print()
	fmt.Println("hashMap.Set(key1, value1)")
	hashMap.Set("key1", "value1")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("hashMap.Set(key2, value2)")
	hashMap.Set("key2", "value2")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("hashMap.Set(key1, value1Updated)")
	hashMap.Set("key1", "value1Updated")

	fmt.Println("")
	fmt.Println("")
	hashMap.Print()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("hashMap.Set(key3, value3)")
	hashMap.Set("key3", "value3")
	// fmt.Println(hashMap.Get("key3"))

	fmt.Println("***")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	hashMap.Print()

	hashMap.Set("key4", "value4")

	fmt.Println("***")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	hashMap.Print()

	// hashMap.Print()

	// fmt.Println(hashMap.Get("key1"))

	hashMap.Remove("key1")

	fmt.Println("***")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	hashMap.Print()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}
