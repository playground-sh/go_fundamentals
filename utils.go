package main

import (
	"fmt"
	CodeOrganization "go_fundamentals/code_organization"
	ControlFlow "go_fundamentals/control_flow"
	BasicDataTypes "go_fundamentals/data_types/basic"
	CompositeDataTypes "go_fundamentals/data_types/composite"
)

const (
	BooleanDataTypes         = "BOOLEAN_DATA_TYPES"
	ConstantDataTypes        = "CONSTANT_DATA_TYPES"
	GoIota                   = "GO_IOTA"
	NumericDataTypes         = "NUMERIC_DATA_TYPES"
	StringDataTypes          = "STRING_DATA_TYPES"
	ArrayAndSliceTypes       = "ARRAY_AND_SLICE_TYPES"
	Pointers                 = "POINTERS"
	Operators                = "OPERATORS"
	Sets                     = "SETS"
	Maps                     = "MAPS"
	Structs                  = "Structs"
	IfElse                   = "IF_ELSE"
	Switch                   = "SWITCH"
	Loops                    = "LOOPS"
	DeferredStatement        = "DEFERRED_STATEMENT"
	Panic                    = "PANIC"
	PanicAndRecover          = "PANIC_AND_RECOVER"
	Functions                = "FUNCTIONS"
	AnonymousFunctions       = "ANONYMOUS_FUNCTIONS"
	GenericCustomConstraints = "GENERIC_CUSTOM_CONSTRAINTS"
	GenericFunctions         = "GENERIC_FUNCTIONS"
	GenericConstraints       = "GENERIC_CONSTRAINTS"
	Methods                  = "METHODS"
	Interfaces               = "INTERFACES"
	Embedding                = "EMBEDDING"
	Composition              = "COMPOSITION"
)

func RunLesson(lesson string) {
	switch lesson {
	case NumericDataTypes:
		// Numeric Types
		BasicDataTypes.SignedIntegers()
		fmt.Println()
		BasicDataTypes.UnsignedIntegers()
		fmt.Println()
		BasicDataTypes.FloatingPoints()
		fmt.Println()
		BasicDataTypes.ComplexNumbers()
		fmt.Println()
		BasicDataTypes.DefaultIntegers()
		fmt.Println()
	case BooleanDataTypes:
		// Boolean Type
		BasicDataTypes.BooleanData()
		fmt.Println()
	case StringDataTypes:
		// String Type
		BasicDataTypes.StringData()
		fmt.Println()
	case ConstantDataTypes:
		// Constants
		BasicDataTypes.Constants()
		fmt.Println()
	case GoIota:
		BasicDataTypes.Iota()
		fmt.Println()
	case ArrayAndSliceTypes:
		// Arrays and Slices
		CompositeDataTypes.ArraysAndSlices()
		fmt.Println()
	case Pointers:
		CompositeDataTypes.Pointers()
		fmt.Println()
	case Operators:
		BasicDataTypes.OperatorsAndComparisons()
		fmt.Println()
	case Sets:
		CompositeDataTypes.SetsDemo()
		fmt.Println()
	case Maps:
		CompositeDataTypes.MapsDemo()
	case Structs:
		CompositeDataTypes.Structs()
	case IfElse:
		ControlFlow.IfElse()
	case Switch:
		ControlFlow.SwitchStatements()
	case Loops:
		ControlFlow.InfiniteLoop()
		fmt.Println()
		ControlFlow.LoopUntil()
		fmt.Println()
		ControlFlow.CounterBasedLoop()
		fmt.Println()
		ControlFlow.RangeBasedLoop()
		fmt.Println()
		ControlFlow.LoopingOverCollections()
		fmt.Println()
	case DeferredStatement:
		ControlFlow.DeferredFileIO()
		fmt.Println()
		ControlFlow.StackingDeferredCalls()
		fmt.Println()
	case Panic:
		ControlFlow.Panic()
		fmt.Println()
	case PanicAndRecover:
		ControlFlow.PanicAndRecover()
		fmt.Println()
	case Functions:
		CodeOrganization.AddDemo()
		fmt.Println()
		CodeOrganization.SwapDemo()
		fmt.Println()
		CodeOrganization.SplitDemo()
		fmt.Println()
		CodeOrganization.SumDemo()
		fmt.Println()
	case AnonymousFunctions:
		CodeOrganization.AnonymousFunction()
		fmt.Println()
		CodeOrganization.AssigningAnonymousFunctions()
		fmt.Println()
		CodeOrganization.AnonymousFunctionWithGoRoutine()
		fmt.Println()
		CodeOrganization.AnonymousFunctionAsClosure()
		fmt.Println()
	case GenericCustomConstraints:
		CodeOrganization.GenericCustomConstraintDemo()
		fmt.Println()
	case GenericFunctions:
		CodeOrganization.GenericFunctionsDemo()
		fmt.Println()
	case GenericConstraints:
		CodeOrganization.GenericFunctionWithConstraints()
		fmt.Println()
	case Methods:
		CodeOrganization.MethodsDemo()
		fmt.Println()
	case Interfaces:
		CodeOrganization.InterfaceDemo()
		fmt.Println()
	case Embedding:
		CodeOrganization.EmbeddingDemo()
		fmt.Println()
	case Composition:
		CodeOrganization.CompositionDemo()
		fmt.Println()
	}
}
