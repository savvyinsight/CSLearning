package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func maxMin() {
	minInt := math.MinInt
	maxInt := math.MaxInt
	fmt.Println(minInt)
	fmt.Println(maxInt)
	// For specific sizes
	// minInt32 := math.MinInt32
	// maxInt32 := math.MaxInt32
	// minInt64 := math.MinInt64
	// maxInt64 := math.MaxInt64

	// Using the actual minimum values
	const (
		MinInt8  = -1 << 7
		MaxInt8  = 1<<7 - 1
		MinInt16 = -1 << 15
		MaxInt16 = 1<<15 - 1
		MinInt32 = -1 << 31
		MaxInt32 = 1<<31 - 1
		MinInt64 = -1 << 63
		MaxInt64 = 1<<63 - 1
	)
}

func sorting() {
	arr := []int{3, 1, 4, 1, 5}
	sort.Ints(arr) // ascending

	// Descending
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	// Custom comparator
	sort.Slice(arr, func(i, j int) bool {
		return arr[i]%2 < arr[j]%2 // evens first
	})

	// Sort structs
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{{"Alice", 30}, {"Bob", 25}}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
}

func hashMap() {
	freq := make(map[string]int)
	freq["apple"] = 5
	freq["banana"]++

	// Check existence
	if val, exists := freq["apple"]; exists {
		// val is the value
		fmt.Println(val)
	}

	// Iteration
	for key, value := range freq {
		fmt.Printf("%s: %d\n", key, value)
	}

	// Default value (int = 0)
	val := freq["nonexistent"] // returns 0
	fmt.Println(val)

	// Delete
	delete(freq, "apple")
}

// max-heap/min-heap
func priorityQueue() {
	// Usage
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	min := (*h)[0] // peek
	fmt.Println(min)
	heap.Pop(h) // remove min
}

// Min-heap (Go's heap is always min-heap for ints)
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // min-heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Max-heap (negate values or use custom comparator)
type MaxHeap []int

func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // max-heap

func set() {
	// Go doesn't have built-in set, use map with bool value
	set := make(map[int]bool)
	set[5] = true
	delete(set, 5)
	if _, exists := set[5]; exists {
		// exists
	}
}

func StackQueue() {
	// Stack
	stack := []int{}
	stack = append(stack, 1)   // push
	top := stack[len(stack)-1] // peek
	fmt.Println(top)
	stack = stack[:len(stack)-1] // pop

	// Queue (using slice)
	queue := []int{}
	queue = append(queue, 1) // enqueue
	front := queue[0]        // peek
	fmt.Println(front)
	queue = queue[1:] // dequeue (inefficient for large queues)

	// Better queue (using list or channels)
	l := list.New()
	l.PushBack(1)             // enqueue
	front1 := l.Front().Value // peek
	fmt.Println(front1)
	l.Remove(l.Front()) // dequeue
}

func binarySeach() {

	/*
			#include <algorithm>
		vector<int> arr = {1, 3, 5, 7, 9};
		auto it = lower_bound(arr.begin(), arr.end(), 5);  // first >= 5
		auto it2 = upper_bound(arr.begin(), arr.end(), 5); // first > 5
		bool found = binary_search(arr.begin(), arr.end(), 5);

	*/

	arr := []int{1, 3, 5, 7, 9}
	// sort.Ints(arr) // must be sorted first

	// Binary search returns index where target would be inserted
	idx := sort.SearchInts(arr, 5) // first index where arr[i] >= 5
	found := idx < len(arr) && arr[idx] == 5
	fmt.Println(found)
	// Custom binary search
	idx = sort.Search(len(arr), func(i int) bool {
		return arr[i] >= 5 // lower_bound equivalent
	})
}

func pair() {

}

// Use struct or multiple return values
type Pair struct {
	First  int
	Second string
}

// p := Pair{1, "hello"}

// Or for simple cases, return multiple values
func getPair() (int, string) {
	return 1, "hello"
}

// a, b := getPair()

func randomRange() {
	// Random int between [0, n) (0 to n-1)
	n := 10
	rand.Intn(n)

	// Random float between [0.0, 1.0)
	rand.Float64()

	// Random float between [min, max)
	// min + rand.Float64()*(max-min)

	// Random int between [min, max] inclusive
	// min + rand.Intn(max-min+1)
}

func runeAndbyte() {
	s := "hello世界"

	// s[i] returns byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] = %v (type: %T)\n", i, s[i], s[i]) // byte (uint8)
	}

	// range returns rune
	for i, ch := range s {
		fmt.Printf("s[%d] = %c (type: %T)\n", i, ch, ch) // rune (int32)
	}
	fmt.Println(len(s))
}

func lengthOfLongestSubstringByte(s string) int {
	charIndex := make(map[byte]int)
	maxLen := 0
	l := 0

	// Using integer loop with s[i] → returns byte
	for r := 0; r < len(s); r++ {
		ch := s[r] // byte

		if idx, exists := charIndex[ch]; exists && idx >= l {
			l = idx + 1
		}
		charIndex[ch] = r
		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}
	return maxLen
}

func lengthOfLongestSubstringRune(s string) int {
	charIndex := make(map[rune]int)
	maxLen := 0
	l := 0

	// Using range → returns rune
	runes := []rune(s) // Convert once if using indices
	for r, ch := range runes {
		if idx, exists := charIndex[ch]; exists && idx >= l {
			l = idx + 1
		}
		charIndex[ch] = r
		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}
	return maxLen
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	pivot := partition(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]

	i := left

	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return res
}
func main() {

}
