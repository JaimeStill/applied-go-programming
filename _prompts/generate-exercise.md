# Generate Exercise Prompt

## Purpose
Create a comprehensive, hands-on exercise following the Agentic Exercise Curriculum System (AECS) methodology that combines clear conceptual understanding with practical implementation.

## Parameters
- **TOPIC**: The specific technical concept being taught (e.g., "Goroutine Communication with Channels", "HTTP Middleware Implementation")
- **CHAPTER_CONTEXT**: The chapter this exercise belongs to and its position in the learning sequence
- **DIFFICULTY**: Target complexity level (Beginner/Intermediate/Advanced)
- **CONCEPTS**: Specific technical concepts this exercise should cover
- **PREREQUISITES**: Required knowledge and skills from previous exercises

## Execution Instructions

Use the exercise template at `_templates/exercise.md` to generate a complete, production-quality exercise.

### Step 1: Concept Analysis
Before generating content, analyze:
1. **Core Learning Goal**: What is the single most important thing learners should understand?
2. **Mental Model**: How should learners think about this concept?
3. **Common Misconceptions**: What mistakes do people typically make?
4. **Real-World Application**: Where and when is this concept used professionally?

### Step 2: Progressive Task Design
Create a 3-part exercise structure:
- **Part 1 (Basic)**: Minimal viable implementation demonstrating the core concept
- **Part 2 (Enhanced)**: More realistic implementation with additional requirements
- **Part 3 (Production)**: Professional-grade implementation with error handling, testing, and optimization

### Step 3: Content Generation
Fill the template with:

**Exercise Metadata:**
- Accurate difficulty assessment and time estimates
- Complete list of concepts covered
- Clear prerequisite mapping
- Proper categorization within curriculum

**Learning Objectives:**
- Specific, measurable outcomes using action verbs
- Mix of conceptual understanding and practical skills
- Testable objectives that can be verified

**Overview and Relevance:**
- Clear explanation of what the exercise accomplishes
- Strong connection to real-world professional scenarios
- Motivation for why this concept matters

**Code Examples:**
- Production-ready, idiomatic code following best practices
- Progressive complexity from basic to advanced patterns
- Comprehensive error handling and edge case coverage
- Clear, educational comments explaining key concepts

**Mental Model Section:**
- Clear conceptual framework for understanding the topic
- Analogies and metaphors that aid comprehension
- Connection to related concepts and broader patterns
- Common thinking traps and how to avoid them

**Task Implementation:**
- Three progressive difficulty levels with clear requirements
- Starter code template that compiles and runs
- Specific success criteria for each level
- Setup instructions and environment requirements

**Common Mistakes:**
- At least 3 typical errors with explanations
- Code examples showing both wrong and right approaches
- Clear prevention strategies
- Debugging guidance

**Real-World Context:**
- Multiple industry applications
- Production considerations and best practices
- Performance implications and optimization opportunities
- Security and reliability concerns where applicable

**Extension Challenges:**
- 3 additional challenges of increasing difficulty
- Optional integrations with other concepts
- Performance optimization opportunities
- Creative applications of the core concept

**Testing and Validation:**
- Comprehensive test cases covering normal and edge cases
- Performance benchmarks where relevant
- Validation checklist for self-assessment
- Debugging tips and troubleshooting guidance

**Reference Solution:**
- Complete, well-documented solutions for all three parts
- Alternative implementation approaches where applicable
- Detailed explanation of design decisions
- Performance and trade-off analysis

### Step 4: Expert Review Integration
The generated exercise should be automatically reviewed by:
- **Domain Expert**: Code quality, technical accuracy, best practices
- **Education Expert**: Learning progression, concept clarity, assessment validity
- **Academic Editor**: Writing quality, structure, consistency

### Step 5: Quality Assurance
Ensure the exercise:
- Compiles and runs correctly
- Has clear, unambiguous requirements
- Provides adequate scaffolding for the target difficulty
- Includes proper error handling and edge cases
- Follows industry best practices and conventions

## Output Requirements

Generate a complete exercise document that includes:
1. Comprehensive metadata and learning objectives
2. Clear conceptual foundation and mental models
3. Progressive three-part implementation tasks
4. Production-quality code examples and solutions
5. Common mistakes and debugging guidance
6. Real-world context and applications
7. Extension challenges and further exploration
8. Complete testing and validation framework

## Success Criteria

The generated exercise should:
- Be immediately implementable by learners
- Provide clear success criteria and validation methods
- Build on previous knowledge while introducing new concepts
- Connect to real-world professional scenarios
- Follow AECS principles of hands-on, progressive learning
- Include all necessary scaffolding and support materials

## Special Considerations

### Code Quality Standards
- All code must compile and run correctly
- Follow language-specific best practices and idioms
- Include proper error handling and resource management
- Use production-appropriate patterns and structures
- Provide clear, educational comments

### Learning Progression
- Build incrementally on previous exercises
- Introduce a manageable number of new concepts per exercise to avoid cognitive overload
- Provide multiple practice opportunities for complex concepts
- Include conceptual bridges to future topics

### Assessment Integration
- Include measurable learning objectives
- Provide clear success criteria for each task level
- Offer self-assessment opportunities
- Enable both formative and summative evaluation

### Accessibility
- Support different learning styles (visual, kinesthetic, reading)
- Provide multiple explanation approaches
- Include debugging and troubleshooting support
- Offer extension opportunities for advanced learners

## File Naming Convention

Generated exercises should be saved as:
- **File Path**: `chapters/##-[chapter-name]/##-[exercise-name].md`
- **Example**: `chapters/01-concurrency/03-worker-pools.md`

## Example Usage

```
Execute generate-exercise.md. TOPIC is "Implementing Worker Pool Pattern with Channels". CHAPTER_CONTEXT is "Chapter 2: Concurrent Patterns - Exercise 6 building on previous channel and goroutine exercises". DIFFICULTY is "Intermediate". CONCEPTS is "worker pools, job distribution, graceful shutdown, error aggregation". PREREQUISITES is "goroutines, channels, select statements, basic synchronization".
```

This will generate a complete exercise at `chapters/02-concurrent-patterns/06-worker-pools.md` ready for learner implementation and expert review.