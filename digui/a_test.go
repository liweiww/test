package main

import (
	"testing"
)

func nextPermutation() {
	i := 0
	l := 0
	nums := []int{1, 3, 2}
	for i = len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			l = i - 1
			break
		}
	}
	if l > 0 {
		r := 0
		for i = len(nums) - 1; i > 0; i-- {
			if nums[i] > nums[l] {
				r = i
				break
			}
		}
		t := nums[l]
		nums[l] = nums[r]
		nums[r] = t
		reverse(nums, l+1, len(nums)-1)
	} else {
		reverse(nums, 0, len(nums)-1)
	}
}
func reverse(nums []int, i int, j int) {
	for p, q := i, j; p < q; p, q = p+1, q-1 {
		t := nums[p]
		nums[p] = nums[q]
		nums[q] = t
	}
}

func TestName(t *testing.T) {
	GenerateParenthesis(2)
}

func GenerateParenthesis(n int) []string {
	return dp(n)[n]
}

func dp(n int) map[int][]string {
	if n == 0 {
		return map[int][]string{0: {""}}
	}
	if n == 1 {
		return map[int][]string{0: {""}, 1: {"()"}}
	}
	lastMap := dp(n - 1)
	var oneRes []string
	for i := 0; i < n; i++ {
		inners := lastMap[i]
		outers := lastMap[n-1-i]
		for _, inner := range inners {
			for _, outer := range outers {
				oneRes = append(oneRes, "("+inner+")"+outer)
			}
		}
	}
	lastMap[n] = oneRes
	return lastMap
}
