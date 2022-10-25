# Specify the base image for the go app.
FROM golang:alpine
# Add Maintainer info
LABEL maintainer="Base"
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
# Specify that we now need to execute any commands in this directory.
WORKDIR /go/src/github.com/postgres-go
# Copy everything from this project into the filesystem of the container.
COPY . .
# Obtain the package needed to run code. Alternatively use GO Modules. 
RUN go get -u github.com/lib/pq

# #Setup hot-reload for dev stage
# RUN go get github.com/githubnemo/CompileDaemon
# RUN go get -v golang.org/x/tools/gopls

# ENTRYPOINT CompileDaemon --build="go build -a -installsuffix cgo -o main ." --command=./main

# Compile the binary exe for our app.
RUN go build -o main .
# Expose port 8080 to the outside world
EXPOSE 8080
# Start the application.
CMD ["./main"]
