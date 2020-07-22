#!/usr/bin/env bash

set -u

version=v0.1.5

APP=jarvis

DOWNLOAD_URL="https://github.com/glepnir/jarvis/releases/download/$version"

exists() {
  command -v "$1" >/dev/null 2>&1
}

download() {
  local from=$1
  local to=$2
  if exists "curl"; then
    curl -fLo "$to" "$from"
  elif exists 'wget'; then
    wget --output-document="$to" "$from"
  else
    echo 'curl or wget is required'
    exit 1
  fi
}

try_download() {
  local asset=$1
  if [ -z "${TMPDIR+x}" ]; then
    rm -f /usr/local/bin/$APP
    download "$DOWNLOAD_URL/$asset" bin/$APP
  else
    local temp=${TMPDIR}/jarvis
    download "$DOWNLOAD_URL/$asset" "$temp"
    mv "$temp" /usr/local/bin/$APP
  fi
  chmod a+x "/usr/local/bin/bin/$APP"
}

main() {
  osname=$(uname -sm)
  case "${osname}" in
      "Linux x86_64")
        try_download "$APP"-x86_64-linux ;;
      "Darwin x86_64")
        try_download "$APP"-x86_64-darwin ;;
      *)
        echo "No prebuilt jarvis binary available for ${osname}."
        exit 1
        ;;
  esac
}

main
