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
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}
	instructions, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("here 0")
		fmt.Println("Error")
		return
	}
	nums, err := utils.UniqInts(args[0])
	if err != nil {
		fmt.Println("here 1")
		fmt.Println("Error")
		return
	}
	stack := []*utils.Stack{utils.NewStack(nums, len(nums)), utils.NewStack(make([]int, 0, len(nums)), 0)}
	operations := strings.Split(string(instructions), "\n")
	l := len(operations) - 1
	if operations[l-1] == "" {
		l--
	}
	operations = operations[:l]
	for _, v := range operations {
		if f, ok := pushswap.Cmd[v]; ok {
			f(stack)
		} else {
			fmt.Println("Error")
			return
		}
	}
	//do I need to compare my result with the result of sort.Ints?
	//if so, I need to modify stack in push-swap pr.
	//also as stack pull actually doesn't delete,
	//stack b isn't empty which violates the requirement
	fmt.Println("Stack a", stack[0].Nums)
	fmt.Println("Stack b", stack[1].Nums)
}
