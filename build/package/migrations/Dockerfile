FROM golang:1.26-alpine

# install goose
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.26.0

WORKDIR /app

# copy migration files
COPY ./migrations .
