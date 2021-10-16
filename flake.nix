{
  description = "My workshop template";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    reveal-md-flake = {
      url = "github:sagikazarmark/reveal-md-flake";
      inputs = {
        nixpkgs.follows = "nixpkgs";
        flake-utils.follows = "flake-utils";
      };
    };
    flake-compat = {
      url = "github:edolstra/flake-compat";
      flake = false;
    };
  };

  outputs = { self, nixpkgs, flake-utils, reveal-md-flake, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ reveal-md-flake.overlay ];
        };
      in {
        devShell = pkgs.mkShell {
          buildInputs = with pkgs; [ git gnumake reveal-md go ];
        };
      });
}
