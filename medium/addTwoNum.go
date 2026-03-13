package medium

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

}

	/* Fail case that larger than even uint64
	suml1 := 0
	suml2 := 0
	count := 1
	for l1 != nil {
		suml1 += l1.Val * count
		if l1.Next != nil {
			l1 = l1.Next
			count *= 10
			fmt.Println(suml1)
		} else {
			break
		}
	}
	fmt.Println(suml1)
	count = 1
	for l2 != nil {
		suml2 += l2.Val * count
		if l2.Next != nil {
			l2 = l2.Next
			count *= 10
			fmt.Println(suml2)
		} else {
			break
		}
	}
	fmt.Println(suml2)

	count = suml1 + suml2
	ans := &ListNode{}
	cur := ans
	for count >= 1 {
		cur.Val = count%10
		count /= 10
		if count >= 1{
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	for ans != nil{
		fmt.Print(ans.Val)
		ans = ans.Next
	}
	
	return ans
	*/
