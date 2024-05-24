# Use the official vscode dev container image based on Ubuntu 22.04
ARG VARIANT="jammy"
FROM mcr.microsoft.com/vscode/devcontainers/base:0-${VARIANT}

# Set the maintainer label
LABEL maintainer="pyy0715@gmail.com"

# Set the timezone
ENV TZ=Asia/Seoul

# Install necessary packages
RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
    apt-get -y install --no-install-recommends locales tzdata vim python3-pip && \
    rm -rf /var/lib/apt/lists/*

# Generate and set the locale
RUN locale-gen ko_KR.UTF-8 && update-locale LANG=ko_KR.UTF-8

# Set the user
USER vscode
