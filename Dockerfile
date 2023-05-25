#########
# Build #
#########
FROM golang:alpine AS server_builder

ARG TARGETARCH
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN apk add build-base

COPY . .

RUN GOOS=linux GOARCH=$TARGETARCH go build -a -o vatprc-queue .

##################
# Build Frontend #
##################
FROM node:18 AS frontend_builder

WORKDIR /build

COPY frontend/ .
RUN npm install
RUN npm run build

##########
# Deploy #
##########

FROM alpine

RUN mkdir -p /config
COPY --from=server_builder /build/vatprc-queue /
COPY --from=server_builder /build/config/config.sample.ini /config/config.ini
COPY --from=frontend_builder /build/dist/ /views/

ENTRYPOINT ["/vatprc-queue"]

EXPOSE 80