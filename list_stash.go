package main

import (
	"fmt"
	"strings"
)

type node struct {
	data interface{}
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

type LinkedList interface {
	Add(data interface{})
	Contains(data interface{}) bool
	IndexOf(data interface{}) int
	InsertAfter(index int, data interface{}) error
	Pop() (interface{}, error)
	Push(data interface{})
	Remove(data interface{}) error
	RemoveByIndex(index int) error
	Set(index int, data interface{}) error
	GetLength() int
	GetHead() (interface{}, error)
	GetTail() (interface{}, error)
	Print()
	SearchIndex(partialData string) (int, interface{})
}

func NewLinkedList() LinkedList {
	return &linkedList{}
}

func (l *linkedList) SearchIndex(partialData string) (int, interface{}) {
	// fmt.Printf("...Searching string \"%v\" in the nodes...\n", partialData)
	pointer := l.head
	for i := 0; i < l.length; i++ {
		str := fmt.Sprintf("%v", pointer.data)
		// fmt.Printf("...Converting data into string...%v \n", str)
		fmt.Printf("...Searching if \"%v\" is contained...%v \n", partialData, strings.Contains(str, partialData))
		if strings.Contains(str, partialData) {
			return i, pointer.data
		} else {
			pointer = pointer.next
		}
	}
	return -1, nil
}

func (l *linkedList) GetHead() (interface{}, error) {
	fmt.Printf("...Getting head of the list...")
	if l.length == 0 {
		err := fmt.Errorf("[error: empty list]")
		return "(null)", err
	}
	return l.head.data, nil
}

func (l *linkedList) GetLength() (len int) {
	return l.length
}

func (l *linkedList) GetTail() (interface{}, error) {
	fmt.Printf("...Getting tail of the list...")
	if l.length == 0 {
		err := fmt.Errorf("[error: empty list]")
		return "(null)", err
	}
	return l.tail.data, nil
}

func (l *linkedList) Print() {
	// fmt.Printf("...Printing...\n")
	fmt.Printf("-------- PRINTING ---------\n")
	pointer := l.head
	for pointer != nil {
		fmt.Printf("%v -> ", pointer.data)
		pointer = pointer.next
	}
	fmt.Printf("(null)\n")
	fmt.Println("Length =", l.length)

	if l.head == nil {
		fmt.Println("Head is nil")
	}
	fmt.Println("Head =", l.head.data)

	if l.tail == nil {
		fmt.Println("Tail is nil")
	}
	fmt.Println("Tail =", l.tail.data)
	fmt.Printf("----------- END -----------\n")
}

func (l *linkedList) Add(data interface{}) {
	fmt.Printf("...Adding \"%v\" to the list...", data)
	node := &node{
		data: data,
	}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.length++
	fmt.Println("[Success!]")
}

func (l *linkedList) Contains(data interface{}) bool {
	fmt.Printf("...Checking if the list contains \"%v\"...\n", data)
	pointer := l.head
	for i := 0; i < l.length; i++ {
		if pointer.data == data {
			return true
		} else {
			pointer = pointer.next
		}
	}
	return false
}

func (l *linkedList) IndexOf(data interface{}) int {
	fmt.Printf("...Searching \"%v\" in the list...\n", data)
	pointer := l.head
	for i := 0; i < l.length; i++ {
		if pointer.data == data {
			return i
		} else {
			pointer = pointer.next
		}
	}
	return -1
}

func (l *linkedList) InsertAfter(index int, data interface{}) error {
	fmt.Printf("...Inserting \"%v\" after node(%v) to the list...", data, index)
	if l.length == 0 {
		err := fmt.Errorf("[error: empty list]")
		fmt.Println(err)
		return err
	}
	if index < 0 {
		err := fmt.Errorf("[invalid index: should be [0 - %v]]", l.length-1)
		fmt.Println(err)
		return err
	}
	if index >= l.length {
		err := fmt.Errorf("[invalid index: exceeded maximum index (%v)]", l.length-1)
		fmt.Println(err)
		return err
	}
	node := &node{
		data: data,
	}
	if index == l.length-1 {
		currentTail := l.tail
		currentTail.next = node
		l.tail = node
		l.length++
	} else {
		pointer := l.head
		for i := 0; i <= index; i++ {
			if i == index {
				node.next = pointer.next
				pointer.next = node
				l.length++
				fmt.Println("[Success!]")
			} else {
				pointer = pointer.next
			}
		}
	}

	return nil
}

func (l *linkedList) Pop() (interface{}, error) {
	fmt.Printf("...Poping the first node...")
	if l.head == nil {
		err := fmt.Errorf("[error: empty list]")
		fmt.Println(err)
		return "", err
	}
	data := l.head.data
	fmt.Printf("\"%v\"...", data)
	l.head = l.head.next
	l.length--
	if l.length == 0 {
		l.tail = nil
	}
	fmt.Println("[Success!]")
	return data, nil
}

func (l *linkedList) Push(data interface{}) {
	fmt.Printf("...Pushing \"%v\" to the list...", data)
	node := &node{
		data: data,
	}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		currentHead := l.head
		l.head = node
		l.head.next = currentHead
	}
	l.length++
	fmt.Println("[Success!]")
}

func (l *linkedList) Remove(data interface{}) error {
	fmt.Printf("...Removing \"%v\" from the list...", data)
	pointer := l.head
	var prevPointer *node
	for i := 0; i < l.length; i++ {
		if pointer.data == data {
			if i == 0 {
				l.head = l.head.next
			} else if i == l.length-1 {
				l.tail = prevPointer
				l.tail.next = nil
			} else {
				prevPointer.next = pointer.next
			}
			l.length--
			if l.length == 0 {
				l.tail = nil
			}
			fmt.Println("[Success!]")
			return nil
		} else {
			prevPointer = pointer
			pointer = pointer.next
		}
	}
	err := fmt.Errorf("[error: \"%v\" not found]", data)
	fmt.Println(err)
	return err
}
func (l *linkedList) RemoveByIndex(index int) error {
	fmt.Printf("...Removing node(%v) from the list...", index)
	if index < 0 {
		err := fmt.Errorf("[invalid index: should be [0 - %v]]", l.length-1)
		fmt.Println(err)
		return err
	}
	if l.length == 0 {
		err := fmt.Errorf("[error: empty list]")
		fmt.Println(err)
		return err
	} else if index >= l.length {
		err := fmt.Errorf("[invalid index: exceeded maximum index (%v)]", l.length-1)
		fmt.Println(err)
		return err
	}

	pointer := l.head
	var prevPointer *node
	if index == 0 {
		l.head = l.head.next
	} else {
		for i := 0; i < l.length; i++ {
			if i == index {
				prevPointer.next = pointer.next
				if index == l.length-1 {
					l.tail = prevPointer
				}
			} else {
				prevPointer = pointer
				pointer = pointer.next
			}
		}
	}
	l.length--
	if l.length == 0 {
		l.tail = nil
	}
	fmt.Println("[Success!]")
	return nil
}

func (l *linkedList) Set(index int, data interface{}) error {
	fmt.Printf("...Setting node(%v) as \"%v\"...", index, data)
	if l.length == 0 {
		err := fmt.Errorf("[error: empty list]")
		fmt.Println(err)
		return err
	}
	if index < 0 {
		err := fmt.Errorf("[invalid index: should be [0 - %v]]", l.length-1)
		fmt.Println(err)
		return err
	}
	if index >= l.length {
		err := fmt.Errorf("[invalid index: exceeded maximum index (%v)]", l.length-1)
		fmt.Println(err)
		return err
	}
	if index == 0 {
		l.head.data = data
	} else if index == l.length-1 {
		l.tail.data = data
	} else {
		pointer := l.head
		for i := 0; i < l.length; i++ {
			if i == index {
				pointer.data = data
			} else {
				pointer = pointer.next
			}
		}
	}
	fmt.Println("[Success!]")
	return nil
}
