FROM golang:1.22

WORKDIR /app

COPY ./task_service/go.mod /app
COPY ./task_service/go.sum /app
RUN go mod download && go mod verify

COPY ./task_service /app
RUN go build -v -o /usr/local/bin/task_service /app

CMD [ "task_service" ]
