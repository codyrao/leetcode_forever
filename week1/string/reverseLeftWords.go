/**
 * @Author root
 * @Description //TODO $
 * @Date $ $
 **/
package main

func reverseLeftWords(s string, n int) string {

	if "" == s || n == 0 {
		return s
	}

	return s[n:] + s[:n]
}
