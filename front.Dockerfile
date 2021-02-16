FROM golang:1.15-buster

WORKDIR /app

COPY . ./

RUN go build -o front main_front/main.go

CMD [ "./front" ]