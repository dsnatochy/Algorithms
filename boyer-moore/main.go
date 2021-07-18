package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func print(last map[uint8]int){
	for key, val := range last {
		fmt.Println(string(key), val)
	}
}

func findBoyerMoore(text, pattern string) int {
	textLength := len(text)
	patternLength := len(pattern)
	if patternLength == 0 {
		return 0
	}
	last := make(map[uint8]int)
	for i := 0; i < textLength; i++ {
		last[text[i]] = -1
	}
	for k := 0; k < patternLength; k++ {
		// rightmost occurrence in pattern is last
		last[pattern[k]]=k
		//print(last)
	}

	textIdx := patternLength - 1    // an index into the text
	patternIdx := patternLength - 1 // an index into the pattern
	for textIdx < textLength {
		if text[textIdx] == pattern[patternIdx] { // a matching char
			if patternIdx == 0{ // entire pattern has been found
				return textIdx
			}
			textIdx--    // otherwise examine previous
			patternIdx-- // characters of text/pattern
		}else{
			textIdx += patternLength - min(patternIdx, 1+last[text[textIdx]]) // case analysis for jump step
			patternIdx = patternLength - 1                                    // restart at end of pattern
		}
	}
	return -1
}

func main() {
	idx := findBoyerMoore("charisma","is")
	fmt.Println(idx)
}
