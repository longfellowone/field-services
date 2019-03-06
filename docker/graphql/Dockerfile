# Postgres hostname must be set to "db" when using with docker compose
FROM golang:alpine as builder
RUN apk add --no-cache git

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/graphqlsvc

FROM scratch
COPY --from=builder /app/graphqlsvc /app/
EXPOSE 8080
ENTRYPOINT ["/app/graphqlsvc"]