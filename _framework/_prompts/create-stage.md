# Create Stage Prompt

## Purpose
Initialize a new Stage independently without requiring existing AECS infrastructure. This prompt can work standalone or integrate with existing Paths, using intelligent analysis and inverse-prompting to ensure comprehensive concept coverage.

## Parameters
- **STAGE_TITLE**: The name of the new Stage (e.g., "Concurrency", "Web Programming", "Database Integration")
- **TARGET_CONCEPTS**: Abstract description of the concepts this Stage should cover (e.g., "concurrent programming patterns", "HTTP server implementation", "data persistence strategies")
- **INSERTION_POINT**: Optional - where in the learning progression this Stage fits (use inverse-prompting if not specified)

## AECS Compliance Requirements

**CRITICAL**: This prompt MUST enforce all four AECS principles:

### Exercise Primacy Enforcement
- Stage MUST center on hands-on building through Exercise sequences
- NO theoretical overviews, explanations, or passive consumption elements
- ALL learning MUST occur through immediate practical implementation
- Every Exercise MUST require active building rather than reading

### Concept Atomicity Validation
- Each Exercise MUST address exactly one Concept in its simplest useful form
- Stage MUST contain only atomic concept implementations
- NO multi-concept Exercises that violate atomicity
- Each Concept MUST be implementable as a single Exercise

### Applied Understanding Requirements
- All Concepts MUST be demonstrated through practical implementation producing tangible results
- NO theoretical content without immediate hands-on application
- Every Exercise MUST require building working examples
- ALL content MUST produce observable, testable outcomes

### Progressive Complexity Verification
- Concepts MUST build incrementally within Stage through proper dependency chains
- Each Exercise MUST depend only on previously mastered Concepts
- Stage MUST prepare for future learning through hands-on mastery progression
- NO learning jumps that skip prerequisite practical implementation

## Execution Instructions

### Step 1: Context Analysis and Integration Planning
Analyze existing infrastructure (if any) and plan integration:

1. **Check for Existing PATH.md**: If present, analyze for integration context
2. **Scan Existing Stages**: Identify prerequisite concepts and progression logic
3. **Determine Placement**: Use INSERTION_POINT or inverse-prompting to establish Stage position
4. **Validate Concept Coverage**: Use inverse-prompting to ensure TARGET_CONCEPTS are comprehensive

### Step 2: Inverse-Prompting for Missing Information
Use intelligent questioning to gather required details:

**If INSERTION_POINT not specified**:
```
"I need to understand where this Stage fits in the learning progression:

**Existing Context**: [Analysis of current PATH.md if present]
**Suggested Placement**: Based on your TARGET_CONCEPTS, this Stage appears to build on [prerequisite concepts] and prepare for [subsequent concepts].

Where should this Stage be positioned? Or should I proceed with my suggested placement?"
```

**If TARGET_CONCEPTS need clarification**:
```
"I want to ensure comprehensive concept coverage for '[STAGE_TITLE]':

**Current Concepts**: [List derived from TARGET_CONCEPTS]
**Potential Gaps**: Are there additional concepts in [domain] that should be included?
**Suggested Additions**: Based on the domain, consider: [intelligent suggestions]

Should I proceed with the current concept list or would you like to add/modify concepts?"
```

### Step 3: AECS-Compliant Concept Breakdown
Transform abstract TARGET_CONCEPTS into atomic, implementable concepts:

1. **Concept Atomicity Analysis**: Break complex descriptions into single-concept units
2. **Implementation Feasibility**: Ensure each concept can be demonstrated through working code
3. **Dependency Mapping**: Establish proper progression within the Stage
4. **Practical Focus**: Define hands-on building tasks for each concept

### Step 4: Stage Structure Generation
Using the stage template at `_framework/_templates/stage.md`, create:

**Stage Organization**:
- **stage.md**: Complete Stage structure with Exercise progression
- **Exercise Sequence**: Ordered list of atomic concepts with building tasks
- **Dependency Chain**: Clear progression showing concept relationships
- **Integration Points**: How this Stage connects to broader Path

**Exercise Planning** (for each concept):
- **Concept**: Single, atomic learning target
- **Implementation Description**: What learners will BUILD
- **Prerequisites**: Previously mastered concepts (from analysis)
- **Tangible Result**: Observable outcome of the Exercise
- **Building Task**: Specific hands-on implementation requirements

### Step 5: Path Integration (if applicable)
If existing PATH.md is present:

1. **Update PATH.md**: Add new Stage to progression
2. **Validate Sequence**: Ensure proper dependency flow
3. **Update References**: Adjust Stage numbering if needed
4. **Maintain Consistency**: Use existing vocabulary and patterns

### Step 6: Standalone Operation (if no existing PATH.md)
If creating Stage independently:

1. **Self-Contained Structure**: Ensure Stage works without external dependencies
2. **Clear Prerequisites**: Document what concepts must be mastered first
3. **Integration Guidance**: Provide instructions for future Path integration
4. **Foundation Preparation**: Set up for subsequent Stage development

## Intelligence Features

### Concept Validation Through Inverse-Prompting
```
"I've identified these atomic concepts for [STAGE_TITLE]:

1. [Concept 1] - [Implementation description]
2. [Concept 2] - [Implementation description]
3. [Concept 3] - [Implementation description]

**Validation Questions**:
- Are these concepts truly atomic (single-concept per Exercise)?
- Is the progression logical for hands-on building?
- Are there missing concepts that learners would need?
- Does each concept produce tangible, testable results?

Please confirm or suggest modifications."
```

### Smart Integration Logic
- **Existing Infrastructure**: Analyze and integrate with current PATH.md
- **Prerequisite Detection**: Scan existing Stages for dependency relationships
- **Numbering Intelligence**: Automatically determine appropriate Stage numbering
- **Consistency Maintenance**: Match existing AECS vocabulary and patterns

### Dependency Chain Validation
- **Internal Dependencies**: Ensure concepts within Stage build properly
- **External Prerequisites**: Identify concepts from previous Stages required
- **Future Preparation**: Consider what concepts this Stage enables
- **Gap Detection**: Identify missing prerequisite concepts

## Output Requirements

Generate complete Stage initialization including:
1. **stage.md**: Complete Stage organization with Exercise sequences
2. **Concept Breakdown**: Atomic concepts with implementation descriptions
3. **Exercise Planning**: Detailed building tasks for each concept
4. **Integration Updates**: PATH.md updates if applicable
5. **Dependency Documentation**: Clear prerequisite and progression chains
6. **AECS Compliance**: All content enforces four core principles

## Success Criteria

The created Stage MUST:
- **Independent Operation**: Work standalone or integrate with existing Paths
- **Concept Atomicity**: Each planned Exercise addresses single concepts
- **Exercise Primacy**: All learning through hands-on building tasks
- **Applied Understanding**: Practical implementation focus throughout
- **Progressive Complexity**: Proper internal and external dependency chains
- **Intelligent Integration**: Smart analysis and inverse-prompting usage
- **AECS Vocabulary**: Consistent Subject→Path→Stage→Concept→Exercise usage

## AECS Violations to Prevent

**Prohibited Elements**:
- Theoretical overviews without immediate hands-on application
- Multi-concept Exercise planning that violates atomicity
- Passive consumption elements in Stage design
- Academic assessment patterns in Exercise planning
- Learning objectives focused on "understanding" rather than "building"

**Required Elements**:
- Exercise sequences requiring immediate practical building
- Single-concept focus per planned Exercise
- Working implementation requirements for all concepts
- Progressive dependency on previously mastered concepts
- Hands-on validation through practical building verification

## Example Usage

**With Existing Infrastructure**:
```
Execute _framework/_prompts/create-stage.md. STAGE_TITLE is "Database Integration". TARGET_CONCEPTS is "database connections, query execution, data modeling, transaction handling". INSERTION_POINT is "after Web Programming stage".
```

**Standalone Initialization**:
```
Execute _framework/_prompts/create-stage.md. STAGE_TITLE is "Concurrency". TARGET_CONCEPTS is "parallel execution patterns, synchronization mechanisms, concurrent data structures".
```

**With Inverse-Prompting**:
```
Execute _framework/_prompts/create-stage.md. STAGE_TITLE is "Performance Optimization". TARGET_CONCEPTS is "profiling and optimization techniques".
```
(Inverse-prompting will ask for placement and concept clarification)

This will create a complete AECS-compliant Stage that can work independently or integrate intelligently with existing infrastructure, using inverse-prompting to ensure comprehensive concept coverage and proper progression.