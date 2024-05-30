#!/bin/bash
VERSION=$(git describe --tags)
BUILD_TIME=$(date -u +'%y-%m-%dT%H:%M:%SZ')

echo "Version - ${VERSION}"
echo "BUILD_TIME - ${BUILD_TIME}"
