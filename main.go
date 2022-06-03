package main

import (
	"log"
	"os"
	pushswap "pushswap/push-swap"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Enter numbers")
	}
	pushswap.RadixSort(args[0])
}
