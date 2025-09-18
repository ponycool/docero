FROM 10.0.0.20/ponycool/alpine:3.22
LABEL maintainer="Docero DOCKER MAINTAINER <pony@ponycool.com>"

ARG APP_PATH=/opt/docero
# 设定环境变量，用于支持中文显示
ENV LANG=zh_CN.UTF-8

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    # 安装LibreOffice
    libreoffice \
    # 中文字体支持
    font-noto-cjk \
    font-dejavu \
    # 语言包支持
    libreoffice-lang-zh_cn \
    openjdk21-jre-headless && \
    # 清理不必要的文件
    rm -rf /var/cache/apk/*

RUN mkdir -p "${APP_PATH}" && \
    mkdir -p "${APP_PATH}/logs" && \
    mkdir -p "${APP_PATH}/uploads" && \
    mkdir -p "${APP_PATH}/converted" && \
    chmod -R 777 "${APP_PATH}/logs" && \
    chmod -R 777 "${APP_PATH}/uploads" && \
    chmod -R 777 "${APP_PATH}/converted"

COPY docero "${APP_PATH}"
COPY web "${APP_PATH}"
COPY config "${APP_PATH}"

WORKDIR "${APP_PATH}"

EXPOSE 8080

CMD ["/opt/docero/docero"]