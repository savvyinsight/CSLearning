package main

import "fmt"

// https://leetcode.com/problems/product-of-array-except-self/

/*
Thinking process:
1. Brute force: For each element, calculate the product of all other elements.
	- TC:O(n^2),SC:O(n)
2. Prefix and Suffix products:
	- prefix[i] = product of all elements to the left of i
	- suffix[i] = product of all elements to the right of i
	- result[i] = prefix[i] * suffix[i]
	- TC:O(n),SC:O(n)
3. Division :
	- Calculate the total product of all elements, then for each element, divide the totaol product.
	- if nums[i] ==0, and there is only one zero in the array, then result[i] = product of all non-zero elemnts,
	- if there are more than one zero, then all result will be zero.
	- TC:O(n),SC:O(1)
*/

// Brute force
func productExceptSelf_1(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		prod := 1
		for j := 0; j < n; j++ {
			if i != j {
				prod *= nums[j]
			}
		}
		res[i] = prod
	}
	return res
}

// prefix and suffix products
func productExceptSelf_2(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// prefix
	prefix := 1
	for i := 0; i < n; i++ {
		res[i] = prefix
		prefix *= nums[i]
	}
	// suffix
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= suffix
		suffix *= nums[i]
	}
	return res
}

// division
func productExceptSelf_3(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	totalProduct := 1
	zeroCount := 0

	for _, num := range nums {
		if num == 0 {
			zeroCount++
			continue
		}
		totalProduct *= num
	}

	for i, num := range nums {
		if zeroCount > 1 {
			res[i] = 0
		} else if zeroCount == 1 {
			if num == 0 {
				res[i] = totalProduct
			} else {
				res[i] = 0
			}
		} else {
			res[i] = totalProduct / num
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 4, 6}
	result := productExceptSelf_3(nums)
	fmt.Println(result)
}
