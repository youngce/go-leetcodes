package solutions



func plusOne(digits []int) []int {




	l:=len(digits)
	if l==0{
		return []int{1}
	}
	p:=digits[l-1]+1
	v:=p/10
	mod:=p%10

	if v>0 {
		return append(plusOne(digits[:l-1]),mod)
	} else {
		digits[l-1]=mod
		return digits
	}
	//for i := l - 1; i < 0;i--  {
	//	p:=digits[i]+1
	//	v:=p/10
	//	mod:=digits[i]%%10
	//}
}