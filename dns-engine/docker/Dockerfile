FROM golang:1.27

WORKDIR /app

COPY . .

RUN go build -o kingdns-dns ./cmd/server

EXPOSE 53/udp
EXPOSE 53/tcp

CMD ["./kingdns-dns"]
