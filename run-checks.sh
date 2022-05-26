#!/bin/bash

set -e

cd "$(dirname "$0")"
BASEDIR=$(pwd)

if ! hash go; then
  echo "please install go and try again"
  exit 1
fi
if ! hash staticcheck; then
  echo "installing staticcheck"
  if ! go install honnef.co/go/tools/cmd/staticcheck@2022.1; then
    echo "failed to install staticcheck"
    exit 1
  fi
fi

for package in common influx2otel otel2influx tests-integration; do
  echo checking ${package}
  cd ${BASEDIR}/${package}
  go mod tidy
  if ! git diff --exit-code -- go.mod go.sum; then
    fail=1
  fi
  if ! go build; then
    fail=1
  fi
  if ! go test; then
    fail=1
  fi
  if [[ -n $(gofmt -s -l . | head -n 1) ]]; then
    fail=1
    gofmt -s -d .
  fi
  if ! go vet; then
    fail=1
  fi
  if ! staticcheck -f stylish; then
    fail=1
  fi
done

echo

if [ -n "$fail" ]; then
  echo "at least one check failed"
  exit 1
else
  echo "all checks OK"
fi