# Generate Curriculum Prompt

## Purpose
Create a comprehensive, exercise-driven curriculum for any technical topic following the Agentic Exercise Curriculum System (AECS) methodology.

## Parameters
- **TOPIC**: The main subject matter (e.g., "Linux Administration for Go Programmers", "Rust Systems Programming", "Kubernetes Operations")
- **TARGET_AUDIENCE**: Intended learners (e.g., "experienced developers", "system administrators", "data engineers")
- **DIFFICULTY_LEVEL**: Overall course complexity (Beginner/Intermediate/Advanced)
- **DURATION**: Expected completion time (e.g., "6 weeks", "3 months", "self-paced")

## Execution Instructions

Use the curriculum template at `_templates/curriculum.md` to generate a complete curriculum structure.

### Step 1: Analysis Phase
Before generating content, analyze:
1. **Domain Scope**: What are the essential concepts, tools, and practices in this domain?
2. **Learning Progression**: What's the logical order for introducing concepts?
3. **Practical Applications**: What real-world projects would demonstrate mastery?
4. **Prerequisites**: What foundational knowledge is required?

### Step 2: Structure Design
Create a logically organized curriculum structure:
- **Part I**: Foundational concepts and basic patterns
- **Part II**: Intermediate techniques and real-world applications  
- **Part III**: Advanced topics and production practices

Structure the curriculum with appropriate number of parts and chapters based on the domain scope and learning progression requirements.

### Step 3: Content Generation
Fill the template with:

**Course Overview:**
- Clear description of what learners will accomplish
- Emphasis on hands-on, exercise-driven approach
- Real-world relevance and career benefits

**Learning Objectives:**
- Specific, measurable outcomes
- Mix of technical skills and conceptual understanding
- Progressive complexity across the curriculum

**Chapter Design:**
For each chapter, provide:
- Compelling title that indicates practical value
- Concise description of the chapter's focus
- Comprehensive key topics list
- Realistic duration estimate
- Clear connection to overall learning progression

**Dependencies:**
- Map prerequisite relationships between chapters
- Identify optional vs. required learning paths
- Consider parallel tracks for different learning styles

### Step 4: Quality Assurance
Ensure the curriculum:
- Builds progressively from simple to complex
- Balances theoretical understanding with practical application
- Includes regular mastery checkpoints
- Provides clear success criteria
- Offers flexibility for different learning paces

### Step 5: Expert Review Integration
The generated curriculum should be automatically reviewed by:
- **Domain Expert**: Technical accuracy and industry relevance
- **Education Expert**: Learning progression and pedagogical effectiveness
- **Academic Editor**: Clarity, consistency, and professional presentation

## Output Requirements

Generate a complete curriculum document that includes:
1. Compelling course title and subtitle
2. Clear learning objectives and prerequisites
3. Well-structured chapter progression across logical parts
4. Realistic time estimates and difficulty indicators
5. Assessment strategy and success criteria
6. Essential resources and tools
7. Community and support resources

## Success Criteria

The generated curriculum should:
- Be immediately actionable for content creation
- Provide clear guidance for chapter development
- Appeal to the target audience
- Demonstrate clear value proposition
- Follow AECS principles of exercise-driven learning
- Include all necessary metadata for curriculum management

## Example Usage

```
Execute generate-curriculum.md. TOPIC is "Cloud-Native Go Development". TARGET_AUDIENCE is "backend developers with Go experience". DIFFICULTY_LEVEL is "Intermediate". DURATION is "8 weeks".
```

This will generate a complete curriculum structure ready for chapter-by-chapter development using the `generate-chapter.md` prompt.