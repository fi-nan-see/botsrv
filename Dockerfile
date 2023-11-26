FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o app cmd/botsrv/main.go

FROM alpine
ARG token
ARG tgSalt
ARG apiUrl
ENV token ${token}
ENV tgSalt ${tgSalt}
ENV apiUrl ${apiUrl}
WORKDIR /build
COPY --from=builder /build/app /build/app
RUN echo -ne '[Server]\nApi="'"$apiUrl"'"\n\n[Bot]\nToken = "'"$token"'"\n' > local.toml
CMD ["./app", "-config=local.toml"]