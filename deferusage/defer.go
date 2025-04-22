package deferusage

func DeferPrinting() {
	// Defer is used to ensure that a function call is performed later in the program
	// Defer statements are executed in LIFO order (Last In First Out)
	defer println("World!")
	defer print("beautiful ")
	defer print(" you ")
	print("Hello ")
	// Output:
	// Hello you beautiful World!
}

func DeferWithPanic() {
	// Defer is also used to handle panic situations
	println("This is a text before panic.")
	defer println("This is a deferred text after panic.")
	panic("This is a panic!")
	println("This text will not be printed.")
	// Output:
	// This is a text before panic.
	// panic: This is a panic!
	// This is a deferred text after panic.
}

func ParameterHandlingWithDefer() {
	// Defer can be used to handle parameters in a function
	// The deferred function will use the value of the parameter at the time of the defer statement
	// Not at the time of execution
	for i := 0; i < 5; i++ {
		defer println(i)
		println(i)
	}
	println("------------------")
	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// ------------------
	// 4
	// 3
	// 2
	// 1
	// 0
}

func ChangingParameterHandlingWithDefer() {
	var text string = "World!"
	defer println(text)
	text = "Mars!"
	print("Hello ")
	println(text)
	// Output:
	// Hello Mars!
	// World!
}
