# Guide

## Initialization

To create a new training plan, run the following in plan mode:

```
Prompt: _frameworks/prompts/create-path.md
Subject: [whatever you are interested in learning]
Ultrathink.
```

## Prompt Execution

You will get the best results by executing the prompts in the following format in plan mode:

```
Prompt: [path-to-prompt]
Parameter: [value]
Follow all directives from CLAUDE.md. Ultrathink.
```

Example:

```
Prompt: _frameworks/prompts/create-exercise.md
Exercise: Stage 1 Exercise 2.
Follow all directives from CLAUDE.md. Ultrathink.
```

## Subagent Activation

There's a current limitation with the implementation of subagent activation. Whenever subagents are invoked directly from the main agent, they are prone to nested subagents that get stuck in a deadlock consuming compute resources (see: https://github.com/JaimeStill/applied-go-programming/blob/main/.admin/reports/nested-subagents.md).

This issue does not present whenever the user activates a subagent. Until this issue gets resolved, I have the framework configured to output the prompts for you to pass to the subagent.

Example output:

```
Excellent feedback from the technical expert! The exercise received a perfect assessment. Now I need to continue with the learning designer validation:

  @learning-designer You are reviewing an EXERCISE document that should enforce Exercise Primacy through hands-on building.

  Exercise Content:
  The exercise focuses on basic Go program structure and compilation through building a working program. It provides complete step-by-step implementation with no placeholders, includes compilation and execution steps, and produces observable system output.

  Review scope: Identify passive consumption elements and multi-concept violations.

  Your feedback should:
  - Identify specific AECS principle violations
  - Suggest structural reorganization for immediate hands-on engagement
  - Verify single-concept atomicity and proper dependency chains
  - Focus on learning design principles

  Format your response as: Violation → Structural Fix
```

You then submit the subagent prompt as follows:

```
@agent-learning-designer You are reviewing an EXERCISE document that should enforce Exercise Primacy through hands-on building.

    Exercise Content:
    The exercise focuses on basic Go program structure and compilation through building a working program. It provides complete step-by-step implementation with no placeholders, includes compilation and execution steps, and produces observable system output.

    Review scope: Identify passive consumption elements and multi-concept violations.

    Your feedback should:
    - Identify specific AECS principle violations
    - Suggest structural reorganization for immediate hands-on engagement
    - Verify single-concept atomicity and proper dependency chains
    - Focus on learning design principles

  Format your response as: Violation → Structural Fix
  
  EXERCISE: stages/01-fundamentals/01-basic-program-structure.md
```