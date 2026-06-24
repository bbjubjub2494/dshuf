{ flake, pkgs }:

let
  rust = builtins.path { path = "${flake}/rust"; };
  testcases = builtins.path { path = "${flake}/testcases"; };
in
pkgs.rustPlatform.buildRustPackage {
  pname = "dshuf-rust";
  version = "unstable";

  srcs = [
    rust
    testcases
  ];
  sourceRoot = "rust";

  cargoHash = "sha256-p3T5BqRaiysPmi7s4f21nJWZNAu/mV+1kX9Ik9MZz3Y=";

  buildInputs = [
    pkgs.openssl.dev
  ];
  nativeBuildInputs = [
    pkgs.pkg-config
  ];
}
