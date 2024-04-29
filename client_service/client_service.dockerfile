FROM golang:1.22

WORKDIR /app

COPY ./client_service/go.mod /app
COPY ./client_service/go.sum /app
RUN go mod download && go mod verify

COPY ./client_service /app
RUN go build -v -o /usr/local/bin/client_service /app

CMD [ "client_service" ]
