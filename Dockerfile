FROM golang:1.26.4 AS development

WORKDIR /app

COPY ./src .

RUN go install github.com/air-verse/air@latest

RUN go mod download 

CMD ["air", "-c", ".air.toml"]

# RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/main.go

# EXPOSE 8000

# CMD ["/server"]
