/**
 * @Author root
 * @Description //TODO $
 * @Date $ $
 **/
package main

func main() {

}


func runningSum(nums []int) []int {
	if nil ==nums {
		return nil
	}
	count:=0
	for i, num := range nums {
		count+=num
		nums[i]=count
	}

	return nums
}