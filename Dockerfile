FROM golang:1.22.6-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o discordify

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /build/discordify /app/

EXPOSE 8888

CMD ["./discordify"]