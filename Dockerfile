FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN mkdir listmonk
RUN mkdir rss
RUN mkdir loadenv

COPY main.go ./
COPY listmonk/listmonk.go ./listmonk/
COPY rss/rss.go ./rss/
COPY loadenv/loadenv.go ./loadenv/
COPY .env ./

RUN go build -o /opensourcegeeks-newsletter

CMD [ "/opensourcegeeks-newsletter" ]