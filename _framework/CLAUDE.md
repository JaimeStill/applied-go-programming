# AECS Framework Implementation Guide

@_framework/README.md for complete methodology and implementation details.

## Core AECS Principles

**Enforce all four principles in every interaction:**

1. **Exercise Primacy**: Learning centers on hands-on building, never passive consumption
2. **Concept Atomicity**: Each Exercise addresses exactly one Concept through practical implementation
3. **Applied Understanding**: Learning occurs through building working examples with tangible results
4. **Progressive Complexity**: Concepts build incrementally through mastered dependency chains

## Content Structure

**Required Hierarchy**: Subject → Path → Stage → Concept → Exercise (+ Appendix parallel)

## Subagent Engagement Protocol

**CRITICAL**: Use Claude's built-in subagent feature by referencing subagents defined in `.claude/agents/`. Do NOT use the Task tool to simulate subagents.

**Discovery Phase**: Examine `.claude/agents/` directory to identify available project subagents and match to required expertise roles.

### Subagent Review Framework

**MANDATORY CONTEXT PROVISION**: Every subagent engagement MUST include:

1. **Artifact Type Declaration**: "You are reviewing a [STAGE|EXERCISE|PATH|APPENDIX] document"
2. **Format Constraints**: Specify exact format requirements and abstraction level
3. **Review Scope Boundaries**: Define what changes are/aren't within scope
4. **Output Format Requirements**: Specify expected feedback format (not artifact generation)

### Scoped Review Patterns

**Technical Review Template**:
```
"You are reviewing a [ARTIFACT_TYPE] document. This artifact should maintain [ABSTRACTION_LEVEL] without [PROHIBITED_ELEMENTS].

Review scope: Identify violations where [SPECIFIC_CRITERIA]. 

Your feedback should:
- List specific violations with line references
- Suggest conceptual fixes that maintain document format
- NOT include implementation code or detailed examples
- Focus only on [DOMAIN_EXPERTISE] aspects

Format your response as: Issue → Conceptual Fix (no code)"
```

**Learning Review Template**:
```
"You are reviewing a [ARTIFACT_TYPE] document that should enforce Exercise Primacy at [ABSTRACTION_LEVEL].

Review scope: Identify passive consumption elements and multi-concept violations.

Your feedback should:
- Identify specific AECS principle violations
- Suggest structural reorganization approaches
- NOT provide detailed content rewrites
- Focus on learning design principles

Format your response as: Violation → Structural Fix (no content)"
```

**Structure Review Template**:
```
"You are reviewing a [ARTIFACT_TYPE] document for AECS vocabulary and hierarchy compliance.

Review scope: AECS vocabulary usage and Subject→Path→Stage→Concept→Exercise hierarchy.

Your feedback should:
- Identify vocabulary deviations with specific examples
- Suggest hierarchy corrections
- NOT rewrite content sections
- Focus only on structural compliance

Format your response as: Deviation → Vocabulary/Structure Fix (no rewrites)"
```

### Feedback Validation Protocol

**MANDATORY**: After receiving subagent feedback, validate:

1. **Scope Compliance**: Feedback addresses only assigned review domain
2. **Format Adherence**: No artifact generation or detailed content provision
3. **Abstraction Maintenance**: Suggestions maintain appropriate document level
4. **AECS Alignment**: Recommendations align with framework principles

**Rejection Criteria**: Reject and re-prompt subagent if feedback includes:
- Implementation code or detailed examples in stage/path documents
- Content rewrites instead of conceptual guidance
- Out-of-scope suggestions beyond assigned expertise
- Artifact generation rather than review feedback

### Subagent Response Format

**Required Response Structure**:
```
## [DOMAIN] Review: [ARTIFACT_TYPE]

### Compliance Assessment
- [X] AECS Principle adherence
- [X] Format appropriateness 
- [X] Scope boundaries

### Identified Issues
1. **Issue**: [Specific violation]
   **Fix**: [Conceptual correction approach]

2. **Issue**: [Specific violation]  
   **Fix**: [Conceptual correction approach]

### Summary
- Total issues: [N]
- Critical fixes needed: [N]
- Format compliance: [PASS/FAIL with explanation]
```

**Important**: Subagents provide structured review feedback only. The main agent creates content and incorporates subagent feedback while maintaining artifact format integrity.

## Artifact Format Constraints

### Stage Document Format Requirements

**Abstraction Level**: High-level exercise descriptions without implementation code
**Required Sections**: Stage Metadata, Exercise Sequence, Progressive Complexity Validation
**Prohibited Elements**: 
- Detailed code examples or implementations
- Step-by-step coding instructions
- Specific syntax or API usage details
- File content or directory structures

**Valid Stage Content**: "Build a calculator using functions" ✓
**Invalid Stage Content**: "Type this code: `func add(a, b int) int { return a + b }`" ✗

### Exercise Document Format Requirements

**Abstraction Level**: Detailed implementation guidance with complete code examples
**Required Sections**: Exercise metadata, building steps, code examples, verification
**Prohibited Elements**:
- Theoretical explanations without practical implementation
- Multi-concept mixing within single exercise
- Passive consumption elements

### Path Document Format Requirements

**Abstraction Level**: Stage progression overview with learning outcomes
**Required Sections**: Path metadata, stage progression, prerequisites, outcomes
**Prohibited Elements**:
- Individual exercise details
- Implementation specifics
- Code examples or syntax

### Appendix Document Format Requirements

**Abstraction Level**: Reference material supporting hands-on implementation
**Required Sections**: Practical examples, quick reference, troubleshooting
**Prohibited Elements**:
- Theoretical concepts without practical context
- Tutorial-style progressive instruction

## Format Validation Checklist

Before finalizing any artifact, verify:

**Stage Documents**:
- [ ] No implementation code present
- [ ] Exercise descriptions are conceptual, not instructional
- [ ] Focus on what will be built, not how to build it
- [ ] Appropriate abstraction level maintained throughout

**Exercise Documents**:
- [ ] Complete implementation guidance provided
- [ ] Working code examples included
- [ ] Single concept focus maintained
- [ ] Hands-on building required throughout

**Path Documents**:
- [ ] Stage-level progression only
- [ ] No exercise-level details
- [ ] Learning outcomes clearly defined
- [ ] Prerequisites properly established

**Appendix Documents**:
- [ ] Practical reference material only
- [ ] No progressive tutorial elements
- [ ] Quick-access format maintained
- [ ] Implementation-focused content

## AECS Violations to Fix Immediately

- Theoretical explanations without hands-on building → Convert to practical implementation exercises
- Multi-concept Exercises → Split into atomic, single-concept implementations  
- Abstract examples → Replace with concrete, working implementations producing observable results
- Dependency leaps → Create proper incremental progression using only previously mastered concepts

## Subagent Prompt Enforcement

**MANDATORY PROMPT STRUCTURE**: All subagent engagements must follow this exact pattern:

```
"You are reviewing a [ARTIFACT_TYPE] document. This [ARTIFACT_TYPE] should maintain [ABSTRACTION_LEVEL] and must not include [PROHIBITED_ELEMENTS].

[ARTIFACT_CONTENT_TO_REVIEW]

Review scope: [SPECIFIC_REVIEW_CRITERIA]

Your feedback must:
- [SPECIFIC_REQUIREMENTS]
- [FORMAT_CONSTRAINTS] 
- [DOMAIN_FOCUS]

Format your response as: [REQUIRED_FORMAT]"
```

**Context Validation**: Before sending to subagent, verify prompt includes:
1. ✓ Artifact type clearly stated
2. ✓ Format constraints explicitly listed
3. ✓ Review scope boundaries defined
4. ✓ Output format requirements specified
5. ✓ Prohibited elements clearly identified

**Response Validation**: After receiving subagent feedback, verify:
1. ✓ No artifact content generation (code, detailed examples, etc.)
2. ✓ Feedback addresses assigned domain only
3. ✓ Suggestions maintain document abstraction level
4. ✓ Required response format followed
5. ✓ No scope creep beyond assigned expertise

**Rejection Protocol**: If subagent response violates format/scope:
1. **Immediately reject** the response
2. **Re-prompt** with enhanced constraints
3. **Explicitly state** what was inappropriate in previous response
4. **Require confirmation** of understanding before proceeding

## Quality Validation

Before completion, verify: Exercise Primacy (hands-on building), Concept Atomicity (single concepts), Applied Understanding (tangible results), Progressive Complexity (proper dependencies), working implementations, and **artifact format integrity maintained throughout subagent review process**.

## Framework Usage

Execute prompts using: `Execute _framework/_prompts/[prompt-name].md. PARAMETER is "[value]".`

- **create-**: Initialize from scratch without infrastructure assumptions
- **generate-**: Derive from existing project infrastructure