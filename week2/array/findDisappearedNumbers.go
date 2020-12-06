package main

func main() {

}

func findDisappearedNumbers(nums []int) []int {

	length := len(nums)
	for i := 0; i < length; i++ {
		index := nums[i]
		if index < 0 {
			index = -index
		}
		index--
		value := nums[index]
		if value > 0 {
			nums[index] = -value
		}
	}

	res := make([]int, 0)
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}

	return res
}
