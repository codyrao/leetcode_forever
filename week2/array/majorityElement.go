/**
 * @Author root
 * @Description //TODO $
 * @Date $ $
 **/
package main

func majorityElement(nums []int) int {

	hashMap := make(map[int]int)
	length := len(nums)
	for _, num := range nums {
		hashMap[num]++
		if hashMap[num] > (length/2) {
			return num
		}
	}
	return -1
}
