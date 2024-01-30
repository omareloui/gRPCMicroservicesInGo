FROM golang:1.21-bullseye

WORKDIR /usr/src/

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY ./proto/go ./proto/go
COPY ./services/order ./services/order
COPY ./lib/go ./lib/go

RUN go mod tidy

CMD ["air", "-c", "services/order/config/.air.toml"]
