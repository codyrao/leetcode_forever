package main

func main() {

}

func reverseString(s []byte) {
	length := len(s)
	if length < 2 {
		return
	}

	left := 0
	right := length - 1
	var tmp byte
	for ; left < right; {
		tmp = s[left]
		s[left] = s[right]
		s[right] = tmp
		left++
		right--
	}

}
