FROM golang:1.22

WORKDIR /app

COPY ./client_service/go.mod ./
COPY ./client_service/go.sum ./
RUN go mod download && go mod verify

COPY ./client_service .
RUN go build -v -o /usr/local/bin/client_service ./...

CMD [ "client_service" ]
