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

  cargoHash = "sha256-nD/FGlZQ1rNBjxauxfJuUWQ+rDJi9Zqp+vaIJLouW+s=";

  buildInputs = [
    pkgs.openssl.dev
  ];
  nativeBuildInputs = [
    pkgs.pkg-config
  ];
}
