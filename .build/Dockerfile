FROM golang:1.23 AS builder

WORKDIR /build

COPY . .

RUN go mod tidy && go mod download && go mod vendor

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o renco .

FROM scratch

COPY --from=builder /build/renco /
COPY --from=builder /build/config /config
COPY --from=builder /build/database /database

EXPOSE 3001