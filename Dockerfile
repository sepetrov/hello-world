FROM golang:1.25.6 AS builder

WORKDIR /app

COPY . .

RUN mkdir -p bin && \
    rm -frv bin/* && \
    go build -o ./bin/hello-world ./cmd/main.go

FROM debian:trixie

COPY --from=builder /app/bin/hello-world /usr/local/bin/entrypoint

EXPOSE 8080

ENTRYPOINT ["entrypoint"]
