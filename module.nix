inputs: {
  config,
  lib,
  pkgs,
  ...
}: let
  inherit (pkgs.stdenv.hostPlatform) system;
  cfg = config.services.lyricsapi;

  package = inputs.self.packages.${system}.default;
  inherit (lib) mkOption mkEnableOption types mkIf;
in {
  options.services.lyricsapi = {
    enable = mkEnableOption "Simple lyrics API";
    package = mkOption {
      type = types.package;
      default = package;
      example = package;
      description = "LyricsAPI package to use";
    };
  };
  config = mkIf cfg.enable {
    systemd.services.lyricsapi = {
      description = "Simple lyrics API";
      wantedBy = ["multi-user.target"];
      wants = ["network.target"];
      after = [
        "network-online.target"
        "NetworkManager.service"
        "systemd-resolved.service"
      ];
      serviceConfig = {
        ExecStart = ''${cfg.package}/bin/lyricsapi'';
        Restart = "always";
      };
    };
  };
}

