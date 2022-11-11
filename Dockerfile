FROM golang:alpine
LABEL maintainer="Base"
WORKDIR /reazon
COPY . .
RUN go build -o main .
EXPOSE 8080
VOLUME [ "../img/avatars" ]
CMD ["./main"]