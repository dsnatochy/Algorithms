package main

import "fmt"

func computeFailKMP(pattern string) []int{
	m := len(pattern)
	fail := make([]int, m) // by default, all overlaps are zero
	j := 1
	k := 0															   //j:3, k:1
	for j < m { // compute fail[j] during this pass, if nonzero          3 < 12
		if pattern[j] == pattern[k] { // k+1 characters match thus far   l == m ?
			fail[j] = k + 1                                           // fail[2] = 0+1
			j++                                                       // j=3
			k++ 													  // k=1
		}else if k > 0 {  // k follows a matching prefix				 1 > 0 ?
				k = fail[k-1]										  // k=fail[0]=0
		}else {
			j++  // no match found starting at j					     1++ (j=2)
		}
	}
	return fail
}

// time complexity O(m+n)
func findKMP(text, pattern string) int{
	n := len(text)
	m := len (pattern)
	if m == 0{
		return 0
	}
	fail := computeFailKMP(pattern)
	j := 0
	k := 0
	for j < n {
		if text[j] == pattern[k] {
			if (k == m - 1){
				return j - m + 1
			}
			j++
			k++
		}else if k > 0 {
				k = fail[k-1]
		}else {
			j++
		}
	}
	return -1
}

func main(){
	text := "atcamalgamaamalgamationstart"
	pattern := "amalgamation"
	fmt.Println("pattern found at position", findKMP(text, pattern))
}