package main

import "fmt"

func isLeft(r rune) bool  {
	return r =='('|| r =='['|| r =='{'
}
func oppose(r rune) rune  {
	switch r {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	default:
		return 0
	}
}
func isValid(s string) bool {
	stack:=[]rune{}
	var lastIndex int
	lastIndex=-1
	for _,r:=range s{

		if isLeft(r){
			lastIndex+=1
			stack=append(stack,r)
			//fmt.Println(stack,lastIndex)
		}else {
			if lastIndex>=0&&r==oppose(stack[lastIndex]){

				//fmt.Println(lastIndex)
				stack=stack[:lastIndex]
				lastIndex-=1
			} else {
				return false
			}
		}
	}
	if len(stack)>0{
		return false
	}
	return true

}
func main() {
	fmt.Println(isValid("()"))
}