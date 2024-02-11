{ pkgs ? import <nixpkgs> { } }:

let
  config = builtins.fromJSON (builtins.readFile ./.config.json);
in
pkgs.mkShell {
  buildInputs = [
    pkgs.go
  ];

  MAIL_FROM = config.mail_from;
  MAIL_PASSWORD = config.mail_password;
  MAIL_SMTP_HOST = config.mail_smtp_host;
  MAIL_SMTP_PORT = config.mail_smtp_port;

  MAIL_TO = config.mail_to;
}
