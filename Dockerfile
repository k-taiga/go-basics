FROM golang:1.20

WORKDIR /app

COPY ./src/section2/go.mod ./src/section2/go.sum ./

RUN go mod download

COPY ./src/section2 ./

EXPOSE 8080
