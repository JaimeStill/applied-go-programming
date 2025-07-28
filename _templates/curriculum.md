# {{TITLE}}: {{SUBTITLE}}

## Course Metadata

- **Target Audience**: {{TARGET_AUDIENCE}}
- **Difficulty Level**: {{DIFFICULTY_LEVEL}}
- **Total Duration**: {{TOTAL_DURATION}}
- **Format**: {{FORMAT}}
- **Last Updated**: {{LAST_UPDATED}}

## Course Overview

{{COURSE_DESCRIPTION}}

## Learning Objectives

By completing this curriculum, you will:

{{#each LEARNING_OBJECTIVES}}
- [ ] {{this}}
{{/each}}

## Prerequisites

{{PREREQUISITES}}

## Course Structure

### Part I: {{PART_1_TITLE}}

{{#each PART_1_CHAPTERS}}
**Chapter {{@index}}: {{title}}**

{{description}}

**Key Topics:**
{{#each topics}}
- {{this}}
{{/each}}

**Estimated Duration:** {{duration}}

---
{{/each}}

### Part II: {{PART_2_TITLE}}

{{#each PART_2_CHAPTERS}}
**Chapter {{@index}}: {{title}}**

{{description}}

**Key Topics:**
{{#each topics}}
- {{this}}
{{/each}}

**Estimated Duration:** {{duration}}

---
{{/each}}

### Part III: {{PART_3_TITLE}}

{{#each PART_3_CHAPTERS}}
**Chapter {{@index}}: {{title}}**

{{description}}

**Key Topics:**
{{#each topics}}
- {{this}}
{{/each}}

**Estimated Duration:** {{duration}}

---
{{/each}}

## Learning Path

```mermaid
graph TD
    A[Prerequisites] --> B[{{PART_1_TITLE}}]
    B --> C[{{PART_2_TITLE}}]
    C --> D[{{PART_3_TITLE}}]
    
    {{#each CHAPTER_DEPENDENCIES}}
    {{prerequisite}} --> {{dependent}}
    {{/each}}
```

## Assessment Strategy

### Progress Tracking
- [ ] **Chapter Completion**: Complete all exercises in each chapter
- [ ] **Mastery Checkpoints**: Pass comprehensive assessments every 3-4 chapters
- [ ] **Portfolio Projects**: Build real-world applications demonstrating integrated skills

### Success Criteria
{{#each SUCCESS_CRITERIA}}
- {{this}}
{{/each}}

## Resources

### Essential Tools
{{#each REQUIRED_TOOLS}}
- **{{name}}**: {{description}}
{{/each}}

### Recommended Reading
{{#each RECOMMENDED_READING}}
- [{{title}}]({{url}}) - {{description}}
{{/each}}

### Community & Support
{{#each SUPPORT_RESOURCES}}
- **{{name}}**: {{description}} - [{{url}}]({{url}})
{{/each}}

---

*This curriculum follows the Agentic Exercise Curriculum System (AECS) methodology, ensuring consistent quality and pedagogical effectiveness through AI agent collaboration.*