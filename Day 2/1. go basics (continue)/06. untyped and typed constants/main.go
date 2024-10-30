package main

import "fmt"

func main(){
    const a float64 = 5.1 // typed constant
	const b = 6.7// untyped constant

    const c float64 = a * b // typed constant
    const str = "Hello" + "Go!" // untyped constant

    const d = 5 > 10
    fmt.Println(d)

    // const x int = 5

    // Error: Cannot divide float by int
    // const y float64 = 2.2 * x

    const x = 5
    const y = 2.2 * x

    fmt.Printf("%T\n", y)
    fmt.Println(a)

    var i int = x // x changes to int
    var j float64 = x // var j float64 = float64(x)
    var p byte = x // var p byte = byte(x)
    fmt.Println(i, j, p)

    const r = 5
    var rr = r
    fmt.Printf("%T\n", rr)
}