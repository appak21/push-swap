package pushswap

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func RadixSort(arg string) {
	res := ""
	nums := parseNums(arg)
	stack := []*Stack{NewStack(nums, len(nums)), NewStack(make([]int, 0, len(nums)), 0)}
	maxNum := stack[0].length - 1
	maxBits := 0
	for maxNum>>maxBits != 0 {
		maxBits++
	}
	for i := 0; i < maxBits; i++ {
		for j := 0; j < len(nums); j++ {
			num := stack[0].Nums[stack[0].length-1]
			if (num>>i)&1 == 1 {
				operation(stack, "ra")
				res += "ra\n"
			} else {
				operation(stack, "pb")
				res += "pb\n"
			}
		}
		// putting into boxes is done
		for stack[1].length != 0 {
			operation(stack, "pa")
		}
	}
	//fmt.Println(stack[0], stack[1])
	fmt.Print(res)
}

func parseNums(arg string) []int {
	input := strings.Split(arg, " ")
	nums := make([]int, len(input))
	for i, n := range input {
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal("Error")
		}
		nums[i] = num
	}
	copiedNums := make([]int, len(nums))
	copy(copiedNums, nums)
	sort.Ints(copiedNums)
	for i, n := range nums {
		nums[i] = binarySearch(copiedNums, n)
	}
	return nums
}

func binarySearch(nums []int, target int) int { //returns the index of the target
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
