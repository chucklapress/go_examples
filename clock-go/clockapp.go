package main

 import (
         "fmt"
         "time"
 )

 func main() {

         now := time.Now()

         fmt.Println("Today : ", now.Format(time.ANSIC))

         hr, min, sec := now.Clock()

         fmt.Printf("Clock : [%d]hour : [%d]minutes : [%d] seconds \n", hr, min, sec)

 }
