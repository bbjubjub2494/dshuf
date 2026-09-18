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

  cargoHash = "sha256-/V37wwjT4LlhKVu/QY5U4YOwEpe8TNq41OYKArM2OhA=";

  buildInputs = [
    pkgs.openssl.dev
  ];
  nativeBuildInputs = [
    pkgs.pkg-config
  ];
}
