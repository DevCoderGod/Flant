FROM golang:1.13

WORKDIR /tarotserv
COPY . .

# ENV CONSUL_HTTP_ADDR=consul:8500

CMD chmod +x tarotserv && ./tarotserv
