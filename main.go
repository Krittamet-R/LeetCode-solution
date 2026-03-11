package main

import (
	"LEETCODE-SOLUTION/easy"
	"fmt"
)

func main() {
	// Test palindrome number
	fmt.Println(easy.IsPalindrome(121))
	fmt.Println(easy.IsPalindrome(-121))
	fmt.Println(easy.IsPalindrome(10))
	// Test two sum
	fmt.Println(easy.TwoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(easy.TwoSum([]int{3, 2, 4}, 6))
	fmt.Println(easy.TwoSum([]int{3, 3}, 6))
}
