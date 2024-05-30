#!/bin/bash
VERSION=$(git describe --tags)
BUILD_TIME=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
IMAGE_NAME=$(basename $PWD)

echo "VERSION - ${VERSION}"
echo "BUILD_TIME - ${BUILD_TIME}"
echo "IMAGE_NAME - ${IMAGE_NAME}"

docker build --build-arg VERSION=${VERSION} --build-arg BUILD_TIME=${BUILD_TIME} --tag ${IMAGE_NAME}:${VERSION} .

echo -e "try it out...\ndocker run --rm -it -p 8080:8080 ${IMAGE_NAME}:${VERSION}"