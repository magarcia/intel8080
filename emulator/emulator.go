package emulator

type ConditionCodes struct {
	z   uint8
	s   uint8
	p   uint8
	cy  uint8
	ac  uint8
	pad uint8
	//uint8_t    z:1
	//uint8_t    s:1;
	//uint8_t    p:1;
	//uint8_t    cy:1;
	//uint8_t    ac:1;
	//uint8_t    pad:3;
}

type State8080 struct {
	a          uint8
	b          uint8
	c          uint8
	d          uint8
	e          uint8
	h          uint8
	l          uint8
	sp         uint16
	pc         uint16
	memory     *[]uint8
	cc         ConditionCodes
	int_enable uint8
}

func UnimplementedInstruction(state *State8080) {
	panic("Error: Unimplemented instruction\n")
}

//func Emulate8080Op(state *State8080) int {
//	var opcode *uint8
//	opcode = state.memory[state.pc]
//}
