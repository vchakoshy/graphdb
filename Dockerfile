FROM dockerhub.ir/golang:1.18-buster AS builder

WORKDIR /app
COPY . .
ENV GO111MODULE=on
RUN go build -o main .

ENTRYPOINT [ "./main", "api"]
EXPOSE 8080
