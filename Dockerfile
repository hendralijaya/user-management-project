FROM golang:alpine
RUN mkdir /app

WORKDIR /golang-dev

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 8000
EXPOSE 8080

ENTRYPOINT CompileDaemon -log-prefix=false -directory="." -build="go build -o usermanagement" -command="./usermanagement"
