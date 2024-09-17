{ pkgs ? import <nixpkgs> {}}:
pkgs.mkShell {
    buildInputs = with pkgs; [
        git
        neovim
        tmux
        htop
        neofetch
        tree
        curl
        wget
        go
        gnumake
    ];
}
