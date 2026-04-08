package main

import (
	"bufio"
	"container/heap"
	"container/list"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

/*
Method			Speed	Ease of Use		Best For
fmt.Scan		Slow	Very easy		Small inputs (<10⁵)
bufio.Scanner	Medium	Easy	Most 	competitions
Custom reader	Fast	Hard	Very 	large inputs (>10⁶)

*/

// input from keyboard
func fmtScan() {
	// 1.fmt.Scan family (Easy but slow)
	var n int
	fmt.Scan(&n) // Reads one int
	fmt.Print("n is :", n)
	var a, b int
	fmt.Scan(&a, &b) // Reads multiple
	fmt.Print("a,b is :", a, b)
	var s string
	fmt.Scan(&s) // Reads string (whitespace-separated)
	fmt.Print("s is :", s)
	//
}

// Reads raw line or tokens with full control
// Buffered, handles large input efficiently
func bufioScan() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords) //space-seperated words
	// Read next token
	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, 10, 10)
	fmt.Println(arr)
	// Reading multiple numbers
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(arr)

	// Reading entire line
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	line := scanner1.Text()
	fmt.Println(line)
}

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
	arr := []int{3, 1, 4, 1, 5, 8, 2, 7}
	sort.Ints(arr) // ascending

	fmt.Println(arr)

	// Descending
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	fmt.Println(arr)

	// Custom comparator
	sort.Slice(arr, func(i, j int) bool {
		return arr[i]%2 < arr[j]%2 // evens first
	})

	fmt.Println(arr)

	// Sort structs
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{{"Alice", 30}, {"Bob", 25}}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}

func sliceCopy() {
	original := []int{1, 2, 3, 4, 5}
	// Method 1: Using built-in copy function
	copied := make([]int, len(original))
	copy(copied, original)

	fmt.Println("Original:", original)
	fmt.Println("Copied:", copied)

	// Method 2: Using append to create a new slice
	copied2 := append([]int(nil), original...)
	fmt.Println("Copied2:", copied2)
}

func packAndUnpack() {
	// Packing multiple values into a struct
	type Point struct {
		X, Y int
	}
	p := Point{X: 3, Y: 4}

	// Unpacking values from a struct
	x, y := p.X, p.Y
	fmt.Printf("x: %d, y: %d\n", x, y)

	// For simple cases, can also return multiple values from a function
	a, b := getPair()
	fmt.Printf("a: %d, b: %s\n", a, b)

	original := []int{1, 2, 3}
	copied := append([]int(nil), original...) // creates a new slice with same contents
	// same as copied := append([]int{},1,2,3) or copied := make([]int, len(original)); copy(copied, original)

	fmt.Println("Original:", original)
	fmt.Println("Copied:", copied)

	// Merge two slices
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	merged := append(slice1, slice2...)
	fmt.Println("Merged:", merged)

	// use sum
	sum := func(nums ...int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}
	nums := []int{1, 2, 3, 4}
	sumResult := sum(nums...) // unpack slice into variadic function, same as sum(1, 2, 3, 4)
	fmt.Println("Sum:", sumResult)
}

func hashMap() {
	freq := make(map[string]int)
	freq["apple"] = 5
	freq["banana"]++
	freq["good"]++
	freq["good"]++

	fmt.Println(freq)

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

	fmt.Println(freq)
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

// ==============================================HEAP========
// Using `container/heap` (standard library), go provides heap interface, that you need to
// implement.

// Min-heap (Go's heap is always min-heap for ints)
// IntHeap implements heap.Interface
type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

// For min-heap priority < , for max-heap priority >
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // min-heap

func (h IntHeap) Swap(i, j int) { // Not need pointer, slice share underlying array
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) { // Need pointer, will slice length changes.(slice header)
	*h = append(*h, x.(int))
}

// Slice header (ptr, len, cap) is passed by value (copy) --->need pointer to change
// Underlying array is shared (reference)

func (h *IntHeap) Pop() interface{} { // Need pointer, will slice length changes.(slice header)
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

func String() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}

	//The %q (quoted) verb will escape any non-printable byte
	// sequences in a string so the output is unambiguous.
	fmt.Printf("\n%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)
}

func runeAndbyte() {
	s := "hello世界"

	fmt.Println(len(s))
	fmt.Println(s[5], s[6], s[7])

	// Indexing and Ranging
	// s[i] returns byte (uint8), not character
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] = %v (type: %T)\n", i, s[i], s[i]) // byte (uint8)
	}

	// range returns rune (int32) and index of rune
	for i, ch := range s {
		fmt.Printf("s[%d] = %c (type: %T)\n", i, ch, ch) // rune (int32)
	}
	// return bytes not characters
	fmt.Println(len(s))

	// actual character count
	fmt.Println(len([]rune(s)))

	// Converting between them
	bytes := []byte(s) // 11 bytes
	runes := []rune(s) // 7 runes (5 latin + 2 CJK)
	fmt.Println(bytes, runes)

	// Strings inmutable, so first need convert to []rune or []byte
	// s[2] = "d" , false
	arr := []byte(s)
	arr[2] = 'k'
	fmt.Println(arr) // returns byte slice, not string

	// if want string, need convert back
	// arr := []rune(s)
	// arr[2] = 'k'
	fmt.Println(string(arr))
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

func joinStrings(sep string, words ...string) string {
	if len(words) == 0 {
		return ""
	}
	res := words[0]
	for i := 1; i < len(words); i++ {
		res += sep + words[i]
	}
	return res
}

func testJoinStrings() {
	words := []string{"good", "boy", "hello"}
	fmt.Println(joinStrings(",", words...))
}
func main() {
	runeAndbyte()
}
