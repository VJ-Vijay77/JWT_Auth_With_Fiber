FROM golang:alpine

RUN mkdir /app

COPY mainfile /app

CMD [/app/mainfile]