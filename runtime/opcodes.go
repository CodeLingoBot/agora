package runtime

// The opcode takes 8 bytes, leaving 256 possible codes.
type Opcode byte

const (
	OP_RET Opcode = iota
	OP_PUSH
	OP_POP
)

var (
	OpNames = [...]string{
		OP_RET:  "RET",
		OP_PUSH: "PUSH",
		OP_POP:  "POP",
	}

	OpLookup = map[string]Opcode{
		"RET":  OP_RET,
		"PUSH": OP_PUSH,
		"POP":  OP_POP,
	}
)

func (ø Opcode) String() string {
	return OpNames[ø]
}