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

Steading uses PostgreSQL, through pgx, sqlc, and goose. PostgreSQL is a commitment, not an implementation detail: Steading doesn't aim to be database-agnostic, and no component keeps a second storage backend possible. [You just need Postgres](https://youjustneedpostgres.com). What a component keeps fixed is its Connect-RPC contract and its tenancy, not its database.[^stack][^postgres]

PostgreSQL offers great types, message queueing, full-text search, vector search, unstructured documents, time series, caching, and more. The most value comes from using all of it, not from supporting every database.[^stack] Use its features fully, such as views, SQL functions, recursive CTEs, row locking, advisory locks, LISTEN/NOTIFY, and custom plans, and aim for queries that are performant, expressive, easy to understand, and clean.[^postgres]

SQLite is probably enough for a monolith on a home server, but Steading's components get the most from multi-process concurrent reads and writes and transactional guarantees. Every component can be a tenant on one shared PostgreSQL server.[^stack]

### The store owns the SQL

In each component, only the store package talks to PostgreSQL. It owns every query, the connections, and LISTEN. The domain holds the rules as decisions and never writes SQL.[^postgres]

The boundary is structural, not left to review. The store doesn't export its raw transaction, so SQL outside the store fails to compile, and [depguard](#golangci-lint) allows pgx only in the store. An exception doesn't get through with a written rationale, since an agent can write a plausible rationale for anything.[^postgres]

### Writing SQL

Agents write raw SQL, and there's no query builder.[^stack] sqlc is the default for plain queries: it generates Go from SQL the way Protocol Buffers generate Connect-RPC bindings.[^stack][^postgres] It's a tool to use where it fits, not a rule to follow everywhere. SQL that uses PostgreSQL features sqlc can't parse lives in `.sql` files inside the store, loaded with `go:embed`. SQL never goes in a Go string literal.[^postgres]

A rule that several queries share is written once, in the schema, as a view or SQL function, so the database holds it. It's never a materialized view or state kept up by triggers: that's a cached derivation, and it drifts. Changing such a rule takes a migration, and a committed schema dump keeps its current definition readable.[^postgres] goose runs migrations.[^stack]

## Ginkgo and Gomega

Tests use Ginkgo and Gomega instead of Go's `testing` package, so a [skift](skift.md) node's requirements become specifications of behavior, and business logic is documented in English. That makes it easier to get clear on what the software should do.[^stack]

## golangci-lint

golangci-lint supports custom rules, such as depguard, which put light guardrails on agents.[^stack] They enforce boundaries agents frequently cross: touching the database directly, crossing layers, or dropping to Go's `testing` package.[^thicket] depguard holds [the store's boundary](#the-store-owns-the-sql).[^postgres]

## hk

hk runs every check and the Git hooks.[^stack] A component's `hk.pkl` lists each tool as its own step, using hk's builtins where they exist, such as `golangci_lint`, `go_vet`, and `buf_lint`, so hk runs the steps in parallel, each on its own files. `hk check --all` runs every check, and CI runs the same command. A `fix` hook applies autofixes. hk doesn't wrap another runner such as `make lint`: behind a single step, hk adds nothing.[^postgres]

Markdown is linted too: one paragraph per line, never hard-wrapped.[^postgres]

## Everything else

- **CLIs:** cobra, pflag, and viper, the best libraries for CLIs in Go.[^stack]
- **Logging:** the standard library's `log/slog`.[^stack]
- **Tool versions:** mise.[^stack]

[^stack]: [Projects interview: tech stack](../sources/202611150814-projects-in-ivys-words.md#tech-stack)
[^postgres]: [PostgreSQL and hk across the stack](../sources/202611250838-postgres-and-hk.md)
[^thicket]: [Projects interview: thicket](../sources/202611150814-projects-in-ivys-words.md#steading-aithicket)
