ARG BIN=artifacts
ARG CMD=./cmd/artifacts

FROM golang:tip-alpine3.23 AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOD=linux go build -o "/out/${BIN}" "${CMD}"

FROM gcr.io/distroless/static:nonroot

ARG BIN
COPY --from=builder "/out/${BIN}" /app

USER nonroot

EXPOSE 8080

CMD ["/app"]
