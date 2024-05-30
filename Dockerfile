FROM golang:1.21.5-alpine AS builder

WORKDIR /app

COPY . .

ARG VERSION
ARG BUILD_TIME
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '--extldflags "--static" -Xmain.Version=$VERSION -Xmain.BuildTime=$BUILD_TIME' -o gin

FROM scratch

WORKDIR /app

COPY --from=builder /app/gin

ENV PORT=8080

EXPOSE 8080

CMD ["./gin"]
