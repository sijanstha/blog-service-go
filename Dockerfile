FROM golang:1.15.3-alpine3.12

EXPOSE 9090

RUN apk update \
    && apk add --no-cache \ 
    mysql-client \
    build-base

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
COPY ./entrypoint.sh /usr/local/bin/entrypoint.sh
RUN /bin/chmod +x /usr/local/bin/entrypoint.sh

RUN go build src/main.go
RUN mv main /usr/local/bin/

CMD ["main"]
ENTRYPOINT ["entrypoint.sh"]