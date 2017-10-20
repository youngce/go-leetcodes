package solutions

import "strings"

func getChar(s string,idx int) byte  {
	return s[idx]
}
func addBinary(a string, b string) string {
	idx_a:=len(a)-1
	idx_b:=len(b)-1
	var s string
	var carry,sum byte

	for idx_a>=0 || idx_b>=0 {

		sum=carry
		if idx_a>=0{
			sum+=a[idx_a]-'0'
			idx_a--
		}
		if idx_b>=0{
			sum+=b[idx_b]-'0'
			idx_b--
		}
		carry=sum/2
		s=string(sum%2+'0')+s



	}
	if carry!=0{
		s=string(sum%2+48)+s
	}

	return s



}