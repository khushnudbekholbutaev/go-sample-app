FROM golang:1.24.3-alpine3.21 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY go-sample-app/ ./go-sample-app/

WORKDIR /app/go-sample-app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch

WORKDIR /app

COPY --from=builder /app/go-sample-app/app .
COPY --from=builder /app/go-sample-app/index.tmpl .
COPY --from=builder /app/go-sample-app/image.tmpl .

EXPOSE 8080
CMD ["./app"]