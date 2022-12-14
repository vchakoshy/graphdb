[![Go](https://github.com/vchakoshy/graphdb/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/vchakoshy/graphdb/actions/workflows/go.yml)

# GraphDB 
GraphDB is an open source In Memory graph database written in Golang, optimized for social networks.

## Run Server 
```bash
go run main.go api
```

rest api http://localhost:8081 

grpc server http://localhost:8080

## Client 

```bash 
go get -u github.com/vchakoshy/graphdb
```

```golang
import "github.com/vchakoshy/graphdb/service"
```

```golang
var opts []grpc.DialOption{
    grpc.WithTransportCredentials(insecure.NewCredentials()),
}

conn, err := grpc.Dial("127.0.0.1:8080", opts...)
if err != nil {
    log.Fatal(err)
}
defer conn.Close()

client := service.NewGraphdbClient(conn)
ctx := context.Background()

// add follow items
client.AddFollow(ctx, &service.Follow{From: 1, To: 2})
client.AddFollow(ctx, &service.Follow{From: 2, To: 3})
client.AddFollow(ctx, &service.Follow{From: 2, To: 4})
client.AddFollow(ctx, &service.Follow{From: 2, To: 5})

// Get friends of friends 
res, err := client.GetFriendsOfFriends(ctx, &service.User{Id: 1})
if err != nil {
    panic(err)
}
for _, i := range res.GetUsers() {
    log.Println("fof of ", 1, "is", i)
}
```
