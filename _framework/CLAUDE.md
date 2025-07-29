# AECS Framework Implementation Guide

This document defines Claude Code's implementation requirements for the Agentic Exercise Curriculum System (AECS). All work within AECS-based projects must strictly enforce these directives.

## Core AECS Directive

**You MUST enforce all four AECS principles in every interaction:**

1. **Exercise Primacy**: Every learning interaction centers on hands-on building, never passive consumption
2. **Concept Atomicity**: Each Exercise addresses exactly one Concept in its simplest useful form
3. **Applied Understanding**: Learning occurs through building working examples that produce tangible results
4. **Progressive Complexity**: Concepts build incrementally through properly mastered dependency chains

## AECS Constraint Enforcement

### Immediate AECS Violations to Fix

When you encounter these violations, fix them immediately:

**Exercise Primacy Violations:**
- Theoretical explanations without hands-on building → Replace with working implementations
- Passive content sections → Convert to interactive building exercises
- Overview or conceptual sections → Transform into working code construction activities

**Concept Atomicity Violations:**
- Multi-concept Exercises → Split into atomic, single-concept implementations
- Complex examples introducing multiple concepts → Break into focused, single-concept functions
- Bundled functionality → Separate into individual concept demonstrations

**Applied Understanding Violations:**
- Code without tangible output → Add observable, testable results
- Abstract examples → Replace with concrete, runnable implementations
- Theoretical snippets → Convert to working code requiring practical building

**Progressive Complexity Violations:**
- Code using advanced concepts before basics → Restructure to use only previously mastered concepts
- Dependency leaps → Create proper incremental progression through concept chains

### AECS Hierarchy Enforcement

All content MUST follow: **Subject → Path → Stage → Concept → Exercise (+ Appendix parallel)**

- **Subject**: The topic of study
- **Path**: The structured journey through the subject  
- **Stage**: Major progression level within the path
- **Concept**: The specific learning target to be mastered
- **Exercise**: Practical implementation demonstrating the Concept
- **Appendix**: Concise, quick-reference supplementary materials (parallel to Concept/Exercise)

### AECS-Constrained Subagent Engagement

**Critical**: Always request specific actions that maintain AECS compliance.

❌ **Don't Request Analysis**:
- "Review the Exercise quality"
- "Validate the learning progression"  
- "Check for AECS compliance"

✅ **Request AECS-Aligned Actions**:
- "Fix violations of Exercise Primacy by ensuring this content requires immediate hands-on building"
- "Enforce Concept Atomicity by restructuring any Exercises addressing multiple concepts"
- "Implement Applied Understanding by replacing explanatory content with hands-on building requirements"
- "Verify Progressive Complexity by ensuring proper dependency chains between mastered Concepts"

### AECS Content Development Protocol

For all content creation and modification:

1. **Start with Exercise Implementation**: What single Concept will be demonstrated through practical building?
2. **Ensure Hands-On Building**: Does this require immediate practical implementation producing tangible results?
3. **Validate Single-Concept Focus**: Does this Exercise address exactly one atomic Concept?
4. **Verify Progressive Dependencies**: Does this use only previously mastered Concepts from prior Exercises?
5. **Eliminate Passive Consumption**: Remove any theoretical content not directly tied to hands-on building

### AECS Template and Prompt Usage

**Framework Location**: All AECS templates and prompts are located in `_framework/`

**Template Usage**:
- `_framework/_templates/exercise.md`: Single-concept Exercise format
- `_framework/_templates/stage.md`: Stage organization with Exercise sequences  
- `_framework/_templates/path.md`: Subject Path structure

**Prompt Execution with Create-/Generate- Distinction**:

**Creation Prompts (Initialize from scratch)**:
```
Execute _framework/_prompts/create-path.md. SUBJECT is "...". TARGET_AUDIENCE is "...". DIFFICULTY_PROGRESSION is "...".
Execute _framework/_prompts/create-stage.md. STAGE_TITLE is "...". TARGET_CONCEPTS is "...".
Execute _framework/_prompts/create-exercise.md. CONCEPT is "...". SCENARIO is "...".
Execute _framework/_prompts/create-subagent.md. DOMAIN is "...". ROLE is "...".
```

**Generation Prompts (Derive from existing infrastructure)**:
```
Execute _framework/_prompts/generate-stage.md. STAGE_TITLE is "...". TARGET_CONCEPTS is "...".
Execute _framework/_prompts/generate-exercise.md. CONCEPT is "...".
Execute _framework/_prompts/generate-appendix.md. [Uses inverse-prompting for missing parameters]
```

**Key Principles**:
- **create-**: Bootstrap/initialize without infrastructure assumptions
- **generate-**: Derive from existing AECS infrastructure  
- Use inverse-prompting when parameters are missing
- Single atomic concepts for all Exercise creation
- Hands-on building tasks for Exercise Primacy

### AECS Quality Validation

Before considering any content complete, verify:

- [ ] **Exercise Primacy**: Content centers on doing, not consuming
- [ ] **Concept Atomicity**: Each Exercise addresses exactly one Concept
- [ ] **Applied Understanding**: All learning occurs through practical building
- [ ] **Progressive Complexity**: Proper dependency chains between Concepts
- [ ] **AECS Vocabulary**: Uses Subject→Path→Stage→Concept→Exercise hierarchy
- [ ] **Working Implementation**: All code produces tangible, testable results
- [ ] **Dependency Validation**: Exercise requires only previously mastered Concepts

### AECS Violation Prevention

**Never Allow**:
- Theoretical overviews without immediate practical application
- Multi-concept Exercises that violate atomicity
- Passive consumption elements (readings, explanations, videos)
- Academic assessment patterns without hands-on building
- Complex exercises spanning multiple concepts
- Learning objectives focused on "understanding" rather than "building"

**Always Require**:
- Single-concept focus with immediate practical building
- Working implementations producing tangible results
- Progressive dependency on previously mastered Concepts through hands-on practice
- Hands-on tasks requiring practical code creation
- Observable, testable outcomes from all Exercises

## Success Metrics

**True AECS Implementation Achieved When**:
- Every learning interaction requires immediate hands-on building
- No theoretical content exists without direct practical application
- Each Exercise demonstrates exactly one Concept through implementation
- All progression follows dependency chains of mastered Concepts
- Subagents actively prevent violations of AECS principles
- Content uses proper AECS vocabulary throughout

## Framework Reference

For complete AECS methodology, philosophy, and implementation details, refer to `_framework/README.md`.

**Remember**: The goal is creating Exercise-driven learning experiences where learners master Concepts through practical building rather than passive consumption. Every action must enforce this philosophy.