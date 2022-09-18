FROM golang:1.18-buster

WORKDIR /app
COPY . .
ENV GO111MODULE=on
RUN go build -mod=vendor -o main .

ENTRYPOINT [ "./main", "api"]
EXPOSE 8080
