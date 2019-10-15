package main

import (
	"math"
)

/*
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/

func maxProfit(prices []int) int {
	length := len(prices)
	max := 0
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			diff := prices[j] - prices[i]
			if max < diff {
				max = diff
			}
		}
	}
	return max
}
func maxProfit2(prices []int) int {
	minPrice, maxProfit := math.MaxInt64, 0
	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		}
		profit := p - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func maxProfit3(prices []int) int {
	minPrice, maxProfit := math.MaxInt64, 0
	for _, p := range prices {
		minPrice = min(p, minPrice)
		maxProfit = max(maxProfit, p-minPrice)
	}
	return maxProfit
}

func maxProfit4(prices []int) int {
	length := len(prices)
	if length == 0 || length == 1 {
		return 0
	}

	profits := make([]int, length-1)
	for i := 1; i < length; i++ {
		profits[i-1] = prices[i] - prices[i-1]
	}

	// maxProfits[i]表示从profits[0]到profits[i]的连续子数组的最大和
	maxProfits := make([]int, length-1)
	maxProfits[0] = profits[0]

	// 记录maxProfits中的最大值
	maxProfit := maxProfits[0]

	for i := 1; i < length-1; i++ {
		maxProfits[i] = max(profits[i], maxProfits[i-1]+profits[i])
		maxProfit = max(maxProfits[i], maxProfit)
	}
	return max(maxProfit, 0)
}

func main() {
	// prices := []int{1, 2}
	// max := maxProfit(prices)
	// max2 := maxProfit2(prices)
	// max3 := maxProfit3(prices)
	// max4 := maxProfit4(prices)
	// fmt.Printf("max4:%d\n", max4)

	nums := []int{2, 1, 4, 9, 8, 5, 10}
	lis(nums)
}
