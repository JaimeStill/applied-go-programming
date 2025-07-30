# Exercise: Basic Program Structure

## Exercise Metadata

- **Stage**: Fundamentals
- **Single Concept**: Basic Program Structure
- **Prerequisites**: Go installation and command-line familiarity
- **Estimated Time**: 20 minutes

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

### Step 2: Execute Program

Run the program:

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

## Verification

Program runs successfully and outputs:
- Go version information 
- Operating system and architecture
- CPU count
- All three program structure elements working: package main, import section, main function


## Path Integration

**Concept Demonstrated**: Basic Program Structure
**Next Concept in Path**: Variables and Type Declarations

---

**Exercise Metadata**
- **Created**: 2025-07-30
- **AECS Compliance**: Verified for Exercise Primacy, Concept Atomicity, Applied Understanding, and Progressive Complexity
- **Implementation Status**: Complete and Verified