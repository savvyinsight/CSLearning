package main

// 900. RLE Iterator
// https://leetcode.com/problems/rle-iterator/
/*
Thinking process:
1. We can use an index to keep track of our current position in the encoding array.
2. For each call to next(n), we check if the current count (encoding[index]) is greater than or equal to n.
   - If it is, we subtract n from the count and return the corresponding value (encoding[index + 1]).
   - If it is not, we subtract the count from n and move to the next pair of count and value by incrementing the index by 2.
3. If we exhaust the encoding array without finding a valid count, we return -1.

Time complexity: O(k) where k is the number of pairs in the encoding array.
Space complexity: O(1) since we are using only a constant amount of extra space.
*/

type RLEIterator struct {
	encoding []int
	index    int
}

func Constructor(encoding []int) RLEIterator {
	return RLEIterator{
		encoding: encoding,
		index:    0,
	}
}

func (this *RLEIterator) Next(n int) int {
	for this.index < len(this.encoding) {
		if this.encoding[this.index] >= n {
			this.encoding[this.index] -= n
			return this.encoding[this.index+1]
		} else {
			n -= this.encoding[this.index]
			this.index += 2
		}
	}
	return -1
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(encoding);
 * param_1 := obj.Next(n);
 */

func main() {
	iterator := Constructor([]int{3, 8, 0, 9, 2, 5})
	println(iterator.Next(2)) // return 8
	println(iterator.Next(1)) // return 8
	println(iterator.Next(1)) // return 5
	println(iterator.Next(2)) // return -1
}
