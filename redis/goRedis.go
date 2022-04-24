package main

import (
  "context"
  "fmt"
  "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// @see https://github.com/go-redis/redis
func main()  {
  rdb := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    Password: "",
    DB: 0,
  })

  err := rdb.Set(ctx, "foo", "1111", 0).Err()
  if err != nil {
    fmt.Println(err)
  }

  val, err := rdb.Get(ctx,"foo").Result()
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println("result = ", val)
}

