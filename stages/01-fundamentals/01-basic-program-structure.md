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


## Path Integration

**Concept Demonstrated**: Basic Program Structure
**Next Concept in Path**: Variables and Type Declarations

---

**Exercise Metadata**
- **Created**: 2025-07-30
- **AECS Compliance**: Verified for Exercise Primacy, Concept Atomicity, Applied Understanding, and Progressive Complexity
- **Implementation Status**: Complete and Verified