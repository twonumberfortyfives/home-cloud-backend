FROM golang:1.26.4 AS builder

WORKDIR /home-cloud-backend

COPY go.mod go.mod ./

RUN go mod download 

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 8080

CMD ["/docker-gs-ping"]
