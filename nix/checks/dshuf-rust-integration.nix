{
  flake,
  system,
  ...
}:
with flake.packages.${system};
  internal-integration.overrideAttrs (_: {
    nativeCheckInputs = [dshuf-rust];
    checkFlags = ["-test.run" "/impl=rust"];
  })
