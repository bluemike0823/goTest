# FROM postgres:latest as postgres

# ENV POSTGRES_USER=postgres


FROM golang:latest

WORKDIR /app

COPY . .

CMD ["go","run","main.go","package-dev"]