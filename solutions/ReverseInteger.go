package main

import (
	"strconv"

	"fmt"
)

func reverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(string(runes))
	return string(runes)
}
func reverse(x int) int {
	var sign string
	sign=""

	var num int
	num=x

	if x<0{
		sign="-"
		num=-x
	}





	if i,err:=strconv.Atoi(sign+reverseStr(strconv.Itoa(num)));err==nil{
		//-2,147,483,648 to 2,147,483,647
		if i>2147483647||i< -2147483648{
			return 0
		}
		return i
	} else {
		fmt.Println(err.Error())
	}

	return 0
}

func main() {
	reverse(1534236469)

}