FROM golang:1.16

# Seems to be a Debian 10 base (is this the official golang image from Docker hub?)
# go binary is installed in /usr/local/go/bin/go

RUN apt-get update -y
RUN export GO111MODULE=on \
    && go get -v -u \
        golang.org/x/tools/gopls \
        github.com/uudashr/gopkgs/v2/cmd/gopkgs \
        github.com/ramya-rao-a/go-outline \
        github.com/go-delve/delve/cmd/dlv \
        honnef.co/go/tools/cmd/staticcheck

CMD [ "/bin/bash" ]
