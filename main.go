package main

import (
	"HashMap/hashMap"
	"fmt"
)

func main() {
	hashMap := hashMap.NewHashMap(2)
	fmt.Println("		hashMap.Print()")
	hashMap.Print()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("		hashMap.Set(key1, value1)")
	hashMap.Set("key1", "value1")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("		hashMap.Set(key2, value2)")
	hashMap.Set("key2", "value2")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("		hashMap.Set(key1, value1Updated)")
	hashMap.Set("key1", "value1Updated")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("		hashMap.Set(key3, value3)")
	hashMap.Set("key3", "value3")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("		hashMap.Print()")
	hashMap.Print()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("		ListKeys()")
	list := hashMap.ListKeys()
	fmt.Println(list)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("		hashMap.Get(key2)")
	fmt.Println(hashMap.Get("key2"))

	fmt.Println("		hashMap.Set(key4, value4)")
	hashMap.Set("key4", "value4")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("		hashMap.Print()")
	hashMap.Print()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("		hashMap.ManualRehash()")
	hashMap.ManualRehash()

	// fmt.Println("")
	// fmt.Println("")
	// fmt.Println("		hashMap.Remove(key1)")
	// hashMap.Remove("key1")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("		hashMap.Print()")
	hashMap.Print()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("		ListKeys()")
	list = hashMap.ListKeys()
	fmt.Println(list)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("		hashMap.Clear()")
	hashMap.Clear()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("		hashMap.Print()")
	hashMap.Print()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}
