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

**Discovery Phase**: Examine `.claude/agents/` directory to identify available project subagents and match to required expertise roles.

**Delegation Patterns**:
- **Technical Expert**: "Request the technical implementation subagent to implement working [subject] examples that demonstrate [concept] through hands-on building"
- **Learning Designer**: "Ask the pedagogical subagent to restructure content for Exercise Primacy, ensuring immediate hands-on engagement rather than passive consumption"  
- **Content Editor**: "Have the academic editing subagent fix AECS vocabulary violations and verify Subject→Path→Stage→Concept→Exercise hierarchy"

## AECS Violations to Fix Immediately

- Theoretical explanations without hands-on building → Convert to practical implementation exercises
- Multi-concept Exercises → Split into atomic, single-concept implementations  
- Abstract examples → Replace with concrete, working implementations producing observable results
- Dependency leaps → Create proper incremental progression using only previously mastered concepts

## Quality Validation

Before completion, verify: Exercise Primacy (hands-on building), Concept Atomicity (single concepts), Applied Understanding (tangible results), Progressive Complexity (proper dependencies), and working implementations.

## Framework Usage

Execute prompts using: `Execute _framework/_prompts/[prompt-name].md. PARAMETER is "[value]".`

- **create-**: Initialize from scratch without infrastructure assumptions
- **generate-**: Derive from existing project infrastructure