package main

import (
	"fmt"
	"os"
)

var exerciseMap = map[string]func(){
	"1":  Run01HelloWorldAndTypes,
	"2":  Run02ConstantDeclarations,
	"3":  Run03SliceCapacityAndGrowth,
	"4":  Run04ForRangeMapsAndFuncs,
	"5":  Run05LabeledBreakAndSwitch,
	"6":  Run06FunctionBasicDivision,
	"7":  Run07FunctionMapCalculator,
	"8":  Run08BinarySearchTree,
	"9":  Run09StructMethodsAndSorting,
	"10": Run10PointersAndOptionalFields,
	"11": Run11CliArgsAndDecimal,
	"12": Run12StructEmbeddingComposition,
	"13": Run13EmbeddedStructMethods,
	"14": Run14TypeAssertions,
	"15": Run15GenericsMapReduceFilter,
	"16": Run16GenericsMapSquare,
	"17": Run17GenericTypeConstraints,
	"18": Run18BenchmarkLoopPrint,
	"19": Run19PanicDeferRecover,
	"20": Run20ErrorWrappingAndUnwrap,
	"21": Run21ErrorJoiningValidation,
	"22": Run22ErrorWrappingDivide,
	"23": Run23PanicRecoverDivide,
	"24": Run24CryptoRandomSeeding,
	"25": Run25BenchmarkTimingAndSorting,
	"26": Run26GenericStackAndSliceSort,
	"27": Run27PointerVsValueReceivers,
	"28": Run28BinarySearch,
	"29": Run29ConcurrencyMutexWaitgroup,
	"30": Run30QueueSlice,
	"31": Run31StackGenericFunctional,
	"32": Run32DsaDataStructures,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("=== Go Learning Exercises Runner ===")
		fmt.Println("Usage: go run . <exercise_number>")
		fmt.Println("Available exercises: 1 to 32")
		fmt.Println("Example:")
		fmt.Println("  go run . 1   (Hello World & Types)")
		fmt.Println("  go run . 7   (Calculator with Op Map)")
		fmt.Println("  go run . 26  (Generic Stack & Slice Sorting)")
		return
	}

	choice := os.Args[1]
	if fn, exists := exerciseMap[choice]; exists {
		fmt.Printf("=== Running Exercise %s ===\n", choice)
		fn()
	} else {
		fmt.Printf("Exercise '%s' not found. Choose a number between 1 and 32.\n", choice)
	}
}
