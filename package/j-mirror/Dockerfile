FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /main

FROM alpine

WORKDIR /

COPY --from=build /main /main

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/main"]
