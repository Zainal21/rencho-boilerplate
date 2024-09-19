FROM golang:1.23 AS builder

ARG BINARY_NAME

WORKDIR /build

COPY . .

RUN go mod tidy && go mod download && go mod vendor

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o ${PROJECT_NAME} .

FROM scratch

COPY --from=builder /build/${PROJECT_NAME} /
COPY --from=builder /build/config /config

EXPOSE 3001