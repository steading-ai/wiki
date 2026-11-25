---
participants: [ivy]
---

# PostgreSQL and hk across the stack

A working session in skift in which ivy settled how Steading components use PostgreSQL and hk, with an agent that had found skift's code and docs drifting from the [tech stack](../pages/tech-stack.md). What follows is what ivy decided, what was passed over and why, and what stayed open.

## PostgreSQL is a commitment

PostgreSQL is a commitment, not an implementation detail, and database agnosticism is off the table. The corollary is to use PostgreSQL's features fully, such as views, SQL functions, recursive CTEs, row locking, advisory locks, LISTEN/NOTIFY, and custom plans, aiming for queries that are performant, expressive, easy to understand, and clean. What a component keeps fixed is its Connect contract and its tenancy, not its database.

skift's docs claimed otherwise: its core principle 8, "Storage, transport, and identity are swappable behind interfaces", and its architecture doc's "keeps a second storage backend possible". Both are being amended in skift.

## Only the store talks to PostgreSQL

skift's domain layer had grown hand-written SQL, through a public field exposing the raw pgx transaction, and a small query builder. The boundary that holds across the stack:

- The store package owns all SQL, the connections, and LISTEN.
- The domain holds the rules as decisions and never writes SQL.
- Enforcement is structural, not reviewer vigilance. The raw transaction is unexported, so SQL outside the store fails to compile, and depguard allows pgx only in the store.

Passed over: a lint that judges whether sqlc could have expressed a query, since no lint can decide that; and a written `nolint` rationale at each exception, because agents write plausible rationales for anything.

## sqlc is a tool

ivy on sqlc: don't cargo-cult it. sqlc stays the default for plain queries. SQL that uses PostgreSQL features sqlc can't parse lives in `.sql` files inside the store, loaded with `go:embed`. SQL never goes in Go string literals.

A rule stated in several queries, such as skift's readiness conditions, is defined once in the schema as a view or SQL function, so the database enforces it being written once. Never as a materialized view or trigger-maintained state, which is cached derivation that drifts. The accepted cost: changing such a rule takes a migration, and a committed schema dump keeps its current definition readable.

## hk is the check runner

skift's hk ran one step, `make lint`, which ran eight tools in series, and CI called `make lint` directly, so hk added nothing. The shape that holds, following hk's own CI and hooks docs:

- `hk.pkl` lists each tool as its own step, using hk's builtins where they exist, such as `golangci_lint`, `go_vet`, `buf_lint`, `jactionlint`, and `rumdl`.
- hk runs the steps in parallel and picks the files for each.
- `hk check --all` runs every check, and CI runs the same command.
- A `fix` hook applies autofixes.

Markdown is linted too: one paragraph per line, no hard wrapping. skift had 888 hard-wrapped lines across 20 files. rumdl, through hk's builtin, is the candidate check.

## What stayed open

- hk 2.13 ships `hk agent hooks --target claude-code`, a Stop hook that runs `hk run check --safe` on changed files before an agent may finish. skift is evaluating it, pending a measurement of how fast it is; whether it ships repo-wide is undecided.
- The performance target for readiness rules stated in the schema. A spike will measure them with `EXPLAIN ANALYZE` at the scale skift's vision sets.
