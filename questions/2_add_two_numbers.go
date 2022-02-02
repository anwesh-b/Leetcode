package leetcode

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result ListNode
	var isL1Nil, isL2Nil bool
	var carry int
	var templ1, templ2, tempRes *ListNode

	carry = 0
	templ1 = l1
	templ2 = l2
	tempRes = &result

	for isL1Nil != true || isL2Nil != true {
		sum := carry
		var nextRes ListNode
		if isL1Nil != true {
			sum += templ1.Val
			if templ1.Next != nil {
				templ1 = templ1.Next
			} else {
				isL1Nil = true
			}
		}
		if isL2Nil != true {
			sum += templ2.Val
			if templ2.Next != nil {
				templ2 = templ2.Next
			} else {
				isL2Nil = true
			}
		}
		carry = sum / 10
		sum = sum % 10
		tempRes.Val = sum
		if isL1Nil != true || isL2Nil != true {
			tempRes.Next = &nextRes
			tempRes = tempRes.Next
		}
	}

	if carry > 0 {
		var nextRes ListNode
		tempRes.Next = &nextRes
		tempRes.Next.Val = carry
	}

	return &result
}

func Q_2() {
	var l1, l2 ListNode
	l1.Val = 9
	var l11, l12, l13, l14, l15, l16 ListNode
	l1.Next = &l11
	l11.Val = 9
	l11.Next = &l12
	l12.Val = 9
	l12.Next = &l13
	l13.Val = 9
	l13.Next = &l14
	l14.Val = 9
	l14.Next = &l15
	l15.Val = 9
	l15.Next = &l16
	l16.Val = 9
	l16.Next = nil

	l2.Val = 9
	var l21, l22, l23 ListNode
	l2.Next = &l21
	l21.Val = 9
	l21.Next = &l22
	l22.Val = 9
	l22.Next = &l23
	l23.Val = 9
	l23.Next = nil

	res := addTwoNumbers(&l1, &l2)

	fmt.Println(res.Next.Next.Next.Next.Next.Next.Next.Next.Val)
}
