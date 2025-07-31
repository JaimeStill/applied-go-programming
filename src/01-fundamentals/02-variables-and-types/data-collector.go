package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func demonstrateScope() {
	localVar := "I'm local to this function"
	fmt.Printf("Inside function - localVar: %v\n", localVar)
}

func main() {
	fmt.Println("=== Go Variables and Types Demo ===")

	// Method 1: var declaration with type
	var name string = "Alice"
	var age int = 30
	var temperature float64 = 98.6
	var isActive bool = true

	fmt.Printf("name: %v (type: %T)\n", name, name)
	fmt.Printf("age: %v (type: %T)\n", age, age)
	fmt.Printf("temperature: %v (type: %T)\n", temperature, temperature)
	fmt.Printf("isActive: %v (type: %T)\n", isActive, isActive)

	// Method 2: var with type inference
	var city = "New York"
	var score = 95.5
	var count = 100

	fmt.Println("\n2. Type Inference:")
	fmt.Printf("city: %v (type: %T)\n", city, city)
	fmt.Printf("score: %v (type: %T)\n", score, score)
	fmt.Printf("count: %v (type: %T)\n", count, count)

	// Method 3: short declaration
	country := "USA"
	rating := 4.5
	total := 250

	fmt.Println("\n3. Short Declaration:")
	fmt.Printf("country: %v (type: %T)\n", country, country)
	fmt.Printf("rating: %v (type: %T)\n", rating, rating)
	fmt.Printf("total: %v (type: %T)\n", total, total)

	// Zero values
	var emptyString string
	var zeroInt int
	var zeroFloat float64
	var falseBool bool

	fmt.Println("\n4. Zero Values:")
	fmt.Printf("emptyString: '%v' (type: %T)\n", emptyString, emptyString)
	fmt.Printf("zeroInt: %v (type: %T)\n", zeroInt, zeroInt)
	fmt.Printf("zeroFloat: %v (type: %T)\n", zeroFloat, zeroFloat)
	fmt.Printf("falseBool: %v (type: %T)\n", falseBool, falseBool)

	// Constants
	const pi float64 = 3.14159
	const maxRetries = 3
	const greeting = "Welcome"

	fmt.Println("\n5. Constants:")
	fmt.Printf("pi: %v (type: %T)\n", pi, pi)
	fmt.Printf("maxRetries: %v (type: %T)\n", maxRetries, maxRetries)
	fmt.Printf("greeting: %v (type: %T)\n", greeting, greeting)

	// Type sizes
	fmt.Println("\n6. Type Sizes:")
	fmt.Printf("size of int: %d bytes\n", unsafe.Sizeof(age))
	fmt.Printf("size of float64: %d bytes\n", unsafe.Sizeof(temperature))
	fmt.Printf("size of bool: %d bytes\n", unsafe.Sizeof(isActive))
	fmt.Printf("size of string: %d bytes\n", unsafe.Sizeof(name))

	// Type conversions
	var intValue int = 42
	var floatValue float64 = float64(intValue)
	var backToInt int = int(floatValue)

	fmt.Println("\n7. Numeric Type Conversions:")
	fmt.Printf("intValue: %v (type: %T)\n", intValue, intValue)
	fmt.Printf("floatValue: %v (type: %T)\n", floatValue, floatValue)
	fmt.Printf("backToInt: %v (type: %T)\n", backToInt, backToInt)

	// String conversions
	stringAge := strconv.Itoa(age)
	parsedAge, _ := strconv.Atoi(stringAge)

	fmt.Println("\n8. Stirng Conversions:")
	fmt.Printf("age as string: '%v' (type: %T)\n", stringAge, stringAge)
	fmt.Printf("parsed back to int: %v (type: %T)\n", parsedAge, parsedAge)

	// Variable scope demonstration
	fmt.Println("\n9. Variable Scope:")
	demonstrateScope()
	fmt.Println("Main function continues after demonstrateScope()")
}
