package main

import (
	"go/ast"
	"strings"
)

/*
Implement a MapSum class with insert, and sum methods.

For the method insert, you'll be given a pair of (string, integer). The string represents the key and the integer represents the value. If the key already existed, then the original key-value pair will be overridden to the new one.

For the method sum, you'll be given a string representing the prefix, and you need to return the sum of all the pairs' value whose key starts with the prefix.

Example 1:

Input: insert("apple", 3), Output: Null
Input: sum("ap"), Output: 3
Input: insert("app", 2), Output: Null
Input: sum("ap"), Output: 5
*/
type MapSum struct {
	Kv map[string]int

}


/** Initialize your data structure here. */
func Constructor() MapSum {

	return MapSum{

		Kv:make(map[string]int),
	}
}


func (this *MapSum) Insert(key string, val int)  {


	this.Kv[key]=val
}


func (this *MapSum) Sum(prefix string) int {

	var sum=0
	for k,v:=range this.Kv{

		if strings.HasPrefix(k,prefix){
			sum+= v
		}
	}
	return sum
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
