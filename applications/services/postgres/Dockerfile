# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

LABEL org.opencontainers.image.source https://github.com/AndreasVikke-School/cph-business-si-exam

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY . .
RUN go build -o ./postgres_service

CMD [ "./postgres_service", "prod" ]