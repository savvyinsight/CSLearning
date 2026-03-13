package main

// 121. Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

/*
Thinking process:
1. Brute force:
   - Iterate through all pairs of days and calculate the profit for each pair.
   - Time complexity: O(n^2)
   - Space complexity: O(1)

2. Two pointer:
   - Use two pointers to track the minimum price and the maximum profit.
   - Iterate through the prices, updating the minimum price and calculating the profit at each step.
   - Time complexity: O(n)
   - Space complexity: O(1)

3. Dynamic programming:
   - Keep track of the minimum price seen so far and the maximum profit at each step.
   - Time complexity: O(n)
   - Space complexity: O(1)
*/

// 1.approach 1: Brute force
func maxProfit_1(prices []int) int {
	res := 0
	for i := 0; i < len(prices); i++ {
		buy := prices[i]
		for j := i + 1; j < len(prices); j++ {
			sell := prices[j]
			res = max(res, sell-buy)
		}
	}
	return res
}

// approach 2: Two pointer
func maxProfit_2(prices []int) int {
	res := 0
	l, r := 0, 1
	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			res = max(res, profit)
		} else {
			l = r
		}
		r++
	}
	return res
}

// approach 3: Dynamic programming
func maxProfit_3(prices []int) int {
	maxP := 0
	minPrice := prices[0]

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxP {
			maxP = price - minPrice
		}
	}
	return maxP
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	println(maxProfit_1(prices)) // Output: 5
	println(maxProfit_2(prices)) // Output: 5
}
