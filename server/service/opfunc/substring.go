package opfunc

import "unicode/utf8"

func LongestSubstring(s string) string {
	//map with visited symbols, malloc max required memory to avoid overallocation
	visited := make(map[rune]int, len(s))
	//max len substring
	maxLen := 0
	//start position current substring
	start := 0
	//start position max substring
	substringStart := 0

	for i := 0; i < len(s); {
		//get char and size
		char, size := utf8.DecodeRuneInString(s[i:])
		//end position symb if size>1
		end := i + size
		//checking for visited char
		if index, ok := visited[char]; ok && index >= start {
			//last visited pos
			start = index + 1
		}
		//last visited pos
		visited[char] = i
		//check len substring
		if end-start > maxLen {
			maxLen = end - start
			substringStart = start
		}
		//next symbol
		i += size
	}

	return s[substringStart : substringStart+maxLen]

}
