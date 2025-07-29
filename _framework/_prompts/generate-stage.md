# Generate Stage Prompt

## Purpose
Create a detailed Stage with Exercise progression for any learning subject, following the Agentic Exercise Curriculum System (AECS) methodology with strict enforcement of all four AECS principles.

## Parameters
<parameters>
<stage_title>The specific Stage name from the Path</stage_title>
<target_concepts>Abstract description of the concepts this Stage should cover (e.g., "concurrent programming patterns", "color theory fundamentals", "knife skills basics")</target_concepts>
</parameters>

## AECS Compliance Requirements

**CRITICAL**: This prompt MUST enforce all four AECS principles:

### Exercise Primacy Enforcement
- Stage MUST center on hands-on building through Exercise sequences
- NO theoretical overviews, explanations, or passive consumption elements
- ALL learning MUST occur through immediate practical implementation
- Every section MUST require active building rather than reading

### Concept Atomicity Validation
- Each Exercise MUST address exactly one Concept in its simplest useful form
- Stage MUST contain atomic concept implementations only
- NO multi-concept Exercises that violate atomicity
- Each Concept MUST be implementable as single Exercise

### Applied Understanding Requirements
- All Concepts MUST be demonstrated through practical implementation producing tangible results
- NO theoretical content without immediate hands-on application
- Every Exercise MUST require building working examples
- ALL content MUST produce observable, testable outcomes

### Progressive Complexity Verification
- Concepts MUST build incrementally within Stage through proper dependency chains
- Each Exercise MUST depend only on previously mastered Concepts through practical building
- Stage MUST prepare for next Stage through hands-on mastery progression
- NO learning jumps that skip prerequisite practical implementation

## Execution Instructions

Use the stage template at `_framework/_templates/stage.md` to generate a comprehensive Stage with Exercise sequences.

### Step 1: PATH.md Analysis and Context Extraction
Before generating content, analyze existing infrastructure:
1. **Read PATH.md**: Extract stage context, numbering, and overall progression
2. **Analyze Stage Position**: Determine prerequisites from previous stages and preparation for next stages
3. **Extract Dependencies**: Identify concepts mastered through previous Stage Exercises
4. **Validate TARGET_CONCEPTS**: Use inverse-prompting to ensure comprehensive concept coverage
5. **Progressive Logic**: Ensure this Stage fits properly in the overall Path progression

### Step 2: AECS-Compliant Concept Analysis
Transform abstract TARGET_CONCEPTS into atomic implementations:
1. **Atomic Concept Breakdown**: What individual concepts can be mastered through single Exercises producing tangible results?
2. **Hands-On Implementation Mapping**: How can each concept be demonstrated through practical building?
3. **Dependency Chain Validation**: What's the proper progression for building each concept from previously mastered ones?
4. **Practical Building Requirements**: What working examples demonstrate each concept through implementation?
5. **Tangible Results Identification**: What observable outcomes prove concept mastery through building?

### Step 3: Exercise-First Sequence Design
Create comprehensive Exercise sequence covering all atomic concepts:
- **Foundation Exercises**: Implement the most basic concepts producing immediate results
- **Building Exercises**: Add complexity through hands-on implementation building on previous work
- **Integration Exercises**: Combine previously built concepts through practical application
- **Applied Exercises**: Demonstrate real-world usage through working implementations

Each Exercise MUST:
- Address exactly one atomic concept
- Require immediate hands-on building
- Produce tangible, testable results
- Build only on previously mastered concepts through practical implementation

### Step 4: AECS-Compliant Content Generation
Fill the template with Exercise-driven content:

**Exercise Sequence:**
For each Exercise, provide:
- **Concept**: Single, atomic concept addressed through practical building
- **Implementation Description**: What learners will BUILD (not learn theoretically)
- **Prerequisites**: Previously mastered concepts through hands-on Exercises
- **Tangible Result**: Observable outcome produced through building
- **Building Task**: Specific hands-on implementation requirements

**Progressive Complexity Validation:**
- **Dependency Chain**: Clear progression showing each Exercise builds on previous practical mastery
- **Concept Mastery Progression**: List of concepts mastered through hands-on implementation
- **Implementation Verification**: Practical validation criteria requiring working examples

### Step 5: AECS Principle Verification
Ensure the Stage:
- **Exercise Primacy**: Every Exercise centers on hands-on building
- **Concept Atomicity**: Each Exercise addresses single concepts only
- **Applied Understanding**: All learning through practical implementation producing results
- **Progressive Complexity**: Proper dependency chains through mastered concept building

### Step 6: MANDATORY AECS-Constrained Expert Review Integration

**REQUIRED**: Identify and engage specific project subagents during Stage generation to ensure AECS compliance:

### Subagent Discovery and Engagement

**Step 6.1**: Examine `.claude/agents/` directory to identify available project subagents and their capabilities.

**Step 6.2**: Match subagents to required expertise roles and engage using natural language delegation:

**Technical Implementation Expert**: Identify and request the technical/implementation subagent to "implement working [subject] examples that demonstrate each concept through hands-on building. Ensure every Exercise produces tangible, testable results using only previously mastered concepts. Fix any violations where learners consume rather than build."

**Learning Design Expert**: Ask the pedagogical/educational subagent to "restructure this Stage content to enforce Exercise Primacy by ensuring learners immediately engage in hands-on [subject] practice rather than theoretical study. Verify single-concept atomicity per Exercise and eliminate any passive consumption elements."

**Content Structure Expert**: Have the academic editing/writing subagent "fix any deviations from AECS vocabulary and structure. Verify that content follows Subject→Path→Stage→Concept→Exercise hierarchy and maintains consistent AECS terminology throughout, ensuring all sections center on doing rather than consuming."

## Output Requirements

Generate a complete Stage document that includes:
1. AECS-compliant Stage structure using proper vocabulary
2. Exercise sequences requiring hands-on building for atomic concepts
3. Progressive complexity through dependency chains of mastered concepts
4. Practical validation through working implementation verification
5. Next Stage preparation through hands-on mastery
6. Proper AECS metadata and hierarchy compliance

## Success Criteria

<success_criteria>
The generated Stage MUST demonstrate:
- **Exercise Primacy**: Every Exercise centers on hands-on building with immediate practical engagement
- **Concept Atomicity**: Each Exercise addresses exactly one concept through focused implementation  
- **Applied Understanding**: All learning occurs through building working examples with tangible, testable results
- **Progressive Complexity**: Clear dependency chains where each Exercise builds only on previously mastered concepts
- **AECS Vocabulary**: Consistent use of Subject→Path→Stage→Concept→Exercise hierarchy
- **Implementation Focus**: Working examples that can be built, run, and verified rather than theoretical explanations
</success_criteria>

## AECS Violations to Prevent

**Prohibited Elements:**
- Theoretical overviews, explanations, or conceptual sections without immediate hands-on building
- Multi-concept Exercises violating atomicity
- Passive consumption elements (readings, summaries, explanations)
- Academic assessment patterns (self-assessment, mastery checkpoints without practical building)
- Learning objectives focused on "understanding" rather than "building"

**Required Elements:**
- Exercise sequences requiring immediate practical implementation
- Single-concept focus per Exercise producing tangible results
- Working implementations that can be run, tested, and verified
- Progressive dependency on previously built and mastered concepts
- Hands-on validation through practical implementation verification

## File Naming Convention

Generated Stage documents should be saved as:
- **File Path**: `stages/##-[stage-name]/stage.md`
- **Example**: `stages/01-fundamentals/stage.md`

## Example Usage

<example>
**Subject**: Digital Photography
**Input**: Execute _framework/_prompts/generate-stage.md. STAGE_TITLE is "Camera Fundamentals". TARGET_CONCEPTS is "aperture control, shutter speed mastery, ISO sensitivity".

**Expected Output**: Stage document with atomic Exercises like:
- Exercise 1: Demonstrate aperture effects through hands-on shooting
- Exercise 2: Build shutter speed control through motion capture practice  
- Exercise 3: Implement ISO adjustment through lighting experiments

Each Exercise produces tangible photo results demonstrating single concepts.
</example>

```
Execute _framework/_prompts/generate-stage.md. STAGE_TITLE is "[Stage Topic]". TARGET_CONCEPTS is "[abstract concept description]".
```