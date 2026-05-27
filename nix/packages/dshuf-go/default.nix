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

  vendorHash = "sha256-7eFIATHpLq8Htul1gfYI7iBK8t1fxqEDOlt36IB49XA=";
}
