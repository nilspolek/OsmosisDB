{ pkgs ? import <nixpkgs> {}}:
pkgs.mkShell {
    packages = with pkgs; [
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
