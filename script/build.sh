#!/usr/bin/env sh

IMAGE=docero
VERSION=latest

if [ -n "$1" ]; then
  VERSION=$1
fi

echo $VERSION
# 启用buildx插件，适用于v19.03+
docker buildx create --use --name larger_log --node larger_log0 --driver-opt env.BUILDKIT_STEP_LOG_MAX_SIZE=10485760

# 检查当前的构建器实例
docker buildx inspect larger_log --bootstrap

# 编译生成
docker buildx build -t 10.0.0.20/kuma/${IMAGE}:${VERSION} --platform=linux/amd64 --push ./