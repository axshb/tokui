{
  description = "tokui";

  inputs = {
    # using a direct tarball url to bypass github api 504 errors
    nixpkgs.url = "https://github.com/NixOS/nixpkgs/archive/nixos-unstable.tar.gz";
  };

  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux"; 
      pkgs = import nixpkgs { inherit system; };
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        buildInputs = with pkgs; [
          go
          gopls
        ];

        shellHook = ''
          echo "🫧 Bubble Tea Environment Loaded"
          go version
        '';
      };
    };
}