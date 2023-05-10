#########
# Build #
#########
FROM golang:alpine AS builder

ARG TARGETARCH
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN apk add build-base

COPY . .

RUN GOOS=linux GOARCH=$TARGETARCH go build -a -o vatprc-queue .

##########
# Deploy #
##########

FROM alpine

RUN mkdir -p /config
COPY --from=builder /build/vatprc-queue /
COPY --from=builder /build/config/config.sample.ini /config/config.ini

ENTRYPOINT ["/vatprc-queue"]

EXPOSE 80