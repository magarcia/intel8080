package cmd

import (
	"fmt"
	"strings"
)

type Instruction struct {
	Code   string
	Params []string
}

func getHexFromByte(code []byte, pc int, length int) string {
	var out string
	i := pc + length
	for i > pc {
		out += fmt.Sprintf("%02x", code[i])
		i -= 1
	}
	return "$" + out
}

func getInstruction(buffer *[]byte, pc int) Instruction {
	code := *buffer

	var instruction Instruction

	switch code[pc] {
	case 0x00:
		instruction.Code = "NOP"
	case 0x01:
		instruction.Code = "LXI"
		instruction.Params = append(instruction.Params, "B")
		instruction.Params = append(instruction.Params, "#"+getHexFromByte(code, pc, 2))
	case 0x02:
		instruction.Code = "STAX"
		instruction.Params = append(instruction.Params, "B")
	case 0x03:
		instruction.Code = "INX"
		instruction.Params = append(instruction.Params, "B")
	case 0x04:
		instruction.Code = "INR"
		instruction.Params = append(instruction.Params, "B")
	case 0x05:
		instruction.Code = "DCR"
		instruction.Params = append(instruction.Params, "B")
	case 0x06:
		instruction.Code = "MVI"
		instruction.Params = append(instruction.Params, "B")
		instruction.Params = append(instruction.Params, "#"+getHexFromByte(code, pc, 1))
	case 0x07:
		instruction.Code = "RLC"
	case 0x08:
		instruction.Code = "NOP"
	case 0x09:
		instruction.Code = "DAD"
		instruction.Params = append(instruction.Params, "B")
	case 0x0a:
		instruction.Code = "LDAX"
		instruction.Params = append(instruction.Params, "B")
	case 0x0b:
		instruction.Code = "DCX"
		instruction.Params = append(instruction.Params, "B")
	case 0x0c:
		instruction.Code = "INR"
		instruction.Params = append(instruction.Params, "C")
	case 0x0d:
		instruction.Code = "DCR"
		instruction.Params = append(instruction.Params, "C")
	case 0x0e:
		instruction.Code = "MVI"
		instruction.Params = append(instruction.Params, "C")
		instruction.Params = append(instruction.Params, "#"+getHexFromByte(code, pc, 1))
	case 0x0f:
		instruction.Code = "RRC"
	case 0x10:
		instruction.Code = "NOP"
	case 0x11:
		instruction.Code = "LXI"
		instruction.Params = append(instruction.Params, "D")
		instruction.Params = append(instruction.Params, "#"+getHexFromByte(code, pc, 2))
	case 0x12:
		instruction.Code = "STAX"
		instruction.Params = append(instruction.Params, "D")
	case 0x13:
		instruction.Code = "INX"
		instruction.Params = append(instruction.Params, "D")
	case 0x14:
		instruction.Code = "INR"
		instruction.Params = append(instruction.Params, "D")
	case 0x15:
		instruction.Code = "DCR"
		instruction.Params = append(instruction.Params, "D")
	case 0x16:
		instruction.Code = "MVI"
		instruction.Params = append(instruction.Params, "D")
		instruction.Params = append(instruction.Params, "#"+getHexFromByte(code, pc, 1))
	case 0x17:
		instruction.Code = "RAL"
	case 0x18:
		instruction.Code = "NOP"
	case 0x19:
		instruction.Code = "DAD"
		instruction.Params = append(instruction.Params, "D")
	case 0x1a:
		instruction.Code = "LDAX"
		instruction.Params = append(instruction.Params, "D")
	case 0x1b:
		instruction.Code = "DCX"
		instruction.Params = append(instruction.Params, "D")
	case 0x1c:
		instruction.Code = "INR"
		instruction.Params = append(instruction.Params, "E")
	case 0x1d:
		instruction.Code = "DCR"
		instruction.Params = append(instruction.Params, "E")
	case 0x1e:
		instruction.Code = "MVI"
		instruction.Params = append(instruction.Params, "E")
		instruction.Params = append(instruction.Params, "#"+getHexFromByte(code, pc, 1))
	case 0x1f:
		instruction.Code = "RAR"
	case 0x20:
		instruction.Code = "RIM"
	case 0x21:
		instruction.Code = "LXI"
		instruction.Params = append(instruction.Params, "H")
		instruction.Params = append(instruction.Params, "#"+getHexFromByte(code, pc, 2))
	case 0x22:
		instruction.Code = "SHLD"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0x23:
		instruction.Code = "INX"
		instruction.Params = append(instruction.Params, "H")
	case 0x24:
		instruction.Code = "INR"
		instruction.Params = append(instruction.Params, "H")
	case 0x25:
		instruction.Code = "DCR"
		instruction.Params = append(instruction.Params, "H")
	case 0x26:
		instruction.Code = "MVI"
		instruction.Params = append(instruction.Params, "H")
		instruction.Params = append(instruction.Params, "#"+getHexFromByte(code, pc, 1))
	case 0x27:
		instruction.Code = "DAA"
	case 0x28:
		instruction.Code = "NOP"
	case 0x29:
		instruction.Code = "DAD"
		instruction.Params = append(instruction.Params, "H")
	case 0x2a:
		instruction.Code = "LHLD"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0x2b:
		instruction.Code = "DCX"
		instruction.Params = append(instruction.Params, "H")
	case 0x2c:
		instruction.Code = "INR"
		instruction.Params = append(instruction.Params, "L")
	case 0x2d:
		instruction.Code = "DCR"
		instruction.Params = append(instruction.Params, "L")
	case 0x2e:
		instruction.Code = "MVI"
		instruction.Params = append(instruction.Params, "L")
		instruction.Params = append(instruction.Params, "#"+getHexFromByte(code, pc, 1))
	case 0x2f:
		instruction.Code = "CMA"
	case 0x30:
		instruction.Code = "SIM"
	case 0x31:
		instruction.Code = "LXI"
		instruction.Params = append(instruction.Params, "SP")
		instruction.Params = append(instruction.Params, "#"+getHexFromByte(code, pc, 2))
	case 0x32:
		instruction.Code = "STA"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0x33:
		instruction.Code = "INX"
		instruction.Params = append(instruction.Params, "SP")
	case 0x34:
		instruction.Code = "INR"
		instruction.Params = append(instruction.Params, "M")
	case 0x35:
		instruction.Code = "DCR"
		instruction.Params = append(instruction.Params, "M")
	case 0x36:
		instruction.Code = "MVI"
		instruction.Params = append(instruction.Params, "M")
		instruction.Params = append(instruction.Params, "#"+getHexFromByte(code, pc, 1))
	case 0x37:
		instruction.Code = "STC"
	case 0x38:
		instruction.Code = "NOP"
	case 0x39:
		instruction.Code = "DAD"
		instruction.Params = append(instruction.Params, "SP")
	case 0x3a:
		instruction.Code = "LDA"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0x3b:
		instruction.Code = "DCX"
		instruction.Params = append(instruction.Params, "SP")
	case 0x3c:
		instruction.Code = "INR"
		instruction.Params = append(instruction.Params, "A")
	case 0x3d:
		instruction.Code = "DCR"
		instruction.Params = append(instruction.Params, "A")
	case 0x3e:
		instruction.Code = "MVI"
		instruction.Params = append(instruction.Params, "A")
		instruction.Params = append(instruction.Params, "#"+getHexFromByte(code, pc, 1))
	case 0x3f:
		instruction.Code = "CMC"
	case 0x40:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "B")
		instruction.Params = append(instruction.Params, "B")
	case 0x41:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "B")
		instruction.Params = append(instruction.Params, "C")
	case 0x42:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "B")
		instruction.Params = append(instruction.Params, "D")
	case 0x43:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "B")
		instruction.Params = append(instruction.Params, "E")
	case 0x44:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "B")
		instruction.Params = append(instruction.Params, "H")
	case 0x45:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "B")
		instruction.Params = append(instruction.Params, "L")
	case 0x46:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "B")
		instruction.Params = append(instruction.Params, "M")
	case 0x47:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "B")
		instruction.Params = append(instruction.Params, "A")
	case 0x48:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "C")
		instruction.Params = append(instruction.Params, "B")
	case 0x49:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "C")
		instruction.Params = append(instruction.Params, "C")
	case 0x4a:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "C")
		instruction.Params = append(instruction.Params, "D")
	case 0x4b:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "C")
		instruction.Params = append(instruction.Params, "E")
	case 0x4c:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "C")
		instruction.Params = append(instruction.Params, "H")
	case 0x4d:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "C")
		instruction.Params = append(instruction.Params, "L")
	case 0x4e:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "C")
		instruction.Params = append(instruction.Params, "M")
	case 0x4f:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "C")
		instruction.Params = append(instruction.Params, "A")
	case 0x50:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "D")
		instruction.Params = append(instruction.Params, "B")
	case 0x51:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "D")
		instruction.Params = append(instruction.Params, "C")
	case 0x52:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "D")
		instruction.Params = append(instruction.Params, "D")
	case 0x53:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "D")
		instruction.Params = append(instruction.Params, "E")
	case 0x54:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "D")
		instruction.Params = append(instruction.Params, "H")
	case 0x55:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "D")
		instruction.Params = append(instruction.Params, "L")
	case 0x56:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "D")
		instruction.Params = append(instruction.Params, "M")
	case 0x57:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "D")
		instruction.Params = append(instruction.Params, "A")
	case 0x58:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "E")
		instruction.Params = append(instruction.Params, "B")
	case 0x59:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "E")
		instruction.Params = append(instruction.Params, "C")
	case 0x5a:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "E")
		instruction.Params = append(instruction.Params, "D")
	case 0x5b:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "E")
		instruction.Params = append(instruction.Params, "E")
	case 0x5c:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "E")
		instruction.Params = append(instruction.Params, "H")
	case 0x5d:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "E")
		instruction.Params = append(instruction.Params, "L")
	case 0x5e:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "E")
		instruction.Params = append(instruction.Params, "M")
	case 0x5f:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "E")
		instruction.Params = append(instruction.Params, "A")
	case 0x60:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "H")
		instruction.Params = append(instruction.Params, "B")
	case 0x61:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "H")
		instruction.Params = append(instruction.Params, "C")
	case 0x62:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "H")
		instruction.Params = append(instruction.Params, "D")
	case 0x63:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "H")
		instruction.Params = append(instruction.Params, "E")
	case 0x64:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "H")
		instruction.Params = append(instruction.Params, "H")
	case 0x65:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "H")
		instruction.Params = append(instruction.Params, "L")
	case 0x66:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "H")
		instruction.Params = append(instruction.Params, "M")
	case 0x67:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "H")
		instruction.Params = append(instruction.Params, "A")
	case 0x68:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "L")
		instruction.Params = append(instruction.Params, "B")
	case 0x69:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "L")
		instruction.Params = append(instruction.Params, "C")
	case 0x6a:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "L")
		instruction.Params = append(instruction.Params, "D")
	case 0x6b:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "L")
		instruction.Params = append(instruction.Params, "E")
	case 0x6c:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "L")
		instruction.Params = append(instruction.Params, "H")
	case 0x6d:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "L")
		instruction.Params = append(instruction.Params, "L")
	case 0x6e:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "L")
		instruction.Params = append(instruction.Params, "M")
	case 0x6f:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "L")
		instruction.Params = append(instruction.Params, "A")
	case 0x70:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "M")
		instruction.Params = append(instruction.Params, "B")
	case 0x71:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "M")
		instruction.Params = append(instruction.Params, "C")
	case 0x72:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "M")
		instruction.Params = append(instruction.Params, "D")
	case 0x73:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "M")
		instruction.Params = append(instruction.Params, "E")
	case 0x74:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "M")
		instruction.Params = append(instruction.Params, "H")
	case 0x75:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "M")
		instruction.Params = append(instruction.Params, "L")
	case 0x76:
		instruction.Code = "HLT"
	case 0x77:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "M")
		instruction.Params = append(instruction.Params, "A")
	case 0x78:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "A")
		instruction.Params = append(instruction.Params, "B")
	case 0x79:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "A")
		instruction.Params = append(instruction.Params, "C")
	case 0x7a:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "A")
		instruction.Params = append(instruction.Params, "D")
	case 0x7b:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "A")
		instruction.Params = append(instruction.Params, "E")
	case 0x7c:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "A")
		instruction.Params = append(instruction.Params, "H")
	case 0x7d:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "A")
		instruction.Params = append(instruction.Params, "L")
	case 0x7e:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "A")
		instruction.Params = append(instruction.Params, "M")
	case 0x7f:
		instruction.Code = "MOV"
		instruction.Params = append(instruction.Params, "A")
		instruction.Params = append(instruction.Params, "A")
	case 0x80:
		instruction.Code = "ADD"
		instruction.Params = append(instruction.Params, "B")
	case 0x81:
		instruction.Code = "ADD"
		instruction.Params = append(instruction.Params, "C")
	case 0x82:
		instruction.Code = "ADD"
		instruction.Params = append(instruction.Params, "D")
	case 0x83:
		instruction.Code = "ADD"
		instruction.Params = append(instruction.Params, "E")
	case 0x84:
		instruction.Code = "ADD"
		instruction.Params = append(instruction.Params, "H")
	case 0x85:
		instruction.Code = "ADD"
		instruction.Params = append(instruction.Params, "L")
	case 0x86:
		instruction.Code = "ADD"
		instruction.Params = append(instruction.Params, "M")
	case 0x87:
		instruction.Code = "ADD"
		instruction.Params = append(instruction.Params, "A")
	case 0x88:
		instruction.Code = "ADC"
		instruction.Params = append(instruction.Params, "B")
	case 0x89:
		instruction.Code = "ADC"
		instruction.Params = append(instruction.Params, "C")
	case 0x8a:
		instruction.Code = "ADC"
		instruction.Params = append(instruction.Params, "D")
	case 0x8b:
		instruction.Code = "ADC"
		instruction.Params = append(instruction.Params, "E")
	case 0x8c:
		instruction.Code = "ADC"
		instruction.Params = append(instruction.Params, "H")
	case 0x8d:
		instruction.Code = "ADC"
		instruction.Params = append(instruction.Params, "L")
	case 0x8e:
		instruction.Code = "ADC"
		instruction.Params = append(instruction.Params, "M")
	case 0x8f:
		instruction.Code = "ADC"
		instruction.Params = append(instruction.Params, "A")
	case 0x90:
		instruction.Code = "SUB"
		instruction.Params = append(instruction.Params, "B")
	case 0x91:
		instruction.Code = "SUB"
		instruction.Params = append(instruction.Params, "C")
	case 0x92:
		instruction.Code = "SUB"
		instruction.Params = append(instruction.Params, "D")
	case 0x93:
		instruction.Code = "SUB"
		instruction.Params = append(instruction.Params, "E")
	case 0x94:
		instruction.Code = "SUB"
		instruction.Params = append(instruction.Params, "H")
	case 0x95:
		instruction.Code = "SUB"
		instruction.Params = append(instruction.Params, "L")
	case 0x96:
		instruction.Code = "SUB"
		instruction.Params = append(instruction.Params, "M")
	case 0x97:
		instruction.Code = "SUB"
		instruction.Params = append(instruction.Params, "A")
	case 0x98:
		instruction.Code = "SBB"
		instruction.Params = append(instruction.Params, "B")
	case 0x99:
		instruction.Code = "SBB"
		instruction.Params = append(instruction.Params, "C")
	case 0x9a:
		instruction.Code = "SBB"
		instruction.Params = append(instruction.Params, "D")
	case 0x9b:
		instruction.Code = "SBB"
		instruction.Params = append(instruction.Params, "E")
	case 0x9c:
		instruction.Code = "SBB"
		instruction.Params = append(instruction.Params, "H")
	case 0x9d:
		instruction.Code = "SBB"
		instruction.Params = append(instruction.Params, "L")
	case 0x9e:
		instruction.Code = "SBB"
		instruction.Params = append(instruction.Params, "M")
	case 0x9f:
		instruction.Code = "SBB"
		instruction.Params = append(instruction.Params, "A")
	case 0xa0:
		instruction.Code = "ANA"
		instruction.Params = append(instruction.Params, "B")
	case 0xa1:
		instruction.Code = "ANA"
		instruction.Params = append(instruction.Params, "C")
	case 0xa2:
		instruction.Code = "ANA"
		instruction.Params = append(instruction.Params, "D")
	case 0xa3:
		instruction.Code = "ANA"
		instruction.Params = append(instruction.Params, "E")
	case 0xa4:
		instruction.Code = "ANA"
		instruction.Params = append(instruction.Params, "H")
	case 0xa5:
		instruction.Code = "ANA"
		instruction.Params = append(instruction.Params, "L")
	case 0xa6:
		instruction.Code = "ANA"
		instruction.Params = append(instruction.Params, "M")
	case 0xa7:
		instruction.Code = "ANA"
		instruction.Params = append(instruction.Params, "A")
	case 0xa8:
		instruction.Code = "XRA"
		instruction.Params = append(instruction.Params, "B")
	case 0xa9:
		instruction.Code = "XRA"
		instruction.Params = append(instruction.Params, "C")
	case 0xaa:
		instruction.Code = "XRA"
		instruction.Params = append(instruction.Params, "D")
	case 0xab:
		instruction.Code = "XRA"
		instruction.Params = append(instruction.Params, "E")
	case 0xac:
		instruction.Code = "XRA"
		instruction.Params = append(instruction.Params, "H")
	case 0xad:
		instruction.Code = "XRA"
		instruction.Params = append(instruction.Params, "L")
	case 0xae:
		instruction.Code = "XRA"
		instruction.Params = append(instruction.Params, "M")
	case 0xaf:
		instruction.Code = "XRA"
		instruction.Params = append(instruction.Params, "A")
	case 0xb0:
		instruction.Code = "ORA"
		instruction.Params = append(instruction.Params, "B")
	case 0xb1:
		instruction.Code = "ORA"
		instruction.Params = append(instruction.Params, "C")
	case 0xb2:
		instruction.Code = "ORA"
		instruction.Params = append(instruction.Params, "D")
	case 0xb3:
		instruction.Code = "ORA"
		instruction.Params = append(instruction.Params, "E")
	case 0xb4:
		instruction.Code = "ORA"
		instruction.Params = append(instruction.Params, "H")
	case 0xb5:
		instruction.Code = "ORA"
		instruction.Params = append(instruction.Params, "L")
	case 0xb6:
		instruction.Code = "ORA"
		instruction.Params = append(instruction.Params, "M")
	case 0xb7:
		instruction.Code = "ORA"
		instruction.Params = append(instruction.Params, "A")
	case 0xb8:
		instruction.Code = "CMP"
		instruction.Params = append(instruction.Params, "B")
	case 0xb9:
		instruction.Code = "CMP"
		instruction.Params = append(instruction.Params, "C")
	case 0xba:
		instruction.Code = "CMP"
		instruction.Params = append(instruction.Params, "D")
	case 0xbb:
		instruction.Code = "CMP"
		instruction.Params = append(instruction.Params, "E")
	case 0xbc:
		instruction.Code = "CMP"
		instruction.Params = append(instruction.Params, "H")
	case 0xbd:
		instruction.Code = "CMP"
		instruction.Params = append(instruction.Params, "L")
	case 0xbe:
		instruction.Code = "CMP"
		instruction.Params = append(instruction.Params, "M")
	case 0xbf:
		instruction.Code = "CMP"
		instruction.Params = append(instruction.Params, "A")
	case 0xc0:
		instruction.Code = "RNZ"
	case 0xc1:
		instruction.Code = "POP"
		instruction.Params = append(instruction.Params, "B")
	case 0xc2:
		instruction.Code = "JNZ"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xc3:
		instruction.Code = "JMP"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xc4:
		instruction.Code = "CNZ"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xc5:
		instruction.Code = "PUSH"
		instruction.Params = append(instruction.Params, "B")
	case 0xc6:
		instruction.Code = "ADI"
		instruction.Params = append(instruction.Params, "#"+getHexFromByte(code, pc, 1))
	case 0xc7:
		instruction.Code = "RST"
		instruction.Params = append(instruction.Params, "0")
	case 0xc8:
		instruction.Code = "RZ"
	case 0xc9:
		instruction.Code = "RET"
	case 0xca:
		instruction.Code = "JZ"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xcb:
		instruction.Code = "NOP"
	case 0xcc:
		instruction.Code = "CZ"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xcd:
		instruction.Code = "CALL"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xce:
		instruction.Code = "ACI"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 1))
	case 0xcf:
		instruction.Code = "RST"
		instruction.Params = append(instruction.Params, "1")
	case 0xd0:
		instruction.Code = "RNC"
	case 0xd1:
		instruction.Code = "POP"
		instruction.Params = append(instruction.Params, "D")
	case 0xd2:
		instruction.Code = "JNC"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xd3:
		instruction.Code = "OUT"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 1))
	case 0xd4:
		instruction.Code = "CNC"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xd5:
		instruction.Code = "PUSH"
		instruction.Params = append(instruction.Params, "D")
	case 0xd6:
		instruction.Code = "SUI"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 1))
	case 0xd7:
		instruction.Code = "RST"
		instruction.Params = append(instruction.Params, "2")
	case 0xd8:
		instruction.Code = "RC"
	case 0xd9:
		instruction.Code = "NOP"
	case 0xda:
		instruction.Code = "JC"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xdb:
		instruction.Code = "IN"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 1))
	case 0xdc:
		instruction.Code = "CC"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xdd:
		instruction.Code = "NOP"
	case 0xde:
		instruction.Code = "SBI"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 1))
	case 0xdf:
		instruction.Code = "RST"
		instruction.Params = append(instruction.Params, "3")
	case 0xe0:
		instruction.Code = "RPO"
	case 0xe1:
		instruction.Code = "POP"
		instruction.Params = append(instruction.Params, "H")
	case 0xe2:
		instruction.Code = "JPO"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xe3:
		instruction.Code = "XTHL"
	case 0xe4:
		instruction.Code = "CPO"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xe5:
		instruction.Code = "PUSH"
		instruction.Params = append(instruction.Params, "H")
	case 0xe6:
		instruction.Code = "ANI"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 1))
	case 0xe7:
		instruction.Code = "RST"
		instruction.Params = append(instruction.Params, "4")
	case 0xe8:
		instruction.Code = "RPE"
	case 0xe9:
		instruction.Code = "PCH"
	case 0xea:
		instruction.Code = "JPE"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xeb:
		instruction.Code = "XCHG"
	case 0xec:
		instruction.Code = "CPE"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xed:
		instruction.Code = "NOP"
	case 0xee:
		instruction.Code = "XRI"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 1))
	case 0xef:
		instruction.Code = "RST"
		instruction.Params = append(instruction.Params, "5")
	case 0xf0:
		instruction.Code = "RP"
	case 0xf1:
		instruction.Code = "POP"
		instruction.Params = append(instruction.Params, "PSW")
	case 0xf2:
		instruction.Code = "JP"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xf3:
		instruction.Code = "DI"
	case 0xf4:
		instruction.Code = "CP"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xf5:
		instruction.Code = "PUSH"
		instruction.Params = append(instruction.Params, "PSW")
	case 0xf6:
		instruction.Code = "ORI"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 1))
	case 0xf7:
		instruction.Code = "RST"
		instruction.Params = append(instruction.Params, "6")
	case 0xf8:
		instruction.Code = "RM"
	case 0xf9:
		instruction.Code = "SPHL"
	case 0xfa:
		instruction.Code = "JM"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xfb:
		instruction.Code = "EI"
	case 0xfc:
		instruction.Code = "CM"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 2))
	case 0xfd:
		instruction.Code = "NOP"
	case 0xfe:
		instruction.Code = "CPI"
		instruction.Params = append(instruction.Params, getHexFromByte(code, pc, 1))
	case 0xff:
		instruction.Code = "RST"
		instruction.Params = append(instruction.Params, "7")
	default:
		panic(fmt.Sprintf("ERROR: Invalid instruction %02x", code[pc]))
	}

	return instruction
}

// Usage:
//
//   rom, err := io.LoadROM("invaders/invaders.h")
//
//   if err != nil {
//       panic(err)
//   }
//	 Disassembler(rom)
//
func Disassembler(rom []byte, showHexCodes bool, numerate bool) {

	pc := 0
	for pc < len(rom) {
		instruction := getInstruction(&rom, pc)

		if numerate {
			fmt.Printf("%04x ", pc)
		}

		if showHexCodes {
			fmt.Printf("%02x ", rom[pc])
			switch len(instruction.Params) {
			case 0:
				fmt.Printf("     \t")
			case 1:
				fmt.Printf("%02x   \t", rom[pc+1])
			case 2:
				fmt.Printf("%02x %02x \t", rom[pc+1], rom[pc+2])
			}
		}
		pc += len(instruction.Params) + 1
		fmt.Printf("%s \t%s\n", instruction.Code, strings.Join(instruction.Params[:], ","))
	}
}
