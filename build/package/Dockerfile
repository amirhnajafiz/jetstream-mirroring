FROM golang:1.17-alpine as build

# Add Maintainer Info
LABEL maintainer="amirhnajafiz"

# Change to app work directory
WORKDIR /app

# Copy all of the files
COPY . .

# Building the files
RUN CGO_ENABLED=0 go build -o ./main

# Building on the scratch base image
FROM scratch

# Copy the main file from last image
COPY --from=build ./app/main ./bin/main

# port
EXPOSE 8080

# Executing the main file
ENTRYPOINT ["/bin/main"]