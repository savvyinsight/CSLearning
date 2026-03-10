package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//347. Top K Frequent Elements
//https://leetcode.com/problems/top-k-frequent-elements/

/*
Thinking process:
1. Sorting
   - We can count the frequency of each element and then sort them based on frequency.
   Finally, we can return the top k elements.
   - Time complexity: O(n log n) due to sorting
   - Space complexity: O(n) for the hash map and sorting
2. Min Heap
   - We can use a min heap to keep track of the top k elements based on frequency.
   We can iterate through the frequency map and maintain a heap of size k.
   - Time complexity: O(n log k) due to heap operations
   - Space complexity: O(n) for the hash map and O(k) for the heap
3. Bucket Sort
   - We can use bucket sort to group elements by their frequency.
   - We can create an array of buckets where the index represents the frequency and
   each bucket contains the elements with that frequency. Finally, we can iterate through
   the buckets in reverse order to get the top k elements.
   - Time complexity: O(n) for counting frequencies and O(n) for iterating through buckets, resulting in O(n) overall
   - Space complexity: O(n) for the hash map and O(n) for the buckets
*/

// approach 1: sorting
func topKFrequent_1(nums []int, k int) []int {
	fre := make(map[int]int)
	for _, num := range nums {
		fre[num]++
	}

	pairs := make([][2]int, 0, len(fre))
	for num, count := range fre {
		pairs = append(pairs, [2]int{num, count})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, pairs[i][0])
	}
	return result
}

// approach 2: min heap TC: O(n k log k) SC: O(n + k) , we can optimize it to O(n log k) by using a more efficient heap implementation
func topKFrequent_2(nums []int, k int) []int {
	fre := make(map[int]int)
	for _, num := range nums {
		fre[num]++
	}

	type pair struct {
		num   int
		count int
	}
	minHeap := make([]pair, 0, k)

	for num, count := range fre {
		minHeap = append(minHeap, pair{num, count})
		if len(minHeap) > k {
			sort.Slice(minHeap, func(i, j int) bool {
				return minHeap[i].count < minHeap[j].count
			})
			minHeap = minHeap[1:]
		}
	}

	result := make([]int, 0, k)
	for _, p := range minHeap {
		result = append(result, p.num)
	}
	return result
}

// apprach 2: min heap with optimized heap operations TC: O(n log k) SC: O(n + k)
// We can implement a more efficient heap using the container/heap package in Go,
// which provides a priority queue implementation. This way, we can maintain the heap property more efficiently and reduce the time complexity to O(n log k).
func topKFrequent_3(nums []int, k int) []int {
	fre := make(map[int]int)
	for _, num := range nums {
		fre[num]++
	}

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for num, count := range fre {
		heap.Push(minHeap, pair{num, count})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	result := make([]int, 0, k)
	for minHeap.Len() > 0 {
		result = append(result, heap.Pop(minHeap).(pair).num)
	}
	return result
}

type pair struct {
	num   int
	count int
}

type MinHeap []pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// approach 3: bucket sort TC: O(n) SC: O(n)
func topKFrequent_4(nums []int, k int) []int {
	fre := make(map[int]int)
	for _, num := range nums {
		fre[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, count := range fre {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		if len(buckets[i]) > 0 {
			result = append(result, buckets[i]...)
		}
	}
	return result[:k]
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent_4(nums, k))
}
