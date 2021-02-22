FROM golang:alpine as builder

WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o /assets/check ./cmd/check
RUN go build -o /assets/in ./cmd/in
RUN go build -o /assets/out ./cmd/out

FROM alpine as resource
COPY --from=builder assets/ /opt/resource/
RUN chmod +x /opt/resource/*

FROM resource
