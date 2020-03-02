FROM golang:alpine as builder
ENV CGO_ENABLED=0
WORKDIR '/app'

COPY . .
RUN go mod tidy
RUN mkdir bin
RUN go build -o ./bin/prod_app
COPY ./view ./bin/view


FROM alpine
WORKDIR '/app'
COPY --from=builder /app/bin .
RUN ls
CMD [ "./prod_app" ]