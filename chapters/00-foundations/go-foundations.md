# Go Foundations for Experienced Programmers

*Essential Go syntax, types, and concepts for developers transitioning from other languages*

## Introduction

Welcome to Go! If you're an experienced programmer coming from Python, Java, JavaScript, C++, or any other language, this chapter will rapidly introduce you to Go's core concepts and syntax. Unlike a beginner's tutorial, we'll focus on Go's unique features and how they differ from what you already know.

By the end of this chapter, you'll understand Go's fundamental building blocks well enough to tackle the concurrency concepts in Chapter 1 and beyond. Think of this as "Chapter 0" - the essential foundation for everything that follows.

## What Makes Go Different

Before diving into syntax, let's establish Go's core philosophy:

- **Simplicity over complexity**: Go deliberately omits many features common in other languages
- **Composition over inheritance**: No classes or inheritance; types are composed through embedding
- **Explicit over implicit**: Clear, readable code is preferred over clever shortcuts
- **Concurrency built-in**: Goroutines and channels are first-class language features
- **Fast compilation**: Designed for rapid build cycles at scale

---

## Abstract: Go's Type System and Memory Model

Go uses static typing with type inference, garbage collection, and a unique approach to object-oriented programming through structs and interfaces. Understanding Go's type system is crucial because it directly impacts how you'll work with concurrency patterns later.

### Key Concepts You'll Master

1. **Value vs Reference Semantics**: When Go copies data and when it shares references
2. **Interface-based Design**: How Go achieves polymorphism without inheritance
3. **Zero Values**: Go's approach to initialization and the absence of null/undefined
4. **Package-based Visibility**: How Go controls access to types and functions

---

## Example: A Complete Go Program

Let's start with a complete program that demonstrates Go's core features:

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
)

// User represents a user in our system
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Created  time.Time `json:"created"`
}

// String implements the Stringer interface
func (u User) String() string {
    return fmt.Sprintf("User{ID: %d, Name: %s}", u.ID, u.Name)
}

// UserService handles user operations
type UserService struct {
    users map[int]User
    nextID int
}

// NewUserService creates a new user service
func NewUserService() *UserService {
    return &UserService{
        users: make(map[int]User),
        nextID: 1,
    }
}

// CreateUser adds a new user
func (s *UserService) CreateUser(name, email string) User {
    user := User{
        ID:      s.nextID,
        Name:    name,
        Email:   email,
        Created: time.Now(),
    }
    s.users[user.ID] = user
    s.nextID++
    return user
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id int) (User, bool) {
    user, exists := s.users[id]
    return user, exists
}

// ServeHTTP implements http.Handler interface
func (s *UserService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        s.handleCreateUser(w, r)
    case http.MethodGet:
        s.handleGetUser(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func (s *UserService) handleCreateUser(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }
    
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    user := s.CreateUser(req.Name, req.Email)
    
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(user); err != nil {
        log.Printf("Error encoding response: %v", err)
    }
}

func (s *UserService) handleGetUser(w http.ResponseWriter, r *http.Request) {
    // Simplified: In real code, you'd parse the ID from the URL
    user, exists := s.GetUser(1)
    if !exists {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func main() {
    service := NewUserService()
    
    // Create some test users
    user1 := service.CreateUser("Alice Johnson", "alice@example.com")
    user2 := service.CreateUser("Bob Smith", "bob@example.com")
    
    fmt.Println("Created users:")
    fmt.Println(user1)
    fmt.Println(user2)
    
    // Start HTTP server
    fmt.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", service); err != nil {
        log.Fatal(err)
    }
}
```

This program demonstrates:
- **Structs** and **methods**
- **Interfaces** (Stringer, http.Handler)
- **JSON serialization**
- **Error handling**
- **Package organization**
- **HTTP server setup**

---

## Concept: Understanding Go's Building Blocks

### 1. Variables and Types

Go is statically typed but offers type inference for convenience:

```go
// Explicit typing
var name string = "Alice"
var age int = 30
var height float64 = 5.7

// Type inference
var name2 = "Bob"        // string inferred
var age2 = 25           // int inferred
var height2 = 6.1       // float64 inferred

// Short declaration (function scope only)
name3 := "Charlie"      // Most common in Go
age3 := 35
height3 := 5.9

// Multiple variables
var x, y, z int = 1, 2, 3
a, b, c := 4, 5, 6

// Constants
const Pi = 3.14159
const (
    StatusActive   = "active"
    StatusInactive = "inactive"
    StatusPending  = "pending"
)
```

**Zero Values** - Go initializes all variables:
```go
var i int       // 0
var f float64   // 0.0
var b bool      // false
var s string    // ""
var p *int      // nil
var slice []int // nil
var m map[string]int // nil
```

### 2. Basic Types

```go
// Numeric types (size in bits)
var i8 int8 = 127           // -128 to 127
var ui8 uint8 = 255         // 0 to 255 (also: byte)
var i16 int16 = 32767
var i32 int32 = 2147483647  // Also: rune (Unicode code point)
var i64 int64 = 9223372036854775807

// Platform-dependent sizes
var i int = 42              // 32 or 64 bits
var ui uint = 42            // 32 or 64 bits

// Floating point
var f32 float32 = 3.14
var f64 float64 = 3.141592653589793

// Other types
var b bool = true
var s string = "Hello, 世界"    // UTF-8 encoded
var r rune = '界'              // Unicode code point
```

### 3. Composite Types

#### Arrays (Fixed Size)
```go
var arr [5]int                    // Array of 5 ints
arr2 := [3]int{1, 2, 3}          // Array literal
arr3 := [...]int{1, 2, 3, 4, 5}  // Compiler counts elements

// Arrays are values, not references!
arr4 := arr2                     // Copies entire array
```

#### Slices (Dynamic Arrays)
```go
var s []int                      // nil slice
s = make([]int, 5)               // Length 5, capacity 5
s = make([]int, 5, 10)           // Length 5, capacity 10

// Slice literals
s2 := []int{1, 2, 3, 4, 5}

// Slice operations
s = append(s, 6, 7, 8)           // Append elements
s3 := s[1:4]                     // Slice from index 1 to 3
s4 := s[:3]                      // From start to index 2
s5 := s[2:]                      // From index 2 to end

fmt.Println(len(s))              // Length
fmt.Println(cap(s))              // Capacity
```

#### Maps (Hash Tables)
```go
var m map[string]int             // nil map (read-only)
m = make(map[string]int)         // Initialize empty map

// Map literal
m2 := map[string]int{
    "apple":  5,
    "banana": 3,
    "orange": 8,
}

// Map operations
m["grape"] = 10                  // Set value
value := m["apple"]              // Get value (0 if not found)
value, ok := m["apple"]          // Get with existence check
delete(m, "banana")              // Delete key

// Iterate over map
for key, value := range m2 {
    fmt.Printf("%s: %d\n", key, value)
}
```

#### Structs (Custom Types)
```go
// Define struct type
type Person struct {
    Name    string
    Age     int
    Email   string
    private string  // Unexported field (package-private)
}

// Create instances
p1 := Person{
    Name:  "Alice",
    Age:   30,
    Email: "alice@example.com",
}

p2 := Person{"Bob", 25, "bob@example.com", ""}  // Positional (discouraged)
p3 := Person{}                                   // Zero-valued fields

// Access fields
p1.Name = "Alicia"
fmt.Println(p1.Age)

// Anonymous structs (useful for one-off data)
point := struct {
    X, Y int
}{10, 20}

config := struct {
    Host string
    Port int
    SSL  bool
}{
    Host: "localhost",
    Port: 8080,
    SSL:  false,
}
```

### 4. Pointers

Go has pointers but no pointer arithmetic:

```go
var x int = 42
var p *int = &x              // p points to x
fmt.Println(*p)              // Dereference: prints 42

*p = 100                     // Change value through pointer
fmt.Println(x)               // x is now 100

// Pointers are useful for:
// 1. Avoiding copies of large structs
// 2. Allowing functions to modify parameters
// 3. Sharing data between goroutines

func modifyPerson(p *Person) {
    p.Name = "Modified"      // No need to dereference for struct fields
}

person := Person{Name: "Original"}
modifyPerson(&person)
fmt.Println(person.Name)     // "Modified"
```

### 5. Functions

Functions in Go are first-class values:

```go
// Basic function
func add(x, y int) int {
    return x + y
}

// Multiple parameters of same type
func multiply(x, y, z int) int {
    return x * y * z
}

// Multiple return values (very common in Go)
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Named return values
func divmod(a, b int) (quotient, remainder int) {
    quotient = a / b
    remainder = a % b
    return  // Naked return (uses named values)
}

// Variadic functions
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

// Usage
result := sum(1, 2, 3, 4, 5)
numbers := []int{1, 2, 3, 4, 5}
result2 := sum(numbers...)  // Spread slice

// Function as variable
var operation func(int, int) int
operation = add
result3 := operation(5, 3)

// Anonymous functions and closures
func makeAdder(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}

addFive := makeAdder(5)
fmt.Println(addFive(3))  // 8
```

### 6. Methods

Methods are functions with receivers:

```go
type Rectangle struct {
    Width, Height float64
}

// Value receiver (operates on a copy)
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver (can modify the original)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Usage
rect := Rectangle{10.0, 5.0}
area := rect.Area()              // Called on value
rect.Scale(2.0)                  // Go automatically takes address
fmt.Printf("New dimensions: %.1f x %.1f\n", rect.Width, rect.Height)

// You can call pointer methods on values and vice versa
// Go handles the conversion automatically
```

### 7. Interfaces

Interfaces define behavior, not data:

```go
// Interface declaration
type Writer interface {
    Write([]byte) (int, error)
}

type Closer interface {
    Close() error
}

// Interface composition
type WriteCloser interface {
    Writer
    Closer
}

// Any type with matching methods implements the interface automatically
type FileWriter struct {
    filename string
}

func (fw FileWriter) Write(data []byte) (int, error) {
    // Implementation would write to file
    fmt.Printf("Writing %d bytes to %s\n", len(data), fw.filename)
    return len(data), nil
}

func (fw FileWriter) Close() error {
    fmt.Printf("Closing %s\n", fw.filename)
    return nil
}

// FileWriter automatically implements WriteCloser
func processFile(wc WriteCloser) {
    wc.Write([]byte("Hello, World!"))
    wc.Close()
}

// Empty interface accepts any type
var any interface{} = 42
any = "hello"
any = []int{1, 2, 3}

// Type assertions
if str, ok := any.(string); ok {
    fmt.Println("It's a string:", str)
}

// Type switches
switch v := any.(type) {
case int:
    fmt.Printf("Integer: %d\n", v)
case string:
    fmt.Printf("String: %s\n", v)
case []int:
    fmt.Printf("Slice of ints: %v\n", v)
default:
    fmt.Printf("Unknown type: %T\n", v)
}
```

---

## Concept: Control Flow and Error Handling

### 1. Conditional Statements

```go
// Basic if statement
if x > 0 {
    fmt.Println("positive")
} else if x < 0 {
    fmt.Println("negative")
} else {
    fmt.Println("zero")
}

// If with initialization (very common pattern)
if err := doSomething(); err != nil {
    return err
}

// err is only in scope within the if block
```

### 2. Loops (Only `for` in Go)

```go
// Traditional for loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// While-style loop
for condition {
    // Do something
    if shouldBreak {
        break
    }
}

// Infinite loop
for {
    // break or return to exit
    if done {
        break
    }
}

// Range loops (over collections)
slice := []string{"a", "b", "c"}
for index, value := range slice {
    fmt.Printf("%d: %s\n", index, value)
}

// Ignore index or value with _
for _, value := range slice {
    fmt.Println(value)
}

for index := range slice {
    fmt.Println("Index:", index)
}

// Range over maps
m := map[string]int{"a": 1, "b": 2}
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

// Range over strings (iterates over runes)
for index, char := range "Hello" {
    fmt.Printf("%d: %c\n", index, char)
}
```

### 3. Switch Statements

```go
// Basic switch (no break needed)
switch day {
case "Monday":
    fmt.Println("Start of work week")
case "Friday":
    fmt.Println("TGIF!")
case "Saturday", "Sunday":  // Multiple values
    fmt.Println("Weekend!")
default:
    fmt.Println("Midweek")
}

// Switch without expression (replaces if-else chains)
switch {
case x < 0:
    fmt.Println("negative")
case x == 0:
    fmt.Println("zero")
case x > 0:
    fmt.Println("positive")
}

// Switch with initialization
switch day := time.Now().Weekday(); day {
case time.Saturday, time.Sunday:
    fmt.Println("Weekend!")
default:
    fmt.Println("Weekday")
}

// Type switch
switch v := value.(type) {
case int:
    fmt.Printf("Integer: %d\n", v)
case string:
    fmt.Printf("String: %s\n", v)
case bool:
    fmt.Printf("Boolean: %t\n", v)
default:
    fmt.Printf("Unknown type: %T\n", v)
}
```

### 4. Error Handling

Go uses explicit error handling instead of exceptions:

```go
// Functions return error as last value
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide %g by zero", a)
    }
    return a / b, nil
}

// Idiomatic error checking
result, err := divide(10, 2)
if err != nil {
    log.Fatal(err)  // or handle appropriately
}
fmt.Printf("Result: %g\n", result)

// Custom error types
type ValidationError struct {
    Field string
    Value interface{}
    Reason string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %s: %s (got %v)", 
        e.Field, e.Reason, e.Value)
}

func validateAge(age int) error {
    if age < 0 {
        return ValidationError{
            Field:  "age",
            Value:  age,
            Reason: "must be non-negative",
        }
    }
    return nil
}

// Error wrapping (Go 1.13+)
import "errors"

func processUser(id int) error {
    user, err := getUser(id)
    if err != nil {
        return fmt.Errorf("failed to process user %d: %w", id, err)
    }
    // Process user...
    return nil
}

// Check for specific errors
err := processUser(123)
if errors.Is(err, ErrUserNotFound) {
    // Handle specific error
}

// Extract wrapped errors
var validationErr ValidationError
if errors.As(err, &validationErr) {
    fmt.Printf("Validation failed: %s\n", validationErr.Field)
}
```

### 5. Defer, Panic, and Recover

```go
// Defer: execute function when surrounding function returns
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // Always executed before return
    
    // Read file...
    data := make([]byte, 1024)
    n, err := file.Read(data)
    if err != nil {
        return err  // file.Close() still called
    }
    
    fmt.Printf("Read %d bytes\n", n)
    return nil
}

// Multiple defers execute in LIFO order
func example() {
    defer fmt.Println("3rd - executed first")
    defer fmt.Println("2nd - executed second")
    defer fmt.Println("1st - executed last")
    fmt.Println("Function body")
}
// Output:
// Function body
// 1st - executed last
// 2nd - executed second
// 3rd - executed first

// Defer with parameters (evaluated immediately)
func deferExample() {
    x := 10
    defer fmt.Println("Deferred:", x)  // Prints 10, not 20
    x = 20
    fmt.Println("Current:", x)         // Prints 20
}

// Panic and recover (use sparingly!)
func riskyOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
        }
    }()
    
    // Something that might panic
    panic("something went wrong!")
}

// More realistic example: recovering from panics in HTTP handlers
func safeHandler(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if r := recover(); r != nil {
                log.Printf("Panic in handler: %v", r)
                http.Error(w, "Internal Server Error", 500)
            }
        }()
        next(w, r)
    }
}
```

---

## Concept: Package System and Visibility

### 1. Package Declaration and Imports

```go
// Every Go file starts with a package declaration
package main  // Executable package (has main function)
package mylib // Library package

// Import statements
import "fmt"           // Standard library
import "net/http"      // Standard library

// Grouped imports (preferred style)
import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "time"
    
    // Third-party packages
    "github.com/gorilla/mux"
    "github.com/lib/pq"
    
    // Aliased imports
    netpoll "internal/poll"
    
    // Dot import (imports names into current namespace - discouraged)
    . "strings"
    
    // Blank import (for side effects only)
    _ "github.com/lib/pq"  // Database driver registration
)
```

### 2. Visibility Rules

Go uses simple capitalization rules for visibility:

```go
package user

import "time"

// Exported (public) - can be used from other packages
type User struct {
    ID       int       // Exported field
    Name     string    // Exported field
    Email    string    // Exported field
    password string    // Unexported field (package-private)
    createdAt time.Time // Unexported field
}

// Exported method
func (u *User) GetName() string {
    return u.Name
}

// Exported method
func (u *User) SetPassword(password string) {
    u.password = hashPassword(password)
}

// Unexported method (package-private)
func (u *User) validateEmail() error {
    // Validation logic
    return nil
}

// Exported function
func CreateUser(name, email string) *User {
    user := &User{
        Name:      name,
        Email:     email,
        createdAt: time.Now(),
    }
    return user
}

// Unexported function (package-private)
func hashPassword(password string) string {
    // Hashing logic
    return password // Simplified
}

// Exported constant
const MaxUsernameLength = 50

// Unexported constant
const defaultTimeout = 30 * time.Second
```

### 3. Package Initialization

```go
package database

import (
    "database/sql"
    "log"
    _ "github.com/lib/pq"  // Driver registration
)

// Package-level variables
var (
    db *sql.DB
    connected bool
)

// init functions run automatically before main()
func init() {
    log.Println("Initializing database package")
    
    // Initialization code
    var err error
    db, err = sql.Open("postgres", "connection-string")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    
    connected = true
}

// Another init function (multiple allowed)
func init() {
    log.Println("Database package initialization complete")
}

// Exported function that uses initialized variables
func GetDB() *sql.DB {
    if !connected {
        log.Fatal("Database not initialized")
    }
    return db
}
```

### 4. Module System (Go 1.11+)

```bash
# Initialize a new module
go mod init github.com/username/project

# Add dependencies
go get github.com/gorilla/mux
go get github.com/lib/pq@v1.10.0  # Specific version

# Update dependencies
go mod tidy

# Vendor dependencies (optional)
go mod vendor
```

Example `go.mod` file:
```go
module github.com/username/myproject

go 1.21

require (
    github.com/gorilla/mux v1.8.0
    github.com/lib/pq v1.10.9
)

require (
    github.com/gorilla/context v1.1.1 // indirect
)
```

---

## Concept: Basic Concurrency Primitives

While we'll cover concurrency in depth in Chapter 1, you need to understand Go's basic concurrency primitives:

### 1. Goroutines

Goroutines are lightweight threads managed by the Go runtime:

```go
package main

import (
    "fmt"
    "time"
)

func sayHello(name string) {
    for i := 0; i < 3; i++ {
        fmt.Printf("Hello, %s! (%d)\n", name, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    // Launch goroutines
    go sayHello("Alice")
    go sayHello("Bob")
    
    // Launch anonymous function as goroutine
    go func() {
        fmt.Println("Anonymous goroutine")
    }()
    
    // Wait for goroutines to complete (simplified)
    time.Sleep(1 * time.Second)
    fmt.Println("Main function ending")
}
```

### 2. Channels

Channels are Go's way of communicating between goroutines:

```go
func main() {
    // Create channels
    messages := make(chan string)         // Unbuffered channel
    numbers := make(chan int, 5)          // Buffered channel (capacity 5)
    
    // Send values to channel (in goroutine to avoid blocking)
    go func() {
        messages <- "Hello"
        messages <- "World"
        close(messages)  // Close channel when done sending
    }()
    
    // Receive values from channel
    for msg := range messages {
        fmt.Println("Received:", msg)
    }
    
    // Buffered channel example
    numbers <- 1
    numbers <- 2
    numbers <- 3
    close(numbers)
    
    for num := range numbers {
        fmt.Println("Number:", num)
    }
}
```

### 3. Select Statement

Select lets you handle multiple channel operations:

```go
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(100 * time.Millisecond)
        ch1 <- "from ch1"
    }()
    
    go func() {
        time.Sleep(200 * time.Millisecond)
        ch2 <- "from ch2"
    }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Received", msg1)
        case msg2 := <-ch2:
            fmt.Println("Received", msg2)
        case <-time.After(500 * time.Millisecond):
            fmt.Println("Timeout!")
        }
    }
}
```

---

## Concept: Essential Toolchain and Testing

### 1. Go Toolchain Commands

```bash
# Module management
go mod init myproject           # Initialize new module
go mod tidy                     # Add missing and remove unused modules
go get github.com/pkg/errors    # Add dependency
go get -u ./...                 # Update all dependencies

# Building and running
go build                        # Build executable in current directory
go build -o myapp               # Build with custom name
go run main.go                  # Build and run immediately
go install                      # Build and install to $GOPATH/bin

# Testing
go test                         # Run tests in current package
go test ./...                   # Run tests in all packages
go test -v                      # Verbose output
go test -cover                  # Show test coverage
go test -race                   # Enable race detector
go test -bench=.                # Run benchmarks

# Code quality
go fmt ./...                    # Format all Go files
go vet ./...                    # Examine code for suspicious constructs
golint ./...                    # Style checker (external tool)

# Documentation
go doc fmt.Println              # Show documentation for specific function
godoc -http=:6060               # Start local documentation server

# Other useful commands
go version                      # Show Go version
go env                          # Show Go environment
go clean                       # Remove object files and cached files
```

### 2. Basic Testing

Go has built-in testing support:

```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

```go
// math_test.go
package math

import (
    "testing"
)

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

// Table-driven test (idiomatic Go testing pattern)
func TestAddTable(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"zero values", 0, 0, 0},
        {"negative numbers", -1, 1, 0},
        {"large numbers", 1000, 2000, 3000},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

func TestDivide(t *testing.T) {
    // Test normal case
    result, err := Divide(10, 2)
    if err != nil {
        t.Errorf("Divide(10, 2) returned error: %v", err)
    }
    if result != 5.0 {
        t.Errorf("Divide(10, 2) = %f; want 5.0", result)
    }
    
    // Test error case
    _, err = Divide(10, 0)
    if err == nil {
        t.Error("Divide(10, 0) should return error")
    }
}

// Benchmark test (functions starting with Benchmark)
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(i, i+1)
    }
}
```

### 3. Code Organization

Standard Go project structure:

```
myproject/
├── go.mod              # Module definition
├── go.sum              # Dependency checksums
├── main.go             # Main package (for executables)
├── README.md           # Project documentation
├── cmd/                # Multiple executables
│   ├── server/
│   │   └── main.go
│   └── client/
│       └── main.go
├── internal/           # Private packages (not importable by others)
│   ├── config/
│   │   └── config.go
│   └── database/
│       └── db.go
├── pkg/                # Public packages (importable by others)
│   ├── api/
│   │   └── handler.go
│   └── models/
│       └── user.go
├── web/                # Web assets
│   ├── static/
│   └── templates/
├── scripts/            # Build and deployment scripts
├── docs/               # Documentation
└── testdata/           # Test fixtures
```

---

## Your Task: Practice Essential Go Patterns

Now that you've seen Go's fundamental concepts, let's reinforce your understanding with a practical exercise that combines multiple concepts.

### Exercise: Build a Task Manager

Create a simple task management system that demonstrates Go's core features:

**Requirements:**
1. Define a `Task` struct with ID, title, description, completed status, and creation time
2. Create a `TaskManager` that stores tasks and provides CRUD operations
3. Implement proper error handling for operations like "task not found"
4. Add methods that satisfy common interfaces (like `String()` for tasks)
5. Include comprehensive tests for your functionality
6. Use Go's built-in JSON marshaling for data persistence

**Starter Code Structure:**

```go
package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "os"
    "time"
)

// Task represents a task in our system
type Task struct {
    // TODO: Define struct fields with appropriate JSON tags
}

// TaskManager manages a collection of tasks
type TaskManager struct {
    // TODO: Define fields for storing tasks
}

// TODO: Implement the following methods:
// - NewTaskManager() *TaskManager
// - (tm *TaskManager) AddTask(title, description string) Task
// - (tm *TaskManager) GetTask(id int) (Task, error)
// - (tm *TaskManager) UpdateTask(id int, title, description string) error
// - (tm *TaskManager) CompleteTask(id int) error
// - (tm *TaskManager) ListTasks() []Task
// - (tm *TaskManager) SaveToFile(filename string) error
// - (tm *TaskManager) LoadFromFile(filename string) error
// - (t Task) String() string

func main() {
    tm := NewTaskManager()
    
    // TODO: Add some sample tasks
    // TODO: Demonstrate CRUD operations
    // TODO: Save to and load from file
    // TODO: Handle errors appropriately
}
```

**Success Criteria:**
- [ ] All struct fields are properly defined with JSON tags
- [ ] TaskManager provides all required CRUD operations
- [ ] Proper error handling for invalid operations
- [ ] Task implements the `Stringer` interface
- [ ] File persistence works correctly
- [ ] Comprehensive tests cover main functionality
- [ ] Code follows Go naming conventions
- [ ] No race conditions (even though we're not using concurrency yet)

**Extension Challenges:**
1. Add task priority levels and sorting
2. Implement task filtering (by completion status, date range, etc.)
3. Add task categories or tags
4. Create a simple CLI interface using command-line arguments
5. Add validation for task fields (e.g., title can't be empty)

**Testing Your Code:**
Create a `task_test.go` file with tests for:
- Creating and retrieving tasks
- Error cases (invalid IDs, etc.)
- File persistence and loading
- Edge cases (empty task manager, duplicate operations, etc.)

This exercise will help you internalize Go's syntax and idioms before we dive into concurrency patterns in the next chapter.

---

## Key Differences for Developers Coming From Other Languages

### From Python
- **Static typing** with compile-time error checking
- **Explicit error handling** instead of exceptions
- **No classes** - use structs and methods instead
- **Compilation step** - faster runtime, deployment of single binary
- **Built-in concurrency** with goroutines and channels

### From Java/C#
- **No inheritance** - use composition and interfaces
- **Implicit interface implementation** (duck typing for interfaces)
- **Multiple return values** common for error handling
- **No generics** (until Go 1.18+), but when available, simpler than Java generics
- **Garbage collection** but more predictable than JVM

### From JavaScript
- **Static typing** with compile-time safety
- **No undefined** - everything has a zero value
- **True concurrency** with goroutines, not event loop
- **Explicit memory management** concepts (pointers, value vs reference)
- **Package system** instead of module bundlers

### From C/C++
- **Garbage collection** - no manual memory management
- **No pointer arithmetic** - pointers are safe
- **Simpler syntax** - less complexity than C++
- **Built-in string type** with UTF-8 support
- **Modern package system** with dependency management

---

## Summary

You now have a solid foundation in Go's essential concepts:

✅ **Types and Variables**: Static typing with inference, zero values, and Go's type system  
✅ **Control Flow**: if, for, switch, and Go's unique patterns  
✅ **Functions and Methods**: First-class functions, methods with receivers  
✅ **Structs and Interfaces**: Composition over inheritance, implicit interface satisfaction  
✅ **Error Handling**: Explicit errors instead of exceptions  
✅ **Packages**: Visibility rules, module system, code organization  
✅ **Basic Concurrency**: Goroutines, channels, and select statements  
✅ **Toolchain**: Essential commands for building, testing, and maintaining Go code  

**What's Next?**

In Chapter 1, we'll build on this foundation to explore Go's concurrency model in depth. You'll learn:
- Advanced goroutine patterns
- Channel-based communication strategies
- Synchronization primitives
- Worker pools and pipeline patterns
- Context for cancellation and timeouts

The concepts you've learned here - especially interfaces, error handling, and basic channel operations - will be crucial for understanding the advanced concurrency patterns ahead.

Remember: Go's philosophy is "less is more." The language intentionally provides fewer features than many other languages, but those features work together elegantly to create robust, concurrent programs. Master these fundamentals, and you'll be ready to harness Go's true power in concurrent programming.