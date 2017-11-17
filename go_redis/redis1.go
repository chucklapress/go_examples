package main

import "github.com/hoisie/redis"

func main() {
    var client redis.Client
    var key = "Chuck"
    client.Set(key, []byte("LaPress"))
    val, _ := client.Get("Chuck")
    println(key, string(val))
}
