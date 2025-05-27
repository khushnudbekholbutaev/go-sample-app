FROM golang:1.24.3-alpine3.21 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch

WORKDIR /app

COPY --from=builder /app/app .

COPY --from=builder /app/index.tmpl .
COPY --from=builder /app/image.tmpl .

EXPOSE 8080

CMD ["./app"]
