package checks

import "fmt"

var (
	ErrInvalidInstruction    = fmt.Errorf("invalid instruction format")
	ErrInvalidRegister       = fmt.Errorf("invalid register reference")
	ErrOutOfRange            = fmt.Errorf("value out of allowed range")
	ErrRegisterParsingFailed = fmt.Errorf("register parsing failed")
	ErrUnsupportedRegister   = fmt.Errorf("unsupported register")
)

type RegisterError struct {
	Register string
	Reason   string
}

func (e *RegisterError) Error() string {
	return fmt.Sprintf("register error '%s': %s", e.Register, e.Reason)
}

type InstructionError struct {
	Instruction string
	Expected    int
	Got         int
}

func (e *InstructionError) Error() string {
	return fmt.Sprintf("instruction error '%s': expected %d parts, got %d", e.Instruction, e.Expected, e.Got)
}

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
