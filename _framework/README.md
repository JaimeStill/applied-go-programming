# Agentic Exercise Curriculum System (AECS)

The AECS is a methodology for creating world-class, Exercise-driven technical learning paths through subagent collaboration. AECS combines Exercise-first development, AECS-constrained subagents, and atomic Concept implementation to produce consistent, high-quality learning experiences where every interaction centers on practical building.

## Table of Contents

1. [Framework Philosophy](#framework-philosophy)
2. [System Architecture](#system-architecture)
3. [Implementation Methodology](#implementation-methodology)
4. [Usage Instructions](#usage-instructions)
5. [Quality Standards](#quality-standards)
6. [Best Practices](#best-practices)
7. [Scalability Guide](#scalability-guide)
8. [Glossary](#glossary)

## Framework Philosophy

### Vocabulary

**Learning Container**
- **Subject** - The topic of study
- **Path** - The structured journey through the subject

**Content Organization**
- **Stage** - Major progression level within the path

**Instructional Units**
- **Concept** - The specific learning target to be mastered
- **Exercise** - Practical implementation demonstrating the Concept
- **Appendix** - Concise, quick-reference supplementary materials

**Complete Hierarchy:**
Subject → Path → Stage → Concept → Exercise (+ Appendix parallel)

### Foundation Principles

**Exercise Primacy**
Every learning interaction centers on doing rather than consuming. The Exercise is the atomic unit of all learning activity, designed to immediately engage learners in practical application.

**Concept Atomicity**
Each Exercise addresses exactly one Concept in its simplest useful form. This ensures focused learning outcomes and prevents cognitive overload while maintaining clear progression markers.

**Progressive Complexity**
Concepts build incrementally across Stages, with each new Concept depending on previously established understanding through practical mastery.

**Applied Understanding**
Learning occurs through building working examples rather than theoretical study. Every Concept must be demonstrated through practical implementation that produces tangible results.

### Structural Design

**Path Architecture**
A Path represents the complete learning journey through a Subject, organized into sequential Stages that mark major progression levels. Each Stage contains related Concepts that are mastered through dedicated Exercise implementations.

**Exercise Structure**
Each Exercise demonstrates a single Concept through practical implementation:
- Immediate hands-on building requirements
- Working code that produces tangible results
- Step-by-step implementation guidance
- Practical validation and verification
- Real-world application context

**Appendix Structure**
Appendices serve as supplementary reference materials parallel to Concept/Exercise relationships. They provide concise, quick-reference information that supports the overall Subject without being tied to specific Concepts or Exercises. Appendices maintain flexible formatting to accommodate various types of reference content.

**Complexity Boundaries**
Exercises maintain strict scope limitations to ensure completion feasibility and focus retention. Concepts are sequenced in dependency chains where mastery of each enables progression to the next, while practical implementations remain standalone and immediately applicable.

**Learning Progression**
Learners advance through the Path by mastering Concepts through their implementing Exercises within each Stage before progressing to subsequent Stages. This ensures foundational understanding while maintaining momentum through practical achievement of concept mastery.

## System Architecture

### Core Components

```
Repository Structure:
├── _framework/              # AECS methodology and templates
│   ├── README.md           # This documentation
│   ├── CLAUDE.md           # AECS-specific Claude directives
│   ├── _templates/         # AECS-compliant content templates
│   └── _prompts/           # Exercise-first generation prompts
├── PATH.md                 # Main Path document (Subject learning journey)
├── .claude/agents/         # AECS-constrained AI agent profiles
└── stages/                 # Exercise-driven learning content
    ├── 01-stage-name/         # Stage directories containing Exercise sequences
    │   ├── stage.md           # Stage organization and Exercise progression
    │   └── 01-exercise.md     # Individual atomic Exercises
    └── appendices/            # Reference materials supporting Exercises
        ├── 01-appendix-name.md
        └── 02-appendix-name.md
```

### AECS Hierarchy

All content strictly follows:
**Subject** → **Path** → **Stage** → **Concept** → **Exercise** (+ Appendix parallel)

### Template System

Templates define AECS-compliant structure and ensure Exercise Primacy:

- **path.md**: Overall Subject Path with Concept-driven progression
- **stage.md**: Stage organization with Concepts and their implementing Exercises
- **exercise.md**: Exercise format demonstrating single Concepts through practical implementation
- **subagent.md**: AECS-constrained AI agent persona definitions
- **No appendix template**: Appendices maintain flexible formatting as reference materials

### Prompt System

The prompt system uses clear create-/generate- vocabulary distinction:

**Creation Prompts (Initialize from scratch)**:
- **create-path.md**: Bootstrap complete AECS projects with auto-generated subagents
- **create-subagent.md**: Initialize domain-specific AECS-constrained agent personas  
- **create-stage.md**: Initialize new Stages independently with inverse-prompting
- **create-exercise.md**: Create Exercises flexibly (standalone/enhancement/integration)

**Generation Prompts (Derive from existing infrastructure)**:
- **generate-stage.md**: Build Stage progressions from existing PATH.md analysis
- **generate-exercise.md**: Generate Exercises using context derivation from existing content
- **generate-appendix.md**: Create reference materials with inverse-prompting for missing parameters

**Review Prompts**:
- **review-content.md**: Conduct AECS compliance verification and improvement

### AECS-Constrained Subagent Collaboration

Specialized subagents ensure AECS compliance across all content:

- **Domain Expert**: Code implementation with Exercise Primacy and Concept Atomicity enforcement
- **CS Professor**: Exercise-driven learning validation with Applied Understanding verification
- **Academic Editor**: AECS structure maintenance with Progressive Complexity validation

## Implementation Methodology

### File Naming Conventions

**Directory Structure**
- **Stages**: `stages/##-[stage-name]/` (e.g., `stages/01-concurrency/`)
- **Stage Organization**: `stages/##-[stage-name]/stage.md`
- **Concept Exercises**: `stages/##-[stage-name]/##-[concept-name].md`
- **Reference Materials**: `stages/appendices/##-[appendix-name].md` (supplementary, parallel to Concept/Exercise relationships)

### Exercise-First Command Structure

Execute prompts using AECS-aligned format:
```
Execute [prompt-name].md. [PARAMETER] is [single-concept/stage/path].
```

### AECS-Compliant Examples

**Bootstrap a complete AECS project:**
```
Execute _framework/_prompts/create-path.md. SUBJECT is "[Any Technical Subject]". TARGET_AUDIENCE is "[target learners]". DIFFICULTY_PROGRESSION is "Fundamentals to [Advanced Level]".
```

**Create a Stage independently:**
```
Execute _framework/_prompts/create-stage.md. STAGE_TITLE is "[Stage Topic]". TARGET_CONCEPTS is "[abstract concept description]".
```

**Create an Exercise flexibly:**
```
Execute _framework/_prompts/create-exercise.md. CONCEPT is "[Single Atomic Concept]". SCENARIO is "[standalone/enhancement/integration]".
```

**Generate from existing infrastructure:**
```
Execute _framework/_prompts/generate-stage.md. STAGE_TITLE is "[Stage Topic]". TARGET_CONCEPTS is "[concepts for implementation]".
```

### Parameter Requirements

Prompts enforce AECS principles through parameter constraints:
- `SUBJECT`: Main topic requiring Exercise-driven mastery
- `CONCEPT`: Single, atomic learning target for one Exercise
- `STAGE`: Progression level containing related Concepts through Exercise sequences
- `DEPENDENCIES`: Previously mastered Concepts (Progressive Complexity)

## Quality Standards

### Built-in AECS Compliance Process

Every generated Exercise undergoes automatic AECS verification:

1. **Exercise Primacy Check**: Ensures immediate hands-on building, no passive consumption
2. **Concept Atomicity Verification**: Confirms single-concept focus per Exercise
3. **Applied Understanding Validation**: Requires practical implementation producing tangible results
4. **Progressive Complexity Audit**: Verifies proper dependency chains between Concepts
5. **AECS Vocabulary Enforcement**: Ensures proper Subject→Path→Stage→Concept→Exercise usage

### AECS Quality Benchmarks

- **Exercise Primacy**: All learning interactions center on doing, not consuming
- **Concept Atomicity**: Each Exercise addresses exactly one Concept in simplest form
- **Applied Understanding**: Every Concept demonstrated through practical building
- **Progressive Complexity**: Clear dependency chains between mastered Concepts
- **Technical Accuracy**: All implementations compile, run, and produce results
- **AECS Vocabulary**: Consistent use of proper hierarchy terminology

### Exercise Format Standard

AECS Exercises follow atomic, practical structure:

#### Required Sections

1. **Single Concept Implementation**
   - One Concept addressed through immediate practical building
   - Working code that produces tangible, testable results
   - No theoretical explanations without direct practical application

2. **Dependency Verification**
   - Clear statement of required previously mastered Concepts
   - Progressive Complexity through proper concept chains
   - Self-contained implementation using only mastered dependencies  

3. **Practical Building Task**
   - Hands-on implementation requirements
   - Immediate engagement in working code creation
   - Tangible outcomes that can be seen, run, or tested

4. **Applied Validation**
   - Practical verification methods (not theoretical assessment)
   - Working code demonstrations
   - Real-world implementation testing

#### AECS Exercise Violations to Prevent

❌ **Prohibited Elements**:
- Multi-concept Exercises violating atomicity
- Theoretical overviews without immediate practical application
- Passive consumption elements (readings, videos, explanations)
- Academic assessment patterns (quizzes, reflections, evaluations)
- Complex exercises spanning multiple concepts

✅ **Required Elements**:
- Single-concept focus with immediate practical building
- Working implementations producing tangible results
- Progressive dependency on previously mastered Concepts
- Hands-on tasks requiring practical code creation

## Best Practices

### Exercise-First Content Generation

- **Start with atomic Exercises** implementing single Concepts through practical building
- **Maintain Concept Atomicity** throughout all Exercise development
- **Build Progressive Complexity** through clear dependency chains between mastered Concepts within and across Stages
- **Ensure Applied Understanding** by requiring tangible implementation results
- **Use discovery-driven requirements** based on practical implementation needs

### AECS-Constrained Subagent Collaboration

#### Critical Best Practices for AECS Compliance

**Problem**: Subagents may revert to traditional academic structures violating AECS principles.

**Solution**: Use AECS-aligned prompting that explicitly enforces Exercise Primacy and Concept Atomicity.

#### AECS-Compliant Subagent Prompting Patterns

❌ **AECS Violations**: "Please review the content structure"
✅ **AECS Compliant**: "Fix violations of Exercise Primacy by ensuring this content requires immediate hands-on building"

❌ **AECS Violations**: "Validate the learning progression" 
✅ **AECS Compliant**: "Enforce Concept Atomicity by restructuring any Exercises addressing multiple concepts"

❌ **AECS Violations**: "Check for consistency"
✅ **AECS Compliant**: "Verify Progressive Complexity by ensuring proper dependency chains between mastered Concepts"

#### AECS-Aligned Subagent Engagement Framework

**1. Enforce Exercise Primacy**
- Use action verbs: "Implement", "Build", "Create", "Demonstrate"
- Avoid consumption verbs: "Explain", "Describe", "Overview", "Introduce"

**2. Maintain Concept Atomicity**
- Focus on single-concept implementation: "Fix multi-concept violations", "Ensure single-concept focus"
- Verify atomic scope: "Split complex Exercises into atomic components"

**3. Require Applied Understanding**
- Request practical building: "Implement working code", "Create tangible results", "Build practical demonstrations"
- Prevent theoretical content: "Replace explanations with hands-on building", "Remove passive consumption elements"

**4. Validate Progressive Complexity**
- Check dependency chains: "Verify this Exercise depends only on previously mastered Concepts"
- Ensure proper progression: "Confirm Progressive Complexity through dependency validation"

#### AECS Subagent Workflow Integration

- **Invoke AECS-constrained subagents** for all content creation and review
- **Specify concrete AECS compliance actions** to prevent academic structure reversion
- **Request practical implementation deliverables** rather than theoretical assessments
- **Iterate based on AECS compliance feedback** to maintain Exercise Primacy
- **Document successful AECS prompting patterns** for consistent compliance

### AECS Prompt Engineering

- **Use Exercise-driven requirements** instead of traditional academic specifications
- **Let practical implementation determine scope** rather than predetermined theoretical coverage  
- **Focus on atomic Concept mastery** through hands-on building
- **Avoid academic quantification** like "3-5 exercises" - use Concept-driven development
- **Emphasize Progressive Complexity** through dependency-based Exercise sequences

### AECS Quality Assurance

- **Test all Exercise implementations** for working, tangible results
- **Validate Progressive Complexity** through proper Concept dependency chains
- **Ensure consistent AECS vocabulary** throughout all content
- **Review for Exercise Primacy compliance** preventing passive consumption

### AECS Maintenance

- **Version control templates and prompts** for AECS compliance tracking
- **Update agent profiles** to maintain AECS constraint boundaries
- **Collect feedback** from Exercise-driven learning for continuous AECS improvement
- **Regular AECS compliance audits** to maintain Exercise Primacy standards

## Scalability Guide

### Adapting AECS for New Subjects

1. **Define Subject-Specific AECS Constraints**
   - Create AECS-constrained agent profiles for the domain
   - Identify atomic Concepts that can be implemented individually
   - Establish Progressive Complexity chains for the Subject
   - Define practical implementation patterns for each Concept

2. **Customize Templates for Exercise Primacy**
   - Adapt Exercise format for Subject-specific hands-on building
   - Ensure templates enforce single-concept atomicity
   - Remove any elements enabling passive consumption
   - Add Subject-specific practical implementation requirements

3. **Create AECS-Compliant Generation Prompts**
   - Write prompts that enforce Exercise Primacy in the domain
   - Include parameter validation for single-concept requirements
   - Ensure prompts invoke AECS-constrained subagents appropriately
   - Prevent traditional academic structures in all prompts

4. **Validate with Atomic Exercise Pilot**
   - Generate sample Exercises using new AECS-constrained templates
   - Verify each Exercise implements exactly one Concept practically
   - Test with domain experts for AECS compliance and effectiveness
   - Refine based on Exercise Primacy and Applied Understanding feedback

### AECS Replication Checklist

- [ ] AECS-constrained Subject expert agents defined in `.claude/agents/`
- [ ] Templates customized for Exercise Primacy and Concept Atomicity
- [ ] Generation prompts enforce Applied Understanding and Progressive Complexity
- [ ] AECS quality standards documented and implemented
- [ ] Pilot Exercises generated and validated for AECS compliance
- [ ] Framework documentation updated with Subject-specific AECS guidance

## Glossary

**AECS**: Agentic Exercise Curriculum System - The complete methodology for creating Exercise-driven technical learning paths through AECS-constrained AI collaboration.

**Exercise Primacy**: Foundation principle where every learning interaction centers on doing rather than consuming, with Exercises as the atomic unit of all learning activity.

**Concept Atomicity**: Each Exercise addresses exactly one Concept in its simplest useful form, ensuring focused learning outcomes and clear progression markers.

**Applied Understanding**: Learning occurs through building working examples rather than theoretical study, with every Concept demonstrated through practical implementation.

**Progressive Complexity**: Concepts build incrementally across Stages, with each new Concept depending on previously established understanding through practical mastery.

**AECS-Constrained Subagents**: Specialized AI assistants with domain expertise bounded to enforce AECS principles and prevent violations of Exercise Primacy, Concept Atomicity, Applied Understanding, and Progressive Complexity.

**Exercise-First Development**: Content creation methodology starting with atomic Exercises implementing single Concepts rather than traditional structural planning.

**Subject→Path→Stage→Concept→Exercise**: The complete AECS hierarchy ensuring proper organization of Exercise-driven learning content, with Appendices serving as parallel supplementary reference materials.

---

AECS represents a systematic approach to creating world-class Exercise-driven technical education. By combining AECS-constrained subagents with Exercise-first development and atomic Concept implementation, AECS enables the creation of consistent, high-quality learning experiences that scale across domains while maintaining the philosophy of learning through practical building rather than passive consumption.