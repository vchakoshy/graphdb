FROM golang:1.18-buster AS builder 

WORKDIR /app
COPY . .
ENV GO111MODULE=on
RUN go build -mod=vendor -o main .

FROM golang:1.18-buster
WORKDIR /app 
COPY --from=builder /app/main ./

ENTRYPOINT [ "./main", "api"]
EXPOSE 8080
