FROM golang:1.15-buster

WORKDIR /app

COPY . ./

RUN go build -o api main_api/main.go

CMD [ "./api" ]