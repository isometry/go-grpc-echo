#!/usr/bin/env sh
set -ex

echo "=== Test Reflection API ==="
grpcurl -plaintext localhost:8080 describe

echo "=== Test Echo Message ==="

grpcurl -plaintext -d  '{"message": "Hello, World"}' localhost:8080 api.v1.Echo/Echo

echo "Success"
