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

    cargoHash = "sha256-RdkF6J7rIBNGjQHvDYuw+bBVPqP+TMylubILxV1pI+4=";

    buildInputs = [
      pkgs.openssl.dev
    ];
    nativeBuildInputs = [
      pkgs.pkg-config
    ];
  }
