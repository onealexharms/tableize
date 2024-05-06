{
  description = "TODO: fill me in";
  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = { self, nixpkgs, flake-utils }:
    (flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        tableize = pkgs.callPackage ./derivation.nix {};
      in {
        packages = {
          default = tableize;
          inherit tableize;
        };
        checks = {
          test = pkgs.runCommandNoCC "tableize-test" {} ''
            mkdir -p $out
            : ${tableize}
          '';
        };
    })) // {
      overlays.default = final: prev: {
        tableize = prev.callPackage ./derivation.nix {};
      };
    };
}
