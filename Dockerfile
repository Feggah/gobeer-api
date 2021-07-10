FROM golang:1.16-alpine AS build_base

RUN apk add build-base

WORKDIR /tmp/gobeer

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/gobeer-api ./web

FROM alpine:3.14

COPY --from=build_base /tmp/gobeer/out/gobeer-api /app/gobeer-api

EXPOSE 4000

CMD ["/app/gobeer-api"]
