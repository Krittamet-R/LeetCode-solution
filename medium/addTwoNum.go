package medium

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carryover := 0
	ans := &ListNode{}
	cur := ans
	for l1 != nil || l2 != nil {
		temp := l1.Val + l2.Val + carryover
		fmt.Printf("temp: %d\n", temp)
		carryover = temp / 10
		temp = temp % 10
		fmt.Printf("carryover: %d, temp: %d\n", carryover, temp)

		cur.Val = temp
		l1 = l1.Next
		l2 = l2.Next

		switch{
		case l1 != nil && l2 != nil:
			cur.Next = &ListNode{}
			cur = cur.Next
		case l1 == nil && l2 != nil:
			for l2 != nil {
				cur.Next = &ListNode{}
				cur = cur.Next
				t := (l2.Val + carryover) % 10
				carryover = (l2.Val + carryover) / 10
				cur.Val = t
				l2 = l2.Next
			}
		case l2 == nil && l1 != nil:
			for l1 != nil {
				cur.Next = &ListNode{}
				cur = cur.Next
				t := (l1.Val + carryover) % 10
				carryover = (l1.Val + carryover) / 10
				cur.Val = t
				l1 = l1.Next
			}
		}
	}
	
	if carryover >= 1 {
		cur.Next = &ListNode{}
		cur = cur.Next
		cur.Val = carryover
	}

	head := ans
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

	return ans
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
