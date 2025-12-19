package composite

import "fmt"

func Pointers() {
	s := "Hello Planet Earth"
	p := &s // address of `s`
	fmt.Println("p:", p)
	fmt.Println("*p:", *p)

	*p = "Welcome Mercury" // `p` holds the address of `s`, not the address of the string
	fmt.Println("s:", s)
	fmt.Println("p:", p) // therefore, the address does not change, since the address of `s` didn't change
	fmt.Println("*p:", *p)
}
