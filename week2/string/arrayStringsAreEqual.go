package main

import "strings"

func main() {

}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	if nil == word1 && nil == word2 {
		return true
	}

	if nil == word1 || nil == word2 {
		return false
	}
	var compare1 strings.Builder
	for _, s1 := range word1 {
		compare1.WriteString(s1)
	}
	var compare2 strings.Builder
	for _, s2 := range word2 {
		compare2.WriteString(s2)
	}

	return compare1.String() == compare2.String()
}
