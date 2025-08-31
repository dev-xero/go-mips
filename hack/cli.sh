#!/bin/bash

echo "> Formatting Go code..."
go fmt ./...

if [ $? -ne 0 ]; then
    echo "[x] Formatting failed, terminating."
    exit 1
fi

echo "> Building the project..."
go build -o ./bin/mips-sim ./cmd/mips/main.go

if [ $? -eq 0 ]; then
    echo -e "> Build successful. Running the simulator...\n"
    ./bin/mips-sim
else
    echo "[x] Build failed."
fi
