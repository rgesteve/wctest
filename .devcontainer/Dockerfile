FROM golang:1.16

# This Dockerfile adds a non-root user with sudo access. Update the “remoteUser” property in
# devcontainer.json to use it. More info: https://aka.ms/vscode-remote/containers/non-root-user.
# FIXME:  This didn't work... am I using the wrong base image?
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Seems to be a Debian 10 base (is this the official golang image from Docker hub?)
# go binary is installed in /usr/local/go/bin/go
# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# Use Docker within a Codespace
# https://docs.microsoft.com/en-us/visualstudio/codespaces/reference/configuring#use-docker-within-a-codespace
RUN groupadd -g 800 docker
RUN usermod -a -G docker vscode

# `make` may be already installed
RUN apt-get install -y make protobuf-compiler

RUN export GO111MODULE=on \
    && go get -v -u \
        golang.org/x/tools/gopls \
        github.com/uudashr/gopkgs/v2/cmd/gopkgs \
        github.com/ramya-rao-a/go-outline \
        github.com/go-delve/delve/cmd/dlv \
        honnef.co/go/tools/cmd/staticcheck

USER $USERNAME

