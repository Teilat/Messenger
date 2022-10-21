# syntax=docker/dockerfile:1
FROM golang:1.18-alpine
# creaitng dir
WORKDIR /messenger

# getting dependencys
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# copy files
COPY . .
# build app
RUN go build -o /Messenger
EXPOSE 8080
# run app
CMD [ "/Messenger" ]