FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o exe .

FROM alpine

WORKDIR /build

COPY --from=builder /build/ /build/exe

CMD ["./exe/exe"]
