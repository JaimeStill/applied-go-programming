# Generate Appendix Prompt

## Purpose
Create circumstantial reference materials that support the main curriculum without following a rigid template structure. Appendices should be practical, focused, and immediately useful to learners.

## Parameters
- **APPENDIX_TYPE**: The nature of the reference material (e.g., "Quick Reference", "Troubleshooting Guide", "Tool Setup", "Comparison Chart")
- **TARGET_NEED**: Specific learner need this appendix addresses
- **SCOPE**: What topics/concepts the appendix should cover
- **AUDIENCE_CONTEXT**: How learners will likely use this material
- **RELATED_CHAPTERS**: Which curriculum chapters this appendix supports

## Execution Instructions

Since appendices are inherently flexible, do not use a rigid template. Instead, create content that serves the specific need identified.

### Step 1: Need Analysis
Before generating content, clarify:
1. **Primary Use Case**: When and why will learners reference this material?
2. **Format Requirements**: What structure best serves this use case? (reference table, troubleshooting flowchart, step-by-step guide, comparison matrix, etc.)
3. **Depth vs. Breadth**: Should this be comprehensive or focused on essentials?
4. **Integration Points**: How does this connect to curriculum exercises?

### Step 2: Structure Selection
Choose the most appropriate structure:
- **Quick Reference**: Tables, command lists, syntax summaries
- **Troubleshooting Guide**: Problem-solution pairs, diagnostic flowcharts
- **Setup Guide**: Step-by-step installation and configuration
- **Comparison**: Side-by-side analysis of tools, patterns, or approaches
- **Deep Dive**: Comprehensive exploration of a specialized topic
- **Cheat Sheet**: Essential information in condensed format

### Step 3: Content Development

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

### Step 4: Usability Optimization
Ensure the appendix is:
- **Scannable**: Easy to find specific information quickly
- **Actionable**: Provides concrete steps or examples
- **Current**: Uses up-to-date practices and tools
- **Complete**: Addresses the identified need comprehensively

### Step 5: Expert Review Integration
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

### Integration with Curriculum
- Reference specific exercises where the appendix is relevant
- Avoid duplicating content that belongs in exercises
- Support the curriculum's learning progression
- Use consistent terminology and conventions

### Accessibility and Usability
- Design for quick reference rather than sequential reading
- Use clear visual hierarchy and formatting
- Include search-friendly headings and keywords
- Consider different learning styles and preferences

## Example Usage

```
Execute generate-appendix.md. APPENDIX_TYPE is "Go Quick Reference for Experienced Programmers". TARGET_NEED is "rapid syntax lookup for developers coming from other languages". SCOPE is "essential Go syntax, patterns, and idioms". AUDIENCE_CONTEXT is "reference during exercises and projects". RELATED_CHAPTERS is "supports all chapters as foundational reference".
```

```
Execute generate-appendix.md. APPENDIX_TYPE is "Debugging Concurrent Go Programs". TARGET_NEED is "troubleshooting race conditions and deadlocks". SCOPE is "common concurrency issues, diagnostic tools, debugging strategies". AUDIENCE_CONTEXT is "problem-solving support during advanced exercises". RELATED_CHAPTERS is "Chapters covering concurrency topics".
```

This will generate focused, practical reference material tailored to the specific learner need identified.