package solutions

func strStr(haystack string, needle string) int {


	needleLen:=len(needle)
	if needleLen==0{
		return 0
	}
	hsLen:=len(haystack)
	for i:=0;i<(hsLen-needleLen)+1 ;i++  {
		if needle==haystack[i:(i+needleLen)]{
			return i
		}
	}

	return -1
}
