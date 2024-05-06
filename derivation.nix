{ lib, buildGoModule, fetchFromGitHub, ... }:

buildGoModule rec {
  pname = "tableize";
  version = "0.1.2";

  src = ./.;

  vendorHash = "sha256:vLm6ZQMw2TXbLqhcCCIRu6Wp9LSAVOTka4h94flkzEw=";

  meta = with lib; {
    description = "Convert JSON and YAML to a tab-delimited table";
    homepage = "https://github.com/onealexharms/tableize";
    license = licenses.publicDomain;
    platforms = platforms.all;
    maintainers = [ maintainers.eraserhd ];
  };
}
