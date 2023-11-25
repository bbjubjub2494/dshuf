{
  description = "generate reproducible random permutations using public randomness";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = inputs @ {flake-parts, ...}:
    flake-parts.lib.mkFlake {inherit inputs;} {
      systems = ["x86_64-linux" "aarch64-linux"];
      perSystem = {pkgs, ...}: {
        packages.dshuf-go = pkgs.buildGoModule rec {
          pname = "dshuf-go";
          version = "unstable";
          # workaround: src is a mandatory argument
          src = builtins.path {
            path = ./go;
            # workaround: a source called go is confused with GOPATH
            name = "src";
          };
          srcs = [
            src
            ./testcases
          ];
          sourceRoot = "src";
          vendorHash = "sha256-k6fgQ022tmhFeNFo5x/7ZK4Pv6bOqh4eI4mKQ5gje9g=";
        };

        packages.dshuf-rust = with pkgs;
          rustPlatform.buildRustPackage {
            pname = "dshuf-rust";
            version = "unstable";
            srcs = [./rust ./testcases];
            sourceRoot = "rust";
            cargoHash = "sha256-9SoJTeGY7zrTipl4p/ZdJbB3BXERXzECx5AYI+1Y8eI=";

            buildInputs = [
              openssl.dev
            ];
            nativeBuildInputs = [
              pkg-config
            ];
          };

        devShells.default = with pkgs;
          mkShell {
            buildInputs = [
              # To compile curl under Rust
              openssl.dev
              pkg-config
            ];
          };
        formatter = pkgs.alejandra;
      };
    };
}
