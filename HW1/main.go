package main

import (
 "fmt"
)

func Sqrt(x float64) float64 {
z:=1.0
for i:=0; i < 10; i++ {
z -= (z*z - x) / (2*z)
}
return z
}

func mySqrt(x float64) (float64, int) {
z:=1.0
prev:=z
prevv:=z
count:=0
for {
prevv = prev
prev = z
z -= (z*z - x) / (2*z)
count++
if z == prevv {
break
}

}
return z, count
}

func main() {
 fmt.Println(mySqrt(2))
}