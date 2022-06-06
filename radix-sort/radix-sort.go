package pushswap

import (
	"pushswap/utils"
)

func RadixSort(nums []int) string {
	res, opCount := "", 0
	nums = utils.UInts(nums)
	stack := []*utils.Stack{utils.NewStack(nums, len(nums)), utils.NewStack(make([]int, 0, len(nums)), 0)}
	maxNum := stack[0].Length - 1
	maxBits := 0
	for maxNum>>maxBits != 0 {
		maxBits++
	}
	for i := 0; i < maxBits; i++ {
		for j := 0; j < len(nums); j++ {
			num := stack[0].Nums[stack[0].Length-1]
			if (num>>i)&1 == 1 {
				Operation(stack, "ra")
				res += "ra\n"
			} else {
				Operation(stack, "pb")
				res += "pb\n"
			}
			opCount++
		}
		// putting into boxes is done
		for stack[1].Length != 0 {
			Operation(stack, "pa")
			res += "pa\n"
			opCount++
		}
	}
	//fmt.Println(stack[0].Nums, opCount)
	return res
}
