package main

import (
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

// Inspects simulator state, exposing it to the client
func (sim *Simulator) InspectState(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(map[string]interface{}{
		"registers":    sim.cpu.Registers,
		"program":      sim.program,
		"current step": sim.cpu.PC,
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
