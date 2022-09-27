FROM golang:1.19-alpine

ENV GOARCH=arm64

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get 
RUN go build .

# This container exposes port 8080 to the outside world
EXPOSE 8090:8090

# Run the executable
CMD ["./eta"]