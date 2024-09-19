{ pkgs ? import <nixpkgs> {} }:

pkgs.stdenv.mkDerivation {
  name = "OsmosisDB";

  src = ./.;

  buildInputs = [ pkgs.go_1_23 pkgs.gnumake ];

  buildPhase = ''
    export GOCACHE=$(mktemp -d)

    # Build the Go project
    make build
  '';

  installPhase = ''
    mkdir -p $out/bin
    cp ./bin/osmosis $out/bin/
  '';

  meta = {
    description = "OsmosisDB is a simple key-value store written in Go.";
    license = pkgs.lib.licenses.mit;
  };
}
