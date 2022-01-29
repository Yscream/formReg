FROM golang:1.17.5

COPY . .

WORKDIR /cmd

CMD ["go", "run", "cmd/main.go"]