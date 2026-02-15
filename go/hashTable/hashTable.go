package main

import "fmt"

// The hash function can be used directly
func hashSimple(str string, cap int) int {
	seed := int(131)
	hash := int(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + int(str[i])
	}
	return hash % cap
}

type HashNode struct {
	key   string
	value interface{}
	next  *HashNode
}

// Please implement the Hash Table data structure and design its key interfaces
type HashTable struct {
	capacity int
	size     int
	table    []*HashNode
}

// Create
func NewHashTable(capacity int) *HashTable {
	return &HashTable{
		capacity: capacity,
		table:    make([]*HashNode, capacity),
	}
}

// Insert key-value
func (ht *HashTable) Insert(key string, value interface{}) {
	index := hashSimple(key, ht.capacity)
	newNode := &HashNode{key: key, value: value}

	if ht.table[index] == nil {
		ht.table[index] = newNode
	} else {
		// Handle collision
		curr := ht.table[index]
		for curr.next != nil {
			if curr.key == key {
				curr.value = value
				return
			}
			curr = curr.next
		}
		curr.next = newNode
	}
	ht.size++
}

// Get
func (ht *HashTable) Get(key string) (interface{}, bool) {
	index := hashSimple(key, ht.capacity)
	curr := ht.table[index]

	for curr != nil {
		if curr.key == key {
			return curr.value, true
		}
		curr = curr.next
	}
	return nil, false
}

// Remove
func (ht *HashTable) Remove(key string) bool {
	index := hashSimple(key, ht.capacity)
	curr := ht.table[index]
	var prev *HashNode

	for curr != nil {
		if curr.key == key {
			if prev == nil {
				ht.table[index] = curr.next
			} else {
				prev.next = curr.next
			}
			ht.size--
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}

// Please add your test code
func main() {
	ht := NewHashTable(10)

	// Insert and Get
	ht.Insert("n1", "bob")
	ht.Insert("n2", 22)
	ht.Insert("n3", "Hello World!")

	if n, ok := ht.Get("n1"); ok {
		fmt.Println("n1:", n)
	}
	if n, ok := ht.Get("n2"); ok {
		fmt.Println("n2:", n)
	}
	if n, ok := ht.Get("n3"); ok {
		fmt.Println("n3:", n)
	}
	fmt.Println("size:", ht.size)

	ht.Remove("n2")
	if n, ok := ht.Get("n2"); !ok {
		fmt.Println("n2 removed")
	} else {
		fmt.Println(n)
	}
	fmt.Println("size:", ht.size)

}
