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

		instruction := cpu.Decode(line)

		fmt.Println("instruction:", instruction)
		fmt.Println("cpu registers:", cpu.Registers)
	}

	fmt.Println("MIPS interpreter finished execution.")
}
