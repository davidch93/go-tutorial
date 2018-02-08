package main

import (
	"fmt"

	"github.com/davidch93/go-tutorial/basic-syntax/syntax"
)

func main() {
	fmt.Printf("Conditional Example:\n")
	syntax.SimpleConditional()

	fmt.Printf("\nLoop Example:\n")
	syntax.SimpleLoop()
	syntax.LoopPrintEvenNumber()
	syntax.MultipleInitialisationAndIncrement()

	fmt.Printf("\nSwitch Example:\n")
	syntax.SimpleSwitchCase()
	syntax.MultipleExpressionInCase()
	syntax.ConditionalExpressionInCase()
	syntax.FallthroughStatement()
	fmt.Printf("\n")
}
