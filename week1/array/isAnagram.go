/**
 * @Author root
 * @Description //TODO $
 * @Date $ $
 **/
package main

func main() {

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := make([]int, 26)
	strLen := len(s)
	for i := 0; i < strLen; i++ {
		hash[s[i]-'a']++
		hash[t[i]-'a']--
	}
	for _, v := range hash {
		if v != 0 {
			return false
		}
	}

	return true

}
