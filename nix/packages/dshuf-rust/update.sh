#!/usr/bin/env bash
set -euo pipefail

(
	cd go
	cargo update
)

nix-update dshuf-rust --flake --version=skip
