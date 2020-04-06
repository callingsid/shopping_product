FROM golang:1.13.9
ENV REPO_URL=github.com/callingsid/shopping_product
ENV GO111MODULE=on

ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/REPO_URL

ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

CMD["./main.exe"]



