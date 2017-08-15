package main

import "github.com/hoisie/redis"

func main() {
    var client redis.Client
    var key = "Elizabeth"
    client.Set(key, []byte("Barr"))
    val, _ := client.Get("Elizabeth")
    println(key, string(val))

}
