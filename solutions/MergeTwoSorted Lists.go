package solutions


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	if l1.Val>l2.Val{
		l2.Next=mergeTwoLists(l1,l2.Next)
		return l2
	} else {
		l1.Next=mergeTwoLists(l1.Next,l2)
		return l1
	}

	return nil


	//ListNode{}

}