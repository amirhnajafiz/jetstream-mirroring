FROM golang:1.20-alpine as build

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

# Executing the main file
CMD ["/bin/main"]