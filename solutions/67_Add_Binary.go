package solutions


func getChar(s string,idx int) byte  {
	return s[idx]
}
func addBinary(a string, b string) string {
	if len(a)<len(b){
		return addBinary(b,a)
	}

	diff:=len(a)-len(b)
	res:=make([]byte,len(a)+1)
	for i:=len(a)-1;i>=diff ;i--  {

		res[i+1]= (a[i]+b[i-diff])-48

	}
	res[1:diff]=[]byte(a[0:diff])

	for i := len(a); i > 0; i-- {

		if i<=diff {
			res[i]=res[i]+a[i-1]

		}
		b:=res[i]-48
		v:=b/2
		mod:=b%2
		res[i]=mod+48
		res[i-1]=res[i-1]+v
	}
	if res[0]==0{
		return string(res[1:])
	}
	res[0]=res[0]+48
	return string(res)



}