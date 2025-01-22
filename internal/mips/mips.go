package mips

import (
	"strconv"
	"strings"
)

type Register uint32

type InstructionType int

// MIPS instruction types can be classified into 3.
const (
	R_TYPE InstructionType = iota // Register type
	I_TYPE                        // Instruction type
	J_TYPE                        // Jump type
)

// CPU abstraction.
type CPU struct {
	Registers [32]Register // 32 32-bit registers
	PC        uint32       // program counter
	Memory    []byte       // simulated 1MB memory
	HI, LO    Register     // 32 bit larger and smaller part of multiplication
}

// MIP instruction abstraction
//
// Due to language constraints, certain fields have been
// expanded to 8-bits from their original 6 or 5-bits in the MIPS spec.
type Instruction struct {
	Type      InstructionType
	Opcode    uint8
	Rs        uint8
	Rt        uint8
	Rd        uint8
	Shamt     uint8
	Funct     uint8
	Immediate uint16
	Address   uint32
}

// Instantiates a new CPU
func NewCPU() *CPU {
	return &CPU{
		Memory: make([]byte, 1024 * 1024), // 1MB
	}
}

// Reads MIPS assembly line and parses it.
func (cpu *CPU) Decode(line string) Instruction {
	parts := strings.Fields(line)

	if len(parts) == 0 {
		return Instruction{}
	}

	op := parts[0]

	switch op {
	case "add":
		return Instruction{
			Type:   R_TYPE,
			Opcode: 0,
			Rd:     parseRegister(parts[1]),
			Rs:     parseRegister(parts[2]),
			Rt:     parseRegister(parts[3]),
			Funct:  0x22,
		}
	}

	return Instruction{}
}

// Parses and returns register values.
func parseRegister(s string) uint8 {
	s = strings.Trim(s, "$s,")
	reg, _ := strconv.Atoi(s)

	return uint8(reg)
}
