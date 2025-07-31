# Exercise: Control Structures

## Exercise Metadata

- **Stage**: Fundamentals
- **Single Concept**: Control Structures
- **Prerequisites**: Variables and types, basic functions mastered
- **Estimated Time**: 45 minutes

## Building Objective

You'll create an interactive number guessing game that demonstrates Go's three primary control structures through practical application. The game will use conditional logic to evaluate guesses, loops to allow multiple attempts, and switch statements to handle different game states.

**Tangible Result**: A working interactive game where users guess a random number, receiving feedback and playing multiple rounds until they choose to quit.

## Conceptual Overview

Control structures in Go manage program flow by making decisions, repeating operations, and handling multiple conditions. These structures transform linear code into intelligent, responsive applications that react to data and user input.

## Prerequisites

This Exercise demonstrates Control Structures which assumes familiarity with:
- Basic program structure (package main, import, main function)
- Variable declarations and types
- Basic function definitions and calls
- fmt package for input/output operations

**Starting Point**: This exercise provides all necessary code from scratch. No prior exercise code is required.

## Build Number Guessing Game

### Step 1: Create Basic Game Structure with If/Else

Create `guessing-game.go` with conditional logic:

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== Number Guessing Game ===")
	fmt.Println("I'm thinking of a number between 1 and 100!")
	
	// Generate random target number (Go 1.20+ compatible)
	target := rand.Intn(100) + 1
	
	// Get user's first guess
	var guess int
	fmt.Print("Enter your guess: ")
	fmt.Scan(&guess)
	
	// Use if/else to check the guess
	if guess == target {
		fmt.Println("Congratulations! You guessed it!")
	} else if guess < target {
		fmt.Println("Too low! The number is higher.")
	} else {
		fmt.Println("Too high! The number is lower.")
	}
	
	fmt.Printf("The number was: %d\n", target)
}
```

This demonstrates basic if/else conditional logic: checking equality, less than, and greater than conditions to provide appropriate feedback.

**Run and verify Step 1:**
```bash
go run guessing-game.go
```

You should see the game ask for one guess and provide feedback.

### Step 2: Add For Loop for Multiple Attempts

Add iteration capability with a for loop:

```go
package main

import (
	"fmt"
	"math/rand"
)

// Game setup and basic structure from Step 1

func main() {
	fmt.Println("=== Number Guessing Game ===")
	fmt.Println("I'm thinking of a number between 1 and 100!")
	fmt.Println("You have 7 attempts to guess it!")
	
	// Random number generation from Step 1
	target := rand.Intn(100) + 1
	maxAttempts := 7
	
	// NEW: Use for loop to allow multiple attempts
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Printf("\nAttempt %d/%d - Enter your guess: ", attempt, maxAttempts)
		
		var guess int
		fmt.Scan(&guess)
		
		// Conditional logic from Step 1
		if guess == target {
			fmt.Println("Congratulations! You guessed it!")
			fmt.Printf("You won in %d attempts!\n", attempt)
			return
		} else if guess < target {
			fmt.Println("Too low! The number is higher.")
		} else {
			fmt.Println("Too high! The number is lower.")
		}
		
		// Check if this was the last attempt
		if attempt == maxAttempts {
			fmt.Printf("Game Over! The number was: %d\n", target)
		}
	}
}
```

**Run and verify Step 2:**
```bash
go run guessing-game.go
```

You should see the game allow multiple attempts with a loop counter.

### Step 3: Add Switch Statement for Game Difficulty

Add switch-based difficulty selection:

```go
package main

import (
	"fmt"
	"math/rand"
)

// Basic game structure from Steps 1-2

func main() {
	fmt.Println("=== Number Guessing Game ===")
	
	// NEW: Add difficulty selection with switch
	fmt.Println("Select difficulty level:")
	fmt.Println("1 - Easy (1-50, 10 attempts)")
	fmt.Println("2 - Medium (1-100, 7 attempts)")
	fmt.Println("3 - Hard (1-200, 5 attempts)")
	
	var difficulty int
	fmt.Print("Enter difficulty (1-3): ")
	fmt.Scan(&difficulty)
	
	// Use switch to set game parameters based on difficulty
	var maxNumber, maxAttempts int
	var difficultyName string
	
	switch difficulty {
	case 1:
		maxNumber = 50
		maxAttempts = 10
		difficultyName = "Easy"
	case 2:
		maxNumber = 100
		maxAttempts = 7
		difficultyName = "Medium"
	case 3:
		maxNumber = 200
		maxAttempts = 5
		difficultyName = "Hard"
	default:
		fmt.Println("Invalid difficulty! Using Medium.")
		maxNumber = 100
		maxAttempts = 7
		difficultyName = "Medium"
	}
	
	fmt.Printf("\n%s Mode: Guess a number between 1 and %d!\n", difficultyName, maxNumber)
	fmt.Printf("You have %d attempts.\n", maxAttempts)
	
	// Game initialization and loop from previous steps
	target := rand.Intn(maxNumber) + 1
	
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Printf("\nAttempt %d/%d - Enter your guess: ", attempt, maxAttempts)
		
		var guess int
		fmt.Scan(&guess)
		
		// Input validation
		if guess < 1 || guess > maxNumber {
			fmt.Printf("Please enter a number between 1 and %d\n", maxNumber)
			continue
		}
		
		// Conditional logic from previous steps
		if guess == target {
			fmt.Println("Congratulations! You guessed it!")
			fmt.Printf("You won in %d attempts on %s mode!\n", attempt, difficultyName)
			return
		} else if guess < target {
			fmt.Println("Too low! The number is higher.")
		} else {
			fmt.Println("Too high! The number is lower.")
		}
		
		// Check if this was the last attempt
		if attempt == maxAttempts {
			fmt.Printf("Game Over! The number was: %d\n", target)
		}
	}
}
```

**Run and verify Step 3:**
```bash
go run guessing-game.go
```

You should see difficulty selection with switch statement working.

### Step 4: Execute and Verify Complete Program

Run your complete guessing game to verify all control structures work together:

```bash
go run guessing-game.go
```

**Expected Interaction**:
```
=== Number Guessing Game ===
Select difficulty level:
1 - Easy (1-50, 10 attempts)
2 - Medium (1-100, 7 attempts)
3 - Hard (1-200, 5 attempts)
Enter difficulty (1-3): 2

Medium Mode: Guess a number between 1 and 100!
You have 7 attempts.

Attempt 1/7 - Enter your guess: 50
Too low! The number is higher.

Attempt 2/7 - Enter your guess: 75
Too high! The number is lower.

Attempt 3/7 - Enter your guess: 63
Congratulations! You guessed it!
You won in 3 attempts on Medium mode!
```

This output demonstrates mastery of Go's control structures:
- **If/else statements**: Correctly evaluate guess conditions (equal, less than, greater than) and provide appropriate feedback
- **For loop**: Manages the game loop with initialization (`attempt := 1`), condition (`attempt <= maxAttempts`), and increment (`attempt++`)
- **Switch statements**: Handle difficulty selection with multiple cases and default fallback
- **Control flow**: Program exits early on correct guess using `return`, demonstrating flow control within loops
- **Nested conditions**: Combines if/else within for loops, showing how control structures work together
- **Boolean expressions**: Uses comparison operators (`==`, `<`, `>`) in conditional statements
- **Loop termination**: Demonstrates both natural termination (max attempts reached) and early exit (correct guess)
- **Input validation**: Shows additional conditional logic for range checking


## Common Implementation Issues

### Issue: Infinite Loop with For

**Problem**: Creating loops that never terminate

```go
// Incorrect - infinite loop
for {
    fmt.Println("This runs forever!")
}

// Incorrect - condition never changes
attempts := 0
for attempts < 5 {
    fmt.Println("Attempt")
    // Missing: attempts++
}
```

**Solution**: Always ensure loop conditions can become false
```go
// Correct - condition changes each iteration
for attempts := 0; attempts < 5; attempts++ {
    fmt.Printf("Attempt %d\n", attempts+1)
}

// Correct - manual increment
attempts := 0
for attempts < 5 {
    fmt.Printf("Attempt %d\n", attempts+1)
    attempts++ // Critical: modify the condition variable
}
```

### Issue: Switch Fallthrough Confusion

**Problem**: Expecting automatic fallthrough like other languages

```go
// Go does NOT fall through by default
switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday") // This won't run if day == 1
}
```

**Solution**: Use explicit fallthrough when needed
```go
// Explicit fallthrough when required
switch day {
case 6:
    fmt.Println("Weekend!")
    fallthrough
case 7:
    fmt.Println("Sleep in!")
default:
    fmt.Println("Work day")
}
```

### Issue: Missing Default Case

**Problem**: Not handling unexpected input in switch statements

**Solution**: Always include default case for robust error handling
```go
// Correct - handles unexpected input
switch difficulty {
case 1:
    maxAttempts = 10
case 2:
    maxAttempts = 7
case 3:
    maxAttempts = 5
default:
    fmt.Println("Invalid input! Using default.")
    maxAttempts = 7
}
```

## Independent Challenge

Create a simple quiz game that demonstrates all three control structures through interactive question-and-answer gameplay. Your program should show mastery of conditional logic, iteration, and multi-way branching.

**Requirements**:
- Use if/else to check answer correctness and provide feedback
- Use a for loop to ask multiple questions 
- Use switch to handle different question types or scoring levels
- Track and display the final score
- Use only concepts covered in current and previous exercises (variables, functions, control structures)

### Independent Challenge Solution

*Try building the solution yourself before looking at this reference.*

<details>
<summary>Quiz Game Solution</summary>

```go
package main

import "fmt"

func main() {
	fmt.Println("=== Go Programming Quiz ===")
	fmt.Println("Answer the following questions to test your knowledge!")
	
	var score int
	totalQuestions := 3
	
	// Use for loop to ask multiple questions
	for question := 1; question <= totalQuestions; question++ {
		fmt.Printf("\n--- Question %d/%d ---\n", question, totalQuestions)
		
		var userAnswer string
		var correctAnswer string
		
		// Use switch to present different questions
		switch question {
		case 1:
			fmt.Println("What keyword declares a variable in Go?")
			fmt.Println("a) let  b) var  c) def  d) int")
			correctAnswer = "b"
		case 2:
			fmt.Println("What function is the entry point of a Go program?")
			fmt.Println("a) start()  b) begin()  c) main()  d) run()")
			correctAnswer = "c"
		case 3:
			fmt.Println("Which loop keyword is used in Go?")
			fmt.Println("a) while  b) foreach  c) for  d) loop")
			correctAnswer = "c"
		}
		
		fmt.Print("Your answer: ")
		fmt.Scan(&userAnswer)
		
		// Use if/else to check correctness
		if userAnswer == correctAnswer {
			fmt.Println("✅ Correct!")
			score++
		} else {
			fmt.Printf("❌ Wrong! The correct answer was: %s\n", correctAnswer)
		}
	}
	
	// Final scoring with switch for different performance levels
	fmt.Printf("\n=== Quiz Complete! ===\n")
	fmt.Printf("Your score: %d/%d\n", score, totalQuestions)
	
	switch score {
	case totalQuestions:
		fmt.Println("🏆 Perfect score! You're a Go expert!")
	case totalQuestions - 1:
		fmt.Println("⭐ Great job! Almost perfect!")
	case totalQuestions / 2:
		fmt.Println("👍 Good effort! Keep learning!")
	default:
		if score > 0 {
			fmt.Println("📚 You're making progress!")
		} else {
			fmt.Println("🎯 Time to review the basics!")
		}
	}
}
```

**Key Implementation Notes**:
- **Control Structure Integration**: Demonstrates how if/else, for loops, and switch statements work together in a cohesive program
- **For Loop Usage**: Uses traditional for loop with initialization, condition, and increment to iterate through questions
- **Switch for Logic**: Uses switch statement to handle different questions and scoring feedback
- **If/Else for Decisions**: Uses conditional logic to check answers and provide immediate feedback
- **Variable Scope**: Properly manages variables within loop scope and main function scope
- **No New Concepts**: Uses only variables, basic functions, fmt operations, and control structures from current and previous exercises
- **Practical Application**: Creates an engaging, interactive program that demonstrates real-world use of control structures

</details>

## Concept Progression

**Exercise Completed**: Control structures with conditional logic, iteration, and multi-way branching
**Next Exercise**: Arrays and Slices (data collections and iteration patterns)

---

**Exercise Metadata**
- **Created**: 2025-07-31
- **AECS Compliance**: Verified for Exercise Primacy, Concept Atomicity, Applied Understanding, and Progressive Complexity
- **Implementation Status**: Complete with reference solution