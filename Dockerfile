FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o app cmd/botsrv/main.go

FROM alpine
WORKDIR /build
COPY --from=builder /build/app /build/app
COPY cfg .
CMD ["./app"]