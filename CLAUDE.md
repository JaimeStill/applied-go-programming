# Go Mastery Curriculum Development

## Project Overview

This repository contains a comprehensive Go programming curriculum designed to take developers from Go fundamentals to production mastery. The curriculum is structured as a book with multiple chapters, each containing learning roadmaps and practical exercises.

## Repository Structure

```
go-mastery/
├── CLAUDE.md                    # This file
├── .claude/
│   └── agents/                  # Project-specific sub agents
│       ├── go-engineer.md       # Go architecture expert
│       ├── cs-professor.md      # CS education expert  
│       └── academic-editor.md   # Learning materials editor
├── chapters/
│   ├── 00-quick-reference/      # Go syntax quick reference
│   ├── 01-concurrency/          # Concurrency and parallelism
│   ├── 02-data-structures/      # Data structures and algorithms
│   └── ...                      # Additional chapters
├── exercises/
│   └── ...                      # Generated exercise files
└── README.md                    # Public-facing documentation
```

## Claude Code Workflow

This project uses Claude Code's sub agents feature to ensure high-quality content generation. Sub agents are automatically invoked based on the task at hand, or you can explicitly request their help.

### Available Sub Agents

1. **go-engineer** - Reviews and generates Go code with production-level expertise
2. **cs-professor** - Ensures concepts are explained clearly with proper learning progression  
3. **academic-editor** - Polishes content for consistency, clarity, and structure

### How Sub Agents Work

- Each sub agent has its own context window, preventing pollution of the main conversation
- They are configured with specific tools and permissions relevant to their expertise
- Claude Code automatically delegates appropriate tasks to them
- You can also explicitly invoke them: "Ask go-engineer to review this concurrent code"

## Content Development Process

### 1. Initial Content Generation

When creating new content:

- Start by outlining the topic and learning objectives
- Generate initial content focusing on accuracy and clarity
- Let Claude Code automatically delegate to relevant sub agents

### 2. Review and Refinement

For comprehensive review:

- "Have go-engineer review all code examples in this chapter"
- "Ask cs-professor to validate the learning progression"
- "Get academic-editor to polish the final content"

### 3. Validation

Before finalizing:

- Ensure all code compiles and runs correctly
- Verify learning objectives are met
- Check formatting and consistency

## Content Standards

### Code Examples

- Go 1.24+ compatible
- Follow official Go style guide
- Self-contained and runnable
- Well-commented with explanations

### Exercise Format

Each exercise must contain:

1. **Overview** - What the exercise covers
2. **Code Examples** - Demonstrating key concepts
3. **Understanding the Concept** - How to think about the topic
4. **Your Task** - Hands-on exercise instructions

### Writing Guidelines

- Clear, concise technical writing
- Progressive complexity
- Real-world examples
- Consistent terminology

## Working with Sub Agents

### Automatic Delegation

Claude Code will automatically use sub agents when you:

- Generate new Go code or exercises
- Request content review or improvements
- Ask about pedagogical approaches
- Need formatting or structure help

### Explicit Invocation Examples

```
"Have go-engineer create a production-ready worker pool implementation"

"Ask cs-professor if this explanation of channels is clear for beginners"  

"Get academic-editor to ensure this chapter follows our standard format"

"Use go-engineer to validate all code examples in the concurrency chapter"
```

### Chaining Sub Agents

For complex tasks, chain multiple agents:

```
"First have cs-professor design the learning flow for the networking chapter, 
then get go-engineer to implement the code examples, and finally have 
academic-editor polish the entire chapter"
```

## Project Guidelines

### For New Chapters

1. Outline learning objectives and prerequisites
2. Generate initial content with focus on your expertise
3. Let sub agents handle their specialized areas
4. Review and integrate all feedback

### For Code Examples

- Prioritize clarity for learning
- Include error handling  
- Show idiomatic Go patterns
- Add benchmarks for performance topics

### For Exercises

- Start simple, increase complexity
- Provide clear success criteria
- Include common mistakes to avoid
- Offer extension challenges

## Quality Checklist

Before considering content complete:

- [ ] Code reviewed by go-engineer
- [ ] Pedagogy validated by cs-professor  
- [ ] Structure polished by academic-editor
- [ ] All examples compile and run
- [ ] Learning objectives are clear
- [ ] Formatting is consistent
- [ ] Cross-references are accurate

## Tips for Effective Sub Agent Use

1. **Let them work independently** - Sub agents have their own context, so provide complete information
2. **Be specific** - Clear requests get better results
3. **Trust automatic delegation** - Claude Code is good at choosing the right agent
4. **Iterate as needed** - Sub agents can refine their own previous work

Remember: The goal is creating the best possible Go learning resource. Use sub agents proactively to ensure technical accuracy, pedagogical effectiveness, and professional presentation.
