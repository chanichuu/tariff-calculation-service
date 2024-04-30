#!/bin/sh
set -e

echo ""
echo "PACKAGING Golang binaries..."
for model in cmd/*; do
    # If directory
    if [ -d "$model" ]; then
        echo ... cmd $(basename $model)
        # If a main.go file exists
        if [ -f "./cmd/$(basename $model)/main.go" ]; then
            GOOS=linux go build -o bootstrap ./cmd/$(basename $model)
            mkdir -p bin/$(basename $model) && zip bin/$(basename $model)/$(basename $model).zip bootstrap

            # For Windows
            # mkdir -p bin/$(basename $model) && 7z a -tzip bin/$(basename $model)/$(basename $model).zip bootstrap

        else
            echo "--- No main.go file found (ignoring) ---"
        fi
    fi
done