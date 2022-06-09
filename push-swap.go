package main

import (
	"fmt"
	"os"
	pushswap "pushswap/instructions"
	pushswap "pushswap/radix-sort"
	"pushswap/utils"
	"sort"
)

func main() {
	var (
		nums []int
		err  error
	)
	if utils.FlagRandPassed() {
		nums, err = utils.GenerateRandomInts()
		if err != nil {
			fmt.Println("Error")
			return
		}
	} else {
		args := os.Args[1:]
		if len(args) != 1 {
			fmt.Println("Enter numbers or use -rand flag to be able to use randomly generated numbers")
			return
		}
		nums, err = utils.UniqInts(args[0])
		if err != nil {
			fmt.Println("Error")
			return
		}
	}
	if sort.IntsAreSorted(nums) {
		return
	}
	file, err := os.Create("result/push-swap-result.txt")
	if err != nil {
		fmt.Println("the named file wasn't created")
		return
	}
	defer file.Close()
	numsToStr := fmt.Sprint(nums)
	file.WriteString("\"" + numsToStr[1:len(numsToStr)-1] + "\"\n")
	instructions := ""
	if len(nums) <= 5 {
		instructions = pushswap.SmallSort(nums)
	} else {
		instructions = pushswap.RadixSort(nums)
	}
	fmt.Print(instructions)
	file.WriteString(instructions)
}
