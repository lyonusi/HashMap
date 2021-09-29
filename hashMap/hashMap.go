package hashMap

import (
	"HashMap/list"
	"fmt"
)

const (
	expandFactor = 0.75
)

type hashMap struct {
	hashMaps     []*hm
	hashMapIndex int
	cap          int
	ocp          int
	privateTool  tools
}

type hm struct {
	singleHashMap []*list.LinkedList
	cap           int
	ocp           int
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
	ManualRehash()
	Remove(key string) error
	Set(key string, value interface{})
	Print()
}

func (h *hashMap) Clear() {
	h.hashMapIndex = 0
	h0 := new(hm)
	h0.cap = h.cap
	h0.singleHashMap = make([]*list.LinkedList, h.cap)
	h.hashMaps = nil
	h.hashMaps = append(h.hashMaps, h0)

	h.hashMaps[h.hashMapIndex] = h0
	h.ocp = 0
}

func (h *hashMap) ManualRehash() {
	list := h.ListKeys()
	numOfKeys := len(list)
	for i := 0; i < numOfKeys; i++ {
		// fmt.Println("Func ManualRehash Log-1 >> i =", i)
		h.Get(list[i])
	}
}

func (h *hashMap) Length() int {
	return h.ocp
}

func (h *hashMap) ListKeys() []string {
	var keyList []string
	for i := h.hashMapIndex; i >= 0; i-- {
		// fmt.Println("Func ListKeys Log-1 >> i =", i)
		for j := 0; j < h.hashMaps[i].cap; j++ {
			// fmt.Println("Func ListKeys Log-2 >> j =", j)
			list := h.hashMaps[i].singleHashMap[j]
			if list != nil {
				entries := *h.hashMaps[i].singleHashMap[j]
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
						entry = iterator.GetData()

						iterator.Next()
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
	if capacity < 16 {
		capacity = 16
	}

	h := &hashMap{
		cap:          capacity,
		hashMapIndex: 0,
		hashMaps: []*hm{{
			cap:           capacity,
			singleHashMap: make([]*list.LinkedList, capacity),
		},
		},
		privateTool: &tool{},
	}
	return h
}

func Expand(h *hashMap) HashMap {
	newH := &hm{
		cap:           2 * h.hashMaps[h.hashMapIndex].cap,
		singleHashMap: make([]*list.LinkedList, 2*h.hashMaps[h.hashMapIndex].cap),
	}
	// newH := new(hm)
	// newH.cap = 2 * h.hashMaps[h.hashMapIndex].cap
	// newH.singleHashMap = make([]*list.LinkedList, newH.cap)

	// fmt.Println("******expand", newH.cap, newH.singleHashMap)

	// h = &hashMap{
	// 	hashMaps:     append(h.hashMaps, newH),
	// 	cap:          newH.cap,
	// 	hashMapIndex: h.hashMapIndex + 1,
	// }

	h.hashMaps = append(h.hashMaps, newH)
	h.cap = newH.cap
	h.hashMapIndex++
	h.hashMaps[h.hashMapIndex] = newH

	// fmt.Println("******expand2", h.cap, h.hashMapIndex, h.hashMaps[h.hashMapIndex])

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
	fmt.Println("hashMap.h --->", h.hashMaps)

	for i := 0; i <= h.hashMapIndex; i++ {
		fmt.Printf(".................... hashMap.hashMaps[%v] ....................\n", i)
		fmt.Printf("                    hashMap.hashMaps[%v].cap = %v\n", i, h.hashMaps[i].cap)
		fmt.Printf("                    hashMap.hashMaps[%v].ocp = %v\n", i, h.hashMaps[i].ocp)
		fmt.Printf("                    hashMap.hashMaps[%v].singleHashMap ---> %v\n", i, h.hashMaps[i].singleHashMap)
		for j := 0; j < h.hashMaps[i].cap; j++ {
			entries := h.hashMaps[i].singleHashMap[j]
			if entries == nil {
				fmt.Printf("                    hashMap.%v.%v : (nil)\n", i, j)
			} else {
				l := *h.hashMaps[i].singleHashMap[j]
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
	hashMapIndex, entryIndex, _ := h.privateTool.searchIfKeyExists(h, key) //Check if key exists in the hashmap

	if hashMapIndex != -1 && entryIndex != -1 { // if key exists...
		// fmt.Println("Func Set Log-1 >> equal? =", currentEntry == newEntry)

		h.Remove(key) // remove current entry
		// fmt.Println("Func Set Log-2 >> removed:", currentEntry)

		h.privateTool.addEntry(h, *newEntry) // add new entry to latest hashmap
		// fmt.Println("Func Set Log-3 >> added:", newEntry)
	} else if (float64(h.ocp) / float64(h.cap)) > expandFactor { // if key does not exist and hash map is occupied > 0.75
		// fmt.Println("Func Set Log-4 >> hashMap.ocp =", h.ocp)

		newH := Expand(h)
		entryIndex := newH.Index(key)
		newEntries := list.NewLinkedList()
		// fmt.Println("Func Set Log-5 >> New entry (expanded) =", newEntry)
		h.hashMaps[h.hashMapIndex].singleHashMap[entryIndex] = &newEntries
		newEntries.Add(newEntry)
		h.ocp++
		// fmt.Println("Func Set Log-6 >> Expanded hashMap.ocp =", h.ocp)

		h.hashMaps[h.hashMapIndex].ocp++
		for i := h.hashMapIndex; i >= 0; i-- {
			fmt.Printf("Func Set Log-7 >> Expanded hashMap.h[%v].ocp = %v\n", i, h.hashMaps[i].ocp)
		}
	} else { // if key does not exist
		h.privateTool.addEntry(h, *newEntry) // add new entry to latest hashmap
	}
}

func (h *hashMap) Get(key string) interface{} {
	searchEntry := &keyValue{Key: key}

	for i := h.hashMapIndex; i >= 0; i-- { //Check if entry exists in the hashmap
		index := HashFunc(key, h.hashMaps[i].cap)
		if h.hashMaps[i].singleHashMap[index] != nil { //if h[i] hashmap is not nil, check if key exists
			entries := *h.hashMaps[i].singleHashMap[index]
			keyExists, currentEntry := entries.Contains(searchEntry)
			// fmt.Println("Func Get Log-1 >> keyExists =", keyExists)

			if keyExists {
				entry, err := h.privateTool.convertToKeyValue(currentEntry)
				if err != nil {
					fmt.Println(err)
					return err
				}
				if i != h.hashMapIndex { // if entry is not in the latest hashmap, move to the latest hashmap h[h.i]
					fmt.Println("Func Get Log-2 >> moved entry to new hash map:", i, "->", h.hashMapIndex)
					h.Remove(key)
					h.privateTool.addEntry(h, entry)
				}
				// fmt.Println("Func Get Log-3 >> entry =", entry)
				return entry.Value
			}
		}
	}
	return fmt.Errorf("not found")
}

func (h *hashMap) Remove(key string) error {
	hIndex, mIndex, _ := h.privateTool.searchIfKeyExists(h, key)
	if hIndex != -1 && mIndex != -1 {
		entries := *h.hashMaps[hIndex].singleHashMap[mIndex]
		_, entry := entries.Contains(&keyValue{Key: key})
		index := entries.IndexOf(entry)
		entries.RemoveByIndex(index)
		if entries.GetLength() == 0 {
			h.hashMaps[hIndex].singleHashMap[mIndex] = nil
		}
		h.ocp--
		h.hashMaps[hIndex].ocp--
	}
	// fmt.Println("Func Remove Log-1 >> removed: ", currentEntry)
	h.privateTool.emptyMapCheck(h)
	// fmt.Println("Func Remove Log-2 >> emptyMapCheck done: ", h.i,"-->", h.h)

	return nil
}

func (t *tool) addEntry(h *hashMap, data keyValue) {
	index := h.Index(data.Key)
	entries := h.hashMaps[h.hashMapIndex].singleHashMap[index]
	if entries == nil { // no collision, create new list
		e := list.NewLinkedList()
		// fmt.Println("Func Set Log-9 >> Empty entries, created new entry =", newEntry)
		h.hashMaps[h.hashMapIndex].singleHashMap[index] = &e
		e.Add(&data)
	} else { // collision, add to current list
		e := *entries
		e.Add(&data)
	}
	h.ocp++
	h.hashMaps[h.hashMapIndex].ocp++
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
	for i := h.hashMapIndex; i >= 0; i-- { //Check if entry exists in the hashmap
		index := HashFunc(key, h.hashMaps[i].cap)
		if h.hashMaps[i].singleHashMap[index] != nil { //if h[i] hashmap is not nil, check if key exists
			entries := *h.hashMaps[i].singleHashMap[index]
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
	if h.hashMapIndex == 0 {
		return fmt.Errorf("only ONE hash map exists")
	}
	for i := h.hashMapIndex; i >= 0; i-- {
		if h.hashMaps[i].ocp == 0 {
			h.hashMaps = append(h.hashMaps[:i], h.hashMaps[i+1:]...)
			h.hashMapIndex--
			// fmt.Println(h.h)
			return nil
		}
	}
	return nil
}
