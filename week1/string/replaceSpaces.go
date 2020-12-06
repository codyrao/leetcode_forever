/**
 * @Author root
 * @Description //TODO $
 * @Date $ $
 **/
package main


func main() {
	replaceSpaces("Mr John Smith    ",13)
}
func replaceSpaces(S string, length int) string {

	chars := make([]rune, length*3)
	index := 0

	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			chars[index] = '%'
			index++
			chars[index] = '2'
			index++
			chars[index] = '0'
			index++
			continue
		}

		chars[index] = rune(S[i])
		index++
	}

	return string(chars[0:index])
}
