FROM golang:1.18-alpine as builder
WORKDIR /app
ADD . .

RUN apk add build-base git
RUN go install github.com/google/wire/cmd/wire@latest
RUN wire ./cmd/...
RUN mkdir build
RUN go build -o build ./cmd/...

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/build/. .

CMD ["./http"]
