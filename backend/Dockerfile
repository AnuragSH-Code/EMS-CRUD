FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/ems ./cmd/api

############ Runtime stage ##########
FROM scratch
COPY --from=builder /bin/ems /ems
COPY .env /app/.env 
WORKDIR /app               

EXPOSE 8080
USER 1001
ENTRYPOINT ["/ems"]
