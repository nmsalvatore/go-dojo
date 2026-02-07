# The Go Dojo — 2026 Training Plan

## Student Background
- 5 years hobbyist programming (Python, JS — on and off)
- Knows Go basics: task tracker CLI, basic CRUD web app
- Focus areas: interface design, error handling, testing discipline, package structure

---

## How Each Session Works

Every session follows this structure:

1. **Cold Open Retrieval (5 min)** — From memory, no notes, no looking things up: summarize what you learned last session and one thing you're still uncertain about. This is non-negotiable.
2. **State your week/assignment** — Where are you? What are you working on?
3. **Work** — Build, debug, discuss, get reviewed.
4. **Predict Before You Run** — Before executing any code you write, state what you expect the output to be. When the prediction is wrong, we stop and figure out why.

### Learning Protocols (Active Throughout)

- **Interleaved retrieval** — Mid-task, I'll stop you and ask you to recall a concept from 2-4 weeks ago. This isn't a pop quiz — it's baked into the workflow.
- **Elaborative interrogation** — After every new concept: *Why does Go do it this way? What problem does this solve? What would happen if Go didn't have this feature?*
- **Concrete before abstract** — I introduce topics by showing you concrete code first. You identify the pattern. Then I name the abstraction.
- **Debugging reps** — Every week includes a short debugging exercise (10-15 lines of broken code). This is a separate skill from writing code and gets trained separately.
- **Format variation** — Not every week is "build a thing." Some weeks involve reading others' code, writing explanations, deliberately breaking things, or debugging exercises.

---

## Semester 1: Foundations (Weeks 1–8)
**Goal:** Think in Go. Write idiomatic code without hesitation.
**Graduation criteria:** Can write a multi-package CLI app with proper error handling, interfaces, tests, and file I/O without referencing documentation for basic operations.

| Week | Topic | Assignment | Tests | Reading |
|------|-------|------------|-------|---------|
| 1 | Functions, multiple returns, error handling, error wrapping (`errors.Is`/`As`, `fmt.Errorf` with `%w`) | **Part A:** CLI that computes min/max/mean/median from args. **Part B:** Unit converter CLI — multi-package, proper error handling with wrapping | Write tests for Part B | Effective Go: Functions, Error handling |
| 2 | Structs, methods, interfaces, `io.Reader`/`io.Writer` | Given a working book inventory system, add a feature that *requires* extracting an interface. Build a log parser that works with any `io.Reader` (test with strings, files, HTTP bodies) | Table-driven tests for both | Effective Go: Interfaces, Go Proverbs (read and discuss) |
| 3 | Slices, maps, pointers, type assertions, type switches | Implement a frequency counter, a stack, and a queue from scratch. Include a sub-task using error wrapping from Week 1 | Benchmark your data structures | Effective Go: Data (slices, maps) |
| 4 | Packages, modules, visibility, project layout, `time` package | Refactor weeks 1-3 into a proper multi-package module. Build a small time-tracking utility (durations, formatting, tickers) | Tests for the time utility, interleaved: write one test that exercises your Week 2 interfaces | Effective Go: Package names |
| 5 | Strings, runes, file I/O, JSON and struct tags | Build a config file parser (JSON) with validation. Read and annotate a section of the `encoding/json` stdlib source | Tests with malformed input. Interleaved: use `io.Reader` from Week 2 | Effective Go: Commentary (how to write Go docs) |
| 6 | Dependency injection, testability, designing for tests | Refactor your Week 5 config parser to use injected dependencies. Write a version that's fully testable with no file system access | Full test coverage using DI patterns. Interleaved: table-driven tests from Week 2 | — |
| 7 | **Consolidation week** | Revisit and refactor Weeks 1-6 code with fresh eyes. Rate your confidence on each topic (skills matrix). Fill gaps. Debugging exercise: I give you broken code from every prior week's topic | Cumulative mini-assessment: one retrieval question per week, one short code exercise | Review your own code from Week 1 — identify what you'd change |
| 8 | **Capstone** | **Flashcard CLI** — deck management, spaced repetition (Leitner or SM-2), progress tracking, persist to disk, multi-package, injected dependencies, full test suite | Comprehensive tests. I review line-by-line and grill you on every decision | — |

### Semester 1 Assessment
- Line-by-line capstone code review
- Debugging exercise: find and fix bugs in broken Go code
- "Explain it back": error handling philosophy in Go vs. exceptions — in your own words
- Skills matrix self-assessment: rate confidence 1-5 on each topic

---

## Semester 2: Intermediate Systems (Weeks 9–20)
**Goal:** Build real networked software. Understand concurrency deeply.
**Graduation criteria:** Can build a concurrent HTTP service with database persistence, middleware, proper context usage, and no race conditions.

| Week | Topic | Assignment | Tests | Reading |
|------|-------|------------|-------|---------|
| 9 | Goroutines, channels, nil channels, channel closing semantics | Implement a concurrent file hasher (hash N files in parallel). Includes debugging exercise: find the deadlock | Tests with race detector (`-race`). Interleaved: use interfaces from Semester 1 | "Concurrency is not Parallelism" (Rob Pike talk) |
| 10 | `select`, buffered channels, channel direction | Build a pipeline: reader → processor → writer using channels. Predict output of 3 channel code snippets before running | Pipeline tests. Interleaved: error wrapping in pipeline stages | Effective Go: Concurrency |
| 11 | `sync` package: Mutex, WaitGroup, Once | Write a thread-safe in-memory cache with expiration. Debugging exercise: I give you code with a race condition — find it with and without `-race` | Concurrent access tests. Interleaved: use DI so cache storage is swappable | — |
| 12 | Context, cancellation, timeouts — full deep dive | Add timeout/cancellation to your pipeline from Week 10. Implement graceful shutdown. Write an explanation (1 page) of how context propagation works | Tests for cancellation behavior. Interleaved: use your cache from Week 11 | `context` package source code (read and annotate) |
| 13 | **Consolidation week** | Revisit Weeks 9-12. Refactor with fresh eyes. Concurrency debugging gauntlet (3 broken programs). Skills matrix update. Retrieval: explain interfaces, error wrapping, DI from Semester 1 | Cumulative assessment: concurrency + Semester 1 recall | — |
| 14 | `encoding/json` deep dive, custom marshalers, struct tags, `net/http` basics | Build a JSON REST API for your flashcard system. Custom JSON marshaling for complex types | API tests with `httptest`. Interleaved: table-driven tests, context usage | `net/http` package overview |
| 15 | Middleware, routing patterns, request validation | Add logging middleware, basic auth, rate limiting to your API. Read and critique a popular Go HTTP middleware library | Middleware tests. Interleaved: use your concurrent cache from Week 11 as a rate limit store | — |
| 16 | Database access with `database/sql`, Docker basics | Swap file storage for SQLite. Containerize your API with Docker. Set up a basic Makefile | Integration tests. Interleaved: DI for database layer (swap between file and DB) | — |
| 17 | API client design, HTTP client patterns, retries, timeouts | Build an API client that consumes a public REST API. Implement retry with backoff | Client tests with mock server. Interleaved: context timeouts from Week 12 | — |
| 18 | Error strategy audit, sentinel errors, error types | Audit all Semester 1-2 code — implement a consistent error strategy across everything. Write an explanation of your error philosophy | — | — |
| 19 | Review, catch-up, prep | Fill gaps, polish code, set up CI for your project (GitHub Actions) | — | — |
| 20 | **Capstone** | **URL shortener service** — HTTP API, SQLite storage, stats tracking, rate limiting, Dockerized, CI pipeline, full test suite | Comprehensive tests including integration. I review architecture and code | — |

### Semester 2 Assessment
- Code review focused on concurrency bugs ("what happens if...")
- Debugging challenge: find a race condition in code I provide
- Architecture review: diagram your URL shortener and defend your choices
- Timed challenge: build an HTTP handler from scratch in 30 minutes
- Skills matrix update

---

## Semester 3: Advanced Patterns (Weeks 21–30)
**Goal:** Write production-grade Go. Design clean APIs. Publish real software.
**Graduation criteria:** Can profile and optimize a service, design a clean public API, and publish a tested, documented Go library.

| Week | Topic | Assignment | Tests | Reading |
|------|-------|------------|-------|---------|
| 21 | Embedding, composition patterns, advanced interface design | Design a plugin system using interfaces and composition. Read and critique an open-source Go project's interface design | Tests for plugin system. Interleaved: recall concurrency patterns | Effective Go: Embedding |
| 22 | Generics — type parameters, constraints | Start with 3 duplicate functions that differ only by type. Identify the pattern. Rewrite cache and data structures with generics | Generic data structure tests. Interleaved: benchmark vs. non-generic versions | Go blog: intro to generics |
| 23 | Profiling: pprof (CPU, memory, goroutine), load testing | Profile your URL shortener under load, find and fix bottlenecks. Write a report on what you found and how you fixed it | Before/after benchmarks. Interleaved: context and graceful shutdown | — |
| 24 | Reflection, code generation (compressed) | Study how `encoding/json` uses reflection. Write a simple struct pretty-printer. Use `go generate` and `stringer` for enums | Tests for pretty-printer. Interleaved: interface design from Week 21 | `reflect` package docs (skim, not memorize) |
| 25 | Package design, API design, documentation | **Publish milestone:** Extract a reusable library from your projects. Publish to GitHub with docs, tests, CI, and a proper README. This is real — someone should be able to `go get` it | Full test suite, CI passing, godoc complete | — |
| 26 | **Consolidation week** | Revisit Semester 3 work. Skills matrix update. Retrieval gauntlet: concepts from all 3 semesters. Read and critique a function from the Go standard library | Cumulative assessment | — |
| 27 | Integration testing, `testcontainers`, test helpers | Write integration tests for your URL shortener that spin up real databases in Docker | Integration test suite. Interleaved: DI patterns, interface design | — |
| 28 | Failure thinking: retries, circuit breakers, timeouts, degradation | Add failure handling to your URL shortener — what happens when the DB is down? When a downstream service is slow? | Chaos tests (deliberately break things). Interleaved: context, error strategy | — |
| 29 | Review, catch-up, prep | Fill gaps, polish projects, prep for capstone | — | — |
| 30 | **Capstone** | **Concurrent web crawler** — depth limits, rate limiting per domain, `robots.txt` respect, deduplication, exportable results, graceful shutdown | Full test suite including concurrency tests. I introduce subtle bugs — you find them via testing/profiling | — |

### Semester 3 Assessment
- "Teach it back": explain goroutine scheduling or the GC as if to a junior dev
- Standard library reading: walk through a function I select from the Go stdlib
- Refactoring challenge: ugly-but-working code → idiomatic Go
- Bug hunt in your crawler

---

## Semester 4: Mastery (Weeks 31–40)
**Goal:** Ship software that others depend on. Operate in a professional context.
**Graduation criteria:** Can design, build, deploy, and operate a production-grade Go service with observability, real-time features, and professional CI/CD.

| Week | Topic | Assignment | Tests | Reading |
|------|-------|------------|-------|---------|
| 31-32 | gRPC and Protocol Buffers | Build a microservice with gRPC, including streaming. Interleaved: integration tests, Docker, CI | gRPC tests, integration tests | gRPC Go quickstart |
| 33-34 | WebSockets, real-time systems | Add real-time features to an existing project (live updates, presence) | WebSocket tests. Interleaved: concurrency, context | — |
| 35-36 | Worker pools, job queues, backpressure | Build a job queue with configurable worker pools, backpressure, and graceful shutdown | Load tests, failure tests. Interleaved: profiling from Week 23 | — |
| 37-38 | Observability: structured logging, metrics, tracing | Instrument your services with OpenTelemetry. Structured logging throughout. Dashboards | Verify instrumentation under load. Interleaved: everything | — |
| 39-40 | **Final Capstone** | Choose one (see below) | Comprehensive | — |

### Final Capstone Options
- **A:** Real-time chat system (WebSocket server, persistence, rooms, presence, graceful shutdown, observability)
- **B:** CLI tool that solves a real problem you have — published to GitHub with docs, tests, CI, releases, and real users if possible
- **C:** Full-stack service: gRPC backend + REST gateway + worker queue + observability — a mini production system
- **D:** Propose your own (subject to approval)

### Final Assessment
- Full code review simulating a real company PR review
- Architecture walkthrough as if onboarding a new team member
- Edge case and failure scenario reasoning — live
- Read and critique a piece of open-source Go code I select
- Final skills matrix — compare to your Week 1 self-assessment

---

## Progress Tracking

### Skills Matrix
Updated at Weeks 7, 13, 20, 26, 30, and 40. Rate yourself 1-5 on each skill:

| Skill | W7 | W13 | W20 | W26 | W30 | W40 |
|-------|----|-----|-----|-----|-----|-----|
| Functions & error handling | | | | | | |
| Structs, methods, interfaces | | | | | | |
| Slices, maps, pointers | | | | | | |
| Package design & layout | | | | | | |
| Testing & benchmarks | | | | | | |
| File I/O, JSON, strings | | | | | | |
| Dependency injection | | | | | | |
| Goroutines & channels | | | | | | |
| sync, context, cancellation | | | | | | |
| HTTP servers & middleware | | | | | | |
| Database access | | | | | | |
| Docker & CI/CD | | | | | | |
| Generics | | | | | | |
| Profiling & optimization | | | | | | |
| API & package design | | | | | | |
| Failure handling | | | | | | |
| gRPC | | | | | | |
| Real-time systems | | | | | | |
| Observability | | | | | | |

---

## Rules of the Dojo

1. **Cold open every session** — summarize last session from memory before anything else
2. **Predict before you run** — state expected output before executing code
3. Push your code for review — feedback will be honest, not gentle
4. Ask questions anytime — but expect questions back, especially "why"
5. If stuck for 30+ minutes, come ask — struggling is good, suffering isn't
6. Progress is tracked — the plan adapts if gaps appear
7. Every assignment includes tests — testing is a habit, not a phase
8. You can take breaks between semesters — momentum matters more than speed
