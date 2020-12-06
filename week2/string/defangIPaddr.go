package main

func main() {

}

func defangIPaddr(address string) string {
	if "" == address {
		return ""
	}
	length := len(address)
	res := make([]byte, length+6)
	index := 0
	for i := 0; i < length; i++ {

		if address[i] == '.' {
			res[index] = '['
			index++
			res[index] = '.'
			index++
			res[index] = ']'
			index++
			continue
		}
		res[index] = address[i]
		index++
	}

	return string(res)
}
