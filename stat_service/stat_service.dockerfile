FROM golang:1.22

WORKDIR /app

COPY ./stat_service/go.mod /app
COPY ./stat_service/go.sum /app
RUN go mod download && go mod verify

COPY ./stat_service /app
RUN go build -v -o /usr/local/bin/stat_service /app

CMD [ "stat_service" ]
