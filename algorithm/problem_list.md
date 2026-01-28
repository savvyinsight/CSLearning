LeetCode Problems

Easy Problems for training program in Go

# 1. List

1. Reverse Linked List(LeetCode 206)
2. Merge List( 21)
3. Cycle Detect(141)
4. Intersection of List(160)



# 2. Array/String

1. Two Sum(1)

2. Merge Sorted Array(88)

3. Valid Palindrome(125)
4. Longest Substring Without Repeating Characters(3)



# 3. Stack and Queue

1. [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)
2. [Min Stack](https://leetcode.com/problems/min-stack/)
3. 

# 4.Tree and Recursion

1. Binary Tree Preorder, Inorder, PostOrder(144, 94, 145)
2. [104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/)
3. [101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/)
4. [102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)

# 5. Sort and Search

1. must can handwrite quickSort and merge sort

   // quickSort

   ```go
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
   ```

   //

2. [704. Binary Search](https://leetcode.com/problems/binary-search/)

3. [153. Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)



# 6. DP

1. [70. Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)
2. [53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)
3. [121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)



# 7. Concurrency

1. ## Alternate printing numbers/letters(using channels or sync)

   ### Solution 1: Using Two Channels (Most Go-idiomatic)

   ```Go
   package main
   
   import (
   	"fmt"
   	"sync"
   )
   
   func main() {
   	numbers := make(chan struct{})
   	letters := make(chan struct{})
   	done := make(chan struct{})
   
   	var wg sync.WaitGroup
   
   	wg.Add(2)
   
   	// Number printer goroutine
   	go func() {
   		defer wg.Done()
   
   		for i := 1; i <= 10; i++ {
   			<-numbers // wait for signal to print number
   			fmt.Printf("%d ", i)
   			letters <- struct{}{} //Signal letter printer
   		}
   	}()
   
   	// Letter printer goroutine
   	go func() {
   		defer wg.Done()
   
   		for i := 0; i < 10; i++ {
   			<-letters // Wait signal for print letter
   			fmt.Printf("%c ", 'a'+i)
   			if i < 9 {
   				numbers <- struct{}{} // signal number printer (except last)
   			}
   		}
   		close(done)
   
   	}()
   
   	// Kickstart the process
   	numbers <- struct{}{}
   
   	<-done // Wait for completion
   	wg.Wait()
   	fmt.Println()
   }
   ```

   ```go	
   //output
   // 1 a 2 b 3 c 4 d 5 e 6 f 7 g 8 h 9 i 10 j
   ```

   ## Solution 2: Using a Single Channel

   ```go
   package main
   
   import (
   	"fmt"
   	"sync"
   )
   
   func main() {
   	ch := make(chan struct{})
   	var wg sync.WaitGroup
   	wg.Add(2)
   
   	go printLetters(ch, &wg)
   	go printNumbers(ch, &wg)
   
   	// Signal start
   	ch <- struct{}{}
   
   	wg.Wait()
   	fmt.Println()
   }
   
   func printNumbers(ch chan struct{}, wg *sync.WaitGroup) {
   	defer wg.Done()
   	for i := 1; i <= 10; i++ {
   		<-ch // Wait for turn
   		fmt.Printf("%d ", i)
   		ch <- struct{}{} // Pass turn
   	}
   }
   
   func printLetters(ch chan struct{}, wg *sync.WaitGroup) {
   	defer wg.Done()
   	for i := 0; i < 10; i++ {
   		<-ch // Wait for turn
   		fmt.Printf("%c ", 'a'+i)
   		if i < 9 {
   			ch <- struct{}{} // Pass turn back (except last)
   		}
   	}
   }
   ```

   ```go
   //output:
   //1 a 2 b 3 c 4 d 5 e 6 f 7 g 8 h 9 i 10 j
   //Note:
   //The squence of : this is matter
   go printLetters(ch, &wg) // Starts first (but doesn't necessarily execute first)
   go printNumbers(ch, &wg)
   
   ```

   

2. ## Implement a simple worker poll

   

3. ## Producer and Consumer Problem

4. ## 
