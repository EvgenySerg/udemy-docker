FROM golang:alpine

WORKDIR '/app'

COPY . .

RUN go mod tidy

RUN go build main.go


cmd ["go", "run", "main.go"]