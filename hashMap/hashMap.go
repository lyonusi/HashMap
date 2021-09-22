package hashMap

import (
	"HashMap/list"
	"fmt"
)

const (
	expandFactor = 0.75
)

type hashMap struct {
	h           []*ht
	i           int
	cap         int
	ocp         int
	privateTool tools
}

type ht struct {
	m   []*list.LinkedList
	cap int
	ocp int
}

type tool struct {
}

type tools interface {
	convertToKeyValue(data interface{}) (keyValue, error)
	searchIfKeyExists(h *hashMap, key string) (int, int, interface{})
	addEntry(h *hashMap, data keyValue)
	emptyMapCheck(h *hashMap) error
}

type keyValue struct {
	Key   string
	Value interface{}
}

type HashMap interface {
	Clear()
	Index(key string) int
	Get(key string) (value interface{})
	Length() int
	ListKeys() []string
	// ManualRehash()
	Remove(key string) error
	Set(key string, value interface{})
	Print()
}

func (h *hashMap) Clear() {
	h.i = 0
	h0 := new(ht)
	h0.cap = h.cap
	h0.m = make([]*list.LinkedList, h.cap)
	h.h = nil
	h.h = append(h.h, h0)

	h.h[h.i] = h0
	h.ocp = 0
}

// func (h *hashMap) ManualRehash() {
// 	fmt.Println(".................... Current number of hash map ....................", i+1)

// 	for i := 0; i < h.i; i++ {
// 		fmt.Printf("hashMap.h[%v].cap = %v\n", i, h.h[i].cap)
// 		fmt.Printf("hashMap.h[%v].ocp = %v\n", i, h.h[i].ocp)

// 		for j := 0; j < h.h[i].cap; j++ {
// 			entries := h.h[i].m[j]
// 			if entries == nil {
// 				fmt.Printf("hashMap.h[%v].%v = (nil)\n", i, j)
// 			} else {
// 				l := *h.h[i].m[j]
// 				fmt.Printf("hashMap.h[%v].%v  ---> \n", i, j)

// 				h.privateTool.moveToNewHash(h, i, j, _)

// 			}
// 		}
// 	}
// }

func (h *hashMap) Length() int {
	return h.ocp
}

func (h *hashMap) ListKeys() []string {
	var keyList []string
	for i := h.i; i >= 0; i-- {
		// fmt.Println("Func ListKeys Log-1 >> i =", i)
		for j := 0; j < h.h[i].cap; j++ {
			// fmt.Println("Func ListKeys Log-2 >> j =", j)
			list := h.h[i].m[j]
			if list != nil {
				entries := *h.h[i].m[j]
				// fmt.Println("Func ListKeys Log-3 >> entries =", entries)
				if entries != nil {
					iterator := entries.GetIterator()
					// fmt.Println("Func ListKeys Log-4 >> iterator.HasNext =", iterator.HasNext())
					entry := iterator.GetData()
					// fmt.Println("Func ListKeys Log-5 >> entry =", entry)
					data, _ := h.privateTool.convertToKeyValue(entry)
					// fmt.Println("Func ListKeys Log-6 >> key =", data.Key)
					keyList = append(keyList, data.Key)
					for iterator.HasNext() {
						iterator.Next()
						entry = iterator.GetData()
						data, _ := h.privateTool.convertToKeyValue(entry)
						keyList = append(keyList, data.Key)
						// fmt.Println("Func ListKeys Log-7 >> next =", next)
					}
				}
			}
		}
	}
	return keyList
}

func (k *keyValue) Compare(d list.Data) (bool, error) {
	compareData, ok := d.(*keyValue)
	if !ok {
		return false, fmt.Errorf("input cannot be converted to keyValue")
	} else {
		return k.Key == compareData.Key, nil
	}
}

func NewHashMap(capacity int) HashMap {
	// if capacity < 16 {
	// 	capacity = 16
	// }
	h := new(hashMap)
	h.cap = capacity
	h.i = 0

	h0 := new(ht)
	h0.cap = capacity
	h0.m = make([]*list.LinkedList, capacity)
	h.h = append(h.h, h0)

	h.h[h.i] = h0
	// fmt.Println("Func NewHasMap Log-1 >> h.m =", h.m)

	h.privateTool = &tool{}
	return h
}

func Expand(h *hashMap) HashMap {
	newH := new(ht)
	newH.cap = 2 * h.h[h.i].cap
	newH.m = make([]*list.LinkedList, newH.cap)
	h.h = append(h.h, newH)

	h.cap = newH.cap
	h.i++
	// fmt.Println("Func Expand Log-1 >> hashMaps.i =", h.i)
	h.h[h.i] = newH
	// fmt.Printf("Func Expand Log-2 >> hashMaps.h[%v] = %v\n", h.i, newH.m)
	return h
}

func HashFunc(str string, cap int) int {
	seed := int(131)
	hash := int(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + int(str[i])
	}
	return hash % cap
	// return 0
}

func (h *hashMap) Print() {
	fmt.Println("======================== PRINTING ========================")
	fmt.Println("hashMap total cap =", h.cap)
	fmt.Println("hashMap total ocp =", h.ocp)
	fmt.Println("hashMap.h --->", h.h)

	for i := 0; i <= h.i; i++ {
		fmt.Printf(".................... hashMap.h[%v] ....................\n", i)
		fmt.Printf("                    hashMap.h[%v].cap = %v\n", i, h.h[i].cap)
		fmt.Printf("                    hashMap.h[%v].ocp = %v\n", i, h.h[i].ocp)
		fmt.Printf("                    hashMap.h[%v].m ---> %v\n", i, h.h[i].m)
		for j := 0; j < h.h[i].cap; j++ {
			entries := h.h[i].m[j]
			if entries == nil {
				fmt.Printf("                    hashMap.%v.%v : (nil)\n", i, j)
			} else {
				l := *h.h[i].m[j]
				fmt.Printf("                    ...hashMap.%v.%v : ", i, j)
				l.Print()
			}
		}
	}
	fmt.Println("------------------- END -------------------")
}

func (h *hashMap) Index(key string) int {
	// fmt.Println("Func Index Log-1 >> HashFunc =", HashFunc(key, h.cap))
	return HashFunc(key, h.cap)
}

func (h *hashMap) Set(key string, value interface{}) {
	newEntry := &keyValue{Key: key, Value: value}
	i, m, _ := h.privateTool.searchIfKeyExists(h, key) //Check if key exists in the hashmap

	if i != -1 && m != -1 { // if key exists...
		// fmt.Println("Func Set Log-1 >> equal? =", currentEntry == newEntry)

		h.Remove(key) // remove current entry
		// fmt.Println("Func Set Log-2 >> removed:", currentEntry)

		h.privateTool.addEntry(h, *newEntry) // add new entry to latest hashmap
		// fmt.Println("Func Set Log-3 >> added:", newEntry)
	} else if (float64(h.ocp) / float64(h.cap)) > expandFactor { // if key does not exist and hash map is occupied > 0.75
		newH := Expand(h)
		index := newH.Index(key)
		newEntries := list.NewLinkedList()
		// fmt.Println("Func Set Log-5 >> New entry (expanded) =", newEntry)
		h.h[h.i].m[index] = &newEntries
		newEntries.Add(newEntry)
		h.ocp++
		// fmt.Println("Func Set Log-6 >> Expanded hashMap.ocp =", h.ocp)

		h.h[h.i].ocp++
		// for i := h.i; i >= 0; i-- {
		// 	fmt.Printf("Func Set Log-7 >> Expanded hashMap.h[%v].ocp = %v\n", i, h.h[i].ocp)
		// }
	} else { // if key does not exist
		h.privateTool.addEntry(h, *newEntry) // add new entry to latest hashmap
	}

	// for i := h.i; i >= 0; i-- { //Check if entry exists in the current hashMap.h[i]
	// 	index := HashFunc(key, h.h[i].cap)
	// 	if h.h[i].m[index] != nil { //if entry list is not nil, check if key exists
	// 		entries := *h.h[i].m[index]
	// 		keyExists, currentEntry := entries.Contains(newEntry)
	// 		fmt.Println("Func Set Log-2 >> keyExists =", keyExists)

	// 		if keyExists { //if key exists, compare new entry with current entry
	// 			// fmt.Println("Func Set Log-2 >> currentEntry =", currentEntry)
	// 			if currentEntry != newEntry { //if not equal, update entry with new entry
	// 				entries.Remove(currentEntry)
	// 				entries.Add(newEntry)
	// 				// fmt.Println("Func Set Log-3 >> entry updated:", newEntry)
	// 			} else { //if equal, do nothing
	// 				// fmt.Println("Func Set Log-4 >> newEntry == currentEntry")
	// 			}

	// 		}
	// 		if !keyExists {
	// 			if (float64(h.ocp) / float64(h.cap)) > expandFactor {
	// 				newH := Expand(h)
	// 				index := newH.Index(key)
	// 				e := list.NewLinkedList()
	// 				// fmt.Println("Func Set Log-5 >> New entry (expanded) =", newEntry)
	// 				h.h[h.i].m[index] = &e
	// 				e.Add(newEntry)
	// 				h.ocp++
	// 				// fmt.Println("Func Set Log-6 >> Expanded hashMap.ocp =", h.ocp)

	// 				h.h[h.i].ocp++

	// 				for i := h.i; i >= 0; i-- {
	// 					// fmt.Printf("Func Set Log-7 >> Expanded hashMap.h[%v].ocp = %v\n", i, h.h[i].ocp)
	// 				}

	// 			} else {
	// 				index := h.Index(key)
	// 				entries := h.h[h.i].m[index]
	// 				if entries == nil {
	// 					e := list.NewLinkedList()
	// 					// fmt.Println("Func Set Log-9 >> Empty entries, created new entry =", newEntry)
	// 					h.h[h.i].m[index] = &e
	// 					e.Add(newEntry)
	// 					h.ocp++
	// 					h.h[h.i].ocp++
	// 				} else {
	// 					// fmt.Printf("Func Set Log-10 >> %v do not exist\n", key)
	// 					entries := h.h[h.i].m[index]
	// 					e := *entries
	// 					e.Add(newEntry)
	// 					h.ocp++
	// 					h.h[h.i].ocp++
	// 				}
	// 			}
	// 		}
	// 	} else { //if entry list is nil, add entry to list
	// 		entries := list.NewLinkedList()
	// 		entries.Add(newEntry)
	// 		h.h[i].m[index] = &entries
	// 		h.ocp++
	// 		h.h[i].ocp++
	// 	}

	// }
}

func (h *hashMap) Get(key string) interface{} {
	searchEntry := &keyValue{Key: key}

	for i := h.i; i >= 0; i-- { //Check if entry exists in the hashmap
		index := HashFunc(key, h.h[i].cap)
		if h.h[i].m[index] != nil { //if h[i] hashmap is not nil, check if key exists
			entries := *h.h[i].m[index]
			keyExists, currentEntry := entries.Contains(searchEntry)
			// fmt.Println("Func Get Log-1 >> keyExists =", keyExists)

			entry, err := h.privateTool.convertToKeyValue(currentEntry)
			if err != nil {
				fmt.Println(err)
				return err
			}
			if keyExists {
				if i != h.i { // if entry is not in the latest hashmap, move to the latest hashmap h[h.i]
					h.Remove(key)
					h.privateTool.addEntry(h, entry)
					fmt.Println("Func Get Log-2 >> moved entry to new hash map:", i, "->", h.i)
				}
				// fmt.Println("Func Get Log-3 >> entry =", entry)
				return entry.Value
			}
		}
	}
	return fmt.Errorf("not found")
}

func (h *hashMap) Remove(key string) error {
	i, m, _ := h.privateTool.searchIfKeyExists(h, key)
	if i != -1 && m != -1 {
		entries := *h.h[i].m[m]
		_, entry := entries.Contains(&keyValue{Key: key})
		index := entries.IndexOf(entry)
		entries.RemoveByIndex(index)
		h.ocp--
		h.h[i].ocp--
	}
	// fmt.Println("Func Remove Log-1 >> removed: ", currentEntry)
	h.privateTool.emptyMapCheck(h)
	// fmt.Println("Func Remove Log-2 >> emptyMapCheck done: ", h.i,"-->", h.h)

	return nil
}

func (t *tool) addEntry(h *hashMap, data keyValue) {
	index := h.Index(data.Key)
	entries := h.h[h.i].m[index]
	if entries == nil { // no collision, create new list
		e := list.NewLinkedList()
		// fmt.Println("Func Set Log-9 >> Empty entries, created new entry =", newEntry)
		h.h[h.i].m[index] = &e
		e.Add(&data)
	} else { // collision, add to current list
		e := *entries
		e.Add(&data)
	}
	h.ocp++
	h.h[h.i].ocp++
}

func (t *tool) convertToKeyValue(data interface{}) (keyValue, error) {
	keyValue, ok := data.(*keyValue)
	if !ok {
		return *keyValue, fmt.Errorf("input cannot be converted to keyValue")
	} else {
		return *keyValue, nil
	}
}

func (t *tool) searchIfKeyExists(h *hashMap, key string) (int, int, interface{}) {
	searchEntry := &keyValue{Key: key}
	for i := h.i; i >= 0; i-- { //Check if entry exists in the hashmap
		index := HashFunc(key, h.h[i].cap)
		if h.h[i].m[index] != nil { //if h[i] hashmap is not nil, check if key exists
			entries := *h.h[i].m[index]
			keyExists, currentEntry := entries.Contains(searchEntry)
			// fmt.Println("Func searchIfKeyExists Log-1 >> keyExists =", keyExists)

			if keyExists {
				// fmt.Println("Func searchIfKeyExists Log-2 >> currentEntry =", currentEntry)
				return i, index, currentEntry
			}
		}
	}
	return -1, -1, nil
}

func (t *tool) emptyMapCheck(h *hashMap) error {
	if h.i == 0 {
		return fmt.Errorf("only ONE hash map exists")
	}
	for i := h.i; i >= 0; i-- {
		if h.h[i].ocp == 0 {
			h.h = append(h.h[:i], h.h[i+1:]...)
			h.i--
			// fmt.Println(h.h)
			return nil
		}
	}
	return nil
}
