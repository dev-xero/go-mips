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

## MIPS Reference Sheet

Full reference [here.](https://uweb.engr.arizona.edu/~ece369/Resources/spim/MIPSReference.pdf)

These are the instructions Go MIPS currently interprets and executes.

| Function      | Instruction       | Effect        | Explanation                                                                 |
| ------------- | ----------------- | ------------- | --------------------------------------------------------------------------- |
| add           | add R1, R2, R3    | R1 = R2 + R3  | Adds the values in two source registers into the destination register       |
| sub           | sub R1, R2, R3    | R1 = R2 - R3  | Subtracts the values in two source registers into the destination register. |
| add immediate | addi R1, R2, #num | R1 = R2 + num | Adds immediate value to value in source register and stores that.           |
| and           | and R1, R2, R3    | R1 = R2 & R3  | Binary and operation on two source registers, result is stored.             |
| or            | or R1, R2, R3     | R1 = R2 \| R3 | Binary or operation on two source registers, result is stored.              |
