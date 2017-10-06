package main

import (
	"fmt"
	"strings"
)

func getCommonPrefix(s1 string, s2 string) string {
	if !strings.HasPrefix(s2,s1){
		return getCommonPrefix(s2,s1[:len(s1)-1])
	} else {
		return s1
	}
}
func longestCommonPrefix(strs []string) string {

	var prefix string


	size:=len(strs)
	if size>0{
		prefix=strs[0]
		tails:=strs[1:]
		fmt.Println(tails)
		for _,str:=range tails {
			prefix=getCommonPrefix(prefix,str)
			if prefix==""{
				break
			}
		}
	}

	return prefix
}
func main() {
	fmt.Println(longestCommonPrefix([]string{"abc"}))
}