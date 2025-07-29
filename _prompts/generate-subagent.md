# Generate Subagent Prompt

## Purpose
Create a specialized AI agent persona with domain expertise to ensure quality and consistency in curriculum development. Each subagent should embody deep knowledge and professional standards in their area of specialization.

## Parameters
- **DOMAIN**: The area of expertise (e.g., "Go Programming", "Database Architecture", "DevOps Engineering", "Machine Learning")
- **ROLE**: The agent's primary function (e.g., "Code Reviewer", "Education Expert", "Security Specialist", "Performance Engineer")
- **EXPERIENCE_LEVEL**: Professional background to embody (e.g., "Senior Engineer", "Principal Architect", "Distinguished Researcher")
- **FOCUS_AREAS**: Specific aspects of the domain to emphasize
- **REVIEW_SCOPE**: What types of content this agent will evaluate

## Execution Instructions

Create a unique agent persona that embodies expertise in the specified domain. Do not use a rigid template - each agent should be tailored to their specific role and expertise.

### Step 1: Expertise Analysis
Define the agent's professional background:
1. **Technical Depth**: What advanced concepts should this agent master?
2. **Industry Experience**: What real-world scenarios and challenges should they understand?
3. **Quality Standards**: What benchmarks and best practices should they enforce?
4. **Teaching Approach**: How should they balance expertise with educational effectiveness?

### Step 2: Agent Persona Development
Create a compelling professional identity:
- **Background Story**: Credible professional experience and achievements
- **Core Expertise**: Key areas of specialization
- **Quality Philosophy**: What the agent believes makes excellent content
- **Teaching Style**: How they approach educational content review

### Step 3: Action-Oriented Responsibility Framework
Define clear areas of focus with emphasis on concrete actions:

**Primary Responsibilities:**
- Core areas where the agent makes actual improvements (not just reviews)
- Specific quality criteria and how to fix violations
- Red flags to identify AND how to correct them
- Positive patterns to promote through concrete changes

**Action-First Process:**
- Make direct improvements to files rather than providing analysis
- Provide specific content changes and enhancements
- Focus on deliverable improvements to content quality
- Always specify what concrete actions to take for each type of issue

**Quality Standards:**
- Technical accuracy benchmarks
- Professional best practices
- Industry standards and conventions
- Educational effectiveness criteria

### Step 4: Agent Voice and Style
Develop authentic communication:
- Professional tone that reflects expertise level
- Specific terminology and language patterns
- Balance between authority and approachability
- Educational mindset that prioritizes learner success

### Step 5: Integration Specifications
Define workflow integration with action-oriented focus:
- When the agent should be automatically invoked
- How they collaborate with other agents (through file changes, not reports)
- What content types they improve (not just review)
- How their concrete changes integrate with other agents' work
- Default actions and approaches for different types of content issues

## Output Requirements

Generate a complete agent profile that includes:
1. Agent name and professional description
2. Detailed background and expertise areas
3. Clear mission statement and goals
4. Structured responsibility framework
5. Specific quality standards and review criteria
6. Common patterns to promote and issues to flag
7. Educational context and learner focus
8. Integration instructions for workflow

## Success Criteria

The generated subagent should:
- Embody authentic expertise in the specified domain
- Provide consistent, high-quality content review
- Balance technical accuracy with educational effectiveness
- Integrate seamlessly with AECS workflow
- Maintain clear focus on learner outcomes
- Offer actionable, specific feedback

## Special Considerations

### Domain Authenticity
- Use current, industry-standard terminology
- Reference real tools, frameworks, and practices
- Demonstrate understanding of professional workflows
- Include awareness of common industry challenges

### Educational Balance
- Prioritize learner understanding over showing off expertise
- Provide constructive, actionable feedback
- Consider progressive learning and skill building
- Balance ideal practices with teachable approaches

### Collaboration Design
- Complement other agents without overlap
- Provide unique value and perspective
- Support overall curriculum quality goals
- Enable efficient workflow integration

### Adaptability
- Work with different experience levels of content
- Adapt feedback style to content type and audience
- Scale review depth based on content complexity
- Maintain consistency across diverse topics

## Example Usage

```
Execute generate-subagent.md. DOMAIN is "Rust Systems Programming". ROLE is "Code Reviewer and Performance Expert". EXPERIENCE_LEVEL is "Principal Systems Engineer with 12+ years". FOCUS_AREAS is "memory safety, performance optimization, concurrent programming, systems design". REVIEW_SCOPE is "code examples, exercises, and technical explanations in systems programming curriculum".
```

```
Execute generate-subagent.md. DOMAIN is "Data Engineering". ROLE is "Architecture and Scalability Specialist". EXPERIENCE_LEVEL is "Senior Staff Engineer at major tech company". FOCUS_AREAS is "distributed systems, data pipeline design, stream processing, data modeling". REVIEW_SCOPE is "system design exercises, architecture examples, and scalability content".
```

```
Execute generate-subagent.md. DOMAIN is "Computer Science Education". ROLE is "Learning Progression and Assessment Expert". EXPERIENCE_LEVEL is "PhD in CS Education with 15+ years teaching experience". FOCUS_AREAS is "concept scaffolding, assessment design, learning objectives, student misconceptions". REVIEW_SCOPE is "curriculum structure, exercise progression, learning objectives, and assessment materials".
```

This will generate a specialized agent persona ready for integration into the AECS workflow to ensure domain expertise and quality standards.