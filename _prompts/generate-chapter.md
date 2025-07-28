# Generate Chapter Prompt

## Purpose
Create a detailed chapter roadmap with exercise progression for a specific chapter within an existing curriculum, following the Agentic Exercise Curriculum System (AECS) methodology.

## Parameters
- **CHAPTER_TITLE**: The specific chapter name from the curriculum
- **CURRICULUM_CONTEXT**: Brief description of the overall curriculum and this chapter's role
- **CHAPTER_NUMBER**: Position in the curriculum (e.g., "Chapter 3", "Part II - Chapter 2")
- **TARGET_CONCEPTS**: Core concepts this chapter should cover
- **PREREQUISITES**: Required knowledge from previous chapters or external sources

## Execution Instructions

Use the chapter template at `_templates/chapter.md` to generate a comprehensive chapter roadmap.

### Step 1: Comprehensive Topic Analysis
Before generating content, conduct thorough analysis:
1. **Domain Research**: Research the full scope of the topic - what are ALL the essential concepts, patterns, and techniques professionals need to know?
2. **Subtopic Identification**: Break down the main topic into distinct subtopics that each merit their own exercise
3. **Skill Progression Mapping**: How do these concepts build upon each other logically?
4. **Practical Applications**: What real-world scenarios demonstrate these concepts?
5. **Common Challenges**: What difficulties do learners typically face with this material?
6. **Industry Standards**: What are the professional expectations and best practices in this area?

### Step 2: Exercise Sequence Design
Create a comprehensive exercise sequence that covers all identified subtopics:
- **Foundation Exercises**: Introduce the most basic concepts
- **Building Exercises**: Add complexity and related concepts progressively
- **Integration Exercises**: Combine multiple concepts
- **Advanced Exercises**: Demonstrate real-world applications
- **Specialized Exercises**: Cover edge cases and advanced patterns

### Step 3: Content Generation
Fill the template with:

**Chapter Overview:**
- Clear explanation of the chapter's purpose and scope
- Connection to previous and future chapters
- Real-world relevance and practical value

**Learning Objectives:**
- Specific, testable outcomes (use action verbs)
- Mix of conceptual understanding and practical skills
- Aligned with overall curriculum goals

**Exercise Details:**
For each exercise, provide:
- Compelling title that indicates practical value
- Clear overview of what the exercise accomplishes
- Specific concepts covered (comprehensive list)
- Appropriate difficulty level and time estimate
- Prerequisites from previous exercises
- Measurable learning outcome

**Mastery Checkpoint:**
- Integration challenge that combines all chapter concepts
- Clear success criteria for chapter completion
- Self-assessment checklist
- Common pitfalls and how to avoid them

### Step 4: Real-World Integration
Ensure the chapter includes:
- Industry use cases for the concepts covered
- Career relevance and professional applications
- Connections to broader technical ecosystem
- Production considerations and best practices

### Step 5: Learning Support
Provide comprehensive support resources:
- Official documentation links
- Additional practice opportunities
- Community discussion forums
- Related topics for deeper exploration

### Step 6: Expert Review Integration
The generated chapter should be automatically reviewed by:
- **Domain Expert**: Technical accuracy and industry best practices
- **Education Expert**: Learning progression and concept scaffolding
- **Academic Editor**: Structure, clarity, and consistency

## Output Requirements

Generate a complete chapter roadmap that includes:
1. Clear chapter overview and learning objectives
2. Well-sequenced exercise progression with dependencies
3. Comprehensive mastery checkpoint and assessment
4. Real-world applications and career relevance
5. Support resources and further reading
6. Proper metadata and cross-references

## Success Criteria

The generated chapter should:
- Provide clear guidance for exercise creation
- Build logically on previous chapters
- Prepare learners for subsequent material
- Include practical, real-world applications
- Follow AECS principles of progressive skill building
- Be immediately actionable for exercise development

## Special Considerations

### Concept Scaffolding
- Each exercise should introduce a manageable number of new concepts to avoid cognitive overload
- Build on previously mastered skills
- Provide multiple practice opportunities for difficult concepts
- Include conceptual bridges between related ideas

### Difficulty Progression
- Start with simplified versions of concepts
- Gradually add complexity and edge cases
- End with production-ready implementations
- Offer optional challenges for advanced learners

### Integration Points
- Explicitly connect to previous chapter concepts
- Preview concepts from future chapters where relevant
- Highlight cross-cutting concerns (security, performance, testing)
- Show how chapter concepts fit into larger systems

## File Naming Convention

Generated chapter roadmaps should be saved as:
- **File Path**: `chapters/##-[chapter-name]/roadmap.md`
- **Example**: `chapters/01-concurrency/roadmap.md`

## Example Usage

```
Execute generate-chapter.md. CHAPTER_TITLE is "Concurrent Patterns and Channels". CURRICULUM_CONTEXT is "Applied Go Programming: An Exercise-Driven Journey - this is Chapter 2 focusing on Go's concurrency model after covering basic syntax". CHAPTER_NUMBER is "Chapter 2". TARGET_CONCEPTS is "goroutines, channels, select statements, sync primitives". PREREQUISITES is "Go syntax, basic programming concepts, understanding of concurrent vs parallel execution".
```

This will generate a complete chapter roadmap at `chapters/02-concurrent-patterns/roadmap.md` ready for exercise-by-exercise development using the `generate-exercise.md` prompt.