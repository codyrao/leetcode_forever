package main

func main() {

}

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	if numRows == 0 {
		return res
	}
	res = append(res, []int{1})
	for i := 1; i < numRows; i++ {
		row := make([]int,0,i+1)

		row = append(row, 1)
		for j := 1; j < i; j++ {
			val := res[i-1][j-1] + res[i-1][j]
			row = append(row, val)
		}

		row = append(row, 1)

		res = append(res, row)
	}

	return res
}
