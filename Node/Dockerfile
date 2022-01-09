FROM golang:1.17-alpine

RUN pwd
RUN mkdir ../home/app
WORKDIR /../home/app
COPY . .

RUN go build main.go
CMD ["./main"]