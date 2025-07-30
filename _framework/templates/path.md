# {{SUBJECT}}: {{SUBTITLE}}

## Path Metadata

- **Subject**: {{SUBJECT}}
- **Target Audience**: {{TARGET_AUDIENCE}}
- **Difficulty Progression**: {{DIFFICULTY_PROGRESSION}}
- **Total Estimated Duration**: {{TOTAL_DURATION}}
- **Last Updated**: {{LAST_UPDATED}}

## Exercise-Driven Learning Path

{{PATH_DESCRIPTION}}

## Practical Mastery Objectives

By completing this Path through hands-on Exercise implementation, you will have built:

{{#each PRACTICAL_OBJECTIVES}}
- [ ] {{this}}
{{/each}}

## Prerequisites

**Required Practical Experience**:
{{#each PREREQUISITES}}
- {{this}}
{{/each}}

## Stage Progression

### {{#each STAGES}}
## Stage {{@index}}: {{title}}

**Exercise Focus**: {{exercise_focus}}

**Concepts Demonstrated Through Exercises**:
{{#each concepts}}
- {{this}}
{{/each}}

**Prerequisites**: {{prerequisites}}
**Estimated Duration**: {{duration}}
**Practical Outcomes**: {{practical_outcomes}}

---
{{/each}}

## Progressive Complexity Map

```mermaid
graph TD
    {{#each STAGE_DEPENDENCIES}}
    {{prerequisite}} --> {{dependent}}
    {{/each}}
```

## Implementation Verification

### Path Completion Criteria

Your Path mastery is demonstrated through:

{{#each MASTERY_EVIDENCE}}
- [ ] {{this}}
{{/each}}

### Portfolio of Built Solutions

Upon Path completion, you will have implemented:

{{#each PORTFOLIO_ITEMS}}
- **{{implementation}}**: {{description}}
{{/each}}

## Essential Building Tools

{{#each REQUIRED_TOOLS}}
- **{{name}}**: {{description}} - Required for practical implementation
{{/each}}

## Supporting Resources

### Implementation References
{{#each IMPLEMENTATION_REFERENCES}}
- [{{title}}]({{url}}) - {{description}}
{{/each}}

### Community Support
{{#each COMMUNITY_RESOURCES}}
- **{{name}}**: {{description}} - [{{url}}]({{url}})
{{/each}}

---

**Path Metadata**
- **Created**: {{CREATION_DATE}}
- **AECS Compliance**: Verified for Exercise Primacy, Concept Atomicity, Applied Understanding, and Progressive Complexity
- **Hierarchy**: Subject→Path→Stage→Concept→Exercise