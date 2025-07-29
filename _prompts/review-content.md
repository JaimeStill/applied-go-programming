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

### Step 3: Action-Oriented Subagent Review Process

**Critical**: Always request specific actions from subagents, not just analysis.

**Single Document Review:**
1. Engage appropriate subagents with action-oriented prompts:
   - **Domain Expert**: "Fix technical issues and improve code quality by editing the files directly"
   - **CS Professor**: "Restructure sections with poor learning progression by updating the content"
   - **Academic Editor**: "Fix formatting and improve readability by editing the files"
2. Request concrete deliverables: improved files, not just feedback reports
3. Specify what each subagent should accomplish through direct file changes
4. Define success criteria: "After your improvements, all code should compile and follow best practices"

**Multi-Document Review:**
1. Review individual documents with action-oriented subagent engagement
2. Have academic-editor standardize formatting across all documents by editing them
3. Have cs-professor restructure content flow between documents by updating the files
4. Request domain expert to ensure consistent technical patterns by improving files
5. Validate changes improve overall learning progression

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

**Action-Oriented Curriculum Review:**
```
Execute review-content.md. CONTENT_TYPE is "complete curriculum". DOCUMENT_PATHS is "curriculum.md, chapters/00-foundations/go-foundations.md, chapters/01-concurrency/01-basic-goroutines.md". REVIEW_FOCUS is "technical fixes and learning progression improvements". CONTEXT is "Use subagents to make actual edits and improvements, not just provide analysis".
```

**Chapter Improvement with Specific Actions:**
```
Execute review-content.md. CONTENT_TYPE is "chapter with exercises". DOCUMENT_PATHS is "chapters/01-concurrency/roadmap.md, chapters/01-concurrency/01-basic-goroutines.md". REVIEW_FOCUS is "fix code issues and restructure poor learning flow". CONTEXT is "Each subagent should use editing tools to make concrete improvements to files".
```

**Important**: This prompt will engage subagents to make actual changes to files, not just provide assessment reports. Always specify what improvements each subagent should make to the content.