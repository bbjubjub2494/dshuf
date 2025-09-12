#!/usr/bin/env bash
set -euo pipefail

(
	cd integration
	go get -u ./...
)

nix-update dshuf-go --flake --version=skip
