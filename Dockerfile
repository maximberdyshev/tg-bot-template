FROM golang:alpine AS builder

WORKDIR /build

COPY cmd/ /build/cmd/
COPY config/ /build/config/
COPY internal/ /build/internal/
COPY go.mod /build/go.mod
COPY go.sum /build/go.sum

RUN go mod download

RUN go build -o tg-bot-template /build/cmd/tg-bot-template

FROM alpine

WORKDIR /app

COPY --from=builder /build/tg-bot-template /app/tg-bot-template
COPY config.yml /app/config.yml

CMD ["/app/tg-bot-template"]
