#!/usr/bin/env bash
set -euo pipefail

(
	cd go
	go get -u ./...
)

nix-update dshuf-go --flake --version=skip
