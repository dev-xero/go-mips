# Go MIPS!

Go interpreter (and simulator) for the MIPS assembly language. Built for MARS defectors and those who prefer a modern user interface.

## MIPS Reference Sheet

Full reference [here.](https://uweb.engr.arizona.edu/~ece369/Resources/spim/MIPSReference.pdf)

These are the instructions Go MIPS currently interprets and executes.

| Function | Instruction | Effect |
| -------- | ----------- | ------ |
| add | add R1, R2, R3 | R1 = R2 + R3 |
| sub | sub R1, R2, R3 | R1 = R2 - R3 |
| add immediate | addi R1, R2, #num | R1 = R2 + num |
| and | and R1, R2, R3 | R1 = R2 & R3 |
| or | or R1, R2, R3 | R1 = R2 | R3 |
