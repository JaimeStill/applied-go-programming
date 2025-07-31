# {{STAGE_TITLE}}: Exercise Progression

<!-- This template generates Stage README.md files for GitHub-standard directory navigation -->

## Stage Metadata

- **Stage Number**: {{STAGE_NUMBER}}
- **Difficulty Level**: {{DIFFICULTY_LEVEL}}
- **Total Exercises**: {{TOTAL_EXERCISES}}
- **Estimated Duration**: {{TOTAL_DURATION}}
- **Prerequisites**: {{STAGE_DEPENDENCIES}}

## Exercise Sequence

This Stage contains the following Exercise progression for related Concepts:

{{#each EXERCISES}}
### {{@index}}. Exercise: Demonstrating {{concept}}

**Practical Implementation**: {{implementation_description}}

**Prerequisites**: {{prerequisites}}
**Estimated Time**: {{duration}}
**Tangible Result**: {{observable_outcome}}

**Building Task**: {{building_task}}

---
{{/each}}

## Progressive Complexity Validation

### Dependency Chain

```mermaid
graph LR
    {{#each EXERCISE_DEPENDENCIES}}
    {{prerequisite}} --> {{current}}
    {{/each}}
```

### Concept Mastery Progression

After completing all Exercises in this Stage, you will have mastered:

{{#each MASTERED_CONCEPTS}}
- [ ] {{this}} (through hands-on implementation)
{{/each}}

## Implementation Verification

### Practical Validation

To verify Stage completion, your implementations should:

{{#each VERIFICATION_CRITERIA}}
- [ ] {{this}}
{{/each}}

### Next Stage Preparation

**Progressive Complexity**: Completing this Stage's Exercises (which demonstrate mastered Concepts) enables these Concepts to be demonstrated in the next Stage:

{{#each ENABLED_CONCEPTS}}
- **{{concept}}**: {{description}}
{{/each}}

**Next Stage**: {{NEXT_STAGE}}

---

**Stage Metadata**
- **Created**: {{CREATION_DATE}}
- **AECS Compliance**: Verified for Exercise Primacy, Concept Atomicity, Applied Understanding, and Progressive Complexity
- **Hierarchy**: Subject→Path→Stage→Concept→Exercise