# Containerize the go application that we have created
# This is the Dockerfile that we will use to build the image
# and run the container

# Start with a base image
FROM golang:1.22-alpine as base

# Install necessary dependencies
RUN apk --no-cache add git

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod ./

# Download all the dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the application
RUN go build -o binary .

#######################################################
# Use a distroless image to run the application, reducing the final image size
FROM gcr.io/distroless/static-debian11

# Set the working directory
WORKDIR /

# Copy the binary from the previous stage
COPY --from=base /app/binary .

# Copy the source files from the previous stage
COPY --from=base /app/src ./src

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the application
CMD ["./binary"]