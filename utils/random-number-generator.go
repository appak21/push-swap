package utils

import (
	"fmt"
	"math/rand"
)

func GenerateRandomInts() ([]int, error) {
	fmt.Println("Numbers will be randomly generated in the range [min,max).\nEnter the minimum, maximum values and the length of numbers in order.\nmin max length")
	input := InputInts()
	if len(input) != 3 {
		return nil, fmt.Errorf("please, check your input for correctness")
	}
	return RandIntsUniq(input[0], input[1], input[2])
}

//generates unique random integers
func RandIntsUniq(min, max, length int) ([]int, error) {
	if min > max {
		return nil, fmt.Errorf("minimum number can't be larger than %v", max)
	}
	if max-min < length {
		return nil, fmt.Errorf("numbers can't be duplicated, try length >= max-min")
	}
	nums := make([]int, length)
	randomList := rand.Perm(max - min + 1)
	for i := 0; i < length; i++ {
		nums[i] = randomList[i] + min
	}
	return nums, nil
}
