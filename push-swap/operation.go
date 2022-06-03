package pushswap

type f func(stack []*Stack)

var cmd = map[string]f{
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

func operation(stack []*Stack, instr string) {
	cmd[instr](stack)
}
