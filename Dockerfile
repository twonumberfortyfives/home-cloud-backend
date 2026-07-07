FROM golang:1.26.4 AS development

WORKDIR /home-cloud-backend

RUN go install github.com/air-verse/air@latest

COPY go.mod go.mod ./

RUN go mod download 

COPY . .

CMD ["air", "-c", ".air.toml"]

# RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/main.go

# EXPOSE 8000

# CMD ["/server"]
