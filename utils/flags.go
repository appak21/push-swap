package utils

import (
	"flag"
)

//checks whether the --rand flag passed or not
func FlagRandPassed() bool {
	var random = flag.Bool("rand", false, "Use -rand to be able to generate random numbers")
	flag.Parse()
	return *random
}
