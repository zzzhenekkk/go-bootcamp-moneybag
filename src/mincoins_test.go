package main

import (
	"reflect"
	"testing"
)

func TestMinCoins(t *testing.T) {
	tests := []struct {
		val    int
		coins  []int
		result []int
	}{
		{13, []int{1, 5, 10}, []int{10, 1, 1, 1}},
		{23, []int{1, 3, 4, 7, 13, 15}, []int{15, 7, 1}},
		{0, []int{1, 5, 10}, []int{}},
		//{3, []int{2}, []int{}},
		{35, []int{1, 5, 10, 25}, []int{25, 10}},
		//{35, []int{25, 10, 5, 1}, []int{25, 10}},
		{35, []int{1, 1, 1, 5, 10, 25}, []int{25, 10}},
	}

	for _, test := range tests {
		if output := minCoins(test.val, test.coins); !reflect.DeepEqual(output, test.result) {
			t.Errorf("11111Test failed: %d with %v, expected %v, got %v", test.val, test.coins, test.result, output)
		}
	}
}

func TestMinCoins2(t *testing.T) {
	tests := []struct {
		val    int
		coins  []int
		result []int
	}{
		{13, []int{1, 5, 10}, []int{10, 1, 1, 1}},
		{23, []int{1, 3, 4, 7, 13, 15}, []int{15, 7, 1}},
		{0, []int{1, 5, 10}, []int{}},
		//{3, []int{2}, []int{}},
		{35, []int{1, 5, 10, 25}, []int{25, 10}},
		{35, []int{25, 10, 5, 1}, []int{25, 10}},
		{35, []int{1, 1, 1, 5, 10, 25}, []int{25, 10}},
	}

	for _, test := range tests {
		if output := minCoins2(test.val, test.coins); !reflect.DeepEqual(output, test.result) {
			t.Errorf("222222Test failed: %d with %v, expected %v, got %v", test.val, test.coins, test.result, output)
		}
	}
}
