FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /vaksin-id-be

EXPOSE 8080

CMD ["/vaksin-id-be"]