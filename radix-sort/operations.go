package pushswap

import "pushswap/utils"

type fn func(stack []*utils.Stack)

var cmd = map[string]fn{
	"pa":  pa,
	"pb":  pb,
	"sa":  sa,
	"sb":  sb,
	"ss":  ss,
	"ra":  ra,
	"rb":  rb,
	"rr":  rr,
	"rra": rra,
	"rrb": rrb,
	"rrr": rrr,
}

func Operation(stack []*utils.Stack, instr string) {
	cmd[instr](stack)
}
