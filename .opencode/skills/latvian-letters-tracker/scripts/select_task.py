#!/usr/bin/env python3
"""Select a Latvian letter-writing exercise from the bundled list."""

from __future__ import annotations

import argparse
import random
import re
from pathlib import Path


TASK_RE = re.compile(r"(?m)^(\d{1,2})\.\s")


def parse_tasks(text: str) -> dict[int, str]:
    matches = list(TASK_RE.finditer(text))
    tasks: dict[int, str] = {}
    for index, match in enumerate(matches):
        number = int(match.group(1))
        start = match.start()
        end = matches[index + 1].start() if index + 1 < len(matches) else len(text)
        tasks[number] = text[start:end].strip()
    return tasks


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("exercise_file", type=Path)
    parser.add_argument("--number", type=int, help="Print a specific exercise number.")
    parser.add_argument("--seed", type=int, help="Use a deterministic random seed.")
    args = parser.parse_args()

    text = args.exercise_file.read_text(encoding="utf-8")
    tasks = parse_tasks(text)
    if not tasks:
        raise SystemExit("No numbered exercises found.")

    if args.number is not None:
        if args.number not in tasks:
            available = f"{min(tasks)}-{max(tasks)}"
            raise SystemExit(f"Exercise {args.number} not found. Available range: {available}.")
        number = args.number
    else:
        rng = random.Random(args.seed)
        number = rng.choice(sorted(tasks))

    print(tasks[number])


if __name__ == "__main__":
    main()
