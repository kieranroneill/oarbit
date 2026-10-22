#!/usr/bin/env bash

# Public: Stops the background Ollama process if it is running.
#
# Environment Variables
#
# OLLAMA_PID - The Ollama process ID.
function _cleanup() {
  if kill -0 "${OLLAMA_PID:-0}" 2>/dev/null; then
    kill "${OLLAMA_PID}" || true
    wait "${OLLAMA_PID}" || true
  fi
}
