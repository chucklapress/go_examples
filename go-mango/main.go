package main

import (
        "fmt"
	"log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

type Person struct {
        Name string
        Address string
}

func main() {
        session, err := mgo.Dial("mongodb://localhost:27017/peeps")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("peeps")
        err = c.Insert(&Person{"Chuck", "29 Butler Springs Road"},
	               &Person{"Summer", "29 Butler Springs Road"})
        if err != nil {
                log.Fatal(err)
        }

        result := Person{}
        err = c.Find(bson.M{"name": "Chuck"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("Address:", result.Address)
}
