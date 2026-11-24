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

Steading uses PostgreSQL, through pgx, sqlc, and goose, and doesn't aim to be database-agnostic: [you just need Postgres](https://youjustneedpostgres.com). PostgreSQL offers great types, message queueing, full-text search, vector search, unstructured documents, time series, caching, and more. The most value comes from using all of it, not from supporting every database. SQLite is probably enough for a monolith on a home server, but Steading's components get the most from multi-process concurrent reads and writes and transactional guarantees. Every component can be a tenant on one shared PostgreSQL server.[^stack]

Agents write raw SQL, and sqlc generates the package that talks to the store, the way Protocol Buffers generate Connect-RPC bindings. goose runs migrations. There's no query builder.[^stack]

## Ginkgo and Gomega

Tests use Ginkgo and Gomega instead of Go's `testing` package, so a [skift](skift.md) node's requirements become specifications of behavior, and business logic is documented in English. That makes it easier to get clear on what the software should do.[^stack]

## golangci-lint

golangci-lint supports custom rules, such as depguard, which put light guardrails on agents.[^stack] They enforce boundaries agents frequently cross: touching the database directly, crossing layers, or dropping to Go's `testing` package.[^thicket]

## Everything else

- **CLIs:** cobra, pflag, and viper, the best libraries for CLIs in Go.[^stack]
- **Logging:** the standard library's `log/slog`.[^stack]
- **Tools and hooks:** mise manages tool versions, and hk runs Git hooks.[^stack]

[^stack]: [Projects interview: tech stack](../sources/202611150814-projects-in-ivys-words.md#tech-stack)
[^thicket]: [Projects interview: thicket](../sources/202611150814-projects-in-ivys-words.md#steading-aithicket)
