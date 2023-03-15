FROM golang:1.20-alpine

WORKDIR /app

COPY docker-config.yaml ./config.yaml
COPY .build/linux/* ./

EXPOSE 8080
CMD ["./messenger"]