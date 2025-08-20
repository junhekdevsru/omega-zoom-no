#!/usr/bin/env bash
set -euo pipefail

if ! command -v buf >/dev/null 2>&1; then
  echo "buf is required. Install from https://buf.build/docs/installation" >&2
  exit 1
fi

buf generate
echo "✅ proto generated"