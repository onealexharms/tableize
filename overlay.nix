self: super: {
  tableize = super.callPackage ./derivation.nix {
    fetchFromGitHub = _: ./.;
  };
}
