#!/usr/bin/env bash
set -euo pipefail

echo "⏳ waiting for Postgres on localhost:5432 ..."
for i in {1..60}; do
  if nc -z localhost 5432 >/dev/null 2>&1; then
    echo "✅ Postgres is up"
    exit 0
  fi
  sleep 1
done

echo "❌ timeout waiting Postgres" >&2
exit 1