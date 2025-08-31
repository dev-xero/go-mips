// ===============================================================
// GO-MIPS Instruction Set Editor and Simulator
// @dev-xero on GitHub
// 2025
// ===============================================================
package validation

import "fmt"

var (
	ErrInvalidInstruction    = fmt.Errorf("invalid instruction format")
	ErrInvalidRegister       = fmt.Errorf("invalid register reference")
	ErrOutOfRange            = fmt.Errorf("value out of allowed range")
	ErrRegisterParsingFailed = fmt.Errorf("register parsing failed")
	ErrUnsupportedRegister   = fmt.Errorf("unsupported register")
)

// ===============================================================
// Validate Instruction Parts
// ---------------------------------------------------------------
// Exists to confirm gotten instruction matches expected form
// ===============================================================
func ValidateInstructionParts(op string, got, expected int) error {
	if got != expected {
		return &InstructionError{
			Instruction: op,
			Expected:    expected,
			Got:         got,
		}
	}
	return nil
}
