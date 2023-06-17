package test

import (
	"server/service/opfunc"
	"testing"
)

func TestSubString(t *testing.T) {
	cases := []string{"abcabcbb", "bbbbb", "pwwkew"}
	answer := []string{"abc", "b", "wke"}
	for i, val := range cases {
		result := opfunc.LongestSubstring(val)

		if result != answer[i] {
			t.Errorf("Test failed %v: result=%v needed = %v", val, result, answer[i])
		}
	}

}
