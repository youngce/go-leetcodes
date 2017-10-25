package solutions


func trailingZeroes(n int) int {

	v:=n/5
	if v==0{
		return 0
	}

	return n/5+trailingZeroes(n/5)





}
