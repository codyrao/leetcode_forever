/**
 * @Author root
 * @Description //TODO $
 * @Date $ $
 **/
package main

func CheckPermutation(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false

	}

	hash := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		hash[s1[i]-'a']++
		hash[s2[i]-'a']--
	}

	for _, v := range hash {
		if v != 0 {
			return false
		}
	}

	return true
}
