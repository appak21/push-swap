package pushswap

import "pushswap/utils"

func BigSort(nums []int) string {
	res := ""
	stack := []*utils.Stack{utils.NewStack(nums), utils.NewStack(make([]int, 0, len(nums)))}
	res = BigInstructions(nums, stack)
	return res
}

func BigInstructions(nums []int, s []*utils.Stack) string {
	//keyNum := 25
	res := ""
	return res
}
