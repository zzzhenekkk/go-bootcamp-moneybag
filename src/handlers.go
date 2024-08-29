package main

import "sort"

// minCoins calculates the minimum number of coins needed for a given amount.
// @Summary Minimum number of coins
// @Description Calculates the minimum number of coins needed for a given amount using the largest denominations first.
// @Param val query int true "Amount to change"
// @Param coins query []int true "Available denominations"
// @Success 200 {array} int "List of coins that sum up to the amount"
// @Router /mincoins [get]
func minCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

// minCoins2 calculates the minimum number of coins needed for a given amount.
// This function handles cases where the array of denominations contains duplicates and is unsorted.
// @Summary Minimum number of coins with optimization
// @Description Calculates the minimum number of coins needed for a given amount, optimizes by removing duplicates and sorting denominations in descending order.
// @Param val query int true "Amount to change"
// @Param coins query []int true "Available denominations"
// @Success 200 {array} int "Optimized list of coins that sum up to the amount"
// @Router /mincoins2 [get]
func minCoins2(val int, coins []int) []int {
	if len(coins) == 0 {
		return []int{}
	}

	// Удаление дубликатов
	coinMap := make(map[int]struct{})
	for _, coin := range coins {
		coinMap[coin] = struct{}{}
	}

	// Создание уникального среза монет
	uniqueCoins := make([]int, 0, len(coinMap))
	for coin := range coinMap {
		uniqueCoins = append(uniqueCoins, coin)
	}
	// Сортировка в порядке убывания
	sort.Sort(sort.Reverse(sort.IntSlice(uniqueCoins)))

	// Логика для нахождения минимального количества монет
	res := make([]int, 0)
	for _, coin := range uniqueCoins {
		for val >= coin {
			val -= coin
			res = append(res, coin)
		}
	}
	return res
}
