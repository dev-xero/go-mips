#!/bin/bash

# Format all Go files
echo "> Formatting Go code..."
go fmt ./...

# Formatting should succeed
if [ $? -ne 0 ]; then
    echo "[x] Formatting failed, terminating."
    exit 1
fi

# Compile the server
echo "> Compiling sever..."
go build -o ./bin/server ./cmd/server/server.go

# Check if the build was successful
if [ $? -eq 0 ]; then
    echo -e "> Build successful. Running the server...\n"
    ./bin/server
else
    echo "[x] Build failed."
fi
