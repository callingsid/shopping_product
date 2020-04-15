FROM golang:1.13.9

ENV GO111MODULE=on

ENV http_proxy 10.127.134.49:3128

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY src $WORKDIR

RUN echo $PWD

RUN echo $ls

RUN go build
EXPOSE 8080
 
CMD ["./main"]
