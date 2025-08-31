#!/bin/bash

echo "> Formatting Go code..."
go fmt ./...

if [ $? -ne 0 ]; then
    echo "[x] Formatting failed, terminating."
    exit 1
fi

echo "> Compiling sever..."
go build -o ./bin/server ./cmd/server/server.go

if [ $? -eq 0 ]; then
    echo -e "> Build successful. Running the server...\n"
    ./bin/server
else
    echo "[x] Build failed."
fi
