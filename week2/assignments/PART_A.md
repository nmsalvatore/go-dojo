# Week 2, Part A: Book Inventory — Interface Extraction

## Setup

Starter code is in `week2/inventory/`. It is a working book inventory system. Read all of it before you change anything.

The system currently:
- Stores books (title, author, year, genre, price)
- Supports adding, removing, and listing books
- Can search books by title or author
- Exports the full inventory to a JSON file on disk

## Your Task

Add a **notification feature**: whenever a book is added or removed, the system should notify someone. The catch — the system needs to support multiple notification methods:

1. Console output (print to stdout)
2. Write to a log file
3. (In tests) Capture notifications in memory so you can assert on them

You must NOT use `if/else` or `switch` on a notification type string. The design must use an interface.

## Requirements

- Identify the correct interface yourself. I am not telling you what it should look like.
- The `Inventory` struct should not know or care which notification method is being used.
- The interface should be small. If it has more than 2 methods, you've over-designed it.
- Write table-driven tests that verify notifications are sent on add and remove.
- Your test notifications must NOT touch the file system or print to stdout.

## What I'm Looking For

- Can you identify where the abstraction boundary belongs?
- Is your interface minimal?
- Does your inventory depend on the interface, not the implementation?
- Are your tests clean and independent of real side effects?

## Starter Code

The starter code will be provided in `week2/inventory/`. Read it first. Understand it. Then plan your changes before writing any code.
