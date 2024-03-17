{
  config,
  pkgs,
  lib,
  ...
}: let
  cfg = config.givc.appvm;
  givc-agent = pkgs.callPackage ../packages/givc-agent.nix {};
in
  with lib; {

    options.givc.appvm = {
      enable = mkEnableOption "Enable givc-appvm module.";

      name = mkOption {
        description = "Service name";
        type = types.str;
        default = "localhost";
      };

      addr = mkOption {
        description = "Address";
        type = types.str;
        default = "127.0.0.1";
      };

      port = mkOption {
        description = "Port";
        type = types.str;
        default = "9000";
      };

      protocol = mkOption {
        description = "Protocol";
        type = types.str;
        default = "tcp";
      };

      applications = mkOption {
        description = ''
        applications
        '';
        type = types.str;
        default = "";
      };

      admin = mkOption {
        description = ''Admin server configuration.
        '';
        type = with types; submodule {
          options = {
            enable = mkEnableOption "Admin module";
            name = mkOption {
              description = "Hostname of admin server";
              type = types.str;
            };

            addr = mkOption {
              description = "Address of admin server";
              type = types.str;
            };

            port = mkOption {
              description = "Port of admin server";
              type = types.str;
            };

            protocol = mkOption {
              description = "Protocol of admin server";
              type = types.str;
            };
          };
        };
        default = {
          enable = true;
          name = "localhost";
          addr = "127.0.0.1";
          port = "9001";
          protocol = "tcp";
        };
      };

      tls = mkOption {
        description = ''TLS options for gRPC connections.
        It is enabled by default to discourage unprotected connections, and requires paths to certificates and key being set.
        If you really want to disable it use 'tls.enable = false;' .
        '';
        type = with types; submodule {
          options = {
            enable = mkEnableOption "Enable TLS";
            caCertPath = mkOption {
              description = "Path to the CA certificate file.";
              type = str;
              default = "";
            };
            certPath = mkOption {
              description = "Path to the service certificate file.";
              type = str;
              default = "";
            };
            keyPath = mkOption {
              description = "Path to the service key file.";
              type = str;
              default = "";
            };
          };
        };
        default = {
          enable = true;
          caCertPath = "";
          certPath = "";
          keyPath = "";
        };
      };

    };

    config = mkIf cfg.enable {

      assertions = [
        {
          assertion = cfg.applications != "";
          message = "A list of services (or targets) is required for this module to run.";
        }
        {
          assertion = !(cfg.tls.enable && (
            cfg.tls.caCertPath == "" || cfg.tls.certPath == "" || cfg.tls.keyPath == ""
          ));
          message = "The TLS option requires paths' to CA certificate, service certificate, and service key.";
        }
      ];

      systemd.user.services."givc-${cfg.name}" = {
        description = "GIVC remote service manager for application VMs.";
        enable = true;
        after = [ "sockets.target" ];
        wantedBy = [ "default.target" ];
        serviceConfig = {
          Type = "exec";
          ExecStart = "${givc-agent}/bin/givc-agent";
          Restart = "always";
          RestartSec = 1;
        };
        environment = {
          "NAME" = "${cfg.name}";
          "ADDR" = "${cfg.addr}";
          "PORT" = "${cfg.port}";
          "PROTO" = "${cfg.protocol}";
          "TLS" = "${trivial.boolToString cfg.tls.enable}";
          "PARENT" = "microvm@${cfg.name}-vm.service";
          "APPLICATIONS" = "${cfg.applications}";
          "ADMIN_SERVER_NAME" = "admin";
          "ADMIN_SERVER_ADDR" = "${cfg.admin.addr}";
          "ADMIN_SERVER_PORT" = "${cfg.admin.port}";
          "ADMIN_SERVER_PROTO" = "${cfg.admin.protocol}";
        } // attrsets.optionalAttrs cfg.tls.enable {
          "CA_CERT" = "${cfg.tls.caCertPath}";
          "HOST_CERT" = "${cfg.tls.certPath}";
          "HOST_KEY" = "${cfg.tls.keyPath}";
        };
      };
      networking.firewall.allowedTCPPorts =
        let
          port = lib.strings.toInt(cfg.port);
        in
          [port];
    };

}