# SPDX-FileCopyrightText: 2021-2024 Winni Neessen <wn@neessen.dev>
#
# SPDX-License-Identifier: MIT

FROM golang:latest@sha256:927112936d6b496ed95f55f362cc09da6e3e624ef868814c56d55bd7323e0959 AS builder
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