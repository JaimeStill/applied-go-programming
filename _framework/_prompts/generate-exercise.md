# Generate Exercise Prompt

## Purpose
Create a comprehensive, hands-on Exercise following the Agentic Exercise Curriculum System (AECS) methodology with strict enforcement of all four AECS principles, ensuring single-concept atomicity and practical implementation focus.

## Parameters
- **CONCEPT**: The single, atomic technical concept being implemented (e.g., "Function Definition", "Loop Implementation", "Class Inheritance")
- **IMPLEMENTATION_FOCUS**: Optional - specific practical building task that demonstrates the concept

## AECS Compliance Requirements

**CRITICAL**: This prompt MUST enforce all four AECS principles:

### Exercise Primacy Enforcement
- Exercise MUST center on hands-on building, not theoretical understanding
- NO passive consumption elements (explanations without immediate implementation)
- ALL learning MUST occur through immediate practical building
- Every section MUST require active implementation rather than reading

### Concept Atomicity Validation
- Exercise MUST address exactly ONE Concept in its simplest useful form
- NO multi-concept implementations that violate atomicity
- Single Concept MUST be implementable through focused hands-on building
- Exercise scope MUST remain atomic and focused

### Applied Understanding Requirements
- Concept MUST be demonstrated through practical implementation producing tangible results
- NO theoretical content without immediate hands-on application
- Exercise MUST require building working examples that can be run and tested
- ALL content MUST produce observable, verifiable outcomes

### Progressive Complexity Verification
- Exercise MUST build only on previously mastered Concepts through practical implementation
- Implementation MUST use only concepts learned through hands-on building in prior Exercises
- NO learning jumps that skip prerequisite practical mastery
- Dependency chain MUST be clearly validated

## Execution Instructions

Use the exercise template at `_framework/_templates/exercise.md` to generate a complete, AECS-compliant Exercise.

### Step 1: Context Derivation and Analysis
Before generating content, derive context from existing infrastructure:
1. **PATH.md Analysis**: Extract stage context and position from CONCEPT analysis
2. **Content Scanning**: Identify prerequisites from existing Exercises and stages
3. **Stage Integration**: Determine which Stage this Exercise belongs to
4. **Dependency Extraction**: Identify previously mastered concepts from infrastructure analysis
5. **Implementation Focus**: Determine appropriate practical building approach if not specified

### Step 2: AECS-Compliant Concept Analysis
Analyze through AECS lens using derived context:
1. **Single Concept Identification**: What ONE atomic concept will be implemented through hands-on building?
2. **Practical Implementation Requirements**: How can this concept be demonstrated through working code that produces tangible results?
3. **Dependency Verification**: What previously mastered concepts through hands-on Exercises are required?
4. **Observable Outcomes**: What specific, testable results prove concept mastery through building?

### Step 3: Single-Concept Implementation Design
Create atomic Exercise structure focusing on single concept:
- **Core Implementation**: Working demonstration of the single concept through hands-on building
- **Practical Application**: Real-world usage of the concept through implementation
- **Implementation Verification**: Validation that the concept works through testing and observation

Each section MUST:
- Focus on the single atomic concept only
- Require immediate hands-on implementation
- Produce tangible, observable results
- Build only on previously mastered concepts

### Step 4: Content Generation
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

**Practical Implementation Focus:**
- Immediate hands-on building requirements
- Working code that demonstrates the concept through implementation
- Tangible results that can be observed and verified
- Step-by-step building instructions for the concept

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

### Step 5: Expert Review Integration
The generated Exercise should be automatically reviewed by AECS-constrained subagents:
- **go-engineer**: Verify single-concept implementation feasibility and hands-on building requirements
- **cs-professor**: Validate Exercise atomicity and practical implementation focus
- **academic-editor**: Ensure AECS vocabulary and structure compliance

### Step 6: AECS Principle Verification
Ensure the Exercise:
- Enforces Exercise Primacy through hands-on building requirements
- Maintains Concept Atomicity with single-concept focus
- Requires Applied Understanding through practical implementation producing results
- Follows Progressive Complexity using only previously mastered concepts through hands-on building
- Uses correct AECS vocabulary: Subject→Path→Stage→Concept→Exercise

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

Generated Exercises should be saved as:
- **File Path**: `stages/##-[stage-name]/##-[single-concept].md`
- **Example**: `stages/01-fundamentals/03-basic-functions.md`

## Example Usage

```
Execute _framework/_prompts/generate-exercise.md. CONCEPT is "[Single Atomic Concept]".
```

This will generate a complete AECS-compliant Exercise with intelligent context derivation, ensuring all four AECS principles are enforced throughout.