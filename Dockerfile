FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o app cmd/botsrv/main.go

FROM alpine
ARG token
ENV token $token
WORKDIR /build
COPY --from=builder /build/app /build/app
COPY cfg/local.toml cfg/local.toml
RUN echo -ne '\n\n[Bot]\nToken = "'"$token"'"' >> cfg/local.toml
CMD ["cat", "cfg/local.toml"]