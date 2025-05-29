{
  description = "Flake to manage my Java workspace.";

  inputs.nixpkgs.url = "nixpkgs/nixpkgs-unstable";

  outputs = inputs: let
    system = "x86_64-linux";
    pkgs = inputs.nixpkgs.legacyPackages.${system};
  in {
    devShell.${system} = pkgs.mkShell rec {
      name = "java-shell";
      buildInputs = with pkgs; [
        jdk11
        dotnetCorePackages.dotnet_9.sdk
        gopls
        go
        rustc
        cargo
        clippy
        pkgs.rust-code-analysis
      ];

      shellHook = ''
        export JAVA_HOME=${pkgs.jdk11}
        PATH="${pkgs.jdk11}/bin:$PATH"
        echo "  Java flake   "
        echo " Get ready to type this script  "
        if command -v zsh > /dev/null 2>&1; then
            echo "Starting zsh, to exit you'll need to exit zsh first and then bash  "
            zsh
        fi;
      '';
    };
  };
}
