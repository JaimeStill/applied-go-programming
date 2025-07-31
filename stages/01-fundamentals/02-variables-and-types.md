# Exercise: Variables and Type Declarations

## Exercise Metadata

- **Stage**: Fundamentals
- **Single Concept**: Variables and Type Declarations
- **Prerequisites**: Basic program structure mastered
- **Estimated Time**: 40 minutes

## What You'll Build

You'll create a comprehensive data collection program that demonstrates Go's variable and type system through practical application. The program will showcase different variable declaration methods, type safety, memory characteristics, and scope rules.

**Tangible Result**: A working Go program that outputs detailed information about variables, their types, memory usage, and behavior, providing a foundation for understanding Go's type system.

## Conceptual Overview

Go's variable and type system forms the foundation of all Go programs. Unlike dynamically typed languages, Go requires explicit or inferred type declarations, providing compile-time safety and runtime performance. Variables in Go can be declared using three distinct syntaxes, each serving different use cases: explicit typing for clarity, type inference for convenience, and short declaration for local scope efficiency.

Understanding Go's type system is crucial because it influences memory layout, performance characteristics, and API design patterns. Go's zero values ensure variables are always initialized to predictable states, eliminating entire classes of bugs common in other languages. Constants provide compile-time optimization opportunities, while type conversions ensure safe data transformation between compatible types.

This exercise builds a data collection program that demonstrates these concepts through practical application, showing how variables and types work together to create maintainable, efficient Go code.

## Foundation Requirements

Verify your Go environment is working:

```bash
go version
```

**Starting Point**: This exercise provides all necessary code from scratch. No prior exercise code is required.

## Build Data Collection Program

### Step 1: Create Variable Declaration Demo

Create `data-collector.go`:

```go
package main

import "fmt"

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
}
```

### Step 2: Add Type Inference and Short Declaration

Extend the program with more declaration patterns:

```go
package main

import "fmt"

func main() {
	fmt.Println("=== Go Variables and Types Demo ===")
	
	// Method 1: var declaration with type
	var name string = "Alice"
	var age int = 30
	var temperature float64 = 98.6
	var isActive bool = true
	
	fmt.Println("\n1. Explicit Type Declaration:")
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
}
```

### Step 3: Add Zero Values, Constants, and Type Sizes

Demonstrate zero values, constants, and memory sizes:

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("=== Go Variables and Types Demo ===")
	
	// Method 1: var declaration with type
	var name string = "Alice"
	var age int = 30
	var temperature float64 = 98.6
	var isActive bool = true
	
	fmt.Println("\n1. Explicit Type Declaration:")
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
}
```

**Note about unsafe.Sizeof()**: The `unsafe` package provides low-level operations that step outside Go's type safety. `Sizeof` returns the size in bytes of the variable's type, not the actual data. For strings, it returns the size of the string header (16 bytes on 64-bit systems), not the string content length.

### Step 4: Complete with Type Conversions and Scope

Add type conversions and demonstrate variable scope:

```go
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
	
	fmt.Println("\n1. Explicit Type Declaration:")
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
	
	fmt.Println("\n8. String Conversions:")
	fmt.Printf("age as string: '%v' (type: %T)\n", stringAge, stringAge)
	fmt.Printf("parsed back to int: %v (type: %T)\n", parsedAge, parsedAge)
	
	// Variable scope demonstration
	fmt.Println("\n9. Variable Scope:")
	demonstrateScope()
	// localVar is not accessible here - it's scoped to the function
	fmt.Println("Main function continues after demonstrateScope()")
}
```

### Step 5: Execute and Verify

Run the complete program:

```bash
go run data-collector.go
```

**Expected Output**:
```
=== Go Variables and Types Demo ===

1. Explicit Type Declaration:
name: Alice (type: string)
age: 30 (type: int)
temperature: 98.6 (type: float64)
isActive: true (type: bool)

2. Type Inference:
city: New York (type: string)
score: 95.5 (type: float64)
count: 100 (type: int)

3. Short Declaration:
country: USA (type: string)
rating: 4.5 (type: float64)
total: 250 (type: int)

4. Zero Values:
emptyString: '' (type: string)
zeroInt: 0 (type: int)
zeroFloat: 0 (type: float64)
falseBool: false (type: bool)

5. Constants:
pi: 3.14159 (type: float64)
maxRetries: 3 (type: int)
greeting: Welcome (type: string)

6. Type Sizes:
size of int: 8 bytes
size of float64: 8 bytes
size of bool: 1 bytes
size of string: 16 bytes

7. Numeric Type Conversions:
intValue: 42 (type: int)
floatValue: 42 (type: float64)
backToInt: 42 (type: int)

8. String Conversions:
age as string: '30' (type: string)
parsed back to int: 30 (type: int)

9. Variable Scope:
Inside function - localVar: I'm local to this function
Main function continues after demonstrateScope()
```

## Implementation Verification

Your data collection program demonstrates mastery when it compiles and runs successfully, producing output that shows:

- All variable declarations working correctly (explicit type, inference, short declaration)
- Proper zero values displayed for each basic Go type (empty strings, zero numbers, false booleans)
- Constants behaving correctly with both typed and untyped examples
- Type conversions executing without errors, showing numeric and string transformations
- Variable scope working as expected, with function-local variables properly isolated
- Memory size information accurately reported using unsafe.Sizeof for different types
- Clean, formatted output displaying both values and their corresponding types (%T verb)
- The program running from start to finish without compilation errors or runtime panics

Execute `go run data-collector.go` and verify the output matches the expected results shown in Step 5.

## Common Implementation Issues

### Issue: Short Declaration Outside Function Scope

**Problem**: Using `:=` at package level causes compilation error

**Solution**: 
```go
// Incorrect - causes compilation error
// country := "USA"  // outside function

// Correct - use var at package level  
var country = "USA"
```

### Issue: Type Conversion Precision Loss

**Problem**: Converting float64 to int truncates decimal values

**Solution**:
```go
var temperature float64 = 98.6
var rounded int = int(temperature + 0.5) // Add 0.5 for rounding
fmt.Printf("temperature: %.1f, rounded: %d\n", temperature, rounded)
```

### Issue: Unused Variable Compilation Error

**Problem**: Declared variables that aren't used cause compilation failures

**Solution**:
```go
// Incorrect - unused variable causes error
var unusedVar string = "test"

// Correct - use blank identifier or remove declaration
_ = unusedVar  // Explicitly ignore
// Or simply don't declare unused variables
```

## Implementation Extension

**Enhancement Exercise**: Create a type-safe configuration system

**Additional Requirements**:
- Define custom types for different configuration values (e.g., `type Port int`, `type DatabaseURL string`)
- Implement validation functions for each custom type
- Create a configuration struct with typed fields
- Demonstrate type safety by preventing invalid assignments
- Show how custom types can have methods attached

## Verification

Program demonstrates:
- Three variable declaration methods (var with type, var with inference, short declaration)
- All basic Go types (string, int, float64, bool)
- Zero values for each type
- Constants (typed and untyped)
- Type sizes in bytes using unsafe.Sizeof
- Type conversions between numeric types
- String conversions using strconv package
- Variable scope within functions
- Type information output using %T verb
- Clean, formatted output showing values and their types

## Reference Implementation

*Try building the solution yourself before looking at this reference.*

<details>
<summary>Complete Reference Implementation</summary>

```go
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
	
	fmt.Println("\n1. Explicit Type Declaration:")
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
	
	fmt.Println("\n8. String Conversions:")
	fmt.Printf("age as string: '%v' (type: %T)\n", stringAge, stringAge)
	fmt.Printf("parsed back to int: %v (type: %T)\n", parsedAge, parsedAge)
	
	// Variable scope demonstration
	fmt.Println("\n9. Variable Scope:")
	demonstrateScope()
	// localVar is not accessible here - it's scoped to the function
	fmt.Println("Main function continues after demonstrateScope()")
}
```

**Key Implementation Notes**:
- All three variable declaration methods demonstrate different use cases and scoping rules
- Zero values provide predictable initialization behavior for all Go types
- Constants are evaluated at compile-time and can be typed or untyped
- The `unsafe` package provides low-level memory introspection capabilities
- Type conversions must be explicit in Go, preventing accidental data loss
- The `strconv` package handles string-to-numeric conversions safely
- Variable scope is function-based, with package-level variables accessible globally
- The `%T` verb shows runtime type information for any value

</details>

## Path Integration

**Concept Demonstrated**: Variables and Type Declarations
**Next Concept in Path**: Basic Functions

---

**Exercise Metadata**
- **Created**: 2025-07-30
- **AECS Compliance**: Verified for Exercise Primacy, Concept Atomicity, Applied Understanding, and Progressive Complexity
- **Implementation Status**: Complete and Verified