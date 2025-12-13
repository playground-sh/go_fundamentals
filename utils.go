package main

import (
	"fmt"
	DataTypes "go_fundamentals/data_types"
)

const (
	BooleanDataTypes   = "BOOLEAN_DATA_TYPES"
	ConstantDataTypes  = "CONSTANT_DATA_TYPES"
	GoIota             = "GO_IOTA"
	NumericDataTypes   = "NUMERIC_DATA_TYPES"
	StringDataTypes    = "STRING_DATA_TYPES"
	ArrayAndSliceTypes = "ARRAY_AND_SLICE_TYPES"
	Pointers           = "POINTERS"
	Operators          = "OPERATORS"
	Sets               = "SETS"
	Maps               = "MAPS"
)

func RunLesson(lesson string) {
	switch lesson {
	case NumericDataTypes:
		// Numeric Types
		DataTypes.SignedIntegers()
		fmt.Println()
		DataTypes.UnsignedIntegers()
		fmt.Println()
		DataTypes.FloatingPoints()
		fmt.Println()
		DataTypes.ComplexNumbers()
		fmt.Println()
		DataTypes.DefaultIntegers()
		fmt.Println()
	case BooleanDataTypes:
		// Boolean Type
		DataTypes.BooleanData()
		fmt.Println()
	case StringDataTypes:
		// String Type
		DataTypes.StringData()
		fmt.Println()
	case ConstantDataTypes:
		// Constants
		DataTypes.Constants()
		fmt.Println()
	case GoIota:
		DataTypes.Iota()
		fmt.Println()
	case ArrayAndSliceTypes:
		// Arrays and Slices
		DataTypes.ArraysAndSlices()
		fmt.Println()
	case Pointers:
		DataTypes.Pointers()
		fmt.Println()
	case Operators:
		DataTypes.OperatorsAndComparisons()
		fmt.Println()
	case Sets:
		DataTypes.SetsDemo()
		fmt.Println()
	case Maps:
		DataTypes.MapsDemo()
	}
}
