FROM golang:alpine

RUN apk update && apk --no-cache add make git unzip

WORKDIR /app

COPY . .

RUN go mod download

CMD  ["make", "start"]