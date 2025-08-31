# Go MIPS

[![Go](https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white&style=for-the-badge)]("https://github.com/dev-xero/go-mips")
[![License](https://img.shields.io/github/license/dev-xero/go-mips?style=for-the-badge&colorA=131820&colorB=FFFFFF&logo=github)]("https://github.com/dev-xero/go-mips")

Basic CPU simulator for the MIPS assembly language. Built for MARS defectors and those who prefer a modern user interface.

## Table Of Contents

-   [Go MIPS](#go-mips)
    -   [Table Of Contents](#table-of-contents)
    -   [File Structure](#file-structure)
    -   [Libraries](#libraries)
    -   [How It Works](#how-it-works)
        -   [How It Works: In a little more depth](#how-it-works-in-a-little-more-depth)
    -   [Running Locally](#running-locally)
        -   [Via the Command Line](#via-the-command-line)
        -   [Via the Browser](#via-the-browser)
    -   [MIPS Reference Sheet](#mips-reference-sheet)

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
│   ├── cpu               # simulated cpu
│   └── validation        # validation
├── public                # public/web files
├── reference             # specifications
└── static                # statically served files
    ├── css               # stylesheets
    ├── img               # images
    ├── js                # javascript
    └── wasm              # web assembly binaries
```

## Libraries

I try to stay as lean as possible when it comes to external libraries, and perhaps it's no surprise that there isn't any in use except from the Official Go WebAssembly execution runtime support, and syscall/js library:

-   [syscall/js](https://pkg.go.dev/syscall/js)
-   [web assembly runtime support](https://go.googlesource.com/go.git/+/refs/tags/go1.17rc1/misc/wasm/wasm_exec.js)

I purposely decided not to use React or any "heavy" frontend libraries because it would be overkill and frankly isn't the focus of this project.

## How It Works

GoMIPs basically simulates a CPU capable of the complete fetch-decode-execute cycle given instructions that follow MIPs specifications.

The software can be run in two modes:

-   via the command line interface (CLI)
-   or via the browser

The command line interface is the most straightforward way to get started. You simply compile the source files from the shell script and type in MIPs mnemonics.

The browser, which is more interesting, relies on WebAssembly to get the same functionality running.

### How It Works: In a little more depth

The core of everything is the CPU abstraction, which is a struct:

```go
type CPU struct {
	Registers [32]Register
	PC        uint32
	Memory    []byte
	HI, LO    Register
}
```

Of course this is a very loose definition, but is sufficient for what we want to achieve. I have 32 32-bit registers, as stated in the official specs, a program counter, and simulated memory (1MB), and two special registers: `HI` and `LO` which store the upper and lower 32-bits from multiplication and division operations. Though unused for now.

I should also mention that some data type sizes, for instance R-type instructions have been expanded to 8-bits (from the originally defined 6-bits) due to language constraints, but shouldn't pose much of an issue.

The CPU instruction is represented via another struct which specifies several important properties such as the Instruction Type, Op Code, and Address.

```go
// ===============================================================
// MIP instruction abstraction
// ---------------------------------------------------------------
// Go abstraction of a MIPS instruction
// Note:
// Due to language constraints, certain fields have been expanded
// to 8-bits from their original 6 or 5-bits in the MIPS spec.
// ===============================================================
type Instruction struct {
	Type      InstructionType
	Opcode    uint8
	Rs        int16
	Rt        int16
	Rd        int16
	Shamt     uint8
	Funct     uint8
	Immediate int16
	Address   uint32
}
```

## Running Locally

-   [Via the Command Line](#via-the-command-line)
-   [Via the Browser](#via-the-browser)

### Via the Command Line

You can compile the source code using the below command:

```sh
go build -o ./bin/mips-sim ./cmd/main/main.go
```

Then run the compiled binary using:

```sh
./bin/mips-sim
```

Or alternatively use the script provided in `/hack/cli.sh` to run additional formatting before building and executing.

### Via the Browser

Prerequisite: As already mentioned, the browser makes use of compiled web assembly. You can use the `/hack/wasm.sh` script to compile and output the binary under the `/wasm` directory.

I wrote a (very) bare-bones server for the html and static files (css/js/wasm) the editor will use. To spin it up use the following command:

```sh
go build -o ./bin/server ./cmd/server/server.go
./bin/server
```

Or use the script `/hack/server.sh` for additional formatting, compilation and running in one go (pun-intended).

The default port is `8080` accessible via localhost.

## MIPS Reference Sheet

Full reference [here.](./reference/MIPS_Instruction_Set.pdf)

These are the instructions Go MIPS currently interprets and executes.

| Function      | Instruction       | Effect        | Explanation                                                                 |
| ------------- | ----------------- | ------------- | --------------------------------------------------------------------------- |
| add           | add R1, R2, R3    | R1 = R2 + R3  | Adds the values in two source registers into the destination register       |
| sub           | sub R1, R2, R3    | R1 = R2 - R3  | Subtracts the values in two source registers into the destination register. |
| add immediate | addi R1, R2, #num | R1 = R2 + num | Adds immediate value to value in source register and stores that.           |
| and           | and R1, R2, R3    | R1 = R2 & R3  | Binary and operation on two source registers, result is stored.             |
| or            | or R1, R2, R3     | R1 = R2 \| R3 | Binary or operation on two source registers, result is stored.              |
