{
  description = "Temporal intro workshop";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    nur = {
      url = "github:sagikazarmark/nur-packages";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, flake-utils, nur, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;

          overlays = [
            (final: prev: {
              reveal-md = nur.packages.${system}.reveal-md;
              decktape = nur.packages.${system}.decktape;
            })
          ];
        };
      in
      rec
      {
        devShells = {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
              git
              gnumake
              reveal-md
              go

              temporalite
              temporal-cli
            ];
          };

          ci = pkgs.mkShell {
            buildInputs = with pkgs; [
              git
              gnumake
              reveal-md
              go
            ];
          };
        };
      });
}
