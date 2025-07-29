# Generate Appendix Prompt

## Purpose
Create concise, quick-reference supplementary materials that serve as parallel instructional units within the AECS hierarchy. Appendices are ancillary to the Subject and exist parallel to Concept→Exercise relationships, providing flexible reference content without predetermined structure.

## Parameters
- **APPENDIX_TYPE**: Optional - The nature of the reference material (e.g., "Quick Reference", "Troubleshooting Guide", "Tool Setup", "Comparison Chart")
- **TARGET_NEED**: Optional - Specific learner need this appendix addresses
- **SCOPE**: Optional - What topics/concepts the appendix should cover
- **AUDIENCE_CONTEXT**: Optional - How learners will likely use this material
- **RELATED_STAGES**: Optional - Which curriculum Stages this appendix supports (as parallel reference material)

**Note**: All parameters are optional. If not specified, inverse-prompting will be used to gather required information with clear guidance on what's needed.

## Execution Instructions

Since appendices are inherently flexible, do not use a rigid template. Instead, create content that serves the specific need identified.

### Step 1: Parameter Completion Through Inverse-Prompting
If any parameters are missing, use inverse-prompting to gather required information:

**Missing Parameter Detection and Guidance**:

**If APPENDIX_TYPE not specified**:
```
"I need to understand what type of appendix to create:

**Common Appendix Types**:
- **Quick Reference**: Syntax summaries, command lists, function references
- **Troubleshooting Guide**: Common errors, diagnostic procedures, solutions
- **Tool Setup**: Installation guides, configuration instructions
- **Comparison Chart**: Technology comparisons, decision matrices
- **Deep Dive**: Advanced topics not covered in main exercises

Based on your project content, I suggest: [intelligent suggestions from content analysis]

What type of appendix would be most helpful?"
```

**If TARGET_NEED not specified**:
```
"I need to understand the specific learner need this appendix addresses:

**Analysis of Your Content**: [Scan existing exercises and identify potential gaps]
**Suggested Needs**: Based on your exercises, learners might benefit from:
- Quick syntax lookup during hands-on building
- Troubleshooting common implementation errors
- Reference for [domain-specific concepts]

What specific need should this appendix address?"
```

**If SCOPE not specified**:
```
"I need to determine the scope of topics to cover:

**Content Analysis**: Your exercises cover these concepts: [list from content scan]
**Scope Options**:
- **Focused**: Cover specific subset (recommended for quick reference)
- **Comprehensive**: Cover all related concepts (good for troubleshooting guides)

What scope would best serve the identified need?"
```

### Step 2: Need Analysis
Using gathered information, clarify:
1. **Primary Use Case**: When and why will learners reference this material?
2. **Format Requirements**: What structure best serves this use case? (reference table, troubleshooting flowchart, step-by-step guide, comparison matrix, etc.)
3. **Depth vs. Breadth**: Should this be comprehensive or focused on essentials?
4. **Integration Points**: How does this connect to curriculum exercises?

### Step 3: Structure Selection
Choose the most appropriate structure:
- **Quick Reference**: Tables, command lists, syntax summaries
- **Troubleshooting Guide**: Problem-solution pairs, diagnostic flowcharts
- **Setup Guide**: Step-by-step installation and configuration
- **Comparison**: Side-by-side analysis of tools, patterns, or approaches
- **Deep Dive**: Comprehensive exploration of a specialized topic
- **Cheat Sheet**: Essential information in condensed format

### Step 4: Content Development

**Opening Section:**
- Clear statement of purpose and intended use
- Target audience and prerequisites
- How to navigate and use the material effectively

**Core Content:**
Structure content according to the chosen format:
- Use clear headings and logical organization
- Include practical code examples where relevant
- Provide cross-references to related curriculum content
- Focus on actionable information over theory

**Quality Elements:**
- Accurate, tested information
- Clear, concise explanations
- Practical examples and use cases
- Current best practices and standards

### Step 5: Usability Optimization
Ensure the appendix is:
- **Scannable**: Easy to find specific information quickly
- **Actionable**: Provides concrete steps or examples
- **Current**: Uses up-to-date practices and tools
- **Complete**: Addresses the identified need comprehensively

### Step 6: Expert Review Integration
The generated appendix should be reviewed by:
- **Domain Expert**: Technical accuracy and completeness
- **Education Expert**: Usefulness for learners and integration with curriculum
- **Academic Editor**: Clarity, organization, and consistency

## Common Appendix Types

### Quick Reference
- Syntax summaries and command lists
- Function/method reference with examples
- Configuration options and parameters
- Keyboard shortcuts and tool commands

### Troubleshooting Guide
- Common error messages and solutions
- Diagnostic procedures and flowcharts
- Performance issues and optimization tips
- Environment-specific problems and fixes

### Setup and Configuration
- Installation procedures for tools and environments
- Configuration best practices
- Environment variable setup
- IDE and editor configuration

### Comparison and Decision Making
- Tool comparison matrices
- Pattern trade-off analysis
- Technology selection criteria
- Performance benchmarking results

### Deep Dive Topics
- Advanced concepts not covered in main curriculum
- Specialized use cases and implementations
- Historical context and evolution
- Research and experimental approaches

## Output Requirements

Generate appendix content that includes:
1. Clear purpose statement and usage instructions
2. Well-organized, scannable content structure
3. Practical examples and actionable information
4. Proper cross-references to curriculum content
5. Current, accurate technical information
6. Appropriate depth for the intended use case

## Success Criteria

The generated appendix should:
- Solve a specific, identifiable learner need
- Be immediately useful without requiring sequential reading
- Complement rather than duplicate curriculum content
- Follow current best practices and standards
- Be maintainable and updatable as technology evolves
- Integrate seamlessly with the overall learning experience

## Special Considerations

### Maintenance Strategy
- Consider how quickly the information will become outdated
- Design for easy updates and revisions
- Use authoritative sources and official documentation
- Include version information and last-updated dates

### Integration with AECS Hierarchy
- Serve as parallel instructional unit alongside Concept→Exercise relationships
- Reference specific Exercises where the appendix provides supplementary support
- Avoid duplicating content that belongs in Exercises demonstrating Concepts
- Support the Path's learning progression without disrupting Exercise Primacy
- Use consistent AECS terminology: Subject→Path→Stage→Concept→Exercise (+ Appendix parallel)

### Accessibility and Usability
- Design for quick reference rather than sequential reading
- Use clear visual hierarchy and formatting
- Include search-friendly headings and keywords
- Consider different learning styles and preferences

## Example Usage

```
Execute _framework/_prompts/generate-appendix.md. APPENDIX_TYPE is "Quick Reference". TARGET_NEED is "rapid syntax lookup". SCOPE is "essential syntax and patterns". AUDIENCE_CONTEXT is "reference during Exercises". RELATED_STAGES is "supports all Stages as foundational reference".
```

```
Execute _framework/_prompts/generate-appendix.md. APPENDIX_TYPE is "Troubleshooting Guide". TARGET_NEED is "common error resolution". SCOPE is "frequent issues and solutions". AUDIENCE_CONTEXT is "problem-solving support". RELATED_STAGES is "advanced implementation stages".
```

**With Inverse-Prompting (Recommended)**:
```
Execute _framework/_prompts/generate-appendix.md.
```
(Will use inverse-prompting to determine all parameters with intelligent suggestions)

This will generate focused, practical reference material tailored to the specific learner need identified.