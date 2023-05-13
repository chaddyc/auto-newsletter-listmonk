FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN mkdir internal


COPY main.go ./
COPY internal/* ./internal/

RUN go build -o /auto-newsletter

CMD [ "/auto-newsletter" ]