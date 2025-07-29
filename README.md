# Applied Go Programming

An Exercise-driven Go programming learning Path for experienced developers using the Agentic Exercise Curriculum System (AECS).

## Overview

This repository contains a comprehensive Go programming Path structured as hands-on Exercises. The Path takes developers from Go fundamentals to production mastery through practical implementation rather than theoretical instruction, where every learning interaction centers on building working code.

## Learning Approach

**Exercise-Driven Mastery**: Every concept is learned through immediate hands-on building of working Go programs. No passive consumption - all learning occurs through practical implementation that produces tangible, testable results.

**Single-Concept Atomicity**: Each Exercise addresses exactly one Go concept in its simplest useful form, ensuring focused learning outcomes and clear progression markers.

**Progressive Complexity**: Concepts build incrementally across Stages, with each new Exercise depending only on previously mastered Go features through hands-on practice.

## Repository Structure

```
applied-go-programming/
├── README.md                    # This documentation
├── PATH.md                      # Complete Go learning path roadmap
├── CLAUDE.md                    # Project implementation guide
├── _framework/                  # AECS methodology (reusable framework)
│   ├── README.md               # Complete AECS methodology
│   ├── CLAUDE.md               # AECS-specific implementation directives
│   ├── _templates/             # AECS-compliant content templates
│   └── _prompts/               # Exercise-first generation prompts
├── .claude/agents/             # AECS-constrained content experts
└── stages/                     # Exercise-driven learning content
    ├── 01-fundamentals/        # Go basics through atomic Exercises
    ├── 02-concurrency/         # Goroutines and concurrent patterns
    ├── 03-data-structures/     # Advanced data handling
    ├── 04-web-programming/     # HTTP servers and web applications
    ├── 05-database-integration/ # Data persistence
    ├── 06-testing-quality/     # Testing and code quality
    ├── 07-performance/         # Optimization and profiling
    ├── 08-deployment/          # Containerization and deployment
    ├── 09-monitoring/          # Observability and debugging
    ├── 10-advanced-patterns/   # Complex architectural patterns
    ├── 11-production-operations/ # Production-ready systems
    └── appendices/             # Go reference materials
```

## Getting Started

### Prerequisites

- Experience with at least one programming language
- Go 1.24+ installed and configured
- Basic command-line familiarity
- Text editor or IDE with Go support

### Learning Path

1. **Review the Path**: Start with `PATH.md` for complete learning journey overview
2. **Begin with Fundamentals**: Start with `stages/01-fundamentals/` for Go basics
3. **Progress Sequentially**: Complete Stages in numbered order
4. **Build Every Exercise**: Each Stage contains hands-on Exercises requiring working Go code
5. **Verify Implementation**: All code must compile, run, and produce observable results

### Stage Organization

Each Stage follows this structure:
- `stage.md`: Stage overview and Exercise sequence
- `01-concept.md`, `02-concept.md`, etc.: Individual atomic Exercises

Each Exercise requires:
- Building working Go code that compiles and runs
- Producing tangible, observable results
- Using only previously mastered Go concepts
- Demonstrating exactly one Go concept practically

## Exercise-Driven Learning

### What Makes This Different

**Traditional Approach**: Read about Go concepts → Study examples → Maybe try coding
**AECS Approach**: Immediately build working Go programs demonstrating each concept

### Learning Outcomes

By completing this Path through hands-on building, you will have implemented:

- **Working Go Applications**: Demonstrating core language features through practical code
- **Concurrent Programs**: Using goroutines, channels, and concurrent patterns
- **Web Services**: HTTP servers and web application backends
- **Database Integration**: Data persistence and management systems
- **Production Systems**: Tested, optimized, deployable Go applications
- **Monitoring Solutions**: Observability and debugging implementations

### Exercise Requirements

Every Exercise in this Path:
- Addresses exactly one Go concept through practical building
- Produces working code that compiles and runs successfully
- Generates observable, testable results
- Uses only Go features mastered in previous Exercises
- Follows Go best practices and idioms

## AECS Framework

This Path uses the **Agentic Exercise Curriculum System** (AECS), a methodology for creating Exercise-driven technical learning experiences.

**Core Principles**:
- **Exercise Primacy**: All learning through hands-on building, never passive consumption
- **Concept Atomicity**: Each Exercise demonstrates exactly one concept
- **Applied Understanding**: Learning occurs through building working examples
- **Progressive Complexity**: Concepts build incrementally through mastered dependencies

**Framework Details**: See `_framework/README.md` for complete AECS methodology.

## Content Development

This Path uses AECS-constrained AI agents to ensure quality and consistency:

- **go-engineer**: Verifies Go code quality and best practices
- **cs-professor**: Validates Exercise-driven learning progression  
- **academic-editor**: Maintains AECS structure and consistency

**Content Generation**: Uses `_framework/_prompts/` for creating new AECS-compliant content.

## Quality Standards

All content meets these standards:

- **AECS Compliance**: Enforces Exercise Primacy through hands-on building
- **Go Compatibility**: All code compiles with Go 1.24+
- **Production Quality**: Follows Go best practices and conventions
- **Practical Focus**: Every Exercise produces working, testable results
- **Progressive Building**: Clear dependency chains between concepts

## Getting Help

- **Exercise Issues**: Each Exercise includes troubleshooting guidance
- **Go Questions**: Reference materials in `stages/appendices/`
- **AECS Methodology**: Complete documentation in `_framework/README.md`

## Status

This Path is actively developed using AECS methodology. Stages are continuously refined based on Exercise-driven learning effectiveness and hands-on implementation success.

**Current Status**: Foundation Stages implemented with atomic Exercise sequences ready for hands-on Go programming mastery.

---

**Start Building**: Begin your Go mastery journey with `stages/01-fundamentals/` - where every concept is learned through immediate hands-on implementation of working Go programs.