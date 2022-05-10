FROM golang:1.17 AS build

WORKDIR /app

COPY . .

RUN go build -o /main

FROM alpine:latest

WORKDIR /app

COPY --from=build /main /main

ENTRYPOINT ["/main"]
