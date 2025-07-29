# Agentic Exercise Curriculum System (AECS)

## Table of Contents

1. [Overview](#overview)
2. [Philosophy](#philosophy)
3. [System Architecture](#system-architecture)
4. [File Naming Conventions](#file-naming-conventions)
5. [Usage Instructions](#usage-instructions)
6. [Quality Standards](#quality-standards)
7. [Exercise Format Standard](#exercise-format-standard)
8. [Scalability Guide](#scalability-guide)
9. [Best Practices](#best-practices)
10. [Success Metrics](#success-metrics)
11. [Future Enhancements](#future-enhancements)
12. [Glossary](#glossary)

## Overview

The Agentic Exercise Curriculum System (AECS) is a methodology for creating world-class, exercise-driven technical curricula through subagent collaboration. AECS combines standardized templates, dynamic generation prompts, and specialized subagents to produce consistent, high-quality learning materials at scale.

## Philosophy

AECS is built on three core principles:

1. **Exercise-Driven Learning**: Knowledge is best retained through hands-on practice. Every concept is immediately reinforced with practical exercises.

2. **Subagent Collaboration**: Specialized subagents bring domain expertise to every piece of content, ensuring technical accuracy, pedagogical effectiveness, and professional presentation.

3. **Systematic Scalability**: Standardized templates and prompts enable rapid replication across different topics, programming languages, and technical domains.

## System Architecture

### Core Components

```
Repository Structure:
├── WORKFLOW.md          # This documentation
├── curriculum.md        # Main curriculum document
├── _templates/          # Standardized content templates
├── _prompts/           # Generation prompts ([action]-[noun].md format)
├── .claude/agents/     # Specialized AI agent profiles
└── chapters/          # Generated curriculum content
    ├── 00-foundations/    # Chapter 0 materials
    ├── 01-chapter-name/   # Chapter materials with roadmap.md
    │   ├── roadmap.md     # Chapter roadmap
    │   └── 01-exercise.md # Individual exercises
    └── a01-appendix.md    # Appendices as standalone files
```

### Template System

Templates define the structure and quality standards for each content type:

- **curriculum.md**: Overall course structure with learning progression
- **chapter.md**: Chapter roadmaps with exercise sequences
- **exercise.md**: Comprehensive exercise format with metadata
- **appendix.md**: Reference material and quick guides
- **subagent.md**: AI agent persona definitions

### Prompt System

Generation prompts use the `[action]-[noun].md` format for command-style execution:

- **generate-curriculum.md**: Creates complete curriculum structure
- **generate-chapter.md**: Builds chapter roadmaps from curriculum context
- **generate-exercise.md**: Generates individual exercises with full content
- **generate-appendix.md**: Creates reference materials
- **generate-subagent.md**: Defines specialized subagent personas
- **review-content.md**: Conducts comprehensive content review using subagents

### Subagent Collaboration

Specialized subagents ensure quality across different dimensions:

- **Technical Expert** (e.g., go-engineer): Code accuracy, performance, best practices
- **Education Expert** (cs-professor): Learning progression, concept scaffolding
- **Editorial Expert** (academic-editor): Consistency, clarity, professional presentation

## File Naming Conventions

### Directory Structure
- **Chapters**: `chapters/##-[chapter-name]/` (e.g., `chapters/01-concurrency/`)
- **Chapter Roadmaps**: `chapters/##-[chapter-name]/roadmap.md`
- **Exercises**: `chapters/##-[chapter-name]/##-[exercise-name].md`
- **Appendices**: `chapters/a##-[appendix-name].md` (e.g., `chapters/a01-go-idioms.md`)

## Usage Instructions

### Basic Command Structure

Execute prompts using this format:
```
Execute [prompt-name].md. [PARAMETER] is [value].
```

### Examples

**Generate a new curriculum:**
```
Execute generate-curriculum.md. TOPIC is "Linux Administration for Go Programmers".
```

**Create a chapter roadmap:**
```
Execute generate-chapter.md. CHAPTER is "Network Programming" from the Linux Administration curriculum.
```

**Generate an exercise:**
```
Execute generate-exercise.md. TOPIC is "TCP Server Implementation" from the Network Programming chapter.
```

### Parameter Injection

Prompts support dynamic parameter injection:
- `TOPIC`: Main subject matter
- `CHAPTER`: Specific chapter context
- `DIFFICULTY`: Beginner/Intermediate/Advanced
- `PREREQUISITES`: Required prior knowledge
- `DURATION`: Estimated completion time

## Quality Standards

### Built-in Review Process

Every generated piece of content goes through automatic quality assurance:

1. **Template Compliance**: Ensures all required sections are present
2. **Agent Review**: Specialized agents validate their domain of expertise
3. **Consistency Check**: Academic editor ensures style and format standards
4. **Integration Test**: Content integrates properly with existing materials

### Quality Benchmarks

- **Technical Accuracy**: Code compiles, runs, and follows best practices
- **Pedagogical Effectiveness**: Clear learning objectives and progressive complexity
- **Professional Presentation**: Consistent formatting, clear writing, proper documentation
- **Practical Relevance**: Real-world applicability and production-ready examples

## Exercise Format Standard

AECS exercises follow a comprehensive structure:

### Required Sections

1. **Metadata**
   - Prerequisites
   - Difficulty level
   - Estimated time
   - Concepts covered
   - Related exercises

2. **Learning Objectives**
   - Clear, measurable outcomes
   - Checkbox format for self-assessment

3. **Overview**
   - Conceptual foundation
   - Why this topic matters
   - Real-world context

4. **Code Examples**
   - Production-ready implementations
   - Proper error handling
   - Idiomatic language patterns
   - Well-commented explanations

5. **Understanding**
   - Mental models
   - How to think about the concept
   - Common misconceptions

6. **Exercise**
   - Hands-on implementation task
   - Multiple difficulty levels
   - Clear success criteria
   - Step-by-step guidance

7. **Common Mistakes**
   - Typical pitfalls
   - How to avoid them
   - Debugging guidance

8. **Real-World Context**
   - Production use cases
   - When to apply this knowledge
   - Industry examples

9. **Extension Challenges**
   - Advanced variations
   - Integration with other concepts
   - Performance optimizations

10. **Testing Guidelines**
    - How to verify correctness
    - Testing patterns
    - Debugging approaches

## Scalability Guide

### Adapting AECS for New Domains

1. **Define Domain Expertise**
   - Create specialized agent profiles for the domain
   - Identify key technical patterns and best practices
   - Establish quality standards specific to the field

2. **Customize Templates**
   - Adapt exercise format for domain-specific needs
   - Add metadata fields relevant to the subject
   - Include domain-specific sections (e.g., security considerations)

3. **Create Generation Prompts**
   - Write prompts that understand domain context
   - Include parameter injection for domain-specific variables
   - Ensure prompts invoke appropriate specialized agents

4. **Validate with Pilot Content**
   - Generate sample curriculum using the new templates/prompts
   - Test with domain experts for accuracy and effectiveness
   - Refine based on feedback and results

### Replication Checklist

- [ ] Domain expert agents defined in `.claude/agents/`
- [ ] Templates customized for domain needs
- [ ] Generation prompts created with proper parameter support
- [ ] Quality standards documented
- [ ] Pilot content generated and validated
- [ ] WORKFLOW.md updated with domain-specific guidance

## Best Practices

### Content Generation

- **Start with curriculum structure** before diving into individual exercises
- **Maintain consistent difficulty progression** throughout chapters
- **Cross-reference related concepts** to build learning connections
- **Include real-world examples** relevant to professional practice
- **Avoid quantification in prompts** - use discovery-driven requirements instead of arbitrary numbers

### Subagent Collaboration

#### **Critical Best Practices for Effective Subagent Engagement**

**Problem**: Subagents often provide analysis without taking action when asked to "review" or "validate" content.

**Solution**: Use action-oriented prompting that explicitly requests concrete changes.

#### **Effective Subagent Prompting Patterns**

❌ **Ineffective**: "Please review the code for best practices"
✅ **Effective**: "Fix any non-idiomatic code patterns you find by editing the files directly"

❌ **Ineffective**: "Validate the learning progression"
✅ **Effective**: "Restructure sections that have poor learning progression and create any missing bridging content"

❌ **Ineffective**: "Check for consistency"
✅ **Effective**: "Standardize all heading formats and fix any formatting inconsistencies in the files"

#### **Action-Oriented Subagent Engagement Framework**

**1. Specify Concrete Actions**
- Use action verbs: "Fix", "Change", "Edit", "Create", "Update"
- Avoid analysis verbs: "Review", "Validate", "Analyze", "Check"

**2. Request Direct Action**
- Focus on outcomes: "Fix the issues", "Improve the content", "Restructure the sections"
- Define what "done" looks like in terms of improved files

**3. Define Deliverables**
- Request specific outputs: "After editing, the file should have..."
- Set measurable success criteria: "Fix all instances of...", "Create sections for..."

**4. Provide Complete Context**
- Give subagents all necessary information in a single prompt
- Include file paths, specific requirements, and success criteria
- Remember subagents have their own context windows

#### **Subagent Workflow Integration**

- **Invoke subagents early and often** in the generation process
- **Be specific about review criteria** when requesting subagent feedback
- **Always specify concrete actions** to prevent analysis-only responses
- **Request concrete deliverables** rather than assessment reports
- **Iterate based on subagent changes** to improve quality
- **Document successful prompting patterns** for future reference

### Prompt Engineering

- **Use discovery-driven requirements** instead of quantified specifications
- **Let domain research determine scope** rather than predetermined limits
- **Focus on comprehensive coverage** based on professional needs
- **Avoid arbitrary numbers** like "3-5 exercises" or "4-6 concepts" in prompts
- **Emphasize thoroughness** over artificial constraints

### Quality Assurance

- **Test all code examples** before finalizing content
- **Validate learning progression** with education experts
- **Ensure consistent style** across all materials
- **Review for accessibility** and inclusive language

### Maintenance

- **Version control templates and prompts** for change tracking
- **Update agent profiles** as domains evolve
- **Collect feedback** from content users for continuous improvement
- **Regular quality audits** to maintain standards

## Success Metrics

### Content Quality

- Code examples compile and run without errors
- Exercises have clear success criteria and are completable
- Learning objectives are measurable and achievable
- Content maintains consistent style and formatting

### Learning Effectiveness

- Progressive difficulty that builds on previous concepts
- Clear mental models that help learners understand concepts
- Practical exercises that reinforce theoretical knowledge
- Real-world relevance that motivates continued learning

### Scalability

- New curricula can be generated rapidly using existing infrastructure
- Templates and prompts work across different domains
- Quality remains consistent as content volume increases
- Maintenance overhead stays manageable as system grows

## Future Enhancements

### Planned Features

- **Automated testing integration** for code examples
- **Interactive exercise validation** through API integration
- **Personalized learning paths** based on prerequisite tracking
- **Community contribution workflows** for content improvement

### Research Areas

- **Adaptive difficulty scaling** based on learner performance
- **Multi-modal content generation** including diagrams and videos
- **Cross-domain knowledge transfer** patterns
- **Automated content quality assessment** metrics

## Glossary

**AECS**: Agentic Exercise Curriculum System - The complete methodology for creating exercise-driven technical curricula through AI collaboration.

**Subagents**: Specialized AI assistants with domain expertise (e.g., go-engineer, cs-professor, academic-editor) that ensure quality across different dimensions of content creation.

**Templates**: Standardized content structure definitions that ensure consistent formatting and quality across all generated materials.

**Generation Prompts**: Command-style instructions following the `[action]-[noun].md` format that create specific types of content (e.g., generate-curriculum.md, generate-exercise.md).

**Exercise-Driven Learning**: A pedagogical approach where knowledge is primarily acquired through hands-on practice rather than passive consumption of theory.

**Discovery-Driven Requirements**: Content specifications based on comprehensive domain research rather than predetermined quantitative limits.

**Command-Style Execution**: The standardized format for invoking generation prompts: "Execute [prompt-name].md. [PARAMETER] is [value]."

---

AECS represents a systematic approach to creating world-class technical education materials. By combining the expertise of specialized subagents with standardized templates and generation processes, AECS enables the rapid creation of consistent, high-quality learning experiences that scale across domains and topics.