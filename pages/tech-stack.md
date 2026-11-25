---
type: topic
summary: The one stack every Steading component is built on, and why.
---
# Tech stack

Every Steading component is built on the same core stack. Consistency makes refactoring easier, keeps the same tough decisions from being made twice, and leads to software defined by a specification of the behavior it should have, instead of everything being reinvented.[^stack]

What you build on top of Steading is up to you. Its interfaces are Connect-RPC, so anyone can build their own component, in any stack, by implementing them.[^stack]

## Go

Go compiles fast and produces a [single static binary](principles.md#composable). The language is easy to learn, its tooling is unmatched, and it is type-checked, memory-safe, and garbage-collected. It is steeped in distributed systems and orchestration software such as Docker, Terraform, and Kubernetes, and `go fmt` leaves little to bikeshed.[^stack]

## Connect-RPC

Connect-RPC is language-agnostic and type-checked, works in the browser and on mobile, and is easy to understand. Its bindings are far nicer than gRPC's, and it avoids the bikeshedding that comes with designing REST APIs.[^stack]

## PostgreSQL

Steading uses PostgreSQL, through pgx, sqlc, goose, and River. PostgreSQL is a commitment, not an implementation detail: Steading doesn't aim to be database-agnostic, and no component keeps a second storage backend possible. [You just need Postgres](https://youjustneedpostgres.com). What a component keeps fixed is its Connect-RPC contract and its tenancy, not its database.[^stack][^postgres]

PostgreSQL offers great types, message queueing, full-text search, vector search, unstructured documents, time series, caching, and more. The most value comes from using all of it, not from supporting every database.[^stack] Use its features fully, such as views, SQL functions, recursive CTEs, row locking, advisory locks, LISTEN/NOTIFY, and custom plans, and aim for queries that are performant, expressive, easy to understand, and clean.[^postgres]

SQLite is probably enough for a monolith on a home server, but Steading's components get the most from multi-process concurrent reads and writes and transactional guarantees. Every component can be a tenant on one shared PostgreSQL server.[^stack]

### Packages

Each component splits its database code across three packages:

- **`store`** owns PostgreSQL: every query, LISTEN, and River's client. Only the store talks to PostgreSQL.[^postgres][^migrations]
- **`domain`** owns the core rules and use cases, held as decisions. It never writes SQL.[^postgres][^migrations]
- **`admin`** owns migrations, bootstrap, and background jobs, and decides when and how to open the database, through the store's constructor. Its operator commands live under the CLI's `admin` subcommand.[^migrations]

The boundaries are structural, not left to review. The store doesn't export its raw transaction, so SQL outside the store fails to compile.[^postgres] [depguard](#golangci-lint) enforces the rest:[^postgres][^migrations]

- `domain` and `admin` may import `store`.
- pgx and River's PostgreSQL drivers are allowed only in `store`.
- River's core package, for workers and the periodic schedule, is allowed only in `store` and `admin`; its APIs carry no pgx types.

An exception doesn't get through with a written rationale, since an agent can write a plausible rationale for anything.[^postgres]

### Writing SQL

Agents write raw SQL, and there's no query builder.[^stack] sqlc is the default for plain queries: it generates Go from SQL the way Protocol Buffers generate Connect-RPC bindings.[^stack][^postgres] It's a tool to use where it fits, not a rule to follow everywhere. SQL that uses PostgreSQL features sqlc can't parse lives in `.sql` files inside the store, loaded with `go:embed`. SQL never goes in a Go string literal.[^postgres]

A rule that several queries share is written once, in the schema, as a view or SQL function, so the database holds it. It's never a materialized view or state kept up by triggers: that's a cached derivation, and it drifts. Changing such a rule takes a migration, and a committed schema dump keeps its current definition readable.[^postgres]

### Migrations

goose runs migrations.[^stack] Their safety comes from tools, not a style guide:[^migrations]

- **Lint:** squawk checks each new migration for statements that take unsafe locks, as an [hk](#hk) step.
- **Code:** goose's session lock keeps two migrations from running at once, the migration connection sets `lock_timeout` and retries a bounded number of times, and a binary refuses a schema newer than itself.
- **Tests:** an upgrade from every released schema, with data in every table, and a check that shipped migrations never change.
- **Review** covers only what tools can't judge: what a default means for existing rows, how big a backfill is, and data that only running code wrote.

The database schema isn't part of a component's [public API](releases.md#what-the-version-covers); [upgrading](releases.md#upgrading) says what a release promises about migrations.[^migrations]

### Background work

River runs a component's own background work as periodic jobs inside its server, with leader election, so a component doesn't need systemd timers, which compose and launchd installs don't have. River keeps its queue in PostgreSQL, and its schema is embedded as goose migrations, so there's one migration system.[^migrations]

## Ginkgo and Gomega

Tests use Ginkgo and Gomega instead of Go's `testing` package, so a [skift](skift.md) node's requirements become specifications of behavior, and business logic is documented in English. That makes it easier to get clear on what the software should do.[^stack]

### Tests run in parallel

Agents work in parallel git worktrees on one machine and run test suites at the same time, so one test run never interferes with another. Everything a run creates gets a name of its own for that run, and is cleaned up afterwards, even when the run fails:[^parallel]

- databases;
- ports: bind to port 0 or to a free port, never a fixed one;
- socket paths;
- compose projects;
- containers, images, volumes, and networks.

Nothing a test uses sits at a shared, well-known name. Test runs may share the podman daemon, but not names within it, such as compose projects and image tags.[^parallel]

## golangci-lint

golangci-lint supports custom rules, such as depguard, which put light guardrails on agents.[^stack] They enforce boundaries agents frequently cross: touching the database directly, crossing layers, or dropping to Go's `testing` package.[^thicket] depguard holds [the store's boundary](#packages).[^postgres]

## hk

hk runs every check and the Git hooks.[^stack] A component's `hk.pkl` lists each tool as its own step, using hk's builtins where they exist, such as `golangci_lint`, `go_vet`, and `buf_lint`, so hk runs the steps in parallel, each on its own files. `hk check --all` runs every check, and CI runs the same command. A `fix` hook applies autofixes. hk doesn't wrap another runner such as `make lint`: behind a single step, hk adds nothing.[^postgres] Its steps also lint new migrations with squawk and check the API against the last release with [`buf breaking`](releases.md#guarding-the-api).[^migrations]

Markdown is linted too: one paragraph per line, never hard-wrapped.[^postgres]

## Everything else

- **CLIs:** cobra, pflag, and viper, the best libraries for CLIs in Go.[^stack]
- **Logging:** the standard library's `log/slog`.[^stack]
- **Tool versions:** mise.[^stack]

[^stack]: [Projects interview: tech stack](../sources/202611150814-projects-in-ivys-words.md#tech-stack)
[^postgres]: [PostgreSQL and hk across the stack](../sources/202611250838-postgres-and-hk.md)
[^migrations]: [Releases, migrations, and background work](../sources/202611251316-releases-and-migrations.md)
[^parallel]: [Tests that run in parallel](../sources/202611251958-parallel-test-runs.md)
[^thicket]: [Projects interview: thicket](../sources/202611150814-projects-in-ivys-words.md#steading-aithicket)
