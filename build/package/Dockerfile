ARG bin=as

FROM golang:tip-alpine3.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build 

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/as .

EXPOSE 8080

CMD ["/as"]
