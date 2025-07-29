# Go Idioms and Patterns Quick Reference

## Purpose and Usage

This appendix provides a comprehensive quick reference for essential Go idioms, patterns, and conventions that experienced Go developers use daily. Use this material as a lookup reference during exercises and projects when you need to check proper Go style, common patterns, or best practices.

**Target Audience**: Developers working through curriculum exercises who need quick access to Go conventions and patterns.

**How to Use**: This reference is designed for quick scanning and lookup. Use the table of contents to jump to specific sections, or search for keywords related to your current need.

## Table of Contents

1. [Error Handling Idioms](#error-handling-idioms)
2. [Nil Checks and Safety](#nil-checks-and-safety)
3. [Interface Patterns](#interface-patterns)
4. [Struct Patterns](#struct-patterns)
5. [Channel Patterns](#channel-patterns)
6. [Function Patterns](#function-patterns)
7. [Package Organization](#package-organization)
8. [Testing Patterns](#testing-patterns)
9. [Performance Idioms](#performance-idioms)
10. [Style Conventions](#style-conventions)

---

## Error Handling Idioms

### Standard Error Check
The most common Go idiom - always check errors immediately.

```go
// DO: Check errors immediately
result, err := someFunction()
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}

// DON'T: Ignore errors or defer checking
result, _ := someFunction() // ignoring error
```

### Error Wrapping with Context
Add context to errors as they bubble up the call stack.

```go
func processFile(filename string) error {
    data, err := os.ReadFile(filename)
    if err != nil {
        return fmt.Errorf("failed to read config file %s: %w", filename, err)
    }
    
    if err := validate(data); err != nil {
        return fmt.Errorf("invalid config in %s: %w", filename, err)
    }
    
    return nil
}
```

### Custom Error Types
Define error types for specific error conditions.

```go
type ValidationError struct {
    Field string
    Value interface{}
    Msg   string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed for field %s (value: %v): %s", 
        e.Field, e.Value, e.Msg)
}

// Usage
func validateAge(age int) error {
    if age < 0 {
        return ValidationError{
            Field: "age",
            Value: age,
            Msg:   "age cannot be negative",
        }
    }
    return nil
}
```

### Sentinel Errors
Define specific error values for expected error conditions.

```go
var (
    ErrNotFound    = errors.New("item not found")
    ErrInvalidInput = errors.New("invalid input")
    ErrTimeout     = errors.New("operation timed out")
)

func findUser(id string) (*User, error) {
    if id == "" {
        return nil, ErrInvalidInput
    }
    
    user, exists := userDB[id]
    if !exists {
        return nil, ErrNotFound
    }
    
    return user, nil
}

// Usage with errors.Is
user, err := findUser("123")
if errors.Is(err, ErrNotFound) {
    // Handle not found case
}
```

---

## Nil Checks and Safety

### Pointer Safety
Always check for nil before dereferencing pointers.

```go
func processUser(u *User) error {
    if u == nil {
        return errors.New("user cannot be nil")
    }
    
    // Safe to use u
    fmt.Printf("Processing user: %s\n", u.Name)
    return nil
}
```

### Interface Nil Checks
Interfaces can be tricky - check both the interface and its underlying value.

```go
func processReader(r io.Reader) error {
    if r == nil {
        return errors.New("reader cannot be nil")
    }
    
    // For interfaces, also check if the underlying value is nil
    if reflect.ValueOf(r).IsNil() {
        return errors.New("reader implementation is nil")
    }
    
    // Safe to use r
    return nil
}
```

### Safe Map Access
Use the comma ok idiom to safely access maps.

```go
// Safe map access
value, ok := myMap[key]
if !ok {
    // Handle missing key
    return errors.New("key not found")
}

// Or with default value
value := myMap[key] // zero value if key doesn't exist
if value == "" {    // check if meaningful
    value = "default"
}
```

### Slice Bounds Checking
Always verify slice bounds before accessing elements.

```go
func safeSliceAccess(slice []string, index int) (string, error) {
    if index < 0 || index >= len(slice) {
        return "", fmt.Errorf("index %d out of bounds for slice of length %d", 
            index, len(slice))
    }
    return slice[index], nil
}
```

---

## Interface Patterns

### Small Interfaces
Define interfaces with minimal methods - prefer composition.

```go
// Good: Small, focused interfaces
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type ReadWriter interface {
    Reader
    Writer
}

// Usage: Accept interfaces, return structs
func processData(r Reader) ([]byte, error) {
    // Implementation
}
```

### Interface Satisfaction
Let types satisfy interfaces implicitly - don't declare implementation.

```go
// Define interface where you use it, not where you implement it
type Notifier interface {
    Notify(message string) error
}

// Types satisfy interfaces implicitly
type EmailNotifier struct {
    Address string
}

func (e EmailNotifier) Notify(message string) error {
    // Send email
    return nil
}

// No explicit "implements" declaration needed
```

### Type Assertions and Type Switches
Safely extract concrete types from interfaces.

```go
// Type assertion with safety check
if emailer, ok := notifier.(*EmailNotifier); ok {
    // Use emailer-specific methods
    emailer.SetSubject("Alert")
}

// Type switch for multiple types
switch n := notifier.(type) {
case *EmailNotifier:
    n.SetSubject("Email Alert")
case *SMSNotifier:
    n.SetCarrier("Verizon")
default:
    // Handle unknown type
    log.Printf("Unknown notifier type: %T", n)
}
```

---

## Struct Patterns

### Struct Initialization
Use struct literals with field names for clarity.

```go
// Good: Named fields
user := User{
    ID:       123,
    Name:     "John Doe",
    Email:    "john@example.com",
    Created:  time.Now(),
}

// Avoid: Positional arguments (fragile)
user := User{123, "John Doe", "john@example.com", time.Now()}
```

### Constructor Functions
Use constructor functions for complex initialization.

```go
func NewUser(name, email string) (*User, error) {
    if name == "" {
        return nil, errors.New("name is required")
    }
    if !isValidEmail(email) {
        return nil, errors.New("invalid email")
    }
    
    return &User{
        ID:      generateID(),
        Name:    name,
        Email:   email,
        Created: time.Now(),
    }, nil
}
```

### Functional Options Pattern
Use functional options for flexible configuration.

```go
type ServerConfig struct {
    Port    int
    Timeout time.Duration
    TLS     bool
}

type ServerOption func(*ServerConfig)

func WithPort(port int) ServerOption {
    return func(c *ServerConfig) {
        c.Port = port
    }
}

func WithTimeout(timeout time.Duration) ServerOption {
    return func(c *ServerConfig) {
        c.Timeout = timeout
    }
}

func WithTLS(enabled bool) ServerOption {
    return func(c *ServerConfig) {
        c.TLS = enabled
    }
}

func NewServer(opts ...ServerOption) *Server {
    config := &ServerConfig{
        Port:    8080,           // defaults
        Timeout: 30 * time.Second,
        TLS:     false,
    }
    
    for _, opt := range opts {
        opt(config)
    }
    
    return &Server{config: config}
}

// Usage
server := NewServer(
    WithPort(9090),
    WithTimeout(60*time.Second),
    WithTLS(true),
)
```

### Builder Pattern
Use builders for complex object construction.

```go
type QueryBuilder struct {
    table   string
    fields  []string
    where   []string
    orderBy string
    limit   int
}

func NewQueryBuilder() *QueryBuilder {
    return &QueryBuilder{}
}

func (qb *QueryBuilder) From(table string) *QueryBuilder {
    qb.table = table
    return qb
}

func (qb *QueryBuilder) Select(fields ...string) *QueryBuilder {
    qb.fields = append(qb.fields, fields...)
    return qb
}

func (qb *QueryBuilder) Where(condition string) *QueryBuilder {
    qb.where = append(qb.where, condition)
    return qb
}

func (qb *QueryBuilder) OrderBy(field string) *QueryBuilder {
    qb.orderBy = field
    return qb
}

func (qb *QueryBuilder) Limit(n int) *QueryBuilder {
    qb.limit = n
    return qb
}

func (qb *QueryBuilder) Build() string {
    // Build SQL query string
    query := "SELECT " + strings.Join(qb.fields, ", ")
    query += " FROM " + qb.table
    if len(qb.where) > 0 {
        query += " WHERE " + strings.Join(qb.where, " AND ")
    }
    if qb.orderBy != "" {
        query += " ORDER BY " + qb.orderBy
    }
    if qb.limit > 0 {
        query += fmt.Sprintf(" LIMIT %d", qb.limit)
    }
    return query
}

// Usage
query := NewQueryBuilder().
    Select("id", "name", "email").
    From("users").
    Where("active = true").
    Where("created > '2023-01-01'").
    OrderBy("name").
    Limit(10).
    Build()
```

---

## Channel Patterns

### Channel Direction
Use directional channels to clarify intent and prevent misuse.

```go
// Producer only sends
func producer(ch chan<- int) {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)
}

// Consumer only receives
func consumer(ch <-chan int) {
    for value := range ch {
        fmt.Println(value)
    }
}

// Coordinator uses bidirectional channel
func coordinator() {
    ch := make(chan int, 5)
    go producer(ch)
    consumer(ch)
}
```

### Channel Closing Pattern
Always close channels from the sender side.

```go
func worker(jobs <-chan Job, results chan<- Result) {
    defer close(results) // Close when done sending
    
    for job := range jobs {
        result := processJob(job)
        results <- result
    }
}
```

### Select Statement Patterns
Use select for non-blocking operations and timeouts.

```go
// Non-blocking send
select {
case ch <- value:
    // Sent successfully
default:
    // Channel full, handle accordingly
    return errors.New("channel busy")
}

// Timeout pattern
select {
case result := <-resultCh:
    return result, nil
case <-time.After(5 * time.Second):
    return nil, errors.New("operation timed out")
}

// Context cancellation
select {
case result := <-resultCh:
    return result, nil
case <-ctx.Done():
    return nil, ctx.Err()
}
```

### Fan-Out/Fan-In Pattern
Distribute work and collect results.

```go
func fanOut(input <-chan Job, workers int) []<-chan Result {
    channels := make([]<-chan Result, workers)
    
    for i := 0; i < workers; i++ {
        ch := make(chan Result)
        channels[i] = ch
        
        go func(ch chan<- Result) {
            defer close(ch)
            for job := range input {
                result := processJob(job)
                ch <- result
            }
        }(ch)
    }
    
    return channels
}

func fanIn(channels ...<-chan Result) <-chan Result {
    out := make(chan Result)
    
    var wg sync.WaitGroup
    wg.Add(len(channels))
    
    for _, ch := range channels {
        go func(ch <-chan Result) {
            defer wg.Done()
            for result := range ch {
                out <- result
            }
        }(ch)
    }
    
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}
```

---

## Function Patterns

### Method Chaining
Design methods to return the receiver for chaining.

```go
type StringBuilder struct {
    data strings.Builder
}

func (sb *StringBuilder) Add(s string) *StringBuilder {
    sb.data.WriteString(s)
    return sb
}

func (sb *StringBuilder) AddLine(s string) *StringBuilder {
    sb.data.WriteString(s + "\n")
    return sb
}

func (sb *StringBuilder) String() string {
    return sb.data.String()
}

// Usage
result := NewStringBuilder().
    Add("Hello ").
    Add("World").
    AddLine("!").
    Add("Goodbye").
    String()
```

### Variadic Functions
Use variadic parameters for flexible APIs.

```go
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

// Usage
fmt.Println(sum(1, 2, 3, 4, 5))
fmt.Println(sum(slice...)) // expand slice
```

### Closure Patterns
Use closures for configuration and state encapsulation.

```go
// Configuration closure
func createValidator(minLength int) func(string) error {
    return func(s string) error {
        if len(s) < minLength {
            return fmt.Errorf("string too short, minimum length is %d", minLength)
        }
        return nil
    }
}

// State encapsulation
func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// Usage
validator := createValidator(5)
counter := createCounter()

err := validator("hi")    // error: too short
count := counter()        // 1
count = counter()         // 2
```

---

## Package Organization

### Package Naming
Use short, descriptive package names.

```go
// Good
package user
package auth
package http

// Avoid
package utilities
package helpers
package common
```

### Internal Packages
Use internal packages to prevent external access.

```go
// File structure:
// myapp/
//   internal/
//     auth/
//       token.go    // Only accessible within myapp
//   user/
//     service.go    // Can use internal/auth
//   main.go
```

### Package Documentation
Document packages with a package comment.

```go
// Package auth provides authentication and authorization utilities
// for the application. It supports JWT tokens, role-based access
// control, and session management.
//
// Basic usage:
//
//   token, err := auth.GenerateToken(userID)
//   if err != nil {
//       log.Fatal(err)
//   }
//
//   claims, err := auth.ValidateToken(token)
//   if err != nil {
//       http.Error(w, "Unauthorized", http.StatusUnauthorized)
//       return
//   }
package auth
```

### Dependency Injection
Use interfaces for dependency injection.

```go
// Define interface in consumer package
type UserRepository interface {
    FindByID(id string) (*User, error)
    Save(user *User) error
}

type UserService struct {
    repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

// Implementation in separate package
type DatabaseUserRepository struct {
    db *sql.DB
}

func (r *DatabaseUserRepository) FindByID(id string) (*User, error) {
    // Database implementation
}
```

---

## Testing Patterns

### Table-Driven Tests
Use table-driven tests for comprehensive coverage.

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"with zero", 0, 5, 5},
        {"negative numbers", -2, -3, -5},
        {"mixed signs", -2, 3, 1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d, want %d", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

### Test Helpers
Create helpers to reduce test boilerplate.

```go
func setupTestDB(t *testing.T) *sql.DB {
    t.Helper() // Mark as helper function
    
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        t.Fatalf("Failed to open test database: %v", err)
    }
    
    t.Cleanup(func() {
        db.Close()
    })
    
    return db
}

func TestUserRepository(t *testing.T) {
    db := setupTestDB(t)
    repo := NewUserRepository(db)
    
    // Test implementation
}
```

### Mock Interfaces
Use interfaces for easy mocking.

```go
type MockUserRepository struct {
    users map[string]*User
    err   error
}

func (m *MockUserRepository) FindByID(id string) (*User, error) {
    if m.err != nil {
        return nil, m.err
    }
    return m.users[id], nil
}

func (m *MockUserRepository) Save(user *User) error {
    if m.err != nil {
        return m.err
    }
    m.users[user.ID] = user
    return nil
}

func TestUserService_GetUser(t *testing.T) {
    mockRepo := &MockUserRepository{
        users: map[string]*User{
            "123": {ID: "123", Name: "John"},
        },
    }
    
    service := NewUserService(mockRepo)
    user, err := service.GetUser("123")
    
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    if user.Name != "John" {
        t.Errorf("Expected name 'John', got '%s'", user.Name)
    }
}
```

---

## Performance Idioms

### Slice Preallocation
Preallocate slices when size is known.

```go
// Good: Preallocate when size is known
items := make([]Item, 0, expectedSize)
for i := 0; i < expectedSize; i++ {
    items = append(items, createItem(i))
}

// Better: Direct indexing when possible
items := make([]Item, expectedSize)
for i := 0; i < expectedSize; i++ {
    items[i] = createItem(i)
}
```

### String Building
Use strings.Builder for efficient string concatenation.

```go
// Good: Use strings.Builder
var builder strings.Builder
builder.Grow(estimatedSize) // Hint at final size
for _, item := range items {
    builder.WriteString(item.String())
    builder.WriteString("\n")
}
result := builder.String()

// Avoid: Repeated string concatenation
var result string
for _, item := range items {
    result += item.String() + "\n" // Creates new string each time
}
```

### Buffer Reuse
Reuse buffers to reduce allocations.

```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return bytes.NewBuffer(make([]byte, 0, 1024))
    },
}

func processData(data []byte) ([]byte, error) {
    buf := bufferPool.Get().(*bytes.Buffer)
    defer func() {
        buf.Reset()
        bufferPool.Put(buf)
    }()
    
    // Use buffer for processing
    buf.Write(data)
    // ... processing logic
    
    return buf.Bytes(), nil
}
```

### Map Preallocation
Preallocate maps when size is known.

```go
// Good: Preallocate with estimated size
m := make(map[string]int, estimatedSize)

// For exact size
items := []string{"a", "b", "c", "d", "e"}
m := make(map[string]int, len(items))
for i, item := range items {
    m[item] = i
}
```

---

## Style Conventions

### Naming Conventions
Follow Go naming conventions consistently.

```go
// Variables: camelCase
var userName string
var maxRetryCount int

// Constants: camelCase (not SCREAMING_SNAKE_CASE)
const defaultTimeout = 30 * time.Second
const maxConcurrentRequests = 100

// Functions: camelCase, exported functions PascalCase
func calculateTotal(items []Item) float64 { }
func ProcessPayment(amount float64) error { } // exported

// Types: PascalCase for exported, camelCase for unexported
type User struct { }        // exported
type config struct { }      // unexported

// Interfaces: -er suffix when possible
type Reader interface { }
type Writer interface { }
type Processor interface { }
```

### Comment Conventions
Write clear, helpful comments.

```go
// Package comment: Describes the package purpose
// Package auth provides JWT token authentication.

// Function comment: Describes what it does, inputs, outputs
// GenerateToken creates a new JWT token for the given user ID.
// It returns the token string and any error encountered during generation.
func GenerateToken(userID string) (string, error) {
    // Implementation comments explain complex logic
    // Use HMAC-SHA256 for signing
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secretKey))
}

// Struct comment: Describes the type and its purpose
// User represents an authenticated user in the system.
type User struct {
    // Field comments for non-obvious fields
    ID       string    `json:"id"`
    Name     string    `json:"name"`
    Created  time.Time `json:"created_at"` // RFC3339 format
}
```

### Import Organization
Organize imports in groups.

```go
import (
    // Standard library
    "context"
    "fmt"
    "net/http"
    "time"

    // Third-party packages
    "github.com/gorilla/mux"
    "github.com/lib/pq"

    // Local packages
    "myapp/internal/auth"
    "myapp/internal/database"
)
```

### Variable Declaration
Use appropriate variable declaration forms.

```go
// Short variable declaration for local variables
user := getCurrentUser()
count := len(items)

// var for zero values and package-level variables
var (
    client   *http.Client
    config   Config
    errCount int64
)

// const for compile-time constants
const (
    maxRetries = 3
    timeout    = 30 * time.Second
)
```

### Line Length and Formatting
Keep lines readable and well-formatted.

```go
// Good: Break long function calls
result, err := processComplexOperation(
    parameter1,
    parameter2,
    parameter3,
    optionalConfig{
        Timeout: 30 * time.Second,
        Retries: 3,
    },
)

// Good: Break long struct literals
user := User{
    ID:       generateID(),
    Name:     "John Doe",
    Email:    "john@example.com",
    Settings: UserSettings{
        Theme:         "dark",
        Notifications: true,
        Language:      "en",
    },
}
```

---

## Quick Reference Summary

### Common Patterns Checklist
- ✅ Always check errors immediately
- ✅ Use descriptive variable and function names
- ✅ Prefer small, focused interfaces
- ✅ Check for nil before dereferencing
- ✅ Use context.Context for cancellation
- ✅ Close channels from sender side
- ✅ Preallocate slices and maps when size is known
- ✅ Use table-driven tests
- ✅ Document exported functions and types
- ✅ Follow import organization conventions

### Anti-Patterns to Avoid
- ❌ Ignoring errors with `_`
- ❌ Large, monolithic interfaces
- ❌ Panic in libraries (use errors instead)
- ❌ Goroutines without proper cleanup
- ❌ Global mutable state
- ❌ String concatenation in loops
- ❌ Not checking channel closure
- ❌ Mixing receiver types in methods
- ❌ Deep nesting (use early returns)
- ❌ Generic variable names (`data`, `info`, `temp`)

This quick reference covers the essential Go idioms and patterns you'll encounter throughout the curriculum. Keep it handy during exercises and refer back when you need to verify the idiomatic way to solve a particular problem in Go.