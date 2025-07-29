# Basic Program Structure and Compilation

## Exercise Metadata

- **Stage**: Fundamentals
- **Single Concept**: Basic Program Structure and Compilation
- **Prerequisites**: Go installation and basic command-line familiarity
- **Estimated Time**: 30 minutes

## Building Challenge

Build a "Hello World" program using Go's basic program structure and compilation process. Construct a program with package declaration, import statements, and main function that produces system information output.

**Tangible Result**: Working Go program that compiles and executes, outputting Go version string, build information, and runtime architecture

## Dependencies Verification

This Exercise uses only these previously mastered Concepts:
- Go installation and development environment setup (mastered through hands-on setup)
- Basic command-line navigation and file creation (mastered through hands-on practice)

**Verify your dependencies by running:**
```bash
go version
which go
```

## Practical Implementation

### Step 1: Initial Setup

Create a new directory for your first Go program:

```bash
mkdir hello-go
cd hello-go
```

Create an empty file named `main.go` and build the basic Go program structure yourself:

```go
// TODO: Add package declaration for an executable program
// TODO: Add import statement for formatted output
// TODO: Add main function that prints "Hello, Go!"
```

**Build this code and verify it compiles:**
```bash
go run main.go
```

### Step 2: Core Implementation

**Your Task**: Build the working implementation that demonstrates Basic Program Structure and Compilation by adding system information output.

Modify your `main.go` file:

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    // Build the implementation that outputs system information
    // Research Go's runtime package for system details
    // Use fmt package for formatted output
}
```

**Requirements**:
- Use package main declaration at the top of the file
- Import both fmt and runtime packages using proper Go import syntax
- Implement main function as the program entry point
- Use fmt.Println for the greeting message
- Use fmt.Printf with formatting to display system information
- Output must show Go version, OS, and architecture information

### Step 3: Verification

**Test your implementation:**

```bash
go run main.go
```

**Expected Output Pattern**:
```
Hello, Go!
Go version: [your Go version]
Operating System: [your OS]
Architecture: [your architecture]
```
*Note: Actual values will vary based on your system configuration*

## Success Criteria

Your implementation is complete when:

- [ ] You implement a program that compiles without errors using `go run main.go`
- [ ] You create output functionality that displays "Hello, Go!" as the first line
- [ ] You implement Go version display using runtime.Version()
- [ ] You construct OS information display using runtime.GOOS
- [ ] You build architecture information display using runtime.GOARCH
- [ ] You implement properly formatted and readable output
- [ ] You construct proper Go package structure with package main
- [ ] You create correctly formatted and necessary import statements

## Common Implementation Issues

### Issue: Package Declaration Missing or Incorrect

**Problem**: Forgetting the package main declaration or using the wrong package name causes compilation errors.

**Solution**: 
```go
// Correct - always start with package main for executable programs
package main

import "fmt"

func main() {
    // your code here
}
```

### Issue: Import Statement Syntax Errors

**Problem**: Incorrect import syntax or missing imports for functions you're using.

**Solution**: 
```go
package main

// Correct import syntax for multiple packages
import (
    "fmt"
    "runtime"
)

// Alternative single import format
// import "fmt"
// import "runtime"

func main() {
    fmt.Println("Using fmt package")
    fmt.Println(runtime.Version())
}
```

### Issue: Main Function Not Properly Defined

**Problem**: Incorrect main function signature or placement outside package main.

**Solution**: 
```go
package main

import "fmt"

// Correct - main function takes no parameters and returns nothing
func main() {
    fmt.Println("Program starts here")
}

// Incorrect examples:
// func main() string { return "wrong" }  // Wrong: main shouldn't return
// func main(args []string) {}            // Wrong: main takes no parameters
```

## Implementation Extension

**Optional Challenge**: Create a compiled binary using build command

**Additional Requirements**:
- Use `go build` to create an executable binary instead of `go run`
- Add simple formatted output variations
- Experiment with different compilation flags

## Reference Implementation

*Try building the solution yourself before looking at this reference.*

<details>
<summary>Reference Implementation</summary>

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("Hello, Go!")
    fmt.Printf("Go version: %s\n", runtime.Version())
    fmt.Printf("Operating System: %s\n", runtime.GOOS)
    fmt.Printf("Architecture: %s\n", runtime.GOARCH)
}
```

**Key Implementation Notes**:
- `package main` declares this as an executable program (not a library)
- `import` statements must come after package declaration and before any other code
- `func main()` is the required entry point for Go programs
- `fmt.Println` automatically adds a newline after output
- `fmt.Printf` allows formatted output with placeholders like %s for strings
- `runtime` package provides access to Go runtime information
- Multiple import statements can be grouped using parentheses for cleaner code

</details>

## Next Exercise

**Progressive Complexity**: This Exercise demonstrates Basic Program Structure and Compilation, which enables the next Concept: **Variables and Type Declarations** to be demonstrated through its Exercise.

**Next Building Task**: Construct a data collection program implementing Go variable declaration patterns and type conversions

---

**Exercise Metadata**
- **Created**: 2025-07-29
- **AECS Compliance**: Verified for Exercise Primacy, Concept Atomicity, Applied Understanding, and Progressive Complexity
- **Implementation Status**: Complete and Ready for Implementation