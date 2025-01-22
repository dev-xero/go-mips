# Go MIPS!

Go interpreter (and simulator) for the MIPS assembly language. Built for MARS defectors and those who prefer a modern user interface.

## MIPS Reference Sheet

Full reference [here.](https://uweb.engr.arizona.edu/~ece369/Resources/spim/MIPSReference.pdf)

These are the instructions Go MIPS currently interprets and executes.

| Function | Instruction | Effect | Explanation |
| -------- | ----------- | ------ | ----------- |
| add | add R1, R2, R3 | R1 = R2 + R3 | Adds the values in two source registers into the destination register |
| sub | sub R1, R2, R3 | R1 = R2 - R3 | Subtracts the values in two source registers into the destination register. |
| add immediate | addi R1, R2, #num | R1 = R2 + num | Adds immediate value to value in source register and stores that. |
| and | and R1, R2, R3 | R1 = R2 & R3 | Binary and operation on two source registers, result is stored. |
| or | or R1, R2, R3 | R1 = R2 \| R3 | Binary or operation on two source registers, result is stored. |
