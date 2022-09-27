FROM golang:1.19-alpine

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/bus_eta_hk

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

ENV GOARCH=arm64
# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN go build .

# This container exposes port 8080 to the outside world
EXPOSE 8090

RUN ls
RUN echo $PATH
RUN pwd

# Run the executable
ENTRYPOINT ["eta"]

EXPOSE 8090