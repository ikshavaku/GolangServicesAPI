#!/bin/sh
set -e

echo "Running migrations..."
make goose-up

echo "Starting app..."
exec ./build

