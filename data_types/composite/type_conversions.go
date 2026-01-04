package composite

import (
	"fmt"
	"unsafe"
)

func NumericTypeConversions() {
	fmt.Println("Numeric Type Conversions:")
	var x int = 42
	var y float64 = float64(x)
	fmt.Printf("x: %d, y = float64(x): %.2f\n", x, y)

	var z uint = uint(y)
	fmt.Printf("z = uint(y): %d\n", z)
}

func StringByteConversion() {
	fmt.Println("String Byte Conversion:")
	var s string
	s = "Hello"
	fmt.Println("String(s):", s)

	b := []byte(s)
	fmt.Println("Byte(s):", b)

	str := string(b) // converting a `byte` back to `string`
	fmt.Println("Back to String:", str)
}

func TypeAliasing() {
	type longitude float64
	type latitude float64

	var x longitude
	x = 1.25

	var y latitude
	y = 1.325

	// the following is an error, although the underlying type is the same
	// x = y
	// you have type-cast/convert the types to do this,
	// x = longitude(y)

	fmt.Printf("(longitude, latitude): (%f, %f)\n", x, y)

	type coordinate struct {
		x longitude
		y latitude
	}

	dhaka := coordinate{23.8041, 90.4152}
	fmt.Printf("Dhaka coordinates: (%f, %f)\n", dhaka.x, dhaka.y)
}

func PointerConversion() {
	var a int
	a = 2168

	ptr := unsafe.Pointer(&a)
	fmt.Println("Pointer to an int:", ptr, &a)
}

// Pointer conversion basics
// # Safe Conversions (The "High-Level" Way)
// These are built-in and checked by the compiler.
// * Numeric to Numeric: `int` ↔ `float64` ↔ `uint32`. (Always allowed, might lose precision).
// * String to Byte Slice: `string(bytes)` or `[]byte(str)`. (Go actually copies the data to ensure safety).
// * Identical Underlying Types: If you have `type ID int`, you can do `ID(myInt)`.
//
// # Unsafe Reinterpretations (The "C" Way)
// Using `unsafe.Pointer`, you can technically convert any pointer to any other pointer.
// `*int` → `unsafe.Pointer` → `*float64`
// `*MyStruct` → `unsafe.Pointer` → `*OtherStruct`
// **But beware:** Just because you can cast the pointer doesn't mean the data makes sense.
// Reinterpreting a float64 as an int64 will give you a giant, nonsensical number because
// the bit-layouts (IEEE 754 vs Two's Complement) are totally different.

func WhyUnsafe() {
	var a int64 = 0x1122334455667788 // Hex value

	// 1. This would FAIL:
	// bytePtr := (*byte)(&a)

	// 2. This WORKS:
	// We go from *int64 -> unsafe.Pointer -> *byte
	bytePtr := (*byte)(unsafe.Pointer(&a))

	fmt.Printf("First byte of memory: %x\n", *bytePtr)
}

// Packet - Practical Example: Casting Structs
type Packet struct {
	ID  uint32
	Val uint32
}

func Reinterpret() {
	p := Packet{ID: 1, Val: 100}

	// Treat the Packet struct as an array of 8 bytes
	// This is essentially (char*)&p in C
	ptr := (*[8]byte)(unsafe.Pointer(&p))
	fmt.Println("Raw bytes of struct:", *ptr)

	// 1. We have our 'ptr' which is a *[8]byte
	// 2. Convert *[8]byte -> unsafe.Pointer
	// 3. Convert unsafe.Pointer -> *Packet
	originalPacket := (*Packet)(unsafe.Pointer(ptr))

	// 4. Dereference to get the value
	fmt.Println("Back to Packet:", *originalPacket)
}
