package main

import (
	//"LeetCode-solution/easy"
	"LeetCode-solution/medium"
	"fmt"
)

func main() {
	/*. EASY
		// Test palindrome number
		fmt.Println(easy.IsPalindrome(121))
		// Test two sum
		fmt.Println(easy.TwoSum([]int{2, 7, 11, 15}, 9))
		// Test Roman to integer
		fmt.Println(easy.RomanToInt("LVIII"))
		// Test Longest common prefix
		fmt.Println(easy.LongestCommonPrefix([]string{""}))
		fmt.Println(easy.IsValid("(]"))
	*/
	l1 := &medium.ListNode{Val: 2}
	l1.Next = &medium.ListNode{Val: 4}
	l1.Next.Next = &medium.ListNode{Val: 3}

	l2 := &medium.ListNode{Val: 5}
	l2.Next = &medium.ListNode{Val: 6}
	l2.Next.Next = &medium.ListNode{Val: 4}

	fmt.Println(medium.AddTwoNumbers(l1, l2))
}
