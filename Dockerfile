FROM golang:1.20

WORKDIR /app

COPY ./src/section1/go.mod ./src/section1/go.sum ./

RUN go mod download

COPY ./src/section1 ./

EXPOSE 8080
