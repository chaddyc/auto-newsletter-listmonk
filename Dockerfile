FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN mkdir listmonk
RUN mkdir rss
RUN mkdir loadenv

COPY main.go ./
COPY listmonk/* ./

RUN go build -o /opensourcegeeks-newsletter

CMD [ "/opensourcegeeks-newsletter" ]