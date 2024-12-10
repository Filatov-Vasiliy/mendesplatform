FROM golang:1.23.2

WORKDIR /profile

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o ./profile cmd/main.go


CMD ["./profile"]