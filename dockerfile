FROM golang:1.18-alpine
COPY .build/linux/* ./
EXPOSE 8080