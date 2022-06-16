package pushswap

import (
	"pushswap/utils"
	"sort"
)

func SmallSort(nums []int) string {
	res := ""
	stack := []*utils.Stack{utils.NewStack(nums), utils.NewStack(make([]int, 0, len(nums)))}
	res = SmallInstructions(nums, stack)
	return res
}

func SmallInstructions(nums []int, s []*utils.Stack) string {
	half := len(nums) / 2
	res := ""
	for {
		a, b := s[0].Nums, s[1].Nums
		if isASorted(a) {
			break
		}
		op := match(aStack(a, half), a, b)
		Operation(s, op)
		res += op + "\n"
	}
	for {
		a, b := s[0].Nums, s[1].Nums
		if len(a) == len(nums) && isASorted(a) {
			break
		}
		ins := bStack(b)
		Operation(s, ins)
		res += ins + "\n"
	}
	return res
}

//sa, ra, rra, pb
func aStack(a []int, half int) string {
	swap := saIdx(a)
	mid, i := len(a)/2, len(a)-1
	if a[i] < half {
		return "pb"
	}
	if swap > 0 {
		if inOrder(a[i-1], a[i]) {
			return "sa"
		} else if swap < mid {
			return "rra"
		}
	} else {
		minNumIdx := smallest(a)
		if minNumIdx == i {
			return "pb"
		} else if minNumIdx < mid || inOrder(a[i], a[0]) {
			return "rra"
		}
	}
	return "ra"
}

//sb, rb, rrb, pa
func bStack(bstack []int) string {
	b := make([]int, len(bstack))
	copy(b, bstack)
	utils.UInts(b)
	swap := sbIdx(b)
	mid, j := len(b)/2, len(b)-1
	if swap > 0 {
		if inOrder(b[j], b[j-1]) {
			return "sb"
		} else if swap < mid {
			return "rrb"
		}
	} else {
		maxNumIdx := biggest(b)
		if maxNumIdx == j {
			return "pa"
		} else if maxNumIdx < mid || inOrder(b[0], b[j]) {
			return "rrb"
		}
	}
	return "rb"
}

func match(ins string, a, b []int) string {
	if len(b) < 2 || isBSorted(b) {
		return ins
	}
	bins := bStack(b)
	switch ins { //ins is an instruction from aStack(a)
	case "pb":
		if bins == "sb" {
			if t := len(a) - saIdx(a); t < 3 {
				if t == 1 {
					return "ss"
				}
				return "ra"
			} else {
				return "sb"
			}
		}
	case "sa":
		if t := len(b) - sbIdx(b); t < 3 {
			if bins == "sb" {
				return "ss"
			}
			return "rb"
		}
	case "ra":
		if bins == "rb" || bins == "sb" && len(b) == 2 {
			return "rr"
		}
	case "rra":
		if bins == "rrb" || bins == "sb" && len(b) == 2 {
			return "rrr"
		}
	}
	return ins
}

func smallest(nums []int) int {
	min, idx := 0x3f3f3f3f, -1
	for i := 0; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
			idx = i
		}
	}
	return idx
}

func biggest(nums []int) int {
	max, idx := -1, -1
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
			idx = i
		}
	}
	return idx
}

func saIdx(nums []int) int {
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i]-nums[i-1] == 1 {
			return i
		}
	}
	return -1
}

func sbIdx(nums []int) int {
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i-1]-nums[i] == 1 {
			return i
		}
	}
	return -1
}

//true if a+1 = b
func inOrder(a, b int) bool {
	return b-a == 1
}

func isASorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]-nums[i+1] != 1 {
			return false
		}
	}
	return true
}

func isBSorted(nums []int) bool {
	return sort.IntsAreSorted(nums)
}

//"4 3 2 1 0" //11
//"2 3 0 1 4" //7
//"4 3 2 1 0 5" //11
//"3 0 4 1 2" //4
//"2 1 3 6 5 8" //must return 9 ins //last step 7

//for 25 => 87, 89, 108
//for 10 => 21, 22, 27
//for 100 => 800-1000
//for 500 => 6784
//for 1000 => 15068
