package main

func main() {

}

var numMap = make(map[string]int)

func init() {
	numMap["I"] = 1
	numMap["V"] = 5
	numMap["X"] = 10
	numMap["L"] = 50
	numMap["C"] = 100
	numMap["D"] = 500
	numMap["M"] = 1000
	numMap["IV"] = 4
	numMap["IX"] = 9
	numMap["XL"] = 40
	numMap["XC"] = 90
	numMap["CD"] = 400
	numMap["CM"] = 900

}

func romanToInt(s string) int {

	res := 0
	length := len(s)
	for i := 0; i < length; {
		if i+2 <= length {
			val, ok := numMap[s[i:i+2]]
			if ok {
				res += val
				i += 2
				continue
			}

		}

		val, _ := numMap[string(s[i])]

		res += val
		i++
	}

	return res
}
