FROM golang:1.22-alpine AS build-stage
RUN apk add --update --no-cache
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./
FROM builder AS test-stage
RUN go test -v ./...
# post build stage
FROM gcr.io/distroless/base-debian12 AS release-stage
WORKDIR /
COPY --from=build-stage /api /api
EXPOSE 5000 
CMD ["/api"]