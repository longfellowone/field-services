FROM golang:alpine as builder
RUN apk add --no-cache git

WORKDIR /builder

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o /builder/supplysvc ./cmd/supplysvc/main.go

FROM alpine:latest
COPY --from=builder /builder/supplysvc /app/

EXPOSE 8080

CMD ["/app/grpcsvc"]