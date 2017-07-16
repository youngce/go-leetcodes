package solutions

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
 }

func getListNodeLen(ls *ListNode) int {
	if ls.Next!=nil{
		return 1+getListNodeLen(ls.Next)
	}
	return 1
}
func extListNode(ls *ListNode, len int) *ListNode {
	//fmt.Print(ls.Val,",")
	var next *ListNode
	if ls.Next != nil{
		next= extListNode(ls.Next,len-1)
	} else if len>1 {
		ls.Next=  &ListNode{
			Val: 0,
		}
		next= extListNode(ls.Next,len-1)
	}
	//fmt.Print("\n")
	return &ListNode{Val:ls.Val,Next:next,}
}
func maxListNodeLen(l1 *ListNode, l2 *ListNode) int{
	l1Len:=getListNodeLen(l1)
	l2Len:=getListNodeLen(l2)
	if l1Len>l2Len{
		return l1Len
	}
	return l2Len
}
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {


	//listNodePrint(l1)
	//listNodePrint(l2)
	sum:= l1.Val+l2.Val
	fmt.Println(getListNodeLen(l2))
	if getListNodeLen(l1)!=getListNodeLen(l2) {
		maxLen := maxListNodeLen(l1,l2)
		fmt.Println("max len",maxLen)
		listNodePrint(extListNode(l2,maxLen))
		return AddTwoNumbers(extListNode(l1,maxLen), extListNode(l2,maxLen))
	}


	if l1.Next==nil && l2.Next==nil {
		var next *ListNode

		if sum>=10 {
			next= &ListNode{
				Val: sum/10,

			}
		}
		return &ListNode {
			Val: sum%10,
			Next: next,

		}
	}



	l1.Next.Val=l1.Next.Val+sum/10




	return &ListNode {
		Val: sum%10,
		Next: AddTwoNumbers(l1.Next,l2.Next),

	}



}
func listNodePrint(l *ListNode) {
	fmt.Print(l.Val,",")

	if l.Next!=nil{

		listNodePrint(l.Next)
	}else {
		fmt.Println()
	}


}
