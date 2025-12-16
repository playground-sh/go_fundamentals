package data_types

import (
	"fmt"
	"math"
	"unsafe"
)

func SignedIntegers() {
	var signedInt8 int8
	signedInt8 = math.MaxInt8
	fmt.Println("math.MaxInt8:", signedInt8)

	signedInt8 = math.MinInt8
	fmt.Println("math.MinInt8:", signedInt8)

	var signedInt16 int16
	signedInt16 = math.MaxInt16
	fmt.Println("math.MaxInt16:", signedInt16)

	signedInt16 = math.MinInt16
	fmt.Println("math.MinInt16:", signedInt16)

	var signedInt32 int32
	signedInt32 = math.MaxInt32
	fmt.Println("math.MaxInt32:", signedInt32)

	signedInt32 = math.MinInt32
	fmt.Println("math.MinInt32:", signedInt32)

	var signedInt64 int64
	signedInt64 = math.MaxInt64
	fmt.Println("math.MaxInt64:", signedInt64)

	signedInt64 = math.MinInt64
	fmt.Println("math.MinInt64:", signedInt64)
}

func UnsignedIntegers() {
	var unsignedInt8 uint8
	unsignedInt8 = math.MaxUint8
	fmt.Println("math.MaxUint8:", unsignedInt8)

	var bytes byte
	bytes = math.MaxUint8
	fmt.Println("max(byte) == math.MaxUint8:", bytes)

	var unsignedInt16 uint16
	unsignedInt16 = math.MaxUint16
	fmt.Println("math.MaxUint16:", unsignedInt16)

	var unsignedInt32 uint32
	unsignedInt32 = math.MaxUint32
	fmt.Println("math.MaxUint32:", unsignedInt32)

	var unsignedInt64 int64
	unsignedInt64 = math.MaxInt64
	fmt.Println("math.MaxInt64:", unsignedInt64)

	fmt.Println("Min value for all unsigned types is: 0")
}

func FloatingPoints() {
	var floatingPoint32 float32
	floatingPoint32 = math.MaxFloat32
	fmt.Printf("math.MaxFloat32: %38f\n", floatingPoint32)

	floatingPoint32 = math.SmallestNonzeroFloat32
	fmt.Printf("math.SmallestNonzeroFloat32: : %.16g\n", floatingPoint32)

	var floatingPoint64 float64
	floatingPoint64 = math.MaxFloat64
	fmt.Println("math.MaxFloat64:", floatingPoint64)

	floatingPoint64 = math.SmallestNonzeroFloat64
	fmt.Printf("math.SmallestNonzeroFloat64: : %.16g\n", floatingPoint64)
}

func ComplexNumbers() {
	var cmplx64 complex64
	cmplx64 = 1.0 + 2.5i
	fmt.Printf("cmplx64: %v (Real part is float32)\n", cmplx64)

	var cmplx128 complex128
	cmplx128 = 10.5 + 50.0i
	fmt.Printf("cmplx128: %v (Real part is float64)\n", cmplx128)

	// Getting the component parts:
	fmt.Printf("Real(cmplx128): %f\n", real(cmplx128))
	fmt.Printf("Imag(cmplx128): %f\n", imag(cmplx128))
}

func DefaultIntegers() {
	// Note: To be strict about the Max/Min limits for 'int' and 'uint',
	// you must use the 'unsafe' package to check the size, as these
	// limits are not constants in the math package.

	var defaultInt int = 42
	fmt.Printf("int size (bytes): %d\n", unsafe.Sizeof(defaultInt))

	var defaultUint uint = 100
	fmt.Printf("uint size (bytes): %d (Min: 0)\n", unsafe.Sizeof(defaultUint))

	var pointerUint uintptr = 0xFFFFFFFF
	fmt.Printf("uintptr: %d\n", pointerUint)
}
