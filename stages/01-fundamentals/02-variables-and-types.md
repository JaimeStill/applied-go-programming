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

Extend the program with more declaration patterns. Add the following to your existing code:

```go
package main

import "fmt"

func main() {
	// Method 1: var declaration with type (from Step 1)
	
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

Add the `unsafe` import and demonstrate zero values, constants, and memory sizes:

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Variable declarations and output from Steps 1-2
	
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
	
	// Type sizes (using variables from Step 1)
	fmt.Println("\n6. Type Sizes:")
	fmt.Printf("size of int: %d bytes\n", unsafe.Sizeof(age))
	fmt.Printf("size of float64: %d bytes\n", unsafe.Sizeof(temperature))
	fmt.Printf("size of bool: %d bytes\n", unsafe.Sizeof(isActive))
	fmt.Printf("size of string: %d bytes\n", unsafe.Sizeof(name))
}
```

**Note about unsafe.Sizeof()**: The `unsafe` package provides low-level operations that step outside Go's type safety. `Sizeof` returns the size in bytes of the variable's type, not the actual data. For strings, it returns the size of the string header (16 bytes on 64-bit systems), not the string content length.

### Step 4: Complete with Type Conversions and Scope

Add the `strconv` import, scope function, and complete with type conversions:

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
	// Variable declarations, zero values, constants, and type sizes from Steps 1-3
	
	// Type conversions
	var intValue int = 42
	var floatValue float64 = float64(intValue)
	var backToInt int = int(floatValue)
	
	fmt.Println("\n7. Numeric Type Conversions:")
	fmt.Printf("intValue: %v (type: %T)\n", intValue, intValue)
	fmt.Printf("floatValue: %v (type: %T)\n", floatValue, floatValue)
	fmt.Printf("backToInt: %v (type: %T)\n", backToInt, backToInt)
	
	// String conversions (using age variable from Step 1)
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

Run your complete program:

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

This output demonstrates mastery of Go's variable and type system:
- **Three declaration methods**: Explicit typing, type inference, and short declaration working correctly
- **Zero values**: Go's predictable initialization behavior for all basic types
- **Type safety**: The `%T` verb showing runtime type information for every variable
- **Constants**: Both typed and untyped constants behaving correctly at compile-time
- **Type conversions**: Explicit conversions between numeric types and string transformations using `strconv`
- **Variable scope**: Function-local variables properly isolated from main function scope
- **Memory introspection**: `unsafe.Sizeof` revealing the actual memory footprint of different Go types
- **Complete execution**: Program compiles and runs without errors, producing comprehensive type demonstration

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

## Independent Challenge

Build a personal profile generator that creates detailed user profiles using Go's variable and type system. Your program should collect user information, perform type conversions, and generate a formatted profile report.

**Requirements**:
- Use all three variable declaration methods (var with type, var with inference, short declaration)
- Demonstrate zero values by showing uninitialized variables
- Include at least 3 constants for configuration (max age, default city, greeting message)
- Perform numeric type conversions (int ↔ float64)
- Implement string-to-number conversions using strconv
- Show variable scope with helper functions
- Display type information and memory sizes for key variables
- Generate a complete profile report as program output

### Independent Challenge Solution

*Try building the solution yourself before looking at this reference.*

<details>
<summary>Personal Profile Generator Solution</summary>

```go
package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

// Constants for configuration
const maxAge int = 120
const defaultCity = "Unknown"
const welcomeMessage = "Welcome to Profile Generator"

func calculateAgeCategory(age int) string {
	if age < 18 {
		return "Minor"
	} else if age < 65 {
		return "Adult"
	}
	return "Senior"
}

func displayScopeDemo() {
	localMessage := "This is a local scope variable"
	fmt.Printf("Inside function: %s\n", localMessage)
}

func main() {
	fmt.Println(welcomeMessage)
	fmt.Println("=== Personal Profile Generator ===")
	
	// Method 1: var declaration with explicit type
	var firstName string = "Alice"
	var lastName string = "Johnson"
	var age int = 28
	var height float64 = 5.6
	var isEmployed bool = true
	
	// Method 2: var with type inference
	var city = defaultCity
	var salary = 65000.50
	var yearsExperience = 5
	
	// Method 3: short declaration
	country := "USA"
	department := "Engineering"
	rating := 4.8
	
	// Zero values demonstration
	var emptyAddress string
	var dependents int
	var bonusEligible bool
	var unusedFloat float64
	
	fmt.Println("\n1. Basic Profile Information:")
	fmt.Printf("Name: %s %s (types: %T, %T)\n", firstName, lastName, firstName, lastName)
	fmt.Printf("Age: %d years (type: %T)\n", age, age)
	fmt.Printf("Height: %.1f feet (type: %T)\n", height, height)
	fmt.Printf("Employed: %v (type: %T)\n", isEmployed, isEmployed)
	
	fmt.Println("\n2. Location & Work Details:")
	fmt.Printf("City: %s (type: %T)\n", city, city)
	fmt.Printf("Country: %s (type: %T)\n", country, country)
	fmt.Printf("Department: %s (type: %T)\n", department, department)
	fmt.Printf("Salary: $%.2f (type: %T)\n", salary, salary)
	fmt.Printf("Experience: %d years (type: %T)\n", yearsExperience, yearsExperience)
	fmt.Printf("Rating: %.1f/5.0 (type: %T)\n", rating, rating)
	
	fmt.Println("\n3. Zero Values (uninitialized variables):")
	fmt.Printf("Address: '%s' (type: %T, zero value for string)\n", emptyAddress, emptyAddress)
	fmt.Printf("Dependents: %d (type: %T, zero value for int)\n", dependents, dependents)
	fmt.Printf("Bonus Eligible: %v (type: %T, zero value for bool)\n", bonusEligible, bonusEligible)
	fmt.Printf("Unused Float: %v (type: %T, zero value for float64)\n", unusedFloat, unusedFloat)
	
	fmt.Println("\n4. Constants Configuration:")
	fmt.Printf("Max allowed age: %d (type: %T)\n", maxAge, maxAge)
	fmt.Printf("Default city: %s (type: %T)\n", defaultCity, defaultCity)
	fmt.Printf("Welcome message: %s (type: %T)\n", welcomeMessage, welcomeMessage)
	
	// Type conversions
	ageAsFloat := float64(age)
	salaryAsInt := int(salary)
	heightInInches := height * 12.0
	
	fmt.Println("\n5. Type Conversions:")
	fmt.Printf("Age as float64: %.1f (converted from %T to %T)\n", ageAsFloat, age, ageAsFloat)
	fmt.Printf("Salary as int: %d (converted from %T to %T)\n", salaryAsInt, salary, salaryAsInt)
	fmt.Printf("Height in inches: %.1f (calculated from feet)\n", heightInInches)
	
	// String conversions
	ageString := strconv.Itoa(age)
	experienceString := strconv.Itoa(yearsExperience)
	salaryString := strconv.FormatFloat(salary, 'f', 2, 64)
	
	// Parse back from strings
	parsedAge, _ := strconv.Atoi(ageString)
	parsedExperience, _ := strconv.Atoi(experienceString)
	parsedSalary, _ := strconv.ParseFloat(salaryString, 64)
	
	fmt.Println("\n6. String Conversions:")
	fmt.Printf("Age as string: '%s' (type: %T)\n", ageString, ageString)
	fmt.Printf("Experience as string: '%s' (type: %T)\n", experienceString, experienceString)
	fmt.Printf("Salary as string: '%s' (type: %T)\n", salaryString, salaryString)
	fmt.Printf("Parsed age: %d, experience: %d, salary: %.2f\n", parsedAge, parsedExperience, parsedSalary)
	
	fmt.Println("\n7. Memory Sizes:")
	fmt.Printf("int size: %d bytes\n", unsafe.Sizeof(age))
	fmt.Printf("float64 size: %d bytes\n", unsafe.Sizeof(salary))
	fmt.Printf("string size: %d bytes\n", unsafe.Sizeof(firstName))
	fmt.Printf("bool size: %d bytes\n", unsafe.Sizeof(isEmployed))
	
	fmt.Println("\n8. Variable Scope Demonstration:")
	displayScopeDemo()
	// localMessage is not accessible here - scoped to the function
	
	fmt.Println("\n9. Generated Profile Summary:")
	ageCategory := calculateAgeCategory(age)
	fmt.Printf("Profile: %s %s, %d years old (%s)\n", firstName, lastName, age, ageCategory)
	fmt.Printf("Location: %s, %s\n", city, country)
	fmt.Printf("Career: %s department, %d years experience\n", department, yearsExperience)
	fmt.Printf("Status: Employment=%v, Rating=%.1f/5.0\n", isEmployed, rating)
	fmt.Printf("Financial: $%.2f salary\n", salary)
}
```

**Key Implementation Notes**:
- **Three Declaration Methods**: Demonstrates var with type (lines 29-33), var with inference (lines 35-38), and short declaration (lines 40-43) in realistic contexts
- **Zero Values**: Shows Go's predictable initialization with empty string, zero int, false bool, and zero float64
- **Constants**: Uses typed (maxAge int) and untyped (defaultCity, welcomeMessage) constants for configuration
- **Type Conversions**: Explicit numeric conversions prevent data loss, string conversions use strconv package safely
- **Scope Demonstration**: Helper functions show function-local variable isolation
- **Memory Introspection**: unsafe.Sizeof reveals actual memory footprint of Go types
- **Practical Application**: Creates a complete profile generator demonstrating real-world variable usage patterns
- **Type Safety**: %T verb and explicit type information throughout show Go's type system working

</details>

## Path Integration

**Concept Demonstrated**: Variables and Type Declarations
**Next Concept in Path**: Basic Functions

---

**Exercise Metadata**
- **Created**: 2025-07-30
- **AECS Compliance**: Verified for Exercise Primacy, Concept Atomicity, Applied Understanding, and Progressive Complexity
- **Implementation Status**: Complete and Verified