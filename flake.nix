# SPDX-FileCopyrightText: Winni Neessen <wn@neessen.dev>
#
# SPDX-License-Identifier: MIT

{
  description = "apg-go binary release package";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};

      version = "1.2.0";
      pkgrel = "1";
    in {
      packages.${system}.default = pkgs.stdenvNoCC.mkDerivation {
        pname = "apg-go";
        inherit version;

        src = pkgs.fetchurl {
          url = "https://github.com/wneessen/apg-go/releases/download/v${version}/apg-go-${version}-${pkgrel}-x86_64.pkg.tar.zst";
          hash = "sha256-Hoqr8Ndm6MpM5f6+p1hQ0gswbzL3Tsk0DaL/XdkDZXk=";
        };

        nativeBuildInputs = [
          pkgs.zstd
          pkgs.gnutar
        ];

        unpackPhase = ''
          tar --use-compress-program=zstd -xf "$src"
        '';

        installPhase = ''
          runHook preInstall

          install -Dm755 usr/bin/apg "$out/bin/apg"

          if [ -f usr/share/licenses/apg-go/LICENSE ]; then
            install -Dm644 usr/share/licenses/apg-go/LICENSE "$out/share/licenses/apg-go/LICENSE"
          fi

          runHook postInstall
        '';

        meta = with pkgs.lib; {
          description = "A modern Automated Password Generator clone";
          homepage = "https://github.com/wneessen/apg-go";
          license = licenses.mit;
          mainProgram = "apg";
          platforms = [ "x86_64-linux" ];
          sourceProvenance = with sourceTypes; [ binaryNativeCode ];
        };
      };

      apps.${system}.default = {
        type = "app";
        program = "${self.packages.${system}.default}/bin/apg";
      };
    };
}
