package main

import (
	"fmt"
	"strings"
)

/*Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.*/

func main(){
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}


func lengthOfLongestSubstring(s string) int {
	if len(s)==0{
		return 0
	}
	var res []string = make([]string,0)
	res = append(res, string(s[0]))
	var max []string = make([]string,0)
	max = res
	/*var temp []string = make([]string,0)*/

	for i := 1; i<len(s); i++ {
		if pos := index(string(s[i]), max); -1 == pos{
			max = append(max, string(s[i]))
		}else {
			max = append(max, string(s[i]))
			max = max[ pos + 1:]
		}
/*
		if len(temp) > len(max){
			max = temp
		}else {
			max = max
		}*/

		if len(res) < len(max){
			res = max
		}else {
			res = res
		}
	}
	return len(res)
}

func index(c string, s []string) int{
	for i := 0; i<len(s); i++ {
		if 0 == strings.Compare(c, s[i]){
			return i
		}
	}
	return -1
}

