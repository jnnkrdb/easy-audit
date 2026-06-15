# ---------------------------------------------------------------------------------------------- Golang
# Building the go binary
FROM golang:1.26.3 AS builder
RUN mkdir -p /github.com/jnnkrdb/easy-audit
WORKDIR /github.com/jnnkrdb/easy-audit
# copy the code files
COPY . ./
# set env vars
ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux
# START BUILD
RUN go mod download && go build -o /easy-audit /github.com/jnnkrdb/easy-audit/cmd/main.go
# ---------------------------------------------------------------------------------------------- Final Alpine
FROM alpine:3.22.0
LABEL org.opencontainers.image.source="https://github.com/jnnkrdb/easy-audit"
LABEL org.opencontainers.image.author="jnnkrdb"
LABEL org.opencontainers.image.description="Audit Store in Go."
WORKDIR /
# install neccessary binaries
RUN apk add --no-cache --update openssl
# Copy the Directory Contents
RUN mkdir /opt/easy-audit &&\
    mkdir /opt/easy-audit/config &&\
    mkdir /opt/easy-audit/data &&\
    mkdir /opt/easy-audit/logs &&\
    mkdir /opt/easy-audit/certs &&\
    mkdir /opt/easy-audit/keys

# create user with home dir
RUN addgroup -S easy-audit && adduser -S easy-audit -H -h /opt/easy-audit -s /bin/sh -G easy-audit -u 3453
# Copy Binary
COPY --from=builder /easy-audit /usr/local/bin/easy-audit
RUN chmod 700 /usr/local/bin/easy-audit &&\
    chmod 700 -R /opt/easy-audit &&\
    chown easy-audit:easy-audit /usr/local/bin/easy-audit &&\
    chown easy-audit:easy-audit -R /opt/easy-audit
# change to required user
USER easy-audit:easy-audit
# set env vars for the binary
ENV EASY_AUDIT_HOME="/opt/easy-audit"
ENV EASY_AUDIT_BINARY_PATH="/usr/local/bin/easy-audit"
# set the entrypoints
ENTRYPOINT ["/bin/sh", "-c"]
CMD [ "easy-audit" ]
ARGS ["--server"]