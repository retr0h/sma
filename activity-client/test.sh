#!/usr/bin/env sh
set -e

echo "=== Add Records ==="
go run cmd/client/main.go --add "overhead press: 70lbs"
go run cmd/client/main.go --add "20 minute walk"

echo "=== Retrieve Records ==="
go run cmd/client/main.go --get 0 | grep "overhead press: 70lbs"
go run cmd/client/main.go --get 1 | grep "20 minute walk"
