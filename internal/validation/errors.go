// ===============================================================
// GO-MIPS Instruction Set Editor and Simulator
// @dev-xero on GitHub
// 2025
// ===============================================================
package validation

import "fmt"

// ===============================================================
// Register Error
// ===============================================================
type RegisterError struct {
	Register string
	Reason   string
}

// Returns register error as formatted string
func (e *RegisterError) Error() string {
	return fmt.Sprintf(
		"register error '%s': %s",
		e.Register,
		e.Reason,
	)
}

// ===============================================================
// Instruction Error
// ===============================================================
type InstructionError struct {
	Instruction string
	Expected    int
	Got         int
}

// Returns instruction error as formatted string
func (e *InstructionError) Error() string {
	return fmt.Sprintf(
		"instruction error '%s': expected %d parts, got %d",
		e.Instruction,
		e.Expected,
		e.Got,
	)
}
