FROM golang:1.8

COPY bin/commango /go/bin
COPY bin/config.toml /go/bin

RUN chmod +x /go/bin/commango

WORKDIR /go/bin

CMD ["/go/bin/commango"]