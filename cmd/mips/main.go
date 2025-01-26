package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dev-xero/go-mips/internal/mips"
)

// Simulator CLI
func main() {
	cpu := mips.NewCPU()

	fmt.Printf("MIPS interpreter started, use 'END' to exit.\n\n")

	// Fetch-Decode-Execute Cycle
	for {
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')

		// Remove excess whitespace
		if strings.TrimSpace(line) == "END" {
			break
		}

		// Decode instruction
		instruction, err := cpu.Decode(line)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Execute this instruction
		err = cpu.Execute(instruction)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Debug
		fmt.Println("instruction:", instruction)
		fmt.Println("cpu registers:", cpu.Registers)
		fmt.Println("")
	}

	fmt.Println("MIPS interpreter finished execution.")
}
