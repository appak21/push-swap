package main

import (
	"fmt"
	"io/ioutil"
	"os"
	pushswap "pushswap/radix-sort"
	"pushswap/utils"
	"strings"
)

func main() {
	instructions, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("KO")
		return
	}
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Enter numbers to sort")
		return
	}
	nums, err := utils.UniqInts(args[0])
	if err != nil {
		fmt.Println("Error")
		return
	}
	stack := []*utils.Stack{utils.NewStack(nums, len(nums)), utils.NewStack(make([]int, 0, len(nums)), 0)}
	operations := strings.Split(string(instructions), "\n")
	for i := 0; i < len(operations)-1; i++ {
		pushswap.Operation(stack, operations[i])
	}
	fmt.Println(stack[0].Nums)
}
