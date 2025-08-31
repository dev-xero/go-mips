// ===============================================================
// GO-MIPS Instruction Set Editor and Simulator
// @dev-xero on GitHub
// 2025
// ===============================================================
package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/dev-xero/go-mips/internal/mips"
)

// ===============================================================
// Simulator State Struct
// ===============================================================
type Simulator struct {
	cpu     mips.CPU
	program []mips.Instruction
}

// ===============================================================
// Reset State
// ---------------------------------------------------------------
// Nukes all register values to 0
// ===============================================================
func (sim *Simulator) ResetState(this js.Value, args []js.Value) interface{} {
	sim.cpu.Memory = make([]byte, 1024*1024) // 1MB
	sim.cpu.PC = 0
	for i := range 32 {
		sim.cpu.Registers[i] = 0
	}
	return js.ValueOf(true)
}

// ===============================================================
// Load Program
// ---------------------------------------------------------------
// Loads program with instructions from client
// ===============================================================
func (sim *Simulator) LoadProgram(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return js.ValueOf(false)
	}

	// We receive "length" instructions from the client, so create a
	// "CPU-native" slice to accommodate them. By CPU-native, I refer
	// to the simulated one
	clientInstructions := args[0]
	length := clientInstructions.Length()
	instructions := make([]mips.Instruction, length)

	// In this step, we have to convert the raw mnemonics into CPU
	// "decoded" formats
	for i := 0; i < length; i++ {
		line := clientInstructions.Index(i).String()
		decodedInstruction, err := sim.cpu.Decode(line)

		// Immediately terminate if there is an error
		if err != nil {
			return js.ValueOf(false)
		}
		instructions[i] = decodedInstruction
	}

	sim.program = instructions
	return js.ValueOf(true)
}

// ===============================================================
// Registers to JS Values
// ---------------------------------------------------------------
// Converts CPU registers into web friendly tags
// ===============================================================
func (sim *Simulator) RegistersToJsValues() map[string]interface{} {
	registers := make(map[string]interface{})

	// -----------------------------------------------------------
	// MIPS Register Mapping Reference
	// -----------------------------------------------------------
	// Bit 0:          $zero
	// Bit 1:          Reserved for "at"
	// Bit 2-3:        Result Value Registers (V)
	// Bit 4-7:        Argument Value Registers (A)
	// Bit 8-15:       Temporary Value Registers (T)
	// Bit 16-23:      Saved Value Registers (S)
	// Bit 24 and 25:  t8 and t9 Registers (T)
	// Bit 26-27:      OS Reserved Registers
	// Bit 28:         Global Pointer (gp)
	// Bit 29:         Stack Pointer (sp)
	// Bit 30:         Frame Pointer (fp)
	// Bit 31:         Return Address (ra)
	for i, _ := range sim.cpu.Registers {
		tag := ""
		if i == 0 {
			tag = "zero"
		}
		if i == 1 {
			tag = "at"
		}
		if i == 2 || i == 3 {
			tag = fmt.Sprintf("v%d", i-2)
		}
		if i >= 4 && i <= 7 {
			tag = fmt.Sprintf("a%d", i-4)
		}
		if i >= 8 && i <= 15 {
			tag = fmt.Sprintf("t%d", i-8)
		}
		if i >= 16 && i <= 23 {
			tag = fmt.Sprintf("s%d", i-16)
		}
		if i == 26 || i == 27 {
			tag = fmt.Sprintf("k%d", i-26)
		}
		if i == 24 {
			tag = "t8"
		}
		if i == 25 {
			tag = "t9"
		}
		if i == 28 {
			tag = "gp"
		}
		if i == 29 {
			tag = "sp"
		}
		if i == 30 {
			tag = "fp"
		}
		if i == 31 {
			tag = "ra"
		}

		registers[fmt.Sprintf("$%s", tag)] = uint32(sim.cpu.Registers[i])
	}
	return registers
}

// ===============================================================
// Inspect State
// ---------------------------------------------------------------
// Inspects simulator state, exposing it to the client
// ===============================================================
func (sim *Simulator) InspectState(this js.Value, args []js.Value) interface{} {
	state := map[string]interface{}{
		"registers":   sim.RegistersToJsValues(),
		"program":     sim.program,
		"currentStep": sim.cpu.PC,
	}

	// Attempt converting this into JSON
	jsonData, err := json.Marshal(state)
	if err != nil {
		return js.Null()
	}

	// Parse in Client JS context
	return js.Global().Get("JSON").Call("parse", string(jsonData))
}

// ===============================================================
// Step
// ---------------------------------------------------------------
// Steps (executes) assembly line by line
// ===============================================================
func (sim *Simulator) Step(this js.Value, args []js.Value) interface{} {
	// Program Counter must not exceed instruction count
	if sim.cpu.PC >= uint32(len(sim.program)) {
		return js.ValueOf(false)
	}

	// Execute current instruction
	currentInst := sim.program[sim.cpu.PC]
	err := sim.cpu.Execute(currentInst)
	if err != nil {
		return js.ValueOf(false)
	}

	// Increment program counter and report CPU state
	sim.cpu.PC += 1
	state := map[string]interface{}{
		"registers":   sim.RegistersToJsValues(),
		"program":     sim.program,
		"currentStep": sim.cpu.PC,
	}

	jsonData, err := json.Marshal(state)
	if err != nil {
		return js.Null()
	}

	// Parse JSON in Client JS context
	return js.Global().Get("JSON").Call("parse", string(jsonData))
}

// ===============================================================
// WASM Entry Point
// ===============================================================
func main() {
	simulator := &Simulator{
		cpu: *mips.NewCPU(),
	}

	// Wasm Mapping for Client use
	js.Global().Set("resetState", js.FuncOf(simulator.ResetState))
	js.Global().Set("loadProgram", js.FuncOf(simulator.LoadProgram))
	js.Global().Set("inspectSimulator", js.FuncOf(simulator.InspectState))
	js.Global().Set("simulatorStep", js.FuncOf(simulator.Step))

	select {}
}
