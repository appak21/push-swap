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
		fmt.Println("Error")
		return
	}
	nums, err := utils.UniqInts(args[0])
	if err != nil {
		fmt.Println("Error")
		return
	}
	stack := []*utils.Stack{utils.NewStack(nums), utils.NewStack(make([]int, 0, len(nums)))}
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
	fmt.Println("Stack a", stack[0].Nums)
	//fmt.Println("Stack b", stack[1].Nums)
}
