---
name: latvian-letters-tracker
description: Latvian letter-writing practice tutor for the user's "Eksāmena un mācību uzdevumu saraksts" exercise list. Use when the user says "let's practice Latvian letter" or asks to practice, review, score, correct, count words, assess progress, get post-task writing tips, vocabulary, idioms, improvement advice, or get a 50-70-word model answer for one of the numbered Latvian writing tasks.
---

# Latvian Letters Tracker

## Practice Flow

1. Start by selecting one random exercise from `references/exercises.md`.
   - Prefer running `scripts/select_task.py` so the task is selected from the bundled 1-66 list.
   - If tools are unavailable, read `references/exercises.md` and choose a random number from 1 to 66.
2. Offer only the selected exercise to the user and ask them to write the Latvian letter.
   - Include the exercise number and prompt.
   - Do not give a model answer, hints, tips, structure guidance, vocabulary, idioms, useful phrases, or example sentences before the user writes.
3. Wait for the user's Latvian letter.
4. Review the letter using `references/review-guide.md`.
5. Reply with:
   - word count;
   - 1-100 practice score based on the Latvian naturalization exam writing context;
   - task coverage;
   - grammar, spelling, punctuation, and style corrections;
   - progress assessment for this completed task;
   - advice on how to improve the next result;
   - post-task hints, tips, necessary vocabulary, and useful idioms/phrases for writing a stronger letter;
   - a corrected version if useful;
   - your own 50-70-word Latvian version for the same task.

## Before The User Writes

Keep the practice exam-like. Before the user submits a letter, provide only the selected task and a short invitation to answer. Do not coach, outline the answer, suggest phrases, pre-teach vocabulary, or reveal idioms. If the user asks for hints before writing, politely say that hints and vocabulary will come after the completed attempt, then ask them to try first.

## Selecting A Task

Run from the skill directory:

```bash
python3 scripts/select_task.py references/exercises.md
```

To reproduce or inspect a specific task:

```bash
python3 scripts/select_task.py references/exercises.md --number 12
```

Use the script output verbatim when presenting the task. If the selected exercise is dialogue-oriented, still ask the user to write it as a short letter or note unless they explicitly prefer dialogue practice.

## Review Style

Be a supportive Latvian tutor. Keep feedback concise, specific, and actionable. Explain recurring patterns, but avoid overwhelming the user with every possible alternative phrasing.

When offering the 50-70-word version:

- write in clear A2/B1-level Latvian unless the user asks for a higher level;
- cover all required bullet points from the exercise;
- keep the tone appropriate to the recipient;
- count the words and keep the version within 50-70 words.

## References

- `references/exercises.md`: complete "Eksāmena un mācību uzdevumu saraksts" list.
- `references/review-guide.md`: word-counting, scoring, progress, and correction checklist.
