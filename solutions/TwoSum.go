package main

import "fmt"

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/


func twoSum(nums []int, target int) []int {

	dic:= make(map[int]int,len(nums))

	for i,n:=range nums{
		//key:=target-n

		if v,ok:=dic[target-n];ok{
			return []int{v,i}
		} else {
			dic[n]=i
		}



	}

	return []int{0,0}
}
func main()  {
	fmt.Println(twoSum([]int{2,5,5,11},10))
	//println()
	nums:=[]int{3,3}
	//t:=make(map[int]int)
	//t[1]=2
	//t[2]=3
	////delete(t,1)
	fmt.Println(nums[1:])
	fmt.Println(twoSum(nums,6))
}
