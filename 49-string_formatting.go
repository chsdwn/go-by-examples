package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	fmt.Printf("struct1: %v \n", p)                     // {1 2}
	fmt.Printf("struct2: %+v \n", p)                    // {x:1 y:2}
	fmt.Printf("struct3: %#v \n", p)                    // main.point{x:1, y:2}
	fmt.Printf("type: %T \n", p)                        // main.point
	fmt.Printf("bool: %t \n", true)                     // true
	fmt.Printf("int: %d \n", 123)                       // 123
	fmt.Printf("bin: %b \n", 14)                        // 1110
	fmt.Printf("char: %c \n", 33)                       // !
	fmt.Printf("hex: %x \n", 456)                       // 1c8
	fmt.Printf("float1: %f \n", 78.9)                   // 78.900000
	fmt.Printf("float2: %e \n", 123400000.0)            // 1.234000e+08
	fmt.Printf("float3: %E \n", 123400000.0)            // 1.234000E+08
	fmt.Printf("str1: %s \n", "\"string\"")             // "string"
	fmt.Printf("str2: %q \n", "\"string\"")             // "\"string\""
	fmt.Printf("str3: %x \n", "hex this")               // 6865782074686973
	fmt.Printf("pointer: %p \n", &p)                    // 0x140000aa010
	fmt.Printf("width1: |%6d|%6d| \n", 12, 345)         // |    12|   345|
	fmt.Printf("width2: |%6.2f|%6.2f| \n", 1.2, 3.45)   // |  1.20|  3.45|
	fmt.Printf("width3: |%-6.2f|%-6.2f| \n", 1.2, 3.45) // |1.20  |3.45  |
	fmt.Printf("width4: |%6s|%6s| \n", "foo", "b")      // |   foo|     b|
	fmt.Printf("width5: |%-6s|%-6s| \n", "foo", "b")    // |foo   |b     |

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s) // a string

	fmt.Fprintf(os.Stderr, "io: an %s \n", "error") // io: an error
}
