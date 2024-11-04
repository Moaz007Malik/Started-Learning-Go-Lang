package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255 // uint8 has maximum length of 255 bytes
	x++               // Here we have incremented x
	fmt.Println(x)    // Let's print x and see how it handles the overflow
	// The output is 0

	// Error: constant 256 overflows int8
	// a := int8(255 + 1)

	var b int8 = 127
	fmt.Println(b + 1) // Let's print b and see how it handles the overflow
	// The output is -128

	b = -128
	b--            // Here we have decremented b
	fmt.Println(b) // Let's print b and see how it handles the overflow
	// The output is 127. It goes back to its maximum value

	f := float32(math.MaxFloat32) // storing maximum value of float32 in f
	fmt.Println(f)

	f = f * 1.2
	fmt.Println(f)
	// This means the float overflow is infinity
}
