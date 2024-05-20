{buildGoModule}:
buildGoModule {
  pname = "lyricsapi";
  version = "0.0.1";

  src = ./.;

  vendorHash = "sha256-Kqg9kCmmOcs8O6baZl+F8w8MSP5kCxRhNqqKiRcyN+s=";

  ldflags = ["-s" "-w"];
}
