# Create Exercise Prompt

## Purpose
Initialize Exercises flexibly with intelligent placement logic. Handles multiple Exercise scenarios: standalone concept demonstration, Stage enhancement for unclear concepts, and new concept integration into existing Stages.

## Parameters
- **CONCEPT**: The single, atomic concept being implemented (e.g., "Function Definition", "Color Mixing", "Knife Grip Technique")
- **SCENARIO**: Optional - type of Exercise (standalone/enhancement/integration) - uses inverse-prompting if not specified
- **TARGET_LOCATION**: Optional - placement hint (uses intelligent defaults and inverse-prompting if needed)

## Exercise Scenarios

### 1. Standalone Exercise
Independent concept demonstration outside the official learning path.
- **Purpose**: Explore concepts not part of main progression
- **Location**: `/standalone/[concept].md`
- **Integration**: Not part of official Path progression

### 2. Enhancement Exercise  
Expound upon concepts not fully understood within existing Stages.
- **Purpose**: Clarify or expand on concepts covered by existing Exercises
- **Location**: `/stages/XX-[stage]/enhancements/[concept].md`
- **Integration**: Supplementary to core Stage progression

### 3. Stage Integration Exercise
Add new concept Exercise to existing Stage with automatic stage.md updates.
- **Purpose**: Extend existing Stage with new concepts
- **Location**: `/stages/XX-[stage]/XX-[concept].md`
- **Integration**: Updates stage.md with new Exercise reference

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
- Exercise MUST build only on previously mastered Concepts (for integration/enhancement)
- Implementation MUST use only concepts learned through hands-on building in prior Exercises
- NO learning jumps that skip prerequisite practical mastery
- Dependency chain MUST be clearly validated (except for standalone exercises)

## Execution Instructions

### Step 1: Scenario Determination and Context Analysis
Analyze the request and determine appropriate Exercise scenario:

1. **Scenario Detection**: If SCENARIO not specified, use inverse-prompting
2. **Context Analysis**: Read existing infrastructure to understand placement options
3. **Concept Validation**: Ensure CONCEPT is truly atomic and implementable
4. **Placement Intelligence**: Determine optimal location based on scenario and content

### Step 2: Inverse-Prompting for Missing Information
Use intelligent questioning to gather required details:

**If SCENARIO not specified**:
```
"I need to understand what type of Exercise you want for '[CONCEPT]':

**Available Scenarios**:
1. **Standalone**: Independent concept exploration outside official learning path
2. **Enhancement**: Expand on existing concept within a Stage (requires existing Exercise)
3. **Integration**: Add new concept to existing Stage (updates stage.md)

Based on your concept, I suggest: [intelligent recommendation based on analysis]

Which scenario fits your needs, or should I proceed with my recommendation?"
```

**If TARGET_LOCATION needs clarification**:
```
"I need to determine the best placement for this Exercise:

**Scenario**: [Determined scenario]
**Concept**: [CONCEPT]
**Analysis**: [Context from existing content]

**Placement Options**:
- [Intelligent suggestions based on content analysis]
- [Default conventions if no project-specific rules found]

Should I use my suggested placement or do you have a preference?"
```

### Step 3: Content Placement Intelligence

#### Default Directory Structure and Logic

**Standalone Exercises**:
- **Location**: `/standalone/[concept].md`
- **Logic**: Always placed in root standalone directory
- **Integration**: None - completely independent

**Enhancement Exercises**:
- **Location**: `/stages/XX-[stage]/enhancements/[concept].md`
- **Logic**: Determine appropriate Stage through content analysis
- **Integration**: Create enhancements/ subdirectory if needed

**Stage Integration Exercises**:
- **Location**: `/stages/XX-[stage]/XX-[concept].md`
- **Logic**: Add to existing Stage, update stage.md automatically
- **Integration**: Full integration with Stage progression

#### Placement Intelligence Process

1. **Check Project Conventions**: Read root CLAUDE.md for project-specific placement rules
2. **Content Analysis**: Analyze existing Stages and Exercises for context
3. **Intelligent Suggestions**: Provide placement options based on concept relationships
4. **Default Fallback**: Use standard AECS conventions if no specific guidance found
5. **User Confirmation**: Use inverse-prompting for ambiguous cases

### Step 4: AECS-Compliant Exercise Generation
Using the exercise template at `_framework/_templates/exercise.md`, create:

**Single Concept Implementation**:
- **Atomic Focus**: Address exactly one concept through practical building
- **Working Implementation**: Implement concept through hands-on building
- **Tangible Results**: Produce observable, testable outcomes
- **Implementation Guidance**: Step-by-step building instructions

**Dependency Management**:
- **Prerequisite Analysis**: Extract dependencies from existing content (for integration/enhancement)
- **Standalone Independence**: No prerequisites for standalone exercises
- **Progressive Building**: Use only previously mastered concepts

**Practical Validation**:
- **Success Criteria**: Clear measures of successful implementation
- **Testing Requirements**: How to verify the working implementation
- **Real-World Context**: Practical applications of the concept

### Step 5: Integration and Updates
Handle integration based on scenario:

**Standalone Exercise**:
- Create exercise file in `/standalone/`
- No additional integration required

**Enhancement Exercise**:
- Create exercise file in `/stages/XX-[stage]/enhancements/`
- Create enhancements/ directory if it doesn't exist
- No stage.md updates required

**Stage Integration Exercise**:
- Create exercise file in `/stages/XX-[stage]/`
- Update corresponding stage.md with new Exercise reference
- Maintain proper Exercise numbering and progression

## Intelligence Features

### Concept Analysis and Validation
```
"Analyzing your concept '[CONCEPT]':

**Atomicity Check**: This concept appears to be [atomic/needs breakdown]
**Implementation Feasibility**: Can be demonstrated through [practical building approach]
**Dependency Analysis**: Requires these previously mastered concepts: [list]

**Validation Questions**:
- Is this concept truly atomic (single focus)?
- Can it be implemented through hands-on building?
- Does it produce tangible, testable results?

Proceeding with Exercise creation or need adjustments?"
```

### Smart Placement Logic
- **Content Relationships**: Analyze concept connections to existing Exercises
- **Stage Context**: Understand which Stage concepts belong to
- **Progression Logic**: Ensure proper dependency flows
- **Convention Following**: Use project-specific or default AECS conventions

### Automatic Integration
- **Directory Creation**: Create necessary subdirectories (enhancements/)
- **File Updates**: Update stage.md for integration scenarios
- **Numbering Logic**: Maintain proper Exercise numbering sequences
- **Reference Consistency**: Ensure all references remain accurate

## MANDATORY AECS-Constrained Exercise Validation

**REQUIRED**: After creating the Exercise content, engage project subagents to review and validate AECS compliance. Subagents provide feedback only - they do NOT create new artifacts.

### Subagent Review Process

**Step 1**: Identify available subagents in `.claude/agents/` directory.

**Step 2**: Engage each subagent to review the created Exercise content:

**Technical Implementation Review**: Request the technical/engineering subagent to "review this Exercise and validate that the concept can be implemented through practical building. Identify any code that won't work correctly and suggest fixes. Verify all examples use only previously mastered concepts."

**Learning Design Review**: Ask the pedagogical/educational subagent to "review this Exercise for Exercise Primacy compliance. Identify any passive consumption elements and suggest how to restructure for hands-on building. Validate single-concept atomicity and proper dependency chains."

**Content Structure Review**: Have the academic editing/writing subagent to "review this Exercise for AECS vocabulary and structure compliance. Identify any formatting violations and suggest fixes to maintain proper Exercise format with single-concept focus."

**Step 3**: Incorporate subagent feedback to refine the Exercise, ensuring all AECS principles are enforced.

## Output Requirements

Generate complete Exercise initialization including:
1. **Exercise File**: Complete exercise.md using AECS template
2. **Placement Logic**: Appropriate location based on scenario
3. **Integration Updates**: stage.md updates for integration scenarios
4. **Directory Management**: Create subdirectories as needed
5. **AECS Compliance**: All content enforces four core principles
6. **Dependency Documentation**: Clear prerequisite identification
7. **Subagent Validation**: Project subagents validate the complete Exercise

## Success Criteria

The created Exercise MUST:
- **Scenario Flexibility**: Handle standalone, enhancement, and integration scenarios
- **Intelligent Placement**: Use smart analysis and inverse-prompting for optimal location
- **Concept Atomicity**: Address exactly one concept through practical building
- **Exercise Primacy**: Center on hands-on implementation producing tangible results
- **Applied Understanding**: Require working implementations that can be demonstrated and tested
- **Progressive Complexity**: Use only previously mastered concepts (where applicable)
- **Automatic Integration**: Update infrastructure files as needed

## AECS Violations to Prevent

**Prohibited Elements**:
- Multi-concept Exercises violating atomicity
- Theoretical content without immediate hands-on application
- Passive consumption elements in Exercise structure
- Academic assessment without practical building
- Learning objectives focused on "understanding" rather than "building"

**Required Elements**:
- Single-concept focus with immediate practical building
- Working implementations producing tangible results
- Hands-on tasks requiring practical implementation and testing
- Clear success criteria based on working implementations
- Proper integration with existing infrastructure

## Example Usage

**Standalone Exercise**:
```
Execute _framework/_prompts/create-exercise.md. CONCEPT is "[Independent Concept]". SCENARIO is "standalone".
```

**Enhancement Exercise**:
```
Execute _framework/_prompts/create-exercise.md. CONCEPT is "[Concept Clarification]". SCENARIO is "enhancement". TARGET_LOCATION is "[related stage]".
```

**Stage Integration**:
```
Execute _framework/_prompts/create-exercise.md. CONCEPT is "[New Concept]". SCENARIO is "integration". TARGET_LOCATION is "[target stage]".
```

**With Inverse-Prompting**:
```
Execute _framework/_prompts/create-exercise.md. CONCEPT is "[Any Concept]".
```
(Will use inverse-prompting to determine scenario and placement)

This will create AECS-compliant Exercises with intelligent placement, automatic integration, and support for all Exercise scenarios while maintaining the simplified directory structure and clear purpose separation.