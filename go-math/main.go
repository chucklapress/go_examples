package main
import (
    "fmt"
    "math"
)
func main() {
    var xpow, ypow, mynum float64
    xpow = 10
    ypow = 2
    mynum = math.Pow(xpow, ypow)
    fmt.Printf("sum %g\n", mynum)
}
