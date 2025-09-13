{
  description = "generate reproducible random permutations using public randomness";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs?ref=nixos-unstable";
    blueprint.url = "github:numtide/blueprint";
    blueprint.inputs.nixpkgs.follows = "nixpkgs";
  };

  nixConfig = {
    extra-substituters = "https://dshuf.cachix.org";
    extra-trusted-public-keys = "dshuf.cachix.org-1:DdwhVB1EbEOhBfcGVRjwJYv+SsA9Iba8u290OoBeCps=";
  };

  outputs =
    inputs:
    inputs.blueprint {
      inherit inputs;
      prefix = "nix";
    };
}
