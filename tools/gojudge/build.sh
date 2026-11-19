#!/bin/sh
# Build the browser Go judge: assets/go-judge.wasm.gz (gzipped, the worker unpacks it) + assets/wasm_exec.js.
# Needs Go 1.22+. Run from anywhere: sh tools/gojudge/build.sh
set -e
cd "$(dirname "$0")"
GOOS=js GOARCH=wasm go build -trimpath -ldflags="-s -w" -o go-judge.wasm .
gzip -9 -n -c go-judge.wasm > ../../assets/go-judge.wasm.gz
rm go-judge.wasm
root="$(go env GOROOT)"
if [ -f "$root/lib/wasm/wasm_exec.js" ]; then cp "$root/lib/wasm/wasm_exec.js" ../../assets/; else cp "$root/misc/wasm/wasm_exec.js" ../../assets/; fi
ls -l ../../assets/go-judge.wasm.gz
