FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go build .

CMD ["app"]

EXPOSE 3000
