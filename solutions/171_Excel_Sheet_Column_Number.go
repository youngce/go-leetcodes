package solutions

import (
	"math"

)

func titleToNumber(s string) int {

	var num int
	l:=len(s)
	for i:=l-1;i>=0;i--{

		v:=s[i]-'A'+1
		if s[i]=='Z'{
			v=26
		}


		num+=int(v)*int(math.Pow(26.0,float64(l-1-i)))
	}
	return num
}