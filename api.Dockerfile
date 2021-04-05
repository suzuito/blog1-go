FROM golang:1.15-buster

WORKDIR /app

COPY . ./

RUN make api.exe

CMD [ "./api.exe" ]