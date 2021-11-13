FROM golang:1.17.3 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY app/ app/
RUN CGO_ENABLED=0 go build -o chproxy app/main.go

FROM scratch
COPY --from=builder /app/chproxy /bin/
ENTRYPOINT ["/bin/chproxy"]
