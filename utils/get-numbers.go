package utils

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//UniqInts returns slice of unique integers from the given string s.
//Only distinct numbers separated with single space are allowed
func UniqInts(s string) ([]int, error) {
	input := strings.Split(s, " ")
	nums := make([]int, len(input))
	existing := make(map[int]bool, len(nums))
	for i, n := range input {
		num, err := strconv.Atoi(n)
		if err != nil || existing[num] {
			return nil, fmt.Errorf("invalid number")
		}
		existing[num] = true
		nums[i] = num
	}
	return nums, nil
}

//white spaces and non-digits are allowed
func GetInts(s string) []int {
	fields := strings.Fields(s)
	nums := make([]int, 0, len(fields))
	for _, f := range fields {
		n, err := strconv.Atoi(f)
		if err == nil {
			nums = append(nums, n)
		}
	}
	return nums
}

//InputInts takes input from the User and returns numbers
func InputInts() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return GetInts(scanner.Text())
}

//UInts changes nums to unsigned integers
//by replacing them with their indexes when they're sorted
func UInts(nums []int) {
	copiedNums := make([]int, len(nums))
	copy(copiedNums, nums)
	sort.Ints(copiedNums)
	for i, n := range nums {
		nums[i] = BinarySearch(copiedNums, n)
	}
}
