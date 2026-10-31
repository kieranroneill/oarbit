#!/usr/bin/env bash

set -euo pipefail
set -a
source ./configs/api/.env.test
source ./configs/postgres/.env.test
source ./configs/redis/.env.test
source ./configs/shared/.env.test
set +a

__DOCKER_PROJECT__=oarbit-test
SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}/utilities/_set_vars.sh"

# Private: Stops the docker services.
function _cleanup() {
  local exit_code=$?

  trap - EXIT

  docker compose \
    -f deployments/compose.test.yml \
    -p "${__DOCKER_PROJECT__}" \
    down \
    --volumes || true

  exit "${exit_code}"
}

# Public: Starts the necessary Docker services and runs integration tests.
#
# Examples
#   ./test_integration.sh
function main {
  _set_vars

  trap _cleanup EXIT

  # start the services
  printf "%b starting docker services... \n" "${INFO_PREFIX}"
  docker compose \
	 	-f ./deployments/compose.test.yml \
	 	-p "${__DOCKER_PROJECT__}" \
	 	--env-file ./configs/postgres/.env.test \
	 	--env-file ./configs/redis/.env.test \
		up \
		-d \
		--wait postgres redis

  # run
	printf "%b running migrations... \n" "${INFO_PREFIX}"
  docker compose \
    -f deployments/compose.test.yml \
    -p "${__DOCKER_PROJECT__}" \
	 	--env-file ./configs/postgres/.env.test \
    run \
    --build \
    --rm migrations_up

  # run tests
  printf "%b running integration tests... \n" "${INFO_PREFIX}"
  go test -tags=integration -count=1 -v ./...
}

# and so, it begins...
main "${@}"
