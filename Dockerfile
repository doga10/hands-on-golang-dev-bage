# --- builder
FROM golang:1.15-alpine3.12 as builder
LABEL maintainer="Douglas Dennys <douglasdennys@yahoo.com>"
WORKDIR /go/src/app
COPY . .
RUN go mod download && go build ./src/main.go

# --- production
FROM alpine:latest
LABEL maintainer="Douglas Dennys <douglasdennys@yahoo.com>"
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/app .
EXPOSE 8080
CMD ["./main"]