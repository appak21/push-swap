package main

import (
	"fmt"
	"os"
	pushswap "pushswap/radix-sort"
	"pushswap/utils"
)

func main() {
	var nums []int
	var err error
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
	fmt.Print(pushswap.RadixSort(nums))
}
