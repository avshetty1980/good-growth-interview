FROM golang:1.22-alpine AS builder
RUN apk --no-cache add ca-certificates git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . ./
RUN CGO_ENABLED=0 go build -o /api
# post build stage
FROM alpine
WORKDIR /root
COPY --from=builder /build/api .
EXPOSE 5000
CMD ["./api"]