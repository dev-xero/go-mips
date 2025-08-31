// ===============================================================
// GO-MIPS Instruction Set Editor and Simulator
// @dev-xero on GitHub
// 2025
// ===============================================================
package mips

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dev-xero/go-mips/internal/validation"
)

type Register int32
type InstructionType int

// ===============================================================
// OP Code to Func Code Mapping
// ===============================================================
var opCodeMap = map[string]uint8{
	"add":  0x20,
	"sub":  0x22,
	"addi": 0x8,
	// "subi": 0x8,
}

// ===============================================================
// MIPS instruction types can be classified into 3
// ---------------------------------------------------------------
// R type: Register type
// I type: Immediate type
// J type: Jump type
// ===============================================================
const (
	R_TYPE InstructionType = iota
	I_TYPE InstructionType = iota
	J_TYPE InstructionType = iota
)

// ===============================================================
// CPU abstraction
// ---------------------------------------------------------------
// The CPU is composed of the following components:
// - Registers: 32 32-bit registers
// - PC:        Program Counter
// - Memory:    Simulated 1MB memory
// - HI:        32-bit larger half of multiplication
// - LO:        lower/smaller half of multiplication
// ===============================================================
type CPU struct {
	Registers [32]Register
	PC        uint32
	Memory    []byte
	HI, LO    Register
}

// ===============================================================
// MIP instruction abstraction
// ---------------------------------------------------------------
// Go abstraction of a MIPS instruction
// Note:
// Due to language constraints, certain fields have been expanded
// to 8-bits from their original 6 or 5-bits in the MIPS spec.
// ===============================================================
type Instruction struct {
	Type      InstructionType
	Opcode    uint8
	Rs        int16
	Rt        int16
	Rd        int16
	Shamt     uint8
	Funct     uint8
	Immediate int16
	Address   uint32
}

// ===============================================================
// Instantiates a new CPU with 1MB memory
// ===============================================================
func NewCPU() *CPU {
	return &CPU{
		Memory: make([]byte, 1024*1024),
	}
}

// ================================================================
// **** DECODING PHASE ****
// ================================================================

// ===============================================================
// Parses and returns register values
// ===============================================================
func parseRegister(str string) (int16, error) {
	str = strings.Trim(str, "$,")

	switch {
	case str == "zero":
		// Special register $zero always contains 0
		return 0, nil

	case strings.HasPrefix(str, "t"):
		// Register t0-t9 starts at 8
		OFFSET := 8
		num, err := strconv.Atoi(strings.TrimPrefix(str, "t"))

		if err != nil {
			return 0, &validation.RegisterError{
				Register: str,
				Reason:   "invalid t-register number",
			}
		}

		if num < 0 || num > 9 {
			return 0, &validation.RegisterError{
				Register: str,
				Reason:   "t-register must be 0-9",
			}
		}

		return int16(OFFSET + num), nil

	case strings.HasPrefix(str, "s"):
		// Register s0-s7 starts at 16
		OFFSET := 16
		num, err := strconv.Atoi(strings.TrimPrefix(str, "s"))

		if err != nil {
			return 0, &validation.RegisterError{
				Register: str,
				Reason:   "invalid s-register number",
			}
		}

		if num < 0 || num > 7 {
			return 0, &validation.RegisterError{
				Register: str,
				Reason:   "s-register must be 0-7",
			}
		}

		return int16(OFFSET + num), nil

	// Unknown/Unsupported register
	default:
		return 0, validation.ErrInvalidRegister
	}
}

// ===============================================================
// General purpose registers parsing
// ===============================================================
func parseRTypeRegisters(regs []string) ([]int16, error) {
	result := make([]int16, len(regs))

	// For each passed register string, produce the equivalent CPU
	// native one
	for i, reg := range regs {
		parsed, err := parseRegister(reg)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to parse register %s : %w",
				reg,
				err,
			)
		}
		result[i] = parsed
	}

	return result, nil
}

// ===============================================================
// General purpose Immediate-type register parsing
// ===============================================================
func parseITypeRegister(regs []string) ([]int16, error) {
	registerEndIndex := len(regs) - 2
	result := make([]int16, len(regs))

	// Parse register parts
	for i, reg := range regs[:registerEndIndex+1] {
		parsed, err := parseRegister(reg)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to parse register %s : %w",
				reg,
				err,
			)
		}
		result[i] = parsed
	}

	// Read immediate parts
	// We'll append this to the results at the end
	finalPart := regs[registerEndIndex+1]
	immediate, err := strconv.Atoi(finalPart)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to parse immediate value '%s' : %w",
			finalPart,
			err,
		)
	}

	result[registerEndIndex+1] = int16(immediate)
	return result, nil
}

// ===============================================================
// Reads MIPS assembly line and then decodes/parses it
// ===============================================================
func (cpu *CPU) Decode(line string) (Instruction, error) {
	parts := strings.Fields(line)

	// Ignore empty lines
	if len(parts) == 0 {
		return Instruction{}, validation.ErrInvalidInstruction
	}

	op := parts[0]

	switch op {
	case "add", "sub":
		// Register type instructions
		// Format: add $rd, $rs, $rt
		if err := validation.ValidateInstructionParts(op, len(parts), 4); err != nil {
			return Instruction{}, err
		}

		regs, err := parseRTypeRegisters(parts[1:])
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
		// Immediate type instruction
		// Format: addi $rt, $rs, #immediate
		if err := validation.ValidateInstructionParts(op, len(parts), 4); err != nil {
			return Instruction{}, err
		}

		regs, err := parseITypeRegister(parts[1:])
		if err != nil {
			return Instruction{}, fmt.Errorf("%s: %w", validation.ErrRegisterParsingFailed.Error(), err)
		}

		return Instruction{
			Type:      I_TYPE,
			Opcode:    0,
			Rs:        regs[1],
			Rt:        regs[0],
			Funct:     opCodeMap[op],
			Immediate: int16(regs[2]),
		}, nil
	}

	return Instruction{}, fmt.Errorf("unsupported instruction: %s", op)
}

// ================================================================
// **** EXECUTION PHASE ****
// ================================================================

// ===============================================================
// Executes Register type instructions only
// ===============================================================
func (cpu *CPU) ExecuteRType(inst Instruction) error {
	switch inst.Funct {

	case opCodeMap["add"]:
		fmt.Printf(
			"[add] r type: values at rs: %d, rt: %d\n",
			cpu.Registers[inst.Rs],
			cpu.Registers[inst.Rt],
		)
		cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs] + cpu.Registers[inst.Rt]

	case opCodeMap["sub"]:
		fmt.Printf(
			"[sub] r type: values at rs: %d, rt: %d\n",
			cpu.Registers[inst.Rs],
			cpu.Registers[inst.Rt],
		)
		cpu.Registers[inst.Rd] = cpu.Registers[inst.Rs] - cpu.Registers[inst.Rt]
	}

	return nil
}

// ===============================================================
// Executes Immediate type instructions only
// ===============================================================
func (cpu *CPU) ExecuteIType(inst Instruction) error {
	switch inst.Funct {

	case opCodeMap["addi"]:
		fmt.Printf(
			"[addi] I-type: values at rt: %d, rs: %d, i: %d\n",
			inst.Rt,
			inst.Rs,
			inst.Immediate,
		)
		cpu.Registers[inst.Rt] = cpu.Registers[inst.Rs] + Register(inst.Immediate)
	}

	return nil
}

// ===============================================================
// Execute MIPS instructions
// ===============================================================
func (cpu *CPU) Execute(inst Instruction) error {
	switch inst.Type {

	case R_TYPE:
		cpu.ExecuteRType(inst)

	case I_TYPE:
		cpu.ExecuteIType(inst)
	}

	return nil
}
