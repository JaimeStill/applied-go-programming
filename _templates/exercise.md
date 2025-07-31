# Exercise: Demonstrating {{CONCEPT}}

## Exercise Metadata

- **Stage**: {{STAGE_NAME}}
- **Single Concept**: {{CONCEPT}}
- **Prerequisites**: {{PREREQUISITE_CONCEPTS}}
- **Estimated Time**: {{DURATION}}

## What You'll Build

{{PRACTICAL_OUTCOME_DESCRIPTION}}

**Tangible Result**: {{OBSERVABLE_RESULT}}

## Conceptual Overview

{{CONCEPTUAL_CONTEXT_1_TO_5_PARAGRAPHS}}

## Foundation Requirements

This Exercise demonstrates {{CONCEPT}} which assumes familiarity with:
{{#each PREREQUISITE_CONCEPTS}}
- {{this}}
{{/each}}

**Starting Point**: This exercise provides all necessary code from scratch. No prior exercise code is required.

## Build {{PRACTICAL_OUTCOME_DESCRIPTION}}

### Initial Step: {{INITIAL_STEP_TITLE}}

{{INITIAL_SETUP_DESCRIPTION}}

Create a new Go program with the following structure:

```go
{{GO_PACKAGE_DECLARATION}}

{{GO_IMPORTS}}

{{INITIAL_GO_IMPLEMENTATION}}
```

{{INITIAL_SETUP_EXPLANATION}}

{{#each STEPS}}
### Step {{@index}}: {{STEP_TITLE}}

{{STEP_DESCRIPTION}}

Add the following Go implementation:

```go
// {{GO_CONTEXT_PRESERVATION_COMMENT}}

{{STEP_GO_IMPLEMENTATION}}
```

{{STEP_EXPLANATION}}

{{/each}}

### Final Step: Execute and Verify

Validate your complete Go implementation:

**Compile and run**:
```bash
go run {{GO_MAIN_FILE}}
```

**Expected output**:
```
{{EXPECTED_GO_OUTPUT}}
```

**Verification steps**:
{{#each GO_VERIFICATION_STEPS}}
1. {{this}}
{{/each}}

{{VERIFICATION_EXPLANATION_AND_GO_CONCEPT_DEMONSTRATION}}

## Common Go Implementation Issues

### Issue: {{COMMON_GO_ISSUE_1}}

**Problem**: {{GO_ISSUE_1_DESCRIPTION}}

**Solution**: 
```go
{{GO_ISSUE_1_SOLUTION_CODE}}
```

**Go Best Practice**: {{GO_ISSUE_1_BEST_PRACTICE}}

### Issue: {{COMMON_GO_ISSUE_2}}

**Problem**: {{GO_ISSUE_2_DESCRIPTION}}

**Solution**: 
```go
{{GO_ISSUE_2_SOLUTION_CODE}}
```

**Go Best Practice**: {{GO_ISSUE_2_BEST_PRACTICE}}

### Issue: Compilation Errors

**Problem**: Common Go compilation issues with this concept

**Solution**: 
```go
{{GO_COMPILATION_FIX_CODE}}
```

**Go Best Practice**: {{GO_COMPILATION_BEST_PRACTICE}}

## Go-Specific Implementation Notes

### Incremental Step Focus

This Exercise follows Go-specific incremental step patterns:

**Critical Context Preserved**:
- Package declaration: `{{GO_PACKAGE_DECLARATION}}`
- Import statements: `{{GO_IMPORTS}}`
- Function signatures: `{{GO_FUNCTION_SIGNATURES}}`
- {{GO_CRITICAL_CONTEXT_ITEMS}}

**Transient Details Summarized**:
- {{GO_TRANSIENT_DETAILS_COMMENTS}}

### Go Idioms and Conventions

**Naming Conventions**:
- {{GO_NAMING_CONVENTIONS}}

**Error Handling**:
```go
{{GO_ERROR_HANDLING_PATTERN}}
```

**Resource Management**:
```go
{{GO_RESOURCE_MANAGEMENT_PATTERN}}
```

## Independent Challenge

{{GO_CHALLENGE_DESCRIPTION}}

**Requirements**:
{{GO_CHALLENGE_REQUIREMENTS}}

### Independent Challenge Solution

*Try building the solution yourself before looking at this reference.*

<details>
<summary>{{GO_CHALLENGE_SUMMARY_TITLE}}</summary>
{{GO_CHALLENGE_SOLUTION}}

{{GO_CHALLENGE_KEY_IMPLEMENTATION_NOTES}}
</details>

## Reference Implementation

*Try building the Go solution yourself before looking at this reference.*

<details>
<summary>Complete Go Reference Implementation</summary>

**main.go**:
```go
{{GO_PACKAGE_DECLARATION}}

{{GO_FULL_IMPORTS}}

{{GO_REFERENCE_IMPLEMENTATION}}
```

**Key Go Implementation Notes**:
{{#each GO_IMPLEMENTATION_NOTES}}
- {{this}}
{{/each}}

**Go Performance Considerations**:
{{#each GO_PERFORMANCE_NOTES}}
- {{this}}
{{/each}}

**Memory Management**:
- {{GO_MEMORY_MANAGEMENT_NOTES}}

</details>

## Path Integration

**Go Concept Demonstrated**: {{CONCEPT}}
**Next Go Concept in Path**: {{NEXT_CONCEPT}}
**Go-Specific Dependencies**: {{GO_CONCEPT_DEPENDENCIES}}

---

**Exercise Metadata**
- **Created**: {{CREATION_DATE}}
- **AECS Compliance**: Verified for Exercise Primacy, Concept Atomicity, Applied Understanding, and Progressive Complexity
- **Go Version Compatibility**: {{GO_VERSION_COMPATIBILITY}}
- **Implementation Status**: {{IMPLEMENTATION_STATUS}}