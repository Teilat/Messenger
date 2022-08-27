# syntax=docker/dockerfile:1
FROM golang:1.18
# creaitng dir
WORKDIR /app
# getting dependencys
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# copy files
COPY *.go ./
COPY ./database ./
COPY ./internal ./
COPY ./webapi ./
COPY config.yaml ./
# build app
RUN go build -o /messanger/backend
# run app
CMD [ "/messanger/backend" ]