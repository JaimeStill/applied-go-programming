# Exercise: Basic Functions

## Exercise Metadata

- **Stage**: Fundamentals
- **Single Concept**: Basic Functions
- **Prerequisites**: Variables and types mastered
- **Estimated Time**: 1 hour

## What You'll Build

You'll create a simple calculator program using basic function definitions. Each arithmetic operation will be implemented as a separate function with parameters and return values, demonstrating Go's function syntax and usage patterns.

**Tangible Result**: A working calculator that performs addition, subtraction, multiplication, and division operations using well-defined functions.

## Conceptual Overview

Functions in Go encapsulate reusable logic into named blocks of code. A basic function accepts input through parameters, performs operations, and returns a result. This fundamental building block enables code organization, reusability, and logical separation of concerns.

Functions in Go use the syntax `func name(parameters) returnType` to define reusable code blocks that accept input and produce output.

## Foundation Requirements

This Exercise demonstrates Basic Functions which assumes familiarity with:
- Basic program structure (package main, import, main function)
- Variable declarations and types
- Type conversions and basic operations

**Starting Point**: This exercise provides all necessary code from scratch. No prior exercise code is required.

## Build Calculator Program

### Step 1: Create Function Definitions

Create `calculator.go` with basic function structure:

```go
package main

import "fmt"

// add takes two integers and returns their sum
func add(a int, b int) int {
    return a + b
}

// subtract takes two integers and returns their difference
func subtract(a int, b int) int {
    return a - b
}

// multiply takes two integers and returns their product
func multiply(a int, b int) int {
    return a * b
}

// divide takes two integers and returns their quotient
func divide(a int, b int) int {
    return a / b
}

func main() {
    fmt.Println("Calculator Functions Demo")
    fmt.Println("========================")
}
```

### Step 2: Add Function Calls and Testing

Add function calls to the main function:

```go
package main

import "fmt"

// Calculator function definitions (add, subtract, multiply, divide) from Step 1

func main() {
    fmt.Println("Calculator Functions Demo")
    fmt.Println("========================")
    
    // Test values
    x := 10
    y := 3
    
    // Call each function and display results
    fmt.Printf("%d + %d = %d\n", x, y, add(x, y))
    fmt.Printf("%d - %d = %d\n", x, y, subtract(x, y))
    fmt.Printf("%d * %d = %d\n", x, y, multiply(x, y))
    fmt.Printf("%d / %d = %d\n", x, y, divide(x, y))
    
    // Demonstrate function calls with different values
    fmt.Println("\nMore calculations:")
    result1 := add(25, 17)
    fmt.Printf("25 + 17 = %d\n", result1)
    
    result2 := multiply(8, 7)
    fmt.Printf("8 * 7 = %d\n", result2)
    
    // Chain function calls
    combined := add(multiply(4, 5), subtract(10, 3))
    fmt.Printf("(4 * 5) + (10 - 3) = %d\n", combined)
}
```

### Step 3: Execute and Verify

Run your complete calculator:

```bash
go run calculator.go
```

**Expected Output**:
```
Calculator Functions Demo
========================
10 + 3 = 13
10 - 3 = 7
10 * 3 = 30
10 / 3 = 3

More calculations:
25 + 17 = 42
8 * 7 = 56
(4 * 5) + (10 - 3) = 27
```

This output demonstrates mastery of Go's basic function concepts:
- **Function declaration**: Each function follows the `func name(param type) returnType` pattern correctly
- **Parameter handling**: Functions accept input values through named parameters (a, b)
- **Return values**: Each function returns a single integer result using the `return` statement
- **Function calls**: The main function successfully calls each calculator function with different arguments
- **Result usage**: Return values are captured in variables (`result1`, `result2`) and used directly in expressions
- **Expression chaining**: Functions can be chained together (`add(multiply(4, 5), subtract(10, 3))`) to create more complex calculations
- **Integer division**: Note that `10 / 3 = 3` demonstrates Go's integer division behavior (truncation)


## Common Implementation Issues

### Issue: Division by Zero

**Problem**: Calling `divide(10, 0)` causes a runtime panic with "integer divide by zero"

**Solution**: 
```go
func divide(a int, b int) int {
    if b == 0 {
        fmt.Println("Error: Division by zero")
        return 0
    }
    return a / b
}
```

### Issue: Multiple Parameters Syntax

**Problem**: Writing `func add(a, b int) int` instead of `func add(a int, b int) int`

**Solution**: 
```go
// Both forms are valid in Go:
func add(a int, b int) int { return a + b }  // Explicit
func add(a, b int) int { return a + b }      // Shortened

// Use the explicit form for clarity when learning
```

### Issue: Forgetting Return Statement

**Problem**: Function doesn't return a value when it should

**Solution**: 
```go
// Incorrect - missing return
func multiply(a int, b int) int {
    result := a * b
    // Forgot to return result
}

// Correct - includes return
func multiply(a int, b int) int {
    result := a * b
    return result
}
```

## Independent Challenge

Create a simple temperature converter that converts between Celsius and Fahrenheit using Go functions. Your program should demonstrate basic function definitions, parameters, return values, and function calls.

**Requirements**:
- Create a function to convert Celsius to Fahrenheit
- Create a function to convert Fahrenheit to Celsius  
- Add a function that doubles a temperature value
- Use only int parameters and return values (as covered in the exercise)
- Test your functions with different temperature values
- Display results using fmt.Printf formatting

### Independent Challenge Solution

*Try building the solution yourself before looking at this reference.*

<details>
<summary>Simple Temperature Converter Solution</summary>

```go
package main

import "fmt"

// celsiusToFahrenheit converts Celsius temperature to Fahrenheit
func celsiusToFahrenheit(celsius int) int {
	return (celsius * 9 / 5) + 32
}

// fahrenheitToCelsius converts Fahrenheit temperature to Celsius
func fahrenheitToCelsius(fahrenheit int) int {
	return (fahrenheit - 32) * 5 / 9
}

// doubleTemperature doubles any temperature value
func doubleTemperature(temp int) int {
	return temp * 2
}

func main() {
	fmt.Println("=== Simple Temperature Converter ===")
	fmt.Println("Demonstrating Basic Go Functions")
	fmt.Println()
	
	// Test temperatures
	testCelsius := 25
	testFahrenheit := 77
	
	fmt.Printf("Test temperatures: %d°C and %d°F\n\n", testCelsius, testFahrenheit)
	
	// Celsius to Fahrenheit conversion
	fahrenheitResult := celsiusToFahrenheit(testCelsius)
	fmt.Printf("Celsius to Fahrenheit: %d°C = %d°F\n", testCelsius, fahrenheitResult)
	
	// Fahrenheit to Celsius conversion
	celsiusResult := fahrenheitToCelsius(testFahrenheit)
	fmt.Printf("Fahrenheit to Celsius: %d°F = %d°C\n", testFahrenheit, celsiusResult)
	
	// Double temperature demonstrations
	doubleCelsius := doubleTemperature(testCelsius)
	doubleFahrenheit := doubleTemperature(testFahrenheit)
	fmt.Printf("Double temperature: %d°C doubled = %d°C\n", testCelsius, doubleCelsius)
	fmt.Printf("Double temperature: %d°F doubled = %d°F\n", testFahrenheit, doubleFahrenheit)
	
	// More function call examples
	fmt.Println("\nAdditional Tests:")
	result1 := celsiusToFahrenheit(0)
	fmt.Printf("Water freezing point: 0°C = %d°F\n", result1)
	
	result2 := fahrenheitToCelsius(32)
	fmt.Printf("Water freezing point: 32°F = %d°C\n", result2)
	
	result3 := doubleTemperature(10)
	fmt.Printf("Double 10 degrees = %d degrees\n", result3)
	
	fmt.Println("\n✅ Temperature converter completed successfully!")
}
```

**Key Implementation Notes**:
- **Function Definitions**: Each function uses the `func name(param type) returnType` syntax from the exercise
- **Parameter Handling**: Functions accept int parameters exactly as shown in the calculator exercise  
- **Return Values**: All functions return int values using the `return` statement
- **Function Calls**: Main function calls each function and stores results in variables
- **Arithmetic Operations**: Uses basic operations (+, -, *, /) within functions as covered
- **No New Concepts**: Only uses concepts from current exercise (functions) and prior exercises (variables, basic program structure, fmt functions)
- **Simple Focus**: Concentrates on basic function usage without advanced concepts like arrays or loops

</details>

## Concept Progression

**Exercise Completed**: Basic function definitions with parameters and return values
**Next Exercise**: Advanced Function Features (multiple parameters, multiple returns, named returns)

---

**Exercise Metadata**
- **Created**: 2025-07-31
- **AECS Compliance**: Verified for Exercise Primacy, Concept Atomicity, Applied Understanding, and Progressive Complexity
- **Implementation Status**: Complete with reference solution