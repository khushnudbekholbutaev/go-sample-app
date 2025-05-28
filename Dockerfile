FROM golang:1.24.3-alpine3.21 AS builder

WORKDIR /app

# git o'rnatish uchun
RUN apk add --no-cache git

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/index.tmpl .
COPY --from=builder /app/image.tmpl .

RUN chmod +x ./app

EXPOSE 8080

CMD ["./app"]
