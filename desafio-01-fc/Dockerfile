FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN apt-get update && \
    apt-get install build-essential -y

COPY . /go/src

RUN cd /go/src

RUN go build -o desafio-01-fc main.go

EXPOSE 8000

CMD ["./desafio-01-fc"]