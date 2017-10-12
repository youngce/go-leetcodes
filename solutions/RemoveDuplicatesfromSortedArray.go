package solutions


//Given a sorted array, remove the duplicates in place such that each element appear only once and return the new length.
//
//Do not allocate extra space for another array, you must do this in place with constant memory.
//
//For example,
//Given input array nums = [1,1,2],
//
//Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively. It doesn't matter what you leave beyond the new length.

func removeDuplicates(nums []int) int {
	dic:=make(map[int]int)

	var size int
	for _,n:=range nums{
		if v,ok:=dic[n];ok{
			dic[n]=v+1

		} else {
			dic[n]=1
			size+=1
		}
	}
	nums=make([]int,size)
	var i int
	for k,_:=range dic{

		nums[i]=k
		i++
	}
	return size


}