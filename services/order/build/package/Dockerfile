FROM golang:1.21-bullseye AS build
WORKDIR /usr/src/

RUN useradd -u 1001 nonroot

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./proto/go/ ./proto/go
COPY ./services/order ./services/order
COPY ./lib/go ./lib/go


RUN go build -v \
  -ldflags="-linkmode external -extldflags -static" \
  -tags netgo \
  -o /usr/local/bin/order \
  ./services/order/cmd/order

###
FROM scratch
WORKDIR /app
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /usr/local/bin/order .

USER nonroot

CMD ["./order"]
