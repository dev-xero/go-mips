# Go MIPS

Basic CPU simulator for the MIPS assembly language. Built for MARS defectors and those who prefer a modern user interface.

## Table Of Contents


## File Structure

```
.
├── bin                   # compiled binaries
├── cmd                   # core go platform modules
│   ├── main              # cli
│   ├── server            # server
│   └── wasm              # web assembly
├── hack                  # shell scripts
├── internal              # internal modules
│   ├── CPU               # simulated cpu
│   └── validation        # validation
├── public                # public/web files
├── reference             # specifications
└── static                # statically served files
    ├── css               # stylesheets
    ├── img               # images
    ├── js                # javascript
    └── wasm              # web assembly binaries
```

## How It Works


## Running Locally


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
