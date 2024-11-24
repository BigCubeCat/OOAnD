FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o tg_bot .

FROM alpine

WORKDIR /build

COPY --from=builder /build/ /build/tg_bot

CMD ["./tg_bot/tg_bot"]
