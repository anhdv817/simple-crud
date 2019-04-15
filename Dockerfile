FROM golang:1.12.4-alpine3.9

RUN apk update && apk add --no-cache git

RUN go get github.com/codegangsta/gin

WORKDIR /simple-crud

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o app

ENTRYPOINT ["./app"]