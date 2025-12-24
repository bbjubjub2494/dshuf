{ flake, pkgs }:

let
  src = builtins.path {
    path = "${flake}/go";
    # workaround: a source called go gets confused for GOPATH
    name = "src";
  };
  testcases = builtins.path { path = "${flake}/testcases"; };
in
pkgs.buildGoModule rec {
  pname = "dshuf-go";
  version = "unstable";

  srcs = [
    src
    testcases
  ];
  sourceRoot = "src";
  # workaround: src is a mandatory argument
  inherit src;

  vendorHash = "sha256-lG3VB4GIz0vWNfxG7fXsRaWK1VmgJxk1SQqP7IH8afE=";
}
