package solutions

import "unicode"
//import (
//	"strings"
//	"unicode"
//)

func isLetterOrDigit(r rune) bool {
	return unicode.IsDigit(r)||unicode.IsLetter(r)
}
func isPalindrome2(s string) bool {

	var headIdx, tailIdx int

	chars:=[]rune(s)
	tailIdx =len(chars)-1

	for headIdx <= tailIdx {
		head:=chars[headIdx]
		tail:=chars[tailIdx]

		if !isLetterOrDigit(head)||!isLetterOrDigit(tail){
			if !isLetterOrDigit(tail){
				tailIdx--

			}
			if !isLetterOrDigit(head){
				headIdx++

			}
			continue
		}
		if unicode.ToLower(head)!=unicode.ToLower(tail){
			return false
		}
		tailIdx--
		headIdx++
	}

	return true
}