# SPDX-FileCopyrightText: 2021-2024 Winni Neessen <wn@neessen.dev>
#
# SPDX-License-Identifier: MIT

FROM golang:latest@sha256:b1e05e2c918f52c59d39ce7d5844f73b2f4511f7734add8bb98c9ecdd4443365 AS builder
RUN mkdir /builddir
ADD . /builddir/
WORKDIR /builddir
RUN go mod tidy
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-w -s -extldflags "-static"' -o apg-go \
    github.com/wneessen/apg-go/cmd/apg

## Create scratch image
FROM scratch
LABEL maintainer="wn@neessen.dev"
COPY ["docker-files/passwd", "/etc/passwd"]
COPY ["docker-files/group", "/etc/group"]
COPY --from=builder ["/etc/ssl/certs/ca-certificates.crt", "/etc/ssl/cert.pem"]
COPY --chown=apg-go ["LICENSE", "/apg-go/LICENSE"]
COPY --chown=apg-go ["README.md", "/apg-go/README.md"]
COPY --from=builder --chown=apg-go --chmod=555 ["/builddir/apg-go", "/apg-go/apg-go"]
WORKDIR /apg-go
USER apg-go
ENTRYPOINT ["/apg-go/apg-go"]