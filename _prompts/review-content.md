# Review Content Prompt

## Purpose
Conduct comprehensive review of curriculum documents using specialized subagents to ensure technical accuracy, pedagogical effectiveness, and structural consistency. Supports both single document and multi-document review workflows.

## Parameters
- **CONTENT_TYPE**: Type of content being reviewed (e.g., "curriculum", "chapter", "exercise", "appendix", "workflow documentation")
- **DOCUMENT_PATHS**: List of documents to review (single or multiple)
- **REVIEW_FOCUS**: Specific aspects to emphasize (e.g., "technical accuracy", "learning progression", "structural consistency")
- **CONTEXT**: Additional context about the content's role in the overall curriculum

## Execution Instructions

### Step 1: Content Analysis
Before engaging subagents, analyze the content to determine:
1. **Content Classification**: What type of educational material is being reviewed?
2. **Technical Domain**: What subject matter expertise is required?
3. **Educational Level**: What pedagogical expertise is needed?
4. **Structural Requirements**: What formatting and consistency standards apply?
5. **Integration Points**: How does this content connect to other materials?

### Step 2: Subagent Selection
Based on content analysis, automatically engage appropriate subagents:

**For Technical Content (exercises, code examples, implementations):**
- **Domain Expert** (e.g., go-engineer, database-engineer): Technical accuracy, best practices, production readiness

**For Educational Content (curricula, chapters, learning materials):**
- **CS Professor**: Learning progression, concept scaffolding, assessment validity

**For All Content:**
- **Academic Editor**: Structure, clarity, consistency, professional presentation

### Step 3: Comprehensive Review Process

**Single Document Review:**
1. Engage appropriate subagents simultaneously for parallel review
2. Collect feedback on technical accuracy, educational effectiveness, and presentation quality
3. Identify specific issues with actionable improvement suggestions
4. Validate that all AECS standards are met

**Multi-Document Review:**
1. Review individual documents as above
2. Analyze flow and progression between documents
3. Verify cross-references and connections are accurate
4. Ensure consistent terminology and conventions across all materials
5. Validate learning progression and prerequisite alignment

### Step 4: Integration and Consistency Validation

**Cross-Document Analysis:**
- Verify prerequisite chains are accurate and complete
- Check that difficulty progression is appropriate
- Ensure consistent use of terminology and conventions
- Validate that learning objectives align across materials

**Link and Reference Verification:**
- Confirm all internal references are accurate
- Verify external links are current and relevant
- Check that cross-references enhance rather than distract from learning

**Metadata Consistency:**
- Ensure all documents follow established formatting standards
- Verify metadata is complete and accurate
- Check that file naming conventions are followed

### Step 5: Quality Assurance Report

Generate comprehensive review report that includes:

**Technical Accuracy Assessment:**
- Code compilation and execution verification
- Best practices compliance
- Industry standard alignment
- Performance and security considerations

**Educational Effectiveness Evaluation:**
- Learning objective clarity and measurability
- Concept scaffolding and progression
- Assessment validity and alignment
- Cognitive load appropriateness

**Structural and Presentation Quality:**
- AECS methodology compliance
- Template adherence
- Formatting consistency
- Professional presentation standards

**Integration and Flow Analysis:**
- Prerequisites and dependencies
- Cross-references and connections
- Terminology consistency
- Overall curriculum coherence

### Step 6: Actionable Recommendations

For each identified issue, provide:
1. **Specific Problem**: What exactly needs improvement
2. **Impact Assessment**: How this affects learning outcomes
3. **Recommended Solution**: Concrete steps to address the issue
4. **Priority Level**: Critical, important, or minor
5. **Implementation Guidance**: How to make the recommended changes

## Output Requirements

Generate a comprehensive review report that includes:
1. Executive summary of overall content quality
2. Detailed findings organized by review criteria
3. Specific, actionable improvement recommendations
4. Priority-ranked action items
5. Validation that AECS standards are met
6. Assessment of readiness for learner use

## Success Criteria

The review process should:
- Ensure technical accuracy and industry best practices
- Validate educational effectiveness and learning progression
- Confirm structural consistency and professional presentation
- Verify appropriate integration and cross-references
- Provide clear, actionable improvement guidance
- Maintain AECS quality standards throughout

## Special Considerations

### Multi-Document Workflows
- Pay special attention to transitions between documents
- Ensure prerequisite knowledge is properly established
- Verify that difficulty progression is smooth and logical
- Check that concepts build appropriately across materials

### Technical Content Focus
- All code examples must compile and execute correctly
- Ensure production-ready patterns and practices
- Validate performance and security considerations
- Confirm adherence to language-specific best practices

### Educational Content Focus
- Learning objectives must be clear and measurable
- Concept introduction should follow pedagogical best practices
- Assessment methods should align with learning goals
- Content should be accessible to the target audience

## Example Usage

```
Execute review-content.md. CONTENT_TYPE is "complete curriculum". DOCUMENT_PATHS is "curriculum.md, chapters/00-foundations/go-foundations.md, chapters/01-concurrency/01-basic-goroutines.md". REVIEW_FOCUS is "comprehensive quality assurance before curriculum launch". CONTEXT is "Final validation of AECS-generated Go programming curriculum for experienced developers".
```

```
Execute review-content.md. CONTENT_TYPE is "chapter with exercises". DOCUMENT_PATHS is "chapters/01-concurrency/roadmap.md, chapters/01-concurrency/01-basic-goroutines.md". REVIEW_FOCUS is "technical accuracy and learning progression". CONTEXT is "Concurrency chapter review focusing on Go best practices and pedagogical effectiveness".
```

This will engage appropriate subagents to conduct thorough review and provide actionable feedback for curriculum improvement.