# Plans

## Learning Path Configurations

Just like a containerized server, there should be parameters that can be injected into the framework that adjusts its behavior. Different metrics that affect the output of the prompts that are core to the AECS framework. Example:

```json
{
  "profile": "consice-programmer",
  "path": {
    "minStages": 8,
    "maxStages": 15,
    "stages": {
      "minExercises": 5,
      "maxExercises": 18,
      "exercises": {
        "sections": [
          "exercise-metadata",
          "overview",
          "description",
          "foundation-requirements",
          "exercise",
          "path-integration",
          "footer"
        ]
      }
    }
  }
}
```