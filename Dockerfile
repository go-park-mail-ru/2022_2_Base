FROM golang:alpine
LABEL maintainer="Base"
WORKDIR /reazon
COPY . .

# #Setup hot-reload for dev stage
# RUN go get github.com/githubnemo/CompileDaemon
# RUN go get -v golang.org/x/tools/gopls

# ENTRYPOINT CompileDaemon --build="go build -a -installsuffix cgo -o main ." --command=./main

RUN go build -o main .

EXPOSE 8080
EXPOSE 5432
CMD ["./main"]