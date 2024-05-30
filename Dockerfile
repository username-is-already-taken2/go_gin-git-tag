# Build the application from source
FROM golang:1.21.5-alpine AS builder

WORKDIR /app

COPY . .

ARG VERSION
ARG BUILD_TIME
RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -mod=vendor -ldflags "--extldflags --static -X main.Version=$VERSION -X main.BuildTime=$BUILD_TIME" -o gin

FROM scratch

WORKDIR /app

COPY --from=builder /app/gin .

ENV PORT=8080 \
    GIN_MODE=release

# Match the PORT ENV
EXPOSE 8080

CMD ["./gin"]
