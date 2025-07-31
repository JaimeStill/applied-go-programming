# Exercise: Basic Program Structure

## Exercise Metadata

- **Stage**: Fundamentals
- **Single Concept**: Basic Program Structure
- **Prerequisites**: Go installation and command-line familiarity
- **Estimated Time**: 20 minutes

## What You'll Build

You'll create a Go program that demonstrates the three essential components of every Go executable: package declaration, imports, and the main function. The program will display system information to prove it's running successfully.

**Tangible Result**: A working Go program that executes and displays your Go version, operating system, architecture, and CPU count.

## Conceptual Overview

Every Go executable program follows a specific structure with three mandatory components. The `package main` declaration tells Go this is an executable program, not a library. The import section brings in external packages that provide functionality beyond the core language. The `func main()` is the entry point where program execution begins - without it, Go won't know where to start running your code.

## Foundation Requirements

Verify Go installation:

```bash
go version
```

Expected output showing Go 1.24+:
```
go version go1.24.0 linux/amd64
```

**Starting Point**: This exercise provides all necessary code from scratch. No prior exercise code is required.

## Build Program Structure

### Step 1: Create Program File

Create `hello-system.go`:

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("=== Go Program Structure Demo ===")
	fmt.Printf("Go Version: %s\n", runtime.Version())
	fmt.Printf("Operating System: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("CPU Count: %d\n", runtime.NumCPU())
}
```

### Step 2: Execute and Verify

Run your complete program:

```bash
go run hello-system.go
```

**Expected Output**:
```
=== Go Program Structure Demo ===
Go Version: go1.24.0
Operating System: linux
Architecture: amd64
CPU Count: 8
```

This output demonstrates mastery of Go's basic program structure:
- **package main**: Defines an executable program (not a library)
- **import section**: Brings in required packages (fmt for printing, runtime for system info)
- **func main()**: Entry point where execution begins
- **Successful execution**: Program compiles and runs, producing system information output

## Independent Challenge

Create a personal Go environment reporter that displays basic information about your Go installation and system. Your program should demonstrate proper Go program structure using only the concepts covered in this exercise.

**Requirements**:
- Use proper package declaration for an executable program
- Import fmt and runtime packages only (as covered in the exercise)
- Implement main function as the program entry point
- Display Go version using runtime.Version()
- Show operating system using runtime.GOOS
- Present architecture using runtime.GOARCH
- Display CPU count using runtime.NumCPU()
- Format output with fmt.Println() and fmt.Printf()
- Include a program title and clear section headers


### Independent Challenge Solution

*Try building the solution yourself before looking at this reference.*

<details>
<summary>Go Environment Reporter Solution</summary>

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("=== Go Environment Reporter ===")
	fmt.Println("Displaying basic system information")
	fmt.Println()
	
	fmt.Println("Go Installation Details:")
	fmt.Printf("Go Version: %s\n", runtime.Version())
	fmt.Printf("Operating System: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("CPU Count: %d\n", runtime.NumCPU())
	
	fmt.Println()
	fmt.Println("Program Structure Verification:")
	fmt.Println("✓ package main - Executable program")
	fmt.Println("✓ import section - fmt and runtime packages")
	fmt.Println("✓ func main() - Program entry point")
	fmt.Println("✓ Output statements - Information display")
	
	fmt.Println()
	fmt.Println("Environment report completed successfully!")
}
```

**Key Implementation Notes**:
- **Package Declaration**: Uses `package main` to create an executable program (as demonstrated in the exercise)
- **Import Section**: Only imports `fmt` and `runtime` packages that were covered in the main exercise
- **Main Function**: Implements `func main()` as the program entry point exactly as shown in the exercise
- **Output Functions**: Uses only `fmt.Println()` and `fmt.Printf()` functions demonstrated in the exercise
- **Runtime Functions**: Uses only the four runtime functions covered: Version(), GOOS, GOARCH, and NumCPU()
- **Program Structure**: Follows the exact three-component structure taught in the exercise
- **No New Concepts**: Avoids any packages, functions, or patterns not yet introduced to learners

</details>

## Path Integration

**Concept Demonstrated**: Basic Program Structure
**Next Concept in Path**: Variables and Type Declarations

---

**Exercise Metadata**
- **Created**: 2025-07-30
- **AECS Compliance**: Verified for Exercise Primacy, Concept Atomicity, Applied Understanding, and Progressive Complexity
- **Implementation Status**: Complete and Verified