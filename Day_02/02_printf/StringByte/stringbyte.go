package StringByte

import "fmt"

// One 字符串和[]byte
func One() {
	str := "Smurfs的格格巫"

	fmt.Printf("%s\n", str)
	fmt.Printf("%q\n", str)
	fmt.Printf("%x\n", str)
	fmt.Printf("%X\n", str)
}
