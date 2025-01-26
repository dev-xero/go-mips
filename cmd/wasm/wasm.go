package main

import (
	"fmt"
	"syscall/js"

	"github.com/dev-xero/go-mips/internal/mips"
)

type Simulator struct {
	cpu     mips.CPU
	program []mips.Instruction
}

// Loads program with instructions from client
func (sim *Simulator) LoadProgram(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return js.ValueOf(false)
	}

	jsArray := args[0]
	length := jsArray.Length()
	instructions := make([]mips.Instruction, length)

	for i := 0; i < length; i++ {
		line := jsArray.Index(i).String()
		inst, err := sim.cpu.Decode(line)

		if err != nil {
			return js.ValueOf(false)
		}

		instructions[i] = inst
	}

	sim.program = instructions

	return js.ValueOf(true)

}

func (sim *Simulator) RegistersToJsValues() map[string]interface{} {
	registers := make(map[string]interface{})

	for i, _ := range sim.cpu.Registers {
		tag := ""

		if i == 0 {
			// r0 is always 0
			tag = "zero"
		} else if i == 1 {
			// reserved for assembler
			tag = "at"

		} else if i == 2 || i == 3 {
			// result value registers
			tag = fmt.Sprintf("v%d", i-2)

		} else if i >= 4 && i <= 7 {
			// argument registers
			tag = fmt.Sprintf("a%d", i-4)

		} else if i >= 8 && i <= 15 {
			// temporary registers
			tag = fmt.Sprintf("t%d", i-8)

		} else if i >= 16 && i <= 23 {
			// saved registers
			tag = fmt.Sprintf("s%d", i-16)

		} else if i == 26 || i == 27 {
			// reserved by os
			tag = fmt.Sprintf("k%d", i-26)

		} else if i == 24 {
			// other temps
			tag = "t8"

		} else if i == 25 {
			// other temps
			tag = "t9"

		} else if i == 28 {
			// global pointer
			tag = "gp"

		} else if i == 29 {
			// stack pointer
			tag = "sp"

		} else if i == 30 {
			// frame pointer
			tag = "fp"

		} else if i == 31 {
			// return address
			tag = "ra"
		}

		registers[fmt.Sprintf("$%s", tag)] = uint32(sim.cpu.Registers[i])
	}

	return registers
}

// Inspects simulator state, exposing it to the client
func (sim *Simulator) InspectState(this js.Value, args []js.Value) interface{} {

	return js.ValueOf(map[string]interface{}{
		"registers": sim.RegistersToJsValues(),
		// "program":      sim.program,
		// "current step": sim.cpu.PC,
	})
}

func main() {
	simulator := &Simulator{
		cpu: *mips.NewCPU(),
	}

	js.Global().Set("loadProgram", js.FuncOf(simulator.LoadProgram))
	js.Global().Set("inspectSimulator", js.FuncOf(simulator.InspectState))

	select {}
}
