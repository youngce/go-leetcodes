package solutions

import "unicode"

func ifModIsZero(mod int) int {
	if mod==0{
		return 26
	}
	return mod
}
func convertToTitle(n int) string {
	var str string
	for n>0{
		mod:=n%26
		if mod==0{
			mod=26
		}
		str=string(mod+'A'-1)+str
		n=(n-1)/26
	}
	return str

}