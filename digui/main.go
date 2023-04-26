package main

import "fmt"

/*
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
	return nil
}

func Transform(l1 *ListNode) *ListNode {
	list := new(ListNode)
	list.Next = nil
	p := new(ListNode)
	for l1 != nil {
		p.Val = l1.Val
		p.Next = list.Next
		list.Next = p
		l1 = l1.Next
	}
	return list
}

//func main() {
//	l1 := ListNode{
//		Val:  3,
//		Next: nil,
//	}
//	l2 := ListNode{
//		Val:  2,
//		Next: &l1,
//	}
//	l3 := ListNode{
//		Val:  1,
//		Next: &l2,
//	}
//	l4 := Transform(&l3)
//	for l4 != nil {
//		fmt.Println(l4.Val)
//		l4 = l4.Next
//	}
//}

func twoSum(nums []int, target int) []int {
	slice := make([]int, 0)
	var i = 0
	var j = 0
	var t = 0
	for i = 0; i < len(nums); i++ {
		t = target - nums[i]
		for j = 0; j < len(nums)/2; j++ {
			if t == nums[j] && i != j {
				slice = append(slice, i, j)
			}
		}
	}
	return slice
}

// 链表逆置函数
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func main() {
	// 创建链表 1 -> 2 -> 3 -> 4 -> 5
	node5 := &ListNode{Val: 5, Next: nil}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	fmt.Println("原链表：")
	printList(node1)

	// 链表逆置
	newHead := reverseList(node1)

	fmt.Println("逆置后的链表：")
	printList(newHead)
}

// 打印链表
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}
