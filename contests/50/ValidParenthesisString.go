package main

import (
	"fmt"

)

/*
Given a string containing only three types of characters: '(', ')' and '*', write a function to check whether this string is valid. We define the validity of a string by these rules:

1. Any left parenthesis '(' must have a corresponding right parenthesis ')'.
2. Any right parenthesis ')' must have a corresponding left parenthesis '('.
3. Left parenthesis '(' must go before the corresponding right parenthesis ')'.
4. '*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.
5. An empty string is also valid.
Example 1:

Input: "()"
Output: True

Example 2:

Input: "(*)"
Output: True

Example 3:

Input: "(*))"
Output: True

Note:

    The string size will be in the range [1, 100].

*/
type condition func(rune) bool
func runesFilter(rs []rune,fn condition) []rune {

	var res []rune
	for i:=0;i<len(rs) ;i++  {

		if fn(rs[i]) {
			res=append(res,rs[i])
		}
	}
	return res

}
func normalize(b byte) int{
	switch b {
	case '(':
		return 1
	case ')':
		return -1
	case '*':
		return 0
	default:
		return -100

	}
}

func checkValidString(s string) bool {


	selected:=string(runesFilter([]rune(s), func(i rune) bool {

		return i=='('||i=='*'||i==')'
	}))
	size:=len(selected)

	//var nb []int
	var zeroSize int
	var sum int
	for i := 0; i < size; i++ {

		v:=normalize(selected[i])
		//nb=append(nb,normalize(selected[i]))

		if sum>0&&v==0{
			sum-=1
		} else {
			sum+=v
		}

		if sum<0{
			return false
		}
		if v==0{
			zeroSize+=1
		}


	}
	if sum==zeroSize{
		return true
	}
	fmt.Println(sum,zeroSize)
	//return true
	return true
}

func main() {
	//str:=string(runesFilter([]rune("(12*3)"), func(i rune) bool {
	//
	//	return i=='('||i=='*'||i==')'
	//}))

	//fmt.Println('(',')','*')

	fmt.Println(checkValidString("(*))"))
}