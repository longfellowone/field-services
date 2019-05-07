FROM golang:alpine as builder
RUN apk add --no-cache git

WORKDIR /builder

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o /builder/grpcsvc ./cmd/grpcsvc/main.go

FROM alpine:latest
COPY --from=builder /builder/grpcsvc /app/

EXPOSE 9090

CMD ["/app/grpcsvc"]