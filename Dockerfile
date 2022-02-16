FROM golang:1.17.5
WORKDIR /app
COPY . /app

CMD ["go", "run", "cmd/main.go"]