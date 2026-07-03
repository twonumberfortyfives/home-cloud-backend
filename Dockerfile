FROM golang:1.26.4 AS builder

WORKDIR /home-cloud-backend

COPY go.mod go.mod ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/main.go

EXPOSE 8080

CMD ["/server"]
