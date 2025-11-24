package main

import (
	"fmt"
	"go_fundamentals/data_types"
)

const (
	BooleanDataTypes   = "BOOLEAN_DATA_TYPES"
	ConstantDataTypes  = "CONSTANT_DATA_TYPES"
	NumericDataTypes   = "NUMERIC_DATA_TYPES"
	StringDataTypes    = "STRING_DATA_TYPES"
	ArrayAndSliceTypes = "ARRAY_AND_SLICE_TYPES"
	Pointers           = "POINTERS"
)

func RunLesson(lesson string) {
	switch lesson {
	case NumericDataTypes:
		// Numeric Types
		data_types.SignedIntegers()
		fmt.Println()
		data_types.UnsignedIntegers()
		fmt.Println()
		data_types.FloatingPoints()
		fmt.Println()
		data_types.ComplexNumbers()
		fmt.Println()
		data_types.DefaultIntegers()
		fmt.Println()
	case BooleanDataTypes:
		// Boolean Type
		data_types.BooleanData()
		fmt.Println()
	case StringDataTypes:
		// String Type
		data_types.StringData()
		fmt.Println()
	case ConstantDataTypes:
		// Constants
		data_types.Constants()
		fmt.Println()
	case ArrayAndSliceTypes:
		// Arrays and Slices
		data_types.ArraysAndSlices()
	case Pointers:
		data_types.Pointers()
	}

}
