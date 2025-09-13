#!/usr/bin/env bash
set -euo pipefail

(
	cd rust
	cargo update
)

nix-update dshuf-rust --flake --version=skip
