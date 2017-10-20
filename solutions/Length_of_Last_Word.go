package solutions

import "strings"

func findWord(terms []string) string {
	l:=len(terms)
	if l==0{
		return ""
	}
	term:=terms[l-1]
	if len(term)==0{
		return findWord(terms[:l-1])
	} else {
		return term
	}

}
func lengthOfLastWord(s string) int {

	terms:=strings.Split(s," ")

	lastTerm:=findWord(terms)
	return len(lastTerm)
}