# Applied Go Programming - Project Implementation Guide

## AECS Framework Compliance

**Primary Directive**: Follow the Agentic Exercise Curriculum System (AECS) methodology as defined in `_framework/CLAUDE.md`. All content development must enforce Exercise Primacy, Concept Atomicity, Applied Understanding, and Progressive Complexity.

## Project Overview

This repository contains a comprehensive Go programming learning Path built using AECS methodology. The Path takes developers from Go fundamentals to production mastery through hands-on Exercise completion, where every learning interaction centers on doing rather than consuming.

## Repository Structure

```
applied-go-programming/
├── CLAUDE.md                    # This file - Project implementation guide
├── PATH.md                      # Go Programming learning path
├── README.md                    # Public-facing documentation
├── _framework/                  # AECS methodology and templates
│   ├── README.md               # Complete AECS methodology
│   ├── CLAUDE.md               # AECS-specific Claude directives
│   ├── _templates/             # AECS-compliant content templates
│   └── _prompts/               # Exercise-first generation prompts
├── .claude/
│   └── agents/                 # AECS-constrained subagents
│       ├── go-engineer.md      # Go code implementation expert
│       ├── cs-professor.md     # Exercise-driven learning expert  
│       └── academic-editor.md  # AECS content structure expert
└── stages/                     # Exercise-driven learning content
    ├── 01-fundamentals/        # Stage 1: Go basics through Exercises
    ├── 02-concurrency/         # Stage 2: Concurrent programming
    ├── ...                     # Additional stages
    └── appendices/             # Go-specific reference materials
```

## Go-Specific AECS Implementation

### Subject Context
- **Subject**: Go Programming
- **Target Audience**: Experienced developers familiar with other languages
- **Approach**: Exercise-driven mastery from fundamentals to production

### Stage Progression

The Go Programming Path follows this progression:

1. **Fundamentals** - Basic Go concepts through atomic Exercises
2. **Concurrency** - Goroutines, channels, and concurrent patterns
3. **Data Structures** - Advanced data handling and algorithms  
4. **Web Programming** - HTTP servers and web applications
5. **Database Integration** - Data persistence and management
6. **Testing & Quality** - Testing strategies and code quality
7. **Performance** - Optimization and profiling
8. **Deployment** - Containerization and deployment
9. **Monitoring** - Observability and debugging
10. **Advanced Patterns** - Complex architectural patterns
11. **Production Operations** - Production-ready systems

### Go-Specific Exercise Requirements

All Go Exercises must:

- **Compile and Run**: Use Go 1.24+ compatible code that compiles without errors
- **Produce Tangible Results**: Generate console output, files, or testable behavior
- **Follow Go Idioms**: Demonstrate idiomatic Go patterns and conventions
- **Build Incrementally**: Each Exercise uses only concepts mastered in previous Exercises
- **Single Concept Focus**: Address exactly one Go concept per Exercise

### Code Quality Standards

- **Go Standards**: Follow official Go style guide and best practices
- **Error Handling**: Include proper error handling patterns
- **Resource Management**: Demonstrate proper resource cleanup
- **Documentation**: Include clear, practical comments explaining implementation
- **Testing**: Provide testable examples where appropriate

## AECS-Constrained Subagent Usage

### Go Engineer
Use for Go-specific technical implementation:
```
"Implement working Go code that demonstrates [single concept] through practical application. Ensure the Exercise produces tangible results and follows Go best practices."
```

### CS Professor  
Use for Exercise-driven learning validation:
```
"Restructure this content to enforce Exercise Primacy by ensuring learners immediately engage in hands-on Go programming. Verify single-concept atomicity."
```

### Academic Editor
Use for AECS structure enforcement:
```
"Fix any deviations from AECS vocabulary and structure. Ensure content follows Subject→Path→Stage→Concept→Exercise hierarchy for Go programming."
```

## Content Placement Conventions

### Exercise Placement Logic

**Standalone Exercises**: `/standalone/[concept].md`
- Independent concept demonstrations outside the official learning path
- Not part of main progression - exploratory or supplementary concepts
- Example: `/standalone/binary-search-implementation.md`

**Enhancement Exercises**: `/stages/XX-[stage]/enhancements/[concept].md`  
- Expound upon concepts not fully understood within existing Stages
- Supplementary to core Stage progression
- Example: `/stages/02-concurrency/enhancements/goroutine-error-handling.md`

**Stage Integration Exercises**: `/stages/XX-[stage]/XX-[concept].md`
- Add new concept Exercises to existing Stages (updates stage.md automatically)
- Full integration with Stage progression
- Example: `/stages/04-web-programming/03-json-marshaling.md`

### Directory Structure
```
applied-go-programming/
├── standalone/                    # Independent exercises (outside official path)
│   └── [concept].md
└── stages/                        # Official learning path
    └── XX-[stage]/
        ├── stage.md              # Auto-updated for integrations
        ├── XX-[concept].md       # Core exercises
        └── enhancements/         # Enhancement subdirectory
            └── [concept].md      # Concept expansions
```

## Content Development Process

### 1. Exercise-First Development

For new Go content:
1. Identify the single Go concept to be mastered
2. Create an Exercise demonstrating this concept through working code
3. Ensure the Exercise produces immediate, observable results
4. Verify the Exercise uses only previously mastered Go concepts

### 2. AECS Compliance Verification

Every Go Exercise must pass:
- **Exercise Primacy**: Centers on writing working Go code
- **Concept Atomicity**: Addresses exactly one Go concept  
- **Applied Understanding**: Requires building functional Go programs
- **Progressive Complexity**: Uses only previously learned Go features

### 3. Go-Specific Validation

Additional Go requirements:
- Code compiles with `go build` or `go run`
- Follows `go fmt` formatting standards
- Passes `go vet` analysis
- Includes practical error handling
- Demonstrates real-world Go usage patterns

## Development Commands

### Content Generation

**Creation Commands (Initialize from scratch)**:
```bash
# Create new Stage independently
Execute _framework/_prompts/create-stage.md. STAGE_TITLE is "[Go Topic]". TARGET_CONCEPTS is "[abstract concept description]"

# Create Exercise with flexible placement
Execute _framework/_prompts/create-exercise.md. CONCEPT is "[Single Go Concept]". SCENARIO is "[standalone/enhancement/integration]"

# Create domain-specific subagent
Execute _framework/_prompts/create-subagent.md. DOMAIN is "Go Programming". ROLE is "Implementation Expert"
```

**Generation Commands (Derive from existing infrastructure)**:
```bash
# Generate Stage from existing PATH.md
Execute _framework/_prompts/generate-stage.md. STAGE_TITLE is "[Go Topic]". TARGET_CONCEPTS is "[concepts]"

# Generate Exercise with context derivation
Execute _framework/_prompts/generate-exercise.md. CONCEPT is "[Single Go Concept]"

# Generate appendix with inverse-prompting
Execute _framework/_prompts/generate-appendix.md

# Review Go content for AECS compliance
Execute _framework/_prompts/review-content.md. CONTENT_TYPE is "go programming content"
```

### Go Code Validation
```bash
# Ensure all Go code compiles
find stages -name "*.md" -exec grep -l "```go" {} \; | xargs -I {} sh -c 'echo "Checking Go code in: {}"'

# Validate Go formatting
go fmt ./...

# Run Go static analysis  
go vet ./...
```

## Quality Assurance

### Go-Specific Checklist

Before considering Go content complete:

- [ ] **AECS Compliance**: Enforces all four AECS principles
- [ ] **Go Compilation**: All code examples compile successfully
- [ ] **Go Standards**: Follows official Go style and conventions
- [ ] **Error Handling**: Includes appropriate error handling patterns
- [ ] **Practical Results**: Produces observable, testable outcomes
- [ ] **Progressive Building**: Uses only previously mastered Go concepts
- [ ] **Real-World Context**: Demonstrates practical Go usage scenarios

### Exercise Validation Process

For each Go Exercise:

1. **Technical Review**: go-engineer verifies Go code quality and best practices
2. **Learning Review**: cs-professor validates Exercise-driven progression  
3. **Structure Review**: academic-editor ensures AECS compliance
4. **Integration Testing**: Verify Exercise fits properly in Stage progression

## Project-Specific Notes

### Go Version Compatibility
- Target Go 1.24+ features and syntax
- Include version-specific notes where relevant
- Ensure backward compatibility considerations are addressed

### Learning Path Focus
- Emphasize practical Go programming over theoretical concepts
- Include real-world scenarios and professional development patterns
- Build towards production-ready Go applications

### Exercise Outcomes
By completing this Path, learners will have built:
- Working Go applications demonstrating core language features
- Concurrent programs using goroutines and channels
- Web services and HTTP applications
- Database-integrated applications
- Tested, production-ready Go systems

## Success Metrics

**Go Programming Path Success Achieved When**:
- Every Exercise requires writing working Go code
- No theoretical Go content exists without hands-on implementation
- Each Exercise demonstrates exactly one Go concept practically
- Progression follows proper Go learning dependencies
- Learners build increasingly complex Go applications
- Final outcomes include production-ready Go systems

**Remember**: This is an Exercise-driven Go programming learning experience. Every interaction must center on building working Go code that demonstrates practical mastery through hands-on implementation.