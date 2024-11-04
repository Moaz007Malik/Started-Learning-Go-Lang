package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s1 := "Hi there Go!"
	fmt.Println(s1)

	fmt.Println("He says: \"Hello!\"")

	fmt.Println(`He says: "Hello!"`) // known as "raw string"

	s2 := `I like \n Go!` //raw string | \n will have no effect
	fmt.Println(s2)

	fmt.Println("Price: 100\nBrand: Nike")
	fmt.Println(`
Price: 100
Brand: Nike
	`)

	// These are same
	fmt.Println(`C:\Users\User\Documents\Go\src\github.com\golang\example\hello.go`)
	fmt.Println("C:\\Users\\User\\Documents\\Go\\src\\github.com\\golang\\example\\hello.go")

	var s3 = "I love " + "Go " + "Programming"
	fmt.Println(s3 + "!")

	fmt.Println("Element at index 0:", s3[0])

	// s3[5] = "x"
	// s3[5] = "x" // Error: Index out of range | cannot assign to s3[5] (neither addressable nor a map index expression)

	// Runes and Strings

	var1, var2 := 'a', 'b'
	fmt.Printf("Type: %T, Value: %d\n", var1, var2)

	str := "țară"
	fmt.Println(len(str))
	fmt.Println("Byte (not rune) at position 1:", str[1])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	fmt.Println("\n" + strings.Repeat("#", 30))

	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i += size
	}

	fmt.Println("\n" + strings.Repeat("#", 30))

	for _, r := range str {
		fmt.Printf("%c", r)
	}

	fmt.Println("\n" + strings.Repeat("#", 30))

	// String Length in Bytes and Runes

	q1 := "goLang"
	fmt.Println(len(q1))

	name := "Condruta"
	fmt.Println(len(name))

	fmt.Println(len("ᛥ"))

	n:= utf8.RuneCountInString("name")
	m:= utf8.RuneCountInString("ᛥ")

	fmt.Println(n,m)


	// Slicing a string

	t1 := "I love Golang!"
	fmt.Println(t1[2:5])

	t2 := "中文维基是世界上"
    fmt.Println(t2[0:2])
	fmt.Println(len(t2))

	rs := []rune(t2)
	fmt.Printf("%T\n", rs)
	fmt.Println(string(rs[0:3]))
}
