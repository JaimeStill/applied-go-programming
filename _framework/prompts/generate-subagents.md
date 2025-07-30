# Generate Subagents Prompt

## Purpose
Generate standardized AECS-compliant subagents for a learning path project by processing subject-agnostic templates and creating specialized subagents for the specified subject domain.

## Parameters
- **SUBJECT**: The subject domain for the learning path (e.g., "Go Programming", "Python Development", "Data Science", "Machine Learning")

## Subagent Generation Requirements

**CRITICAL**: This prompt generates the three standardized AECS subagents required for all learning paths:

### 1. Technical Expert (`technical-expert.md`)
- **Role**: Domain-specific technical implementation validation
- **Specialization**: [SUBJECT] expertise with hands-on building focus
- **AECS Enforcement**: Exercise Primacy through practical implementation

### 2. Learning Designer (`learning-designer.md`) 
- **Role**: Exercise-driven learning progression validation
- **Specialization**: AECS pedagogy with subject-agnostic principles
- **AECS Enforcement**: Concept Atomicity and Progressive Complexity

### 3. Content Editor (`content-editor.md`)
- **Role**: AECS vocabulary and structure compliance
- **Specialization**: Technical writing with AECS framework adherence
- **AECS Enforcement**: Applied Understanding and proper hierarchy

## Template Processing

### Step 1: Template Source Identification
Process templates from:
- `_framework/templates/agents/technical-expert.md`
- `_framework/templates/agents/learning-designer.md`
- `_framework/templates/agents/content-editor.md`

### Step 2: Subject Parameterization
Replace all `[SUBJECT]` placeholders with the specified subject:
- Domain expertise references → [SUBJECT] expertise
- Technical examples → [SUBJECT]-specific implementations
- Best practices → [SUBJECT] conventions and patterns

### Step 3: Subagent Generation
Create specialized subagents in `.claude/agents/`:
- `.claude/agents/technical-expert.md`
- `.claude/agents/learning-designer.md`
- `.claude/agents/content-editor.md`

## AECS Compliance Requirements

### Subject Specialization Standards
- **Technical Expert**: Must demonstrate deep [SUBJECT] domain knowledge
- **Learning Designer**: Must maintain AECS principles regardless of subject
- **Content Editor**: Must enforce AECS vocabulary with [SUBJECT] context

### Framework Integration
- **Standardized Names**: All learning paths use identical subagent names
- **Predictable References**: Framework can safely reference @technical-expert, @learning-designer, @content-editor
- **Template Flexibility**: Templates adapt to any subject through parameterization

## Generation Output

### Success Criteria
Generated subagents must:
- [ ] Maintain all AECS principle enforcement from templates
- [ ] Specialize appropriately for the [SUBJECT] domain
- [ ] Use standardized names enabling framework references
- [ ] Include proper YAML frontmatter with tools and descriptions
- [ ] Eliminate recursive spawning through direct reference capability

### Validation Requirements
After generation, verify:
- [ ] Direct subagent references work: @technical-expert, @learning-designer, @content-editor
- [ ] No nested subagent spawning occurs during engagement
- [ ] Subject-specific expertise is properly reflected
- [ ] AECS principle enforcement remains intact

## Usage Instructions

Execute this prompt with:
```
Execute _framework/prompts/generate-subagents.md. SUBJECT is "[Domain Name]".
```

Examples:
- `SUBJECT is "Go Programming"`
- `SUBJECT is "Python Development"`  
- `SUBJECT is "Machine Learning"`
- `SUBJECT is "Data Science"`

## Framework Integration

This prompt is designed to be called automatically during path creation via `_framework/prompts/create-path.md` to ensure all learning paths have properly configured, standardized subagents that eliminate the nested spawning problem while maintaining subject-specific expertise.

---

**Prompt Metadata**
- **Created**: 2025-07-30
- **Purpose**: Eliminate nested subagent spawning through standardized template processing
- **AECS Compliance**: Enforces all four AECS principles through specialized subagent generation