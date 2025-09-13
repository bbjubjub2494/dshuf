{
  flake,
  pkgs,
}: let
  rust = builtins.path {path = "${flake}/rust";};
  testcases = builtins.path {path = "${flake}/testcases";};
in
  pkgs.rustPlatform.buildRustPackage {
    pname = "dshuf-rust";
    version = "unstable";

    srcs = [rust testcases];
    sourceRoot = "rust";

    cargoHash = "sha256-lvHApkdp1coOF57+OnKCzV9o+i7XmIcG7qYz9u/dhBg=";

    buildInputs = [
      pkgs.openssl.dev
    ];
    nativeBuildInputs = [
      pkgs.pkg-config
    ];
  }
