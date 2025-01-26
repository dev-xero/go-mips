package mips

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dev-xero/go-mips/internal/checks"
)

type Register uint32

type InstructionType int

// OP -> FUNC code map
var opCodeMap = map[string]uint8{
	"add":  0x20,
	"sub":  0x22,
	"addi": 0x8,
}

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

// MIP instruction abstraction.
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

// func (inst *Instruction) String() string {
// 	return fmt.Sprintf()
// }

// General purpose register parsing
func parseRegisters(regs []string) ([]uint8, error) {
	result := make([]uint8, len(regs))

	for i, reg := range regs {
		parsed, err := parseRegister(reg)
		if err != nil {
			return nil, fmt.Errorf("failed to parse register %s : %w", reg, err)

		}
		result[i] = parsed
	}

	return result, nil
}

// General purpose immediate-type parsing
func parseRegistersImmediate(regs []string) ([]uint8, error) {
	registerEndIndex := len(regs) - 2
	result := make([]uint8, registerEndIndex+2)

	// Parse register parts
	for i, reg := range regs[:registerEndIndex+1] {
		parsed, err := parseRegister(reg)
		if err != nil {
			return nil, fmt.Errorf("failed to parse register %s : %w", reg, err)
		}
		result[i] = parsed
	}

	// Read immediate parts
	// We'll append this to the results at the end
	finalPart := regs[registerEndIndex+1]
	immediate, err := strconv.Atoi(finalPart)

	if err != nil {
		return nil, fmt.Errorf("failed to parse immediate value '%s' : %w", finalPart, err)
	}

	result[registerEndIndex+1] = uint8(immediate)

	return result, nil
}

// Parses and returns register values.
func parseRegister(s string) (uint8, error) {
	s = strings.Trim(s, "$,")

	switch {
	// Special register $zero always contains 0
	case s == "zero":
		return 0, nil

	// Register t0-t9 starts at 8
	case strings.HasPrefix(s, "t"):
		num, err := strconv.Atoi(strings.TrimPrefix(s, "t"))
		if err != nil {
			return 0, &checks.RegisterError{Register: s, Reason: "invalid t-register number"}
		}
		if num < 0 || num > 9 {
			return 0, &checks.RegisterError{Register: s, Reason: "t-register must be 0-9"}
		}
		return uint8(8 + num), nil

	// Register s0-s7 starts at 16
	case strings.HasPrefix(s, "s"):
		num, err := strconv.Atoi(strings.TrimPrefix(s, "s"))
		if err != nil {
			return 0, &checks.RegisterError{Register: s, Reason: "invalid s-register number"}
		}
		if num < 0 || num > 7 {
			return 0, &checks.RegisterError{Register: s, Reason: "s-register must be 0-7"}
		}
		return uint8(16 + num), nil

	// Unknown register
	default:
		return 0, checks.ErrInvalidRegister
	}
}

// Instantiates a new CPU.
func NewCPU() *CPU {
	return &CPU{
		Memory: make([]byte, 1024*1024), // 1MB
	}
}

// Reads MIPS assembly line and parses it.
func (cpu *CPU) Decode(line string) (Instruction, error) {
	parts := strings.Fields(line)

	if len(parts) == 0 {
		return Instruction{}, checks.ErrInvalidInstruction
	}

	op := parts[0]

	switch op {
	case "add", "sub":
		if err := checks.ValidateInstructionParts(op, len(parts), 4); err != nil {
			return Instruction{}, err
		}

		regs, err := parseRegisters(parts[1:])
		if err != nil {
			return Instruction{}, fmt.Errorf("register parsing failed: %w", err)
		}

		return Instruction{
			Type:   R_TYPE,
			Opcode: 0,
			Rd:     regs[0],
			Rs:     regs[1],
			Rt:     regs[2],
			Funct:  opCodeMap[op],
		}, nil
	case "addi":
		if err := checks.ValidateInstructionParts(op, len(parts), 4); err != nil {
			return Instruction{}, err
		}

		regs, err := parseRegistersImmediate(parts[1:])
		if err != nil {
			return Instruction{}, fmt.Errorf("%s: %w", checks.ErrRegisterParsingFailed.Error(), err)
		}

		fmt.Println("regs i:", regs)

		return Instruction{
			Type:   R_TYPE,
			Opcode: 0,
			Rd:     regs[0],
			Rs:     regs[1],
			Rt:     regs[2],
			Funct:  opCodeMap[op],
		}, nil
	}

	return Instruction{}, fmt.Errorf("unsupported instruction: %s", op)
}
