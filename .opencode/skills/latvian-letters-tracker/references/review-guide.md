# Latvian Letter Review Guide

## Word Count

- Count Latvian letter words as whitespace-separated lexical tokens after ignoring punctuation around them.
- Count greetings, closings, names, standalone numbers, and dates as words if they appear in the user's answer.
- Do not count punctuation-only marks, bullets, or empty lines.
- State the total and whether it is below, within, or above the requested range when a range is provided.

## Task Coverage

Check whether the answer covers each required point in the selected exercise. Mention missing points briefly and suggest what to add.

## Naturalization Exam Scoring Context

Use this as an approximate practice score, not an official exam result. The official naturalization language test checks writing as one of four language skills. Under Cabinet Regulation No. 973, each skill is assessed out of 25 points and the language test is passed when the candidate gets at least 16 of 25 for each skill. PMLP describes the writing task as a short letter about an everyday topic.

Sources checked on 2026-07-06:

- PMLP, "Pilsonības likumā noteiktās pārbaudes": https://www.pmlp.gov.lv/lv/pilsonibas-likuma-noteiktas-parbaudes
- Cabinet Regulation No. 973 on Likumi.lv: https://likumi.lv/doc.php?id=260430

For user-facing practice, always report a 1-100 point score. Convert it to an estimated exam-style writing score in parentheses when useful:

```text
Practice score: 72/100 (about 18/25; likely pass level for writing practice)
```

Use `round(practice_score / 4)` for the approximate 25-point conversion.

Use this 100-point practice rubric:

- Task fulfillment and content: 30 points. Award for covering all prompt bullets, giving concrete details, and making the purpose clear.
- Organization and letter format: 15 points. Award for greeting, logical order, paragraphing, closing, and appropriate formal or informal register.
- Grammar: 25 points. Award for correct cases, agreement, verb forms, prepositions, negation, and sentence structure.
- Vocabulary and naturalness: 15 points. Award for appropriate everyday vocabulary, collocations, and avoiding Russian or English calques.
- Spelling, diacritics, and punctuation: 15 points. Award for garumzīmes, mīkstinājuma zīmes, commas, capitalization, and readable mechanics.

Interpret scores:

- 86-100: strong exam-ready response;
- 70-85: likely pass level, with focused improvements needed;
- 64-69: near pass threshold, but fragile;
- 50-63: understandable but below stable pass level;
- 1-49: major content or language gaps.

Do not overstate certainty. Use phrases like "practice estimate" and "likely" because only the exam commission gives official results.

## Correction Checklist

Prioritize errors that affect exam readability:

- spelling and diacritics: garumzīmes, mīkstinājuma zīmes, missing or extra letters;
- noun and adjective agreement: gender, number, case;
- verb forms: tense, person, reflexive forms, negation with `ne-`;
- prepositions and cases: `uz`, `pie`, `ar`, `par`, `no`, `līdz`, `pēc`, `priekš`;
- word order and sentence completeness;
- punctuation: commas in subordinate clauses, addresses, greetings, and lists;
- register: formal letters to institutions versus informal letters to friends or family;
- vocabulary choice and calques from Russian or English.

## Feedback Format

Use this structure after the user submits a letter:

```markdown
Word count: N words.

Practice score: N/100 (about M/25; ...).

Task coverage: ...

Corrections:
- Original -> Better: short reason.
- Original -> Better: short reason.

Progress assessment:
...

Advice for next time:
- ...
- ...

Post-task writing tips, vocabulary, and idioms:
- Tips: ...
- Useful vocabulary: `...` - ...
- Useful phrases/idioms: `...` - ...

Corrected version:
...

My 50-70-word version:
...
```

If there are many errors, group them by pattern instead of listing every occurrence. Preserve the user's intended meaning whenever possible.

## Progress Assessment

After every completed task, include a brief progress assessment. If earlier attempts exist in the current conversation, compare against them using score, word count, task coverage, and recurring error patterns. If no prior attempts exist, mark the result as the baseline.

Track progress in the conversation, not in files, unless the user explicitly asks to save a log. Do not invent previous scores or long-term history.

In the assessment, identify:

- current level: exam-ready, likely pass, near pass, developing, or needs foundation work;
- strongest area in this response;
- one or two highest-impact weaknesses;
- one concrete practice target for the next letter.

Advice must be specific and immediately usable, for example: "Before writing, list the five required points and check them off," or "Practice prepositions with cases: `uz + accusative`, `pie + genitive`, `ar + accusative/instrumental meaning`."

## Post-Task Tips, Vocabulary, And Idioms

Provide all hints, tips, necessary vocabulary, idioms, useful phrases, and strategy advice only after the user has completed the task. Do not provide this material together with the initial exercise prompt.

After reviewing the completed letter, add a compact post-task learning block:

- 2-4 writing tips tailored to the selected task and the user's mistakes;
- 5-10 useful Latvian words or collocations for that task, with short English explanations if helpful;
- 3-6 natural phrases, connectors, or idioms that could improve the letter;
- one short note on when each formal or informal phrase is appropriate.

Keep vocabulary exam-useful and realistic. Prefer practical phrases such as `Es vēlētos Jūs informēt...`, `Lūdzu paziņojiet...`, `Jau iepriekš pateicos par palīdzību`, `Priecāšos Jūs redzēt`, `Ar cieņu`, and `Sirsnīgi sveicieni` over ornate expressions.

## Model Answer Rules

- Write 50-70 words unless the user asks for a different length.
- Use natural Latvian with appropriate greetings and closing.
- Cover every exercise bullet point.
- Keep vocabulary practical and exam-friendly.
- After the model answer, include the word count in parentheses.
