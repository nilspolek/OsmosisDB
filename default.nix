{ pkgs ? import <nixpkgs> {} }:

pkgs.stdenv.mkDerivation {
  name = "OsmosisDB";

  src = ./.;

  buildInputs = [ pkgs.go_1_23 pkgs.gnumake ];

  buildPhase = ''
    export GOCACHE=$(mktemp -d)

    # Create a Go project directory in the build environment
    mkdir -p $TMPDIR/src/github.com/nilspolek/osmosisdb
    cp -r ./* $TMPDIR/src/github.com/nilspolek/osmosisdb
    cd $TMPDIR/src/github.com/nilspolek/osmosisdb

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
