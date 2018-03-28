package main

import "fmt"

type ListNode struct {
	 Val int
	 Next *ListNode
 }

 func main(){
 	var l1, l2, l1_start, l2_start *ListNode = new(ListNode), new(ListNode),new(ListNode),new(ListNode)

 	var val1 []int = []int{2,4,3}
 	var val2 []int = []int{5,6,4}

	 l1_start = l1
	 l2_start = l2
 	for i := 0; i <len(val1); i++ {
		var temp1 *ListNode = newListNode(val1[i])
		l1.Next = temp1
		l1 = l1.Next
	}

	for i := 0; i <len(val2); i++ {
		var temp2 *ListNode = newListNode(val2[i])
		l2.Next = temp2
		l2 = l2.Next
	}
 	var l3 *ListNode = addTwoNumbers(l1_start.Next, l2_start.Next)
 	for l3 != nil {
		fmt.Print(l3.Val)
		l3 = l3.Next
	}
 }

 func newListNode(val int) *ListNode{
 	var listNode *ListNode
 	listNode.Val = val
 	return listNode
 }

 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
 	sum := 0
 	var l3 *ListNode = new(ListNode)
 	l3_start := l3
 	for l1 != nil || l2 != nil {
 		sum /= 10
 		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil{
			sum += l2.Val
			l2 = l2.Next
		}
		var temp *ListNode = newListNode(sum % 10)
		l3.Next = temp
		l3 = l3.Next
	}
	if sum /10 ==1{
		var temp *ListNode = newListNode(1)
		l3.Next = temp
	}
	return l3_start.Next
 }