package pushswap

import (
	"pushswap/utils"
)

//RadixSort returns a string of instructions.
//Its result is correct if there's no negative number.
func RadixSort(nums []int) string {
	res := ""
	stack := []*utils.Stack{utils.NewStack(nums), utils.NewStack(make([]int, 0, len(nums)))}
	maxNum := len(stack[0].Nums) - 1
	maxBits := 0
	for maxNum>>maxBits != 0 {
		maxBits++
	}
	for i := 0; i < maxBits; i++ {
		for j := 0; j < len(nums); j++ {
			num := stack[0].Nums[len(stack[0].Nums)-1]
			if (num>>i)&1 == 1 {
				Operation(stack, "ra")
				res += "ra\n"
			} else {
				Operation(stack, "pb")
				res += "pb\n"
			}
		}
		for len(stack[1].Nums) != 0 {
			Operation(stack, "pa")
			res += "pa\n"
		}
	}
	return res
}
