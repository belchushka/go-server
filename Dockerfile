FROM golang:1.19

RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash && apt install migrate -y

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go

ENV GIN_MODE=release

CMD ["./main"]
