{
  pkgs,
  flake,
}:
pkgs.buildGoModule {
  name = "dshuf-integration";
  src = "${flake}/integration";

  vendorHash = "sha256-1p3dCLLo+MTPxf/Y3zjxTagUi+tq7nZSj4ZB/aakJGY=";
  doCheck = false; # checks are run in nix/checks
}
