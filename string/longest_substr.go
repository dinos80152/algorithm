package string

import "bytes"

func LengthOfLongestSubstring(s string) int {
	var subStr, topStr []byte
	for i := 0; i < len(s); i++ {
		char := s[i]
		index := bytes.IndexByte(subStr, char)
		if index == -1 {
			subStr = append(subStr, char)
		} else {
			subStr = append(subStr[index+1:], char)
		}
		if len(subStr) > len(topStr) {
			topStr = subStr
		}
	}
	return len(topStr)
}
