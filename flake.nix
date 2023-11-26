{
  description = "generate reproducible random permutations using public randomness";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  inputs.hercules-ci-effects.url = "github:hercules-ci/hercules-ci-effects";

  outputs = inputs @ {flake-parts, ...}:
    flake-parts.lib.mkFlake {inherit inputs;} {
      imports = [
        inputs.hercules-ci-effects.flakeModule
      ];
      systems = ["x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin"];
      # NOTE: I do not have runners for darwin
      herculesCI.ciSystems = ["x86_64-linux" "aarch64-linux"];

      hercules-ci.flake-update.enable = true;
      hercules-ci.flake-update.when.dayOfWeek = "Sat";

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
