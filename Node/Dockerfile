# Todo 이미지 경량화 작업

FROM golang:1.17-alpine

RUN mkdir ../home/app
WORKDIR /../home/app
COPY . .

RUN go build main.go
CMD ["./main"]