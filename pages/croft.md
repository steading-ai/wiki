---
type: project
status: concept
repository: steading-ai/croft
summary: One interface for executing tasks anywhere.
---
# croft

croft is Steading's backend-agnostic sandbox runtime: one interface for executing tasks anywhere.[^croft] It is the fence that contains your agents. It compartmentalizes them and sets them up to work wherever you mean to run them: an ephemeral VM on Daytona, an EC2 or Firecracker instance, a Docker container, a self-hosted VM or container, or bare metal.[^croft] It is meant for other kinds of computing tasks too, not only agents.[^croft]

## Status

croft doesn't exist yet. [mill](mill.md) holds a lot of its responsibilities, and the gaps croft is meant to close are in mill: one Linux server can't build Mac or iOS apps, crashes building an Unreal Engine game, and can't give a graphics card to parallel projects that each need one for tests or local models.[^croft] Harness bindings are implemented in various other projects.[^croft]

## Why

Docker, Podman, Incus, KVM, and Kubernetes each solve shipping and running images in their own way. croft doesn't remake a container or VM runtime; it offers a single interface for executing tasks on any of them.[^croft] It is motivated by skyrocketing hardware costs, which make self-hosting hard to sustain when work bursts past the resources you have.[^croft] You set up backends, and croft works out where each task runs and with what degree of isolation.[^croft]

Budgets are out of scope for croft. The underlying providers have their own limits, budgets, and handling for resource exhaustion, and are the likelier place to decide them.[^croft]

## Harness bindings

croft ships the harness bindings: the communication layer around agent harnesses like Claude Code, OpenCode, Pi, and Prime. The bindings speak ACP or another protocol for communicating with agents.[^croft] [The gateway](gateway.md) may live in croft for a while.[^croft]

[^croft]: [Projects interview: croft](../sources/202611150814-projects-in-ivys-words.md#steading-aicroft)
