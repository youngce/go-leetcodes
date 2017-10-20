package solutions

func deleteDuplicates(head *ListNode) *ListNode {

	if head==nil ||head.Next==nil{
		return head
	}
	if head.Val==head.Next.Val{
		return deleteDuplicates(head.Next)
	} else{
		head.Next=deleteDuplicates(head.Next)
		return head
	}

}
