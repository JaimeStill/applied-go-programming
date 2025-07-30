# Create Path Prompt

## Purpose
Bootstrap a complete AECS project from scratch for any technical Subject. This is the genesis prompt that initializes the entire learning Path infrastructure, including auto-generation of domain-specific subagents and project structure.

## Parameters
- **SUBJECT**: The main topic of study (e.g., "Go Programming", "Italian Cooking", "Digital Photography")
- **TARGET_AUDIENCE**: Intended learners (e.g., "experienced developers", "system administrators", "data engineers")
- **DIFFICULTY_PROGRESSION**: Learning complexity progression (Fundamentals to Advanced/Expert/Production)

## Project Initialization Requirements

**CRITICAL**: This prompt creates a complete AECS project from scratch with no infrastructure assumptions.

### Bootstrap Complete Project Structure
- **PATH.md**: Create main learning Path document
- **stages/**: Initialize exercise-driven learning content directory
- **standalone/**: Create directory for independent exercises outside official path
- **.claude/agents/**: Auto-generate domain-specific AECS-constrained subagents
- **Project files**: Create CLAUDE.md, README.md with project-specific guidance

### Auto-Generate Standardized Subagents
Automatically create the three standardized AECS-constrained subagents using `_framework/prompts/generate-subagents.md`:

**Required Standardized Subagents**:
- **technical-expert**: Subject-specific technical implementation validation
- **learning-designer**: AECS pedagogy and learning progression validation
- **content-editor**: AECS vocabulary and structure compliance validation

**Auto-Generation Process**:
1. Execute: `_framework/prompts/generate-subagents.md. SUBJECT is "[SUBJECT]".`
2. Templates from `_framework/templates/agents/` are processed with [SUBJECT] parameter
3. Three standardized subagents created in `.claude/agents/` directory
4. Eliminates recursive spawning through predictable naming and direct user invocation

**Benefits of Standardized Approach**:
- **Eliminates Recursive Spawning**: Predictable subagent names prevent Task tool issues
- **Framework Consistency**: All learning paths use identical subagent references  
- **Subject Flexibility**: Templates adapt to any domain through parameterization
- **Direct User Invocation**: Subagents work via `@agent-[name]` syntax without nesting
- **Session Independence**: Subagents must exist at session start to be recognized

## AECS Compliance Requirements

**CRITICAL**: All generated content MUST enforce all four AECS principles:

### Exercise Primacy Enforcement
- Path MUST center on hands-on building, not theoretical consumption
- Every Stage MUST contain Exercise sequences requiring practical implementation
- NO passive learning elements (overviews, explanations, readings)
- ALL learning MUST occur through immediate hands-on building

### Concept Atomicity Validation
- Each Exercise MUST address exactly one Concept in its simplest useful form
- NO multi-concept Exercises that violate atomicity
- Stage progression MUST be through atomic concept mastery
- Every Concept MUST be implementable as single Exercise

### Applied Understanding Requirements
- All learning MUST occur through building working examples producing tangible results
- NO theoretical content without immediate practical application
- Every Concept MUST be demonstrated through hands-on implementation
- ALL content MUST require learners to build rather than consume

### Progressive Complexity Verification
- Concepts MUST build incrementally across Stages through dependency chains
- Each Stage MUST depend only on previously mastered Concepts through practical building
- NO learning jumps that skip prerequisite hands-on mastery
- Exercise sequences MUST follow proper dependency progression

## Execution Instructions

### Step 1: Domain Analysis and Subagent Generation
Before creating content, analyze the SUBJECT and auto-generate appropriate subagents:

1. **Identify Technical Domain**: Determine specific expertise needed (programming language, technology stack, operations domain)
2. **Generate Domain Expert**: Create technical specialist using `create-subagent.md`
3. **Generate Pedagogical Expert**: Create cs-professor with domain context
4. **Generate Content Editor**: Create academic-editor with subject awareness
5. **Save Agent Profiles**: Write all generated agents to `.claude/agents/` directory

### Step 2: Project Structure Initialization
Create complete project directory structure:

```
[project-name]/
├── PATH.md                      # Main learning path (generated)
├── CLAUDE.md                    # Project-specific directives (generated)
├── README.md                    # Public documentation (generated)
├── _framework/                  # AECS framework (existing)
├── .claude/agents/              # Auto-generated subagents
│   ├── [domain]-engineer.md    # Technical specialist
│   ├── cs-professor.md          # Pedagogical expert
│   └── academic-editor.md       # Content editor
├── stages/                      # Exercise-driven learning content
│   └── 01-fundamentals/         # Initial stage (generated)
│       └── stage.md
└── standalone/                  # Independent exercises (created empty)
```

### Step 3: AECS-Compliant Content Generation
Using the path template at `_framework/templates/path.md`, generate:

**Exercise-Driven Learning Path**:
- Clear description of what learners will BUILD (not learn theoretically)
- Emphasis on immediate hands-on implementation
- Tangible results and working applications as outcomes

**Practical Mastery Objectives**:
- Specific implementations learners will BUILD through Exercises
- Working applications and systems as concrete outcomes
- Portfolio of built solutions demonstrating practical mastery

**Stage Design**:
For each Stage, provide:
- **Exercise Focus**: What practical building activities center this Stage
- **Concepts Demonstrated Through Exercises**: List concepts demonstrated via Exercise implementation
- **Prerequisites**: Previously mastered concepts through hands-on Exercises
- **Practical Outcomes**: Tangible results produced through Exercise completion

### Step 4: Initialize Foundation Stage
Create Stage 1: Fundamentals with:
- **stage.md**: Stage organization and Exercise progression
- **Fundamental Concepts**: Most basic concepts for the SUBJECT
- **Atomic Exercises**: Single-concept implementations
- **Working Examples**: Immediate practical building tasks

### Step 5: Generate Project Documentation
Create project-specific files:

**CLAUDE.md**: Project implementation guide with:
- AECS framework compliance directive
- Subject-specific implementation requirements
- Generated subagent usage instructions
- Content development processes

**README.md**: Public-facing documentation with:
- Exercise-driven learning approach explanation
- Getting started instructions
- Repository structure overview
- Learning outcomes and requirements

## Auto-Subagent Generation Templates

### Domain Expert Generation
```
Execute _framework/prompts/create-subagent.md. 
DOMAIN is "[SUBJECT domain]". 
ROLE is "AECS-Constrained [Domain] Implementation Expert". 
EXPERIENCE_LEVEL is "Senior [Domain] Engineer with 10+ years and Exercise-driven methodology expertise". 
FOCUS_AREAS is "[domain-specific technical areas] through hands-on building, practical implementation via working examples, [domain] best practices via Exercise sequences".
```

### CS Professor Generation
```
Execute _framework/prompts/create-subagent.md. 
DOMAIN is "Computer Science Education". 
ROLE is "AECS-Constrained Exercise-Driven Learning Expert". 
EXPERIENCE_LEVEL is "PhD in CS Education with 15+ years specializing in Exercise Primacy methodology". 
FOCUS_AREAS is "Exercise-driven concept scaffolding in [SUBJECT], hands-on assessment design, practical building objectives, preventing passive consumption misconceptions in [SUBJECT] learning".
```

### Academic Editor Generation
```
Execute _framework/prompts/create-subagent.md. 
DOMAIN is "Technical Writing and AECS Structure". 
ROLE is "AECS-Constrained Content Structure Expert". 
EXPERIENCE_LEVEL is "Senior Technical Editor with 15+ years and AECS methodology expertise". 
FOCUS_AREAS is "AECS structure maintenance for [SUBJECT], Exercise Primacy enforcement during editing, [SUBJECT] technical accuracy, Progressive Complexity validation".
```

## MANDATORY AECS-Constrained Path Validation

**REQUIRED**: After creating the Path content and subagents, engage the newly created subagents to review and validate AECS compliance. Subagents provide feedback only - they do NOT create new artifacts.

### Subagent Review Process

**Step 1**: Identify the newly created subagents in `.claude/agents/` directory.

**Step 2**: Engage each newly created subagent to review the generated Path content:

**Technical Implementation Review**: Request the newly created technical/engineering subagent to "review this Path content and validate that all technical concepts are accurate and buildable. Identify any Stage progressions that don't build proper foundations and suggest fixes. Verify all proposed Exercises can produce tangible implementations."

**Learning Design Review**: Ask the newly created pedagogical/educational subagent to "review this Path for Exercise Primacy compliance throughout all Stages. Identify any passive consumption elements and suggest restructuring. Validate single-concept atomicity and proper dependency chains."

**Content Structure Review**: Have the newly created academic editing/writing subagent to "review this Path for AECS vocabulary and structure compliance. Identify any hierarchy violations and suggest fixes to maintain Subject→Path→Stage→Concept→Exercise format throughout."

**Step 3**: Incorporate subagent feedback to refine the Path, ensuring all AECS principles are enforced.

## Output Requirements

Generate complete project initialization including:
1. **PATH.md**: Complete learning path with AECS-compliant structure
2. **Project Structure**: All directories and initial files
3. **Auto-Generated Subagents**: Domain-specific AECS-constrained agents
4. **Foundation Stage**: Stage 1 with fundamental concepts and exercises
5. **Project Documentation**: CLAUDE.md and README.md
6. **AECS Compliance**: All content enforces four core principles
7. **Subagent Validation**: Newly created subagents validate the complete Path

## Success Criteria

The created project MUST:
- **Complete Self-Sufficiency**: All infrastructure generated from scratch
- **Domain-Appropriate Subagents**: Auto-generated technical expertise
- **Exercise Primacy**: All learning through hands-on building
- **Concept Atomicity**: Single-concept Exercises throughout
- **Applied Understanding**: Practical implementation focus
- **Progressive Complexity**: Proper dependency chains
- **AECS Vocabulary**: Consistent Subject→Path→Stage→Concept→Exercise usage

## AECS Violations to Prevent

**Prohibited Elements**:
- Theoretical overviews without immediate hands-on application
- Multi-concept Exercises violating atomicity
- Passive consumption elements (readings, videos, explanations)
- Academic assessment patterns (quizzes, tests, evaluations)
- Building objectives focused on "understanding" rather than "implementing"
- TODO comments, placeholders, or incomplete examples in Exercise planning
- "Requirements" lists or "Your Task" sections instead of complete step-by-step examples

**Required Elements**:
- Complete project structure initialization
- Auto-generated domain-specific subagents
- Exercise sequences requiring immediate practical building
- Single-concept focus per Exercise
- Working implementations producing tangible results
- Progressive dependency on previously built concepts

## Example Usage

```
Execute _framework/prompts/create-path.md. SUBJECT is "[Any Technical Subject]". TARGET_AUDIENCE is "[target learner profile]". DIFFICULTY_PROGRESSION is "Fundamentals to [Mastery Level]".
```

This will bootstrap a complete AECS project with auto-generated subagents, full directory structure, PATH.md, initial fundamentals stage, and all supporting documentation ready for immediate Exercise-driven development.