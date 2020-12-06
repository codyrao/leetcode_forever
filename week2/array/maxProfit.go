package main

import (
	"math"
)

func maxProfit(prices []int) int {

	min := math.MaxInt32
	maxProfit := 0
	length := len(prices)
	for i := 0; i < length; i++ {

		price := prices[i]
		if price < min {
			min = price
			continue
		}

		if price-min > maxProfit {
			maxProfit = price - min
		}

	}

	return maxProfit
}
