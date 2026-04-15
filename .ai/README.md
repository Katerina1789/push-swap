# Personal AI Collaboration Guide

This folder exists to help the team use AI support in a consistent and professional way during the development of `push-swap`.

## Purpose

Use the AI assistant as a project support tool, not as a replacement for engineering judgment. Every request and every generated change must stay inside the scope of the school assignment.

Before asking for help, make sure the request is related to:
- the Go core logic and CLI functionality
- the specific project requirements
- validation and error handling
- documentation, testing, or project organization

If a request does not support the assignment requirements, it should not be adopted into the project.

## Required Workflow

When working with the AI assistant in this repository, follow this rule:

1. Ask for a response.
2. After the response, append a short English entry to your own AI log file.
3. If the work affects a task, update the related task file status and checklist before closing the work session.

Each teammate should create and maintain a personal log file inside `.ai/`.

Examples:
- `.ai/hmim.ai.txt`
- `.ai/kkasdana.ai.txt`
- `.ai/ebasou.ai.txt`

Do not rely on one shared log for the whole team. Each person is responsible for updating their own file after every meaningful AI-assisted interaction.

Each log entry should include:

- **Date**: [Today's Date]
- **Model Used**: [e.g., Claude 3.5 Sonnet / Gemini Code Assist]
- **Discussion Summary**: [Brief summary of interaction between the user prompt and the gent's proposal or response]
- **Action Taken**: [What files were generated or modified]
- **Affected Area/Task Card**: [e.g., .docs creation, TASK-01]

This gives the team a lightweight and traceable history of AI-assisted decisions and keeps collaboration transparent at the individual level.

## Required Team Discipline

The following rules are mandatory for all teammates:

- each teammate must maintain their own file inside `.ai/`
- each teammate must update their personal AI log after meaningful AI-assisted work
- each teammate must review changes before committing
- each teammate must use clean and focused commit messages
- each teammate must update the relevant `.tasks` file when a task moves forward
- task checklists and task status must reflect reality, not intention

If work is completed but the log or task file is not updated, the work is considered undocumented and incomplete from a team-process perspective.

## How To Use AI Properly

The final responsibility always stays with the team.
