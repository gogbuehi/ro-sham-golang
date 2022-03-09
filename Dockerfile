FROM golang:1.17.8-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod init app/main
RUN go build -o main .
CMD ["/app/main"]