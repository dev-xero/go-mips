package mips

type Register uint32

/* 
	CPU abstraction
	
	Registers - 32 32-bit registers as per MIPS spec.
	PC - Program Counter.
	Memory - Simulated 1MB of memory.
	HI, LO - 32 bit.

*/
type CPU struct {
	Registers [32]Register
	PC        uint32
	Memory    []byte
	HI, LO    Register
}
