package main

import "fmt"

func main() {
	// Call the while-style loop example
	whileStyleLoop()

	fmt.Println("-----")

	// Call the standard for-loop example
	standardForLoop()

	fmt.Println("-----")

	// Call the range loop example
	rangeloop()
}

// This function demonstrates how Go's 'for' loop
// can act exactly like a 'while' loop in other languages.
// It only checks the condition before each iteration.
func whileStyleLoop() {
	secondsLeft := 3

	for secondsLeft > 0 {
		fmt.Println(secondsLeft, "seconds remaining...")
		secondsLeft-- // We must manually change the variable, or the loop never ends!
	}

	fmt.Println("Liftoff! 🚀")
}

// This function demonstrates a standard 'for' loop.
// A standard 'for' loop has three parts:
// 1. Initialization (i := 1): Runs once before the loop starts. We set a counter variable.
// 2. Condition (i <= 5): Checked before every iteration. If true, the loop runs. If false, it stops.
// 3. Post-statement (i++): Runs at the end of every iteration. We increase the counter by 1.
//
// There are no while or do-while loops in Go;
// the 'for' loop serves all these purposes.
func standardForLoop() {
	fmt.Println("Starting the countdown...")

	for i := 1; i <= 5; i++ {
		// This code block runs repeatedly as long as i <= 5
		fmt.Println("Count is:", i)
	}

	fmt.Println("Loop finished!")
}

func rangeloop() {
	for i := range 5 {
		fmt.Println("Count is", i)
	}
}
