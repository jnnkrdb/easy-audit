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
RUN mkdir -p /ea-bin
RUN go mod download 
RUN go build -a -o /ea-bin/easy-audit /github.com/jnnkrdb/easy-audit/cmd/easy-audit/main.go
RUN go build -a -o /ea-bin/eactl /github.com/jnnkrdb/easy-audit/cmd/eactl/main.go
# ---------------------------------------------------------------------------------------------- Final Alpine
FROM alpine:3.24.1
LABEL org.opencontainers.image.source="https://github.com/jnnkrdb/easy-audit"
LABEL org.opencontainers.image.author="jnnkrdb"
LABEL org.opencontainers.image.description="Audit Store in Go."
WORKDIR /
# install neccessary binaries
RUN apk add --no-cache --update openssl
# Copy the Directory Contents
RUN mkdir /opt/easy-audit &&\
    mkdir /opt/easy-audit/bin &&\
    mkdir /opt/easy-audit/config &&\
    mkdir /opt/easy-audit/data &&\
    mkdir /opt/easy-audit/home
# create user with home dir
RUN addgroup -S easy-audit && adduser -S easy-audit -H -h /opt/easy-audit/home -s /bin/sh -G easy-audit -u 3453
# Copy Binary
COPY --from=builder /ea-bin/* /opt/easy-audit/bin/
RUN chmod 700 -R /opt/easy-audit &&\
    chmod 755 -R /opt/easy-audit/bin &&\
    chown easy-audit:easy-audit -R /opt/easy-audit
# change to required user
USER easy-audit:easy-audit
# set the entrypoints
ENTRYPOINT ["/bin/sh", "-c"]
CMD [ "easy-audit"]