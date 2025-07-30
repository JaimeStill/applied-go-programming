# Nested Subagent Problem Report

**Date**: 2025-07-30  
**Issue**: Recursive subagent spawning causing session timeouts  
**Priority**: Critical - Blocks framework functionality

## Problem Description

During Exercise generation using the AECS framework, subagent engagement results in nested subagent spawning that causes 500+ second timeouts and session deadlocks.

### Observed Behavior

When attempting to engage subagents for Exercise validation, the following pattern appears:

```
[subagent](executing task)...
   [subagent](executing task)...
```

This indicates recursive subagent spawning where:
1. Framework attempts to engage a subagent
2. Subagent spawns another instance of itself
3. Process continues indefinitely
4. Session eventually times out after ~500 seconds

### Reproduction Steps

1. Execute `_framework/prompts/generate-exercise.md`
2. Framework reaches subagent engagement phase
3. Attempts to engage project-specific subagents (go-engineer, cs-professor, academic-editor)
4. Recursive spawning occurs
5. Session freezes with no observable progress

## Root Cause Analysis

### Primary Issue: Framework/Project Subagent Name Mismatch

The `_framework/` system expects to find subagents but doesn't define their names:

**Framework Expectation** (`_framework/CLAUDE.md:58`):
```
Use Claude's built-in subagent feature by referencing subagents defined in `.claude/agents/`. 
Do NOT use the Task tool to simulate subagents.
```

**Framework Discovery** (`_framework/CLAUDE.md:60`):
```
Examine `.claude/agents/` directory to identify available project subagents 
and match to required expertise roles.
```

**The Gap**: Framework discovers project-specific names (`go-engineer`, `cs-professor`, `academic-editor`) but has no standard way to reference them.

### Secondary Issue: Task Tool Fallback

When direct subagent references fail (due to name mismatches), the system appears to fall back to the Task tool, which creates recursive spawning:

1. Framework tries to engage subagent directly → Fails
2. System falls back to Task tool → Launches subagent-type agent
3. That agent internally tries to engage subagent → Uses Task tool
4. Creates infinite loop: Task→Subagent→Task→Subagent→...

### Context Overload

Current approach sends entire Exercise documents (200+ lines) to subagents, creating:
- Excessive token usage
- Extended processing times
- Increased likelihood of recursive behavior

## Impact Assessment

### Immediate Impact
- **Framework Unusability**: Cannot complete Exercise generation
- **Development Blocked**: No new content can be created
- **Poor User Experience**: 500+ second waits with no feedback
- **Token Waste**: Excessive consumption during failed attempts

### Long-term Impact
- **Framework Reliability**: Unpredictable subagent behavior
- **Scalability Issues**: Problem compounds with complex content
- **Maintenance Burden**: Manual workarounds required

## Solution Analysis

### Option 1: Dynamic Discovery with Direct References
**Approach**: Discover agents at runtime, use direct `@agent-name` references
- ✓ Maintains project flexibility
- ✗ Still vulnerable to naming mismatches
- ✗ Runtime discovery complexity

### Option 2: Standardized Subagent Templates (RECOMMENDED)
**Approach**: Create subject-agnostic templates, scaffold into standard names
- ✓ Predictable naming: `@technical-expert`, `@learning-designer`, `@content-editor`
- ✓ Subject flexibility through parameterization
- ✓ Eliminates recursive spawning risk
- ✓ Framework can safely reference known names
- ✓ Templates already exist in framework

### Option 3: Self-Validation Fallback
**Approach**: Skip subagent engagement for simple content
- ✓ Eliminates spawning issues
- ✗ Reduces validation quality
- ✗ Violates framework requirements

## Recommended Solution: Standardized Subagent Templates

### Implementation Strategy

1. **Template Organization**: Move existing templates to `_framework/templates/agents/`
2. **Subject Parameterization**: Update templates to accept `{{SUBJECT}}` parameter
3. **Standard Names**: Scaffold as `technical-expert`, `learning-designer`, `content-editor`
4. **Framework Integration**: Update all framework references to use standard names
5. **Path Creation**: Automatically scaffold subagents during path creation

### Benefits

- **Eliminates Recursive Spawning**: Predictable, known subagent names
- **Maintains Flexibility**: Templates adapt to any subject through parameterization
- **Framework Stability**: No runtime discovery or fallback mechanisms needed
- **Reduced Complexity**: Single standard for all learning paths
- **Backward Compatibility**: Existing templates can be easily converted

## Next Steps

1. Document this problem comprehensively ✓
2. Create standardized template structure ✓
3. Parameterize templates for subject flexibility ✓
4. Generate Go-specialized subagents from templates ✓
5. Test direct subagent references without recursive spawning ✓
6. Update entire framework infrastructure to use standardized approach ✓
7. Clean up non-standard subagents (pending)

## Critical Discovery: Task Tool vs Direct Invocation Behavior

**Key Finding**: Claude Code subagent activation behaves differently depending on invocation method, with Task tool usage causing recursive spawning while direct user invocation works cleanly.

**Evidence**:

### Working Method: Direct User Invocation
```
@agent-technical-expert [prompt]
```
**Result**: 
```
● technical-expert(Review Go code for AECS compliance)
  ⎿  Done (0 tool uses · 15.0k tokens · 9.9s)
```
✅ **Clean execution, no nesting, completes successfully**

### Problematic Method: Task Tool Invocation (from Claude)
```python
Task(subagent_type="technical-expert", prompt="[prompt]")
```
**Result**:
```
learning-designer(Review Exercise structure for AECS compliance)
  ⎿  Listed 10 paths (ctrl+r to expand)
     learning-designer(Review Exercise structure compliance)
     +1 more tool use
```
❌ **Recursive spawning, nested subagent loops, session timeouts after ~500 seconds**

### Additional Discovery: Session Initialization Requirement
- Claude Code loads and indexes agents at session initialization
- Agents created during a session are not recognized until the session is restarted
- Both old and new agents appear in `claude agents list` after restart and can be invoked by users

**Root Cause**: The Task tool appears to trigger Claude Code's internal subagent spawning mechanism, which then attempts to spawn another subagent instance, creating infinite recursion. Direct user invocation with `@agent-[name]` syntax bypasses this mechanism entirely.

## Success Criteria

- [ ] No recursive subagent spawning during Exercise generation
- [ ] Subagent responses complete within reasonable time (< 30 seconds)  
- [ ] Framework works consistently across different subjects
- [✓] Standardized subagent templates created and properly structured
- [ ] Clean session validation passes for Exercise generation (requires restart)

---

## GitHub Issue Template for Claude Code Repository

### Title
**Task Tool Subagent Invocation Causes Recursive Spawning and Session Timeouts**

### Labels
- `bug`
- `subagents`
- `task-tool`

### Description

**Problem Summary**
When Claude uses the Task tool to invoke subagents (`Task(subagent_type="subagent-name")`), it creates recursive subagent spawning that leads to infinite loops and session timeouts (~500 seconds).

**Expected Behavior**
Task tool subagent invocation should work cleanly like direct user invocation with `@agent-[name]` syntax.

**Actual Behavior**
Task tool invocation creates nested subagent loops:
```
subagent-name(task description)
  ⎿  Listed 10 paths (ctrl+r to expand)
     subagent-name(nested task)
     +1 more tool use
```

**Reproduction Steps**
1. Create a subagent in `.claude/agents/` with proper YAML frontmatter
2. Have Claude invoke the subagent using `Task(subagent_type="subagent-name", prompt="test")`
3. Observe recursive spawning and eventual timeout

**Working Workaround**
Direct user invocation works correctly:
```
@agent-subagent-name [prompt]
```
Results in clean execution without nesting.

**Environment**
- Claude Code version: 1.0.63
- Platform: Linux (Pop_OS! 24.04)
- Subagent files: Properly structured with YAML frontmatter

**Impact**
- Blocks programmatic subagent usage in frameworks
- Causes sessions to get stuck and max CPU usage
- Prevents automated workflows that depend on subagent validation

**Additional Context**
This issue prevents building educational frameworks that rely on subagent validation workflows. The recursive spawning suggests the Task tool may be triggering Claude Code's internal subagent mechanism inappropriately.

---

**Report Status**: Appended report to existing GitHub Issue: [#4744 - Agent Execution Timeout: Persistent Hanging During Complex Tasks](https://github.com/anthropics/claude-code/issues/4744)
**Next Action**: Submit GitHub issue and complete framework overhaul