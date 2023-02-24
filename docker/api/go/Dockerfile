FROM golang:latest

WORKDIR /var/www

RUN go install github.com/cosmtrek/air@v1.40.4

CMD ["air", "-c", ".air.toml"]
