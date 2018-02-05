package syntax

import (
	"fmt"
)

// SimpleLoop example.
func SimpleLoop() {
	i := 0
	for i <= 10 {
		fmt.Printf("%d ", i)
		i += 2
	}
	fmt.Printf("\n")
}

// LoopPrintEvenNumber example.
func LoopPrintEvenNumber() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")
}

// MultipleInitialisationAndIncrement example.
func MultipleInitialisationAndIncrement() {
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}
