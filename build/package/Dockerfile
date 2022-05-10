FROM golang:1.17 AS build

WORKDIR /app

COPY . .

RUN go build -o /main

FROM scratch

WORKDIR /app

COPY --from=build /main .

ENTRYPOINT ["./main"]
