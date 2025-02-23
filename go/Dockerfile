FROM golang:1.24.0-bullseye

RUN apt-get update && apt-get install -y build-essential

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

CMD [ "tail", "-f", "/dev/null" ]