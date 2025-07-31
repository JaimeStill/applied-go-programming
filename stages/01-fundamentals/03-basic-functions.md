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

## Prerequisites

Prerequisites for mastering Basic Functions:
- Basic program structure (package main, import, main function)
- Variable declarations and types
- Type conversions and basic operations

**Starting Point**: This exercise provides all necessary code from scratch. No prior exercise code is required.

## Practical Implementation

### Step 1: Initial Setup

Create the foundational structure for this exercise:

```go
package main

import (
    "fmt"
)

// Calculator function definitions
func add(a, b int) int      { return a + b }
func subtract(a, b int) int { return a - b }
func multiply(a, b int) int  { return a * b }
func divide(a, b int) int    { return a / b }

func main() {
    // Test our calculator functions
    fmt.Println("Calculator Functions Demo")
    fmt.Println("========================")
}
```

This establishes our program structure with the main package and fmt import for output operations.

### Step 2: Complete Implementation

Follow these steps to demonstrate Basic Functions:

```go
package main

import (
    "fmt"
)

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

Each function demonstrates the basic pattern: accept parameters, perform an operation, return a result. The main function shows various ways to call these functions and use their return values.

### Step 3: Verification

Execute your implementation to verify it works:

```bash
go build -o calculator main.go
./calculator
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

The output confirms that each function operates correctly, accepting parameters and returning the expected results. Note that integer division truncates decimal values.

## Mastery Verification

Complete this Exercise by achieving:

1. **Function Declaration**: Each function follows the `func name(param type) returnType` pattern
2. **Parameter Handling**: Functions accept input values through parameters
3. **Return Values**: Each function returns a single result value
4. **Function Calls**: The main function successfully calls each calculator function
5. **Result Usage**: Return values are captured and used in variables or expressions
6. **Compilation**: The program compiles without errors using `go build`
7. **Execution**: Running the program produces the expected calculator output

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

## Extension Exercise

Build additional calculator functions by adding this code:

```go
// Add these functions after the basic arithmetic functions
func remainder(a, b int) int {
    if b == 0 {
        fmt.Println("Warning: Division by zero, returning 0")
        return 0
    }
    return a % b
}

func power(base, exponent int) int {
    result := 1
    for i := 0; i < exponent; i++ {
        result = multiply(result, base)
    }
    return result
}
```

Then add these test cases to your main function:
```go
// Add to main function
fmt.Printf("%d %% %d = %d\n", x, y, remainder(x, y))
fmt.Printf("%d ^ %d = %d\n", 2, 5, power(2, 5))
```

## Complete Exercise Solution

*Try building the solution yourself before looking at this complete solution.*

<details>
<summary>Reference Implementation</summary>

```go
package main

import (
    "fmt"
)

func add(a int, b int) int {
    return a + b
}

func subtract(a int, b int) int {
    return a - b
}

func multiply(a int, b int) int {
    return a * b
}

func divide(a int, b int) int {
    if b == 0 {
        fmt.Println("Warning: Division by zero, returning 0")
        return 0
    }
    return a / b
}

// Enhancement functions
func remainder(a int, b int) int {
    if b == 0 {
        fmt.Println("Warning: Division by zero, returning 0")
        return 0
    }
    return a % b
}

func power(base int, exponent int) int {
    result := 1
    for i := 0; i < exponent; i++ {
        result = multiply(result, base)
    }
    return result
}

func main() {
    fmt.Println("Enhanced Calculator Functions")
    fmt.Println("============================")
    
    // Basic operations
    x, y := 10, 3
    fmt.Printf("%d + %d = %d\n", x, y, add(x, y))
    fmt.Printf("%d - %d = %d\n", x, y, subtract(x, y))
    fmt.Printf("%d * %d = %d\n", x, y, multiply(x, y))
    fmt.Printf("%d / %d = %d\n", x, y, divide(x, y))
    fmt.Printf("%d %% %d = %d\n", x, y, remainder(x, y))
    
    // Power function
    base, exp := 2, 5
    fmt.Printf("%d ^ %d = %d\n", base, exp, power(base, exp))
    
    // Edge cases
    fmt.Println("\nEdge cases:")
    fmt.Printf("10 / 0 = %d\n", divide(10, 0))
    fmt.Printf("10 %% 0 = %d\n", remainder(10, 0))
    fmt.Printf("5 ^ 0 = %d\n", power(5, 0))
}
```

**Key Implementation Notes**:
- Functions follow single responsibility principle - each does one operation
- Error handling prevents runtime panics from division by zero
- The power function demonstrates function reuse by calling multiply
- All functions maintain single parameter and single return value pattern
- Edge cases are handled gracefully with appropriate warnings

</details>

## Concept Progression

**Exercise Completed**: Basic function definitions with parameters and return values
**Next Exercise**: Advanced Function Features (multiple parameters, multiple returns, named returns)

---

**Exercise Metadata**
- **Created**: 2025-07-31
- **AECS Compliance**: Verified for Exercise Primacy, Concept Atomicity, Applied Understanding, and Progressive Complexity
- **Implementation Status**: Complete with reference solution