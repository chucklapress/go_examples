package main

import (
        "fmt"
	"log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

type Person struct {
        Name string
        Phone string
}

func main() {
        session, err := mgo.Dial("mongodb://localhost:27017/new")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior to check all nodes.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("new").C("people")
        err = c.Insert(&Person{"Chuck", "864 304 0561"},
	               &Person{"MB", "864 770 3690"})
        if err != nil {
                log.Fatal(err)
        }

        result := Person{}
        err = c.Find(bson.M{"name": "MB"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("Phone:", result.Phone)
}
