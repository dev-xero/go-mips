package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dev-xero/go-mips/internal/mips"
)

// Simulator entry point.
func main() {
	cpu := mips.NewCPU()

	fmt.Printf("MIPS interpreter started, use 'END' to exit.\n\n")

	for {
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')

		if strings.TrimSpace(line) == "END" {
			break
		}

		instruction, err := cpu.Decode(line)
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
