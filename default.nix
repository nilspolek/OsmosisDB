{ pkgs ? import <nixpkgs> {} }:

pkgs.stdenv.mkDerivation {
  name = "OsmosisDB";

  src = ./.;

  buildInputs = [ pkgs.go pkgs.gnumake ];

  buildPhase = ''
    export GOCACHE=$(mktemp -d)

    # Create a Go project directory in the build environment
    mkdir -p $TMPDIR/src/
    cp -r ./* $TMPDIR/src/
    cd $TMPDIR/src/

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
