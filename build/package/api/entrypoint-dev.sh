#!/bin/sh

. ./_set_vars.sh

# Public: Injects the version and runs the Go core server using CompileDaemon.
#
# $1 - [optional] A version to inject, otherwise the version from the VERSION file is read.
#
# Examples
#
#   ./bin/dev.sh # reads the version in the VERSION file
#   ./bin/dev.sh "1.2.3"
#
# Returns exit code 0.
main() {
  _set_vars

  version=$(cat ./VERSION)

  # if the version argument exists, use it
  if [ -n "${1-}" ]; then
    version=$1
  fi

  export VERSION="$version"

  printf '%s starting api server...\n' "$INFO_PREFIX"

  CompileDaemon \
    -build="go build -o .bin/api cmd/api/main.go" \
    -command=".bin/api"

  exit 0
}

# and so, it begins...
main "$@"
