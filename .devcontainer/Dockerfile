FROM golang:1.16

# This Dockerfile adds a non-root user with sudo access. Update the “remoteUser” property in
# devcontainer.json to use it. More info: https://aka.ms/vscode-remote/containers/non-root-user.
# FIXME:  This didn't work... am I using the wrong base image?
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Seems to be a Debian 10 base (is this the official golang image from Docker hub?)
# go binary is installed in /usr/local/go/bin/go
RUN apt-get update -y

# Use Docker within a Codespace
# https://docs.microsoft.com/en-us/visualstudio/codespaces/reference/configuring#use-docker-within-a-codespace
RUN groupadd -g 800 docker
RUN usermod -a -G docker vscode

RUN export GO111MODULE=on \
    && go get -v -u \
        golang.org/x/tools/gopls \
        github.com/uudashr/gopkgs/v2/cmd/gopkgs \
        github.com/ramya-rao-a/go-outline \
        github.com/go-delve/delve/cmd/dlv \
        honnef.co/go/tools/cmd/staticcheck

CMD [ "/bin/bash" ]
