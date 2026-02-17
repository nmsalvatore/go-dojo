# Week 2, Part B: Log Parser — io.Reader

## Your Task

Build a log parser that reads structured log lines and extracts information from them.

The parser must accept an `io.Reader` as its input source. It must NOT accept a filename, a string, or a byte slice directly. The caller wraps their data source in an `io.Reader` and hands it to your parser.

## Log Format

Each line follows this format:

```
2026-02-17T14:30:00Z [LEVEL] message text here
```

Levels are: `INFO`, `WARN`, `ERROR`.

Lines that don't match this format should be skipped (not cause an error).

## Requirements

### Core
- Parse log lines into structured data (timestamp, level, message)
- Filter entries by level (e.g., "give me only ERRORs")
- Count entries by level
- Return all entries between two timestamps

### Design
- The parser takes an `io.Reader`. Nothing else.
- No global state. No package-level variables.
- Proper error handling — use your Week 1 skills. If a custom error type makes sense, use one.

### Tests
- Table-driven tests
- Test with `strings.NewReader` for inline test data
- Test with an actual file (`os.Open` returns an `io.Reader`)
- Test with malformed input — blank lines, garbage data, missing fields
- Test the timestamp range filter with edge cases

## What I'm Looking For

- Do you understand why `io.Reader` is powerful here?
- Can you write code that genuinely doesn't care where its input comes from?
- Are your tests thorough, especially around bad input?
- Is your package structure clean?

## Hint

`bufio.Scanner` wraps an `io.Reader` and gives you line-by-line reading. You'll want it.
