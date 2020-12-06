/**
 * @Author root
 * @Description //TODO $
 * @Date $ $
 **/
package main

func moveZeroes(nums []int) {

	length := len(nums)
	tmp := 0
	j := 0
	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			continue
		}

		tmp = nums[j]
		nums[j] = nums[i]
		nums[i] = tmp
		j++

	}
}
