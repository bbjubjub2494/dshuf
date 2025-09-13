{ flake, system, ... }:

with flake.packages.${system};
internal-integration.overrideAttrs (_: {
  nativeCheckInputs = [ dshuf-go ];
  checkFlags = [
    "-test.run"
    "/impl=go"
  ];

  doCheck = true;
})
