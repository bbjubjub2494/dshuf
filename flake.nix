{
  description = "generate reproducible random permutations using public randomness";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = inputs @ {flake-parts, ...}:
    flake-parts.lib.mkFlake {inherit inputs;} {
      systems = ["x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin"];

      perSystem = {pkgs, ...}: let
        implementations.dshuf-go = pkgs.buildGoModule rec {
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
          vendorHash = "sha256-GbicbJmyrQRiTW4byF3lkM01fMhclkIdSgub/C4rzQw=";
        };

        implementations.dshuf-rust = with pkgs;
          rustPlatform.buildRustPackage {
            pname = "dshuf-rust";
            version = "unstable";
            srcs = [./rust ./testcases];
            sourceRoot = "rust";
            useFetchCargoVendor = true;
            cargoHash = "sha256-RdkF6J7rIBNGjQHvDYuw+bBVPqP+TMylubILxV1pI+4=";

            buildInputs = [
              openssl.dev
            ];
            nativeBuildInputs = [
              pkg-config
            ];
          };
      in {
        packages = implementations;

        devShells.default = with pkgs;
          mkShell {
            buildInputs = [
              go
              cargo

              # To update hashes in nix files
              nix-update

              # To compile curl under Rust
              openssl.dev
              pkg-config
            ];
          };

        formatter = pkgs.alejandra;

        checks = pkgs.lib.genAttrs ["go" "rust"] (impl:
          pkgs.buildGoModule {
            name = "dshuf-${impl}-check";
            src = ./integration;

            vendorHash = "sha256-1p3dCLLo+MTPxf/Y3zjxTagUi+tq7nZSj4ZB/aakJGY=";
            nativeCheckInputs = [implementations."dshuf-${impl}"];
            checkFlags = ["-test.run" "/impl=${impl}"];
          });
      };
    };
}
