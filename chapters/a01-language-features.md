# Language Features Quick Reference

## Purpose and Usage

This appendix provides a comprehensive quick reference for Go language features and syntax. Use this as a lookup guide when working through exercises and projects, especially when you need to check specific syntax, understand type behavior, or review language constructs.

**Target Audience:** Developers working through the curriculum who need quick access to Go language syntax and semantics.

**How to Use:** Designed for rapid lookup and reference. Jump to specific sections using the table of contents, or search for keywords related to your current need.

## Table of Contents

1. [Variables and Types](#variables-and-types)
2. [Basic Types](#basic-types)
3. [Composite Types](#composite-types)
4. [Pointers](#pointers)
5. [Functions](#functions)
6. [Methods](#methods)
7. [Interfaces](#interfaces)
8. [Control Flow](#control-flow)
9. [Error Handling](#error-handling)
10. [Defer, Panic, and Recover](#defer-panic-and-recover)
11. [Concurrency Primitives](#concurrency-primitives)
12. [Channels](#channels)
13. [Context Package](#context-package)

---

## Variables and Types

### Variable Declaration

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

// Zero value declarations
var i int               // 0
var f float64           // 0.0
var b bool              // false
var s string            // ""
```

### Constants

```go
const Pi = 3.14159

// Constant blocks
const (
    StatusActive   = "active"
    StatusInactive = "inactive"
    StatusPending  = "pending"
)

// Typed constants
const MaxRetries int = 3
const Timeout time.Duration = 30 * time.Second

// Iota for enumerations
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)
```

### Zero Values

Go initializes all variables to their zero value:

```go
var i int               // 0
var f float64           // 0.0
var b bool              // false
var s string            // ""
var p *int              // nil
var slice []int         // nil
var m map[string]int    // nil
var ch chan int         // nil
var fn func()           // nil
```

---

## Basic Types

### Numeric Types

```go
// Signed integers
var i8 int8 = 127           // -128 to 127
var i16 int16 = 32767       // -32768 to 32767
var i32 int32 = 2147483647  // Also: rune (Unicode code point)
var i64 int64 = 9223372036854775807

// Unsigned integers
var ui8 uint8 = 255         // 0 to 255 (also: byte)
var ui16 uint16 = 65535
var ui32 uint32 = 4294967295
var ui64 uint64 = 18446744073709551615

// Platform-dependent sizes
var i int = 42              // 32 or 64 bits
var ui uint = 42            // 32 or 64 bits
var ptr uintptr = 0         // Large enough to store pointer

// Floating point
var f32 float32 = 3.14
var f64 float64 = 3.141592653589793

// Complex numbers
var c64 complex64 = 1 + 2i
var c128 complex128 = 1 + 2i
```

### String and Character Types

```go
// Strings (UTF-8 encoded)
var s string = "Hello, 世界"
var raw string = `Raw string
with newlines
and "quotes"`

// Runes (Unicode code points)
var r rune = '界'              // rune is alias for int32
var r2 rune = '\u754c'         // Unicode escape

// Bytes (alias for uint8)
var b byte = 'A'               // byte is alias for uint8
```

### Boolean Type

```go
var b bool = true
var b2 bool = false
var b3 bool                    // false (zero value)

// Boolean operations
result := true && false        // false
result2 := true || false       // true
result3 := !true              // false
```

---

## Composite Types

### Arrays (Fixed Size)

```go
// Array declarations
var arr [5]int                    // Array of 5 ints, zero-valued
arr2 := [3]int{1, 2, 3}          // Array literal
arr3 := [...]int{1, 2, 3, 4, 5}  // Compiler counts elements

// Arrays are values, not references
arr4 := arr2                     // Copies entire array
arr4[0] = 100                    // arr2[0] is still 1

// Multidimensional arrays
var matrix [3][3]int
matrix[0][0] = 1
```

### Slices (Dynamic Arrays)

```go
// Slice declarations
var s []int                      // nil slice
s = make([]int, 5)               // Length 5, capacity 5
s = make([]int, 5, 10)           // Length 5, capacity 10

// Slice literals
s2 := []int{1, 2, 3, 4, 5}

// Slice operations
s = append(s, 6, 7, 8)           // Append elements
s = append(s, s2...)             // Append another slice

// Slicing operations
s3 := s[1:4]                     // Elements 1, 2, 3
s4 := s[:3]                      // Elements 0, 1, 2
s5 := s[2:]                      // From element 2 to end
s6 := s[:]                       // Copy of entire slice

// Length and capacity
fmt.Println(len(s))              // Number of elements
fmt.Println(cap(s))              // Capacity of underlying array

// Copy slices
dest := make([]int, len(s))
copy(dest, s)                    // Copy s to dest
```

### Maps (Hash Tables)

```go
// Map declarations
var m map[string]int             // nil map (read-only)
m = make(map[string]int)         // Initialize empty map
m = make(map[string]int, 100)    // With initial capacity hint

// Map literals
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

// Check if key exists
if value, exists := m["key"]; exists {
    fmt.Printf("Found: %d\n", value)
}

// Iterate over map
for key, value := range m2 {
    fmt.Printf("%s: %d\n", key, value)
}

// Clear a map
for k := range m {
    delete(m, k)
}
```

### Structs (Custom Types)

```go
// Struct definition
type Person struct {
    Name    string
    Age     int
    Email   string
    private string  // Unexported field (package-private)
}

// Struct creation
p1 := Person{
    Name:  "Alice",
    Age:   30,
    Email: "alice@example.com",
}

p2 := Person{"Bob", 25, "bob@example.com", ""}  // Positional (discouraged)
p3 := Person{}                                   // Zero-valued fields
p4 := new(Person)                               // Returns *Person

// Field access
p1.Name = "Alicia"
fmt.Println(p1.Age)

// Anonymous structs
point := struct {
    X, Y int
}{10, 20}

// Struct embedding (composition)
type Employee struct {
    Person              // Embedded struct
    ID       int
    Department string
}

emp := Employee{
    Person: Person{Name: "John", Age: 35},
    ID:     12345,
    Department: "Engineering",
}

// Promoted fields
fmt.Println(emp.Name)           // Access embedded field directly
```

---

## Pointers

Go has pointers but no pointer arithmetic:

```go
// Basic pointer operations
var x int = 42
var p *int = &x              // p points to x
fmt.Println(*p)              // Dereference: prints 42

*p = 100                     // Change value through pointer
fmt.Println(x)               // x is now 100

// Pointer to struct
type Point struct { X, Y int }
pt := Point{1, 2}
ptr := &pt
ptr.X = 10                   // Automatic dereferencing for struct fields
// Equivalent to: (*ptr).X = 10

// New function
p2 := new(int)               // Returns *int, value is 0
*p2 = 42

// Pointers and functions
func modifyValue(p *int) {
    *p = 100
}

value := 42
modifyValue(&value)          // value is now 100

// Nil pointers
var p3 *int                  // nil pointer
if p3 == nil {
    fmt.Println("p3 is nil")
}
```

---

## Functions

### Function Declaration and Calls

```go
// Basic function
func add(x, y int) int {
    return x + y
}

// Multiple parameters of same type
func multiply(x, y, z int) int {
    return x * y * z
}

// Multiple return values
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
```

### Function Types and Variables

```go
// Function as variable
var operation func(int, int) int
operation = add
result := operation(5, 3)

// Function type declaration
type BinaryOp func(int, int) int

var op BinaryOp = add
result2 := op(10, 20)

// Anonymous functions
func() {
    fmt.Println("Anonymous function")
}()

// Function closures
func makeAdder(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}

addFive := makeAdder(5)
fmt.Println(addFive(3))  // 8
```

### Higher-Order Functions

```go
// Functions taking functions as parameters
func apply(f func(int) int, value int) int {
    return f(value)
}

double := func(x int) int { return x * 2 }
result := apply(double, 5)  // 10

// Functions returning functions
func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

triple := makeMultiplier(3)
fmt.Println(triple(4))  // 12
```

---

## Methods

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

// Method with multiple parameters
func (r Rectangle) Contains(x, y float64) bool {
    return x >= 0 && x <= r.Width && y >= 0 && y <= r.Height
}

// Usage
rect := Rectangle{10.0, 5.0}
area := rect.Area()              // Called on value
rect.Scale(2.0)                  // Go automatically takes address
contains := rect.Contains(15.0, 8.0)

// Method expressions
var scaleFunc func(*Rectangle, float64) = Rectangle.Scale
scaleFunc(&rect, 1.5)

// Method values
scaleBound := rect.Scale
scaleBound(0.5)
```

### Method Sets

```go
type Counter struct {
    value int
}

// Method with value receiver
func (c Counter) Value() int {
    return c.value
}

// Method with pointer receiver
func (c *Counter) Increment() {
    c.value++
}

// A *Counter has methods with both receiver types
// A Counter has only methods with value receivers
```

---

## Interfaces

### Interface Declaration and Implementation

```go
// Interface declaration
type Writer interface {
    Write([]byte) (int, error)
}

type Reader interface {
    Read([]byte) (int, error)
}

// Interface composition
type ReadWriter interface {
    Reader
    Writer
}

// Any type with matching methods implements the interface automatically
type StringWriter struct {
    content string
}

func (sw *StringWriter) Write(data []byte) (int, error) {
    sw.content += string(data)
    return len(data), nil
}

// StringWriter automatically implements Writer
func writeData(w Writer, data []byte) error {
    _, err := w.Write(data)
    return err
}

// Usage
sw := &StringWriter{}
writeData(sw, []byte("Hello"))
```

### Empty Interface and Type Assertions

```go
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

### Common Interfaces

```go
// Stringer interface
type Stringer interface {
    String() string
}

type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}

// Error interface
type error interface {
    Error() string
}

type CustomError struct {
    message string
}

func (e CustomError) Error() string {
    return e.message
}
```

---

## Control Flow

### Conditional Statements

```go
// Basic if statement
if x > 0 {
    fmt.Println("positive")
} else if x < 0 {
    fmt.Println("negative")
} else {
    fmt.Println("zero")
}

// If with initialization
if err := doSomething(); err != nil {
    return err
}
// err is only in scope within the if block

// Multiple conditions
if x > 0 && y > 0 {
    fmt.Println("both positive")
}

if x > 0 || y > 0 {
    fmt.Println("at least one positive")
}
```

### Loops (Only `for` in Go)

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
    if shouldContinue {
        continue
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
for index, char := range "Hello, 世界" {
    fmt.Printf("%d: %c\n", index, char)
}

// Range over channels
ch := make(chan int)
go func() {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)
}()

for value := range ch {
    fmt.Println(value)
}
```

### Switch Statements

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

// Fallthrough (rarely used)
switch x {
case 1:
    fmt.Print("one ")
    fallthrough
case 2:
    fmt.Print("two ")
}
```

---

## Error Handling

### Basic Error Handling

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
    log.Fatal(err)
}
fmt.Printf("Result: %g\n", result)
```

### Custom Error Types

```go
// Error type implementation
type ValidationError struct {
    Field  string
    Value  interface{}
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
```

### Error Wrapping (Go 1.13+)

```go
import "errors"

// Wrap errors with context
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

// Unwrap errors
cause := errors.Unwrap(err)
```

### Error Patterns

```go
// Sentinel errors
var (
    ErrNotFound    = errors.New("not found")
    ErrInvalidInput = errors.New("invalid input")
    ErrTimeout     = errors.New("operation timed out")
)

// Error handling with multiple operations
func processData() error {
    data, err := loadData()
    if err != nil {
        return fmt.Errorf("loading data: %w", err)
    }
    
    if err := validateData(data); err != nil {
        return fmt.Errorf("validating data: %w", err)
    }
    
    if err := saveData(data); err != nil {
        return fmt.Errorf("saving data: %w", err)
    }
    
    return nil
}

// Error aggregation
type MultiError []error

func (me MultiError) Error() string {
    if len(me) == 0 {
        return "no errors"
    }
    if len(me) == 1 {
        return me[0].Error()
    }
    return fmt.Sprintf("%s (and %d more errors)", me[0].Error(), len(me)-1)
}
```

---

## Defer, Panic, and Recover

### Defer Statement

```go
// Basic defer usage
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // Always executed before return
    
    // Read file operations...
    return nil
}

// Multiple defers execute in LIFO order
func example() {
    defer fmt.Println("3rd - executed first")
    defer fmt.Println("2nd - executed second")
    defer fmt.Println("1st - executed last")
    fmt.Println("Function body")
}

// Defer with parameters (evaluated immediately)
func deferExample() {
    x := 10
    defer fmt.Println("Deferred:", x)  // Prints 10, not 20
    x = 20
    fmt.Println("Current:", x)         // Prints 20
}

// Defer with closure
func deferClosure() {
    x := 10
    defer func() {
        fmt.Println("Deferred:", x)  // Prints 20 (captured by closure)
    }()
    x = 20
    fmt.Println("Current:", x)
}
```

### Panic and Recover

```go
// Panic stops normal execution
func riskyOperation() {
    if somethingWrong {
        panic("something went wrong!")
    }
}

// Recover from panic
func safeOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
        }
    }()
    
    riskyOperation()  // This might panic
    fmt.Println("This won't be reached if panic occurs")
}

// Recover only works in deferred functions
func handlePanic() {
    defer func() {
        if r := recover(); r != nil {
            switch err := r.(type) {
            case string:
                fmt.Printf("Panic with string: %s\n", err)
            case error:
                fmt.Printf("Panic with error: %v\n", err)
            default:
                fmt.Printf("Panic with unknown type: %v\n", err)
            }
        }
    }()
    
    // Dangerous operations here
}

// Re-panic after handling
func conditionalRecover() {
    defer func() {
        if r := recover(); r != nil {
            if shouldRecover(r) {
                fmt.Printf("Handled panic: %v\n", r)
            } else {
                panic(r)  // Re-panic
            }
        }
    }()
    
    // Operations that might panic
}
```

### Common Defer Patterns

```go
// Resource cleanup
func processFile(filename string) error {
    mu.Lock()
    defer mu.Unlock()
    
    // Critical section
    return nil
}

// Transaction handling
func dbTransaction() error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer func() {
        if err != nil {
            tx.Rollback()
        } else {
            tx.Commit()
        }
    }()
    
    // Database operations
    return nil
}

// Timing operations
func timedOperation() {
    start := time.Now()
    defer func() {
        fmt.Printf("Operation took: %v\n", time.Since(start))
    }()
    
    // Long running operation
}
```

---

## Concurrency Primitives

### Goroutines

Goroutines are lightweight threads managed by the Go runtime:

```go
// Launch a goroutine
go func() {
    fmt.Println("Running in goroutine")
}()

// Goroutine with parameters
go func(name string, count int) {
    for i := 0; i < count; i++ {
        fmt.Printf("Hello %s %d\n", name, i)
        time.Sleep(100 * time.Millisecond)
    }
}("Alice", 3)

// Goroutine with named function
func sayHello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

go sayHello("Bob")

// Wait for goroutines (basic - use sync.WaitGroup for real code)
time.Sleep(1 * time.Second)
```

### Sync Package Primitives

```go
import "sync"

// WaitGroup - wait for multiple goroutines
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Printf("Worker %d done\n", id)
    }(i)
}

wg.Wait()  // Wait for all goroutines to complete

// Mutex - mutual exclusion
var mu sync.Mutex
var counter int

func increment() {
    mu.Lock()
    defer mu.Unlock()
    counter++
}

// RWMutex - allows multiple readers or single writer
var rwmu sync.RWMutex
var data map[string]int

func read(key string) int {
    rwmu.RLock()
    defer rwmu.RUnlock()
    return data[key]
}

func write(key string, value int) {
    rwmu.Lock()
    defer rwmu.Unlock()
    data[key] = value
}

// Once - execute function exactly once
var once sync.Once

func initialize() {
    once.Do(func() {
        fmt.Println("Initialized once")
    })
}

// Atomic operations
import "sync/atomic"

var atomicCounter int64

func atomicIncrement() {
    atomic.AddInt64(&atomicCounter, 1)
}

func getAtomicValue() int64 {
    return atomic.LoadInt64(&atomicCounter)
}
```

---

## Channels

Channels are Go's primary mechanism for communication between goroutines:

### Channel Basics

```go
// Channel creation
ch := make(chan int)           // Unbuffered channel
buffered := make(chan int, 5)  // Buffered channel (capacity 5)

// Send and receive
ch <- 42        // Send value to channel
value := <-ch   // Receive value from channel

// Receive with ok check
value, ok := <-ch
if !ok {
    fmt.Println("Channel closed")
}

// Close channel
close(ch)

// Check if channel is closed
select {
case v, ok := <-ch:
    if !ok {
        fmt.Println("Channel closed")
    } else {
        fmt.Printf("Received: %d\n", v)
    }
default:
    fmt.Println("No value available")
}
```

### Channel Directions

```go
// Send-only channel
func sender(ch chan<- int) {
    ch <- 42
    close(ch)
}

// Receive-only channel
func receiver(ch <-chan int) {
    for value := range ch {
        fmt.Println(value)
    }
}

// Bidirectional channel (can be passed to both)
ch := make(chan int)
go sender(ch)
receiver(ch)
```

### Channel Patterns

```go
// Fan-out pattern
func fanOut(input <-chan int, workers int) []<-chan int {
    outputs := make([]<-chan int, workers)
    for i := 0; i < workers; i++ {
        output := make(chan int)
        outputs[i] = output
        go func() {
            defer close(output)
            for value := range input {
                output <- value * 2  // Process the value
            }
        }()
    }
    return outputs
}

// Fan-in pattern
func fanIn(inputs ...<-chan int) <-chan int {
    output := make(chan int)
    var wg sync.WaitGroup
    
    for _, input := range inputs {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for value := range ch {
                output <- value
            }
        }(input)
    }
    
    go func() {
        wg.Wait()
        close(output)
    }()
    
    return output
}

// Pipeline pattern
func pipeline() {
    // Stage 1: Generate numbers
    numbers := make(chan int)
    go func() {
        defer close(numbers)
        for i := 1; i <= 5; i++ {
            numbers <- i
        }
    }()
    
    // Stage 2: Square numbers
    squares := make(chan int)
    go func() {
        defer close(squares)
        for num := range numbers {
            squares <- num * num
        }
    }()
    
    // Stage 3: Print results
    for square := range squares {
        fmt.Println(square)
    }
}

// Worker pool pattern
func workerPool(jobs <-chan int, results chan<- int, workers int) {
    var wg sync.WaitGroup
    
    for i := 0; i < workers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for job := range jobs {
                // Process job
                results <- job * 2
            }
        }()
    }
    
    go func() {
        wg.Wait()
        close(results)
    }()
}
```

### Select Statement

```go
// Select for non-blocking channel operations
select {
case msg1 := <-ch1:
    fmt.Println("Received from ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received from ch2:", msg2)
case ch3 <- 42:
    fmt.Println("Sent to ch3")
default:
    fmt.Println("No channel ready")
}

// Select with timeout
select {
case result := <-ch:
    fmt.Println("Received:", result)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout!")
}

// Select in loop for multiplexing
for {
    select {
    case msg := <-ch1:
        if msg == "" {
            return  // Exit condition
        }
        fmt.Println("Ch1:", msg)
    case msg := <-ch2:
        fmt.Println("Ch2:", msg)
    case <-time.After(30 * time.Second):
        fmt.Println("No activity for 30 seconds")
        return
    }
}
```

---

## Context Package

The context package provides a way to carry deadlines, cancellation signals, and request-scoped values across API boundaries:

### Context Creation

```go
import "context"

// Background context (typically used at main)
ctx := context.Background()

// TODO context (placeholder)
ctx := context.TODO()

// Context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Context with deadline
deadline := time.Now().Add(10 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()

// Context with cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// Context with value
ctx := context.WithValue(context.Background(), "userID", 12345)
```

### Context Usage

```go
// Function accepting context
func doWork(ctx context.Context) error {
    // Check if context is done
    select {
    case <-ctx.Done():
        return ctx.Err()  // context.Canceled or context.DeadlineExceeded
    default:
        // Continue work
    }
    
    // Long-running operation with context check
    for i := 0; i < 1000; i++ {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            // Do work
            time.Sleep(10 * time.Millisecond)
        }
    }
    
    return nil
}

// HTTP request with context
func makeRequest(ctx context.Context, url string) (*http.Response, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }
    
    client := &http.Client{}
    return client.Do(req)
}

// Database query with context
func queryUser(ctx context.Context, db *sql.DB, userID int) (*User, error) {
    query := "SELECT name, email FROM users WHERE id = ?"
    row := db.QueryRowContext(ctx, query, userID)
    
    var user User
    err := row.Scan(&user.Name, &user.Email)
    return &user, err
}
```

### Context Patterns

```go
// Cancellation propagation
func parentWork() error {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    // Start multiple workers
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            if err := workerTask(ctx, id); err != nil {
                fmt.Printf("Worker %d failed: %v\n", id, err)
                cancel()  // Cancel all other workers
            }
        }(i)
    }
    
    wg.Wait()
    return nil
}

// Timeout handling
func timeoutExample() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    result := make(chan string, 1)
    go func() {
        // Simulate slow operation
        time.Sleep(3 * time.Second)
        result <- "completed"
    }()
    
    select {
    case res := <-result:
        fmt.Println("Result:", res)
    case <-ctx.Done():
        fmt.Println("Operation timed out:", ctx.Err())
    }
}

// Context values (use sparingly)
type contextKey string

const userIDKey contextKey = "userID"

func withUserID(ctx context.Context, userID int) context.Context {
    return context.WithValue(ctx, userIDKey, userID)
}

func getUserID(ctx context.Context) (int, bool) {
    userID, ok := ctx.Value(userIDKey).(int)
    return userID, ok
}

// Middleware pattern with context
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID := extractUserID(r)  // Extract from token, cookie, etc.
        ctx := withUserID(r.Context(), userID)
        next(w, r.WithContext(ctx))
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    if userID, ok := getUserID(r.Context()); ok {
        fmt.Fprintf(w, "Hello user %d", userID)
    } else {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
    }
}
```