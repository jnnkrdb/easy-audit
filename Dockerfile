# ---------------------------------------------------------------------------------------------- Golang
# Building the go binary
FROM golang:1.26.3-alpine AS builder
RUN apk add --no-cache --update gcc musl-dev
#musl-dev util-linux-dev
RUN mkdir -p /github.com/jnnkrdb/easy-audit
WORKDIR /github.com/jnnkrdb/easy-audit
# copy the code files
COPY . ./
# set env vars
ENV CGO_ENABLED=1
# ENV GOARCH=amd64
ENV GOOS=linux
# START BUILD
RUN go mod download 
RUN go build -a -o /eactl /github.com/jnnkrdb/easy-audit/cmd/eactl/main.go
# ---------------------------------------------------------------------------------------------- Final Alpine
FROM alpine:3.24.1
LABEL org.opencontainers.image.source="https://github.com/jnnkrdb/easy-audit"
LABEL org.opencontainers.image.author="jnnkrdb"
LABEL org.opencontainers.image.description="Audit Store in Go."
WORKDIR /
# install neccessary binaries
RUN apk add --no-cache --update openssl
# create user with home dir
RUN mkdir /var/easy-audit 
RUN addgroup -S easy-audit && adduser -S easy-audit -H -h /var/easy-audit -s /bin/sh -G easy-audit -u 3453
RUN chmod 700 -R /var/easy-audit && chown easy-audit:easy-audit -R /var/easy-audit
VOLUME /var/easy-audit
# Copy Binary
COPY --from=builder --chmod=755 --chown=easy-audit:easy-audit /eactl /usr/local/bin/eactl
# change to required user
USER easy-audit:easy-audit
# set the entrypoints
ENTRYPOINT ["/bin/sh", "-c"]
CMD [ "eactl", "serve"]