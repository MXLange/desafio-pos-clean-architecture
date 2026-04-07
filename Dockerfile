FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache build-base

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o /app/bin/server ./cmd/server

FROM alpine:3.21

WORKDIR /app

RUN apk add --no-cache libgcc libstdc++

COPY --from=builder /app/bin/server /app/server
COPY --from=builder /app/.env.local /app/.env
COPY --from=builder /app/internal/infra/db/migrations /app/internal/infra/db/migrations

RUN mkdir -p /app/data

VOLUME ["/app/data"]

CMD ["/bin/sh", "-c", "mkdir -p /app/data/internal/infra/db && cp /app/.env /app/data/.env && rm -rf /app/data/internal/infra/db/migrations && cp -R /app/internal/infra/db/migrations /app/data/internal/infra/db/migrations && cd /app/data && /app/server"]
