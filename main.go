package main

import "github.com/go-redis/redis"
import "fmt"
import "os"
import "strconv"

func main() {

  maxAge,err := strconv.Atoi(os.Args[1])
  if err != nil {
    panic(err)
  }

  maxAgeInHours = maxAge * 24

  fmt.Println("maxAgeInHours =", maxAgeInHours)
  
  redisdb := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
  
  iter := redisdb.Scan(0, "", 0).Iterator()
  for iter.Next() {
    var keyName = iter.Val()
    duration, err := redisdb.ObjectIdleTime(keyName).Result()
    if err != nil {
      panic(err)
    }
    var idleTime = int(duration.Hours())
    if idleTime > maxAgeInHours {
      fmt.Println("Deleting key", keyName, "because idle time is", idleTime, "hours")
      _, err := redisdb.Unlink(keyName).Result()
      if err != nil {
        panic(err)
      }
    }
  }

  if err := iter.Err(); err != nil {
    panic(err)
  }

}
