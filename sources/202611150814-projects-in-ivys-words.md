---
participants: [ivy]
---

# Steading's projects, in ivy's words

An agent interviewed ivy, one repository at a time, about what each Steading project is, what it does, and why it exists, then about the principles they share, and ivy's corrections while proofreading the pages drafted from it. The answers are ivy's exact words, kept verbatim at ivy's request; the questions are the agent's.

## steading-ai/steading

### Round 1

**Interviewer:** In your own words: what is it, what does it do, and why did you make it? And how does it relate to `stead.ing`, both the repository and the domain, and to "Steading" as the organization?

**ivy:** I created steading-ai/steading when I realized I had a need for observability in the form of a web UI. I was experimenting with figuring out a UI that ties the entire suite of software together and it seemed like the easiest way to get started rather than building a bespoke UI to each one. I intentionally chose monolith to prove out the concept. It's basically a Connect-RPC client and React frontend. The unique things:

- The main point of input from a human is an inbox. It collects things needing your action. An agent flags that it needs your attention or decision or taste. You discuss with a thicket agent (we'll talk about that later)--in short, an assistant with tools to manipulate the backlog, investigate, gather details.
- A UI to view metrics from past mill runs (more on that later)--so you took your backlog and had a fleet of agents implement your vision. What did it cost? What models and skills did they use? What was the success rate? How much did it cost in terms of money and tokens?
- A UI to view active and past runs, drill down into the details. It live streams output from the running agent harnesses so you can inspect and if you need to, steer from a side panel with a thicket agent.

I didn't implement any form of backlog visibility. DAGs are difficult to reason about from a human perspective and I didn't get there. I actually found it extremely useful on its own without any form of UI. That said, it goes a long way towards building trust and I think the bigger value is having that assistant with tools to help you talk it out.

I didn't achieve the full vision and I have reservations about continuing on it in its current form. I think it would be better to decompose it into a reusable package of UI components and assets to allow the individual components to serve their pieces of the UI independently. There's questions on my mind about embedding... it creates a new challenge of how to render when someone only has a portion of the system available and being able to interlink and render with the best usability becomes a challenge when it's all divided up.

Now, I've been thinking about the brand, the overall vision, and the potential of building a community of people who want exactly this system and architecture. That, coupled with some recent learnings on how to make more intuitive UI... I kind of want to rebuild it.

### Round 2

**Interviewer:**

1. You describe two directions: decomposing it into a reusable package of UI components that each project serves itself, and rebuilding it, informed by the brand, the community, and what you've learned about UI. Is the rebuild the decomposition, a new monolith, or undecided?
2. Right now, should someone reading `steading-ai/steading` treat it as the live UI, a proof of concept that's being kept but is frozen, or something on its way out?
3. It's a Connect-RPC client of what? Does it call each project's services directly, or is there a steading backend in between? And which system owns the inbox items?
4. "Steading" is the organization, this repository, and nearly `stead.ing`. Does "Steading" also mean the suite or the product a community would adopt? And what is the `stead.ing` repository?

**ivy:** The key thing with the web UI's architecture is that it uses exactly the same underlying APIs as the CLI and stdio MCP server. It was inherently a way to iterate on enhancing the ergonomics of the API to accomodate future apps. It was always disposable and pushing on things like, "I want live streaming updates on everything happening in the system, but filtered so it isn't just a firehose of data," is what led to creating something that transports to other renditions of the same concept. More frontends and clients!

On rebuilding vs. decomposing, I think the main takeaway is to consider what the most successful parts of the UI were. I think it's a rebrand. It was very... AI slop. Which is to say, the UI had an absence of life and taste. It was purely functional and while aspects of its functionality were beautiful, it didn't feel like something I'd want to use.

While it's on my mind: Push notifications. I absolutely want push notifications to let me know when a part of my work becomes blocked.

The repository is private. We should probably rename it and archive it. 🤔 I haven't decided yet. But the main thing is that one of the next phases of work could be to ensure that every member of the product lineup has its own bundled web UI. There's a principle here and we should come back to overall project principles: single static binary. It's easier to maintain when your container image is a scratch base or when your deployed binary has everything it needs. You can fingerprint it with a checksum and replace it easily.

On which backend it talks to: It's a client, not a server or proxy. It obeys CORS and the user's web browser makes requests to the individual services. In *my* particular deployment, each service is a Tailscale service with tagged ACLs. Locally, those services listen on Unix sockets. The CLI and MCP support either. The MCP can run through the CLI or you can point your agent at an HTTP(S)-based MCP served by the individual services. This gives flexibility depending on what kind of deployment and harnesses you have running or if you want to make your own custom tooling or clients.

The project and suite are called Steading. The website is <https://stead.ing/>. So *this* particular repo (steading) is poorly named because it implies that it's the entire system. It isn't and the idea of renaming it to something like "console" or "web-ui" doesn't mesh well with my mental model of ease of adoption. I want people to be able to pick any single component and deploy it on its own and get immediate value. Users can't currently do that and it's a major problem to adoption and onboarding.

### Round 3

**Interviewer:**

1. The inbox draws on agents across services and a thicket conversation; live steering joins mill runs with a thicket panel. If each product serves its own UI, which one owns the inbox, and what does someone who deployed only mill see when a run needs them?
2. Is "the CLI" one CLI for the whole suite, `steading-ai/cli`? Does the standalone principle mean each component ships its own CLI and MCP, or that the shared CLI works against any subset of services?
3. Do push notifications belong to a component, or are they a want with no owner yet? And does "blocked" mean an inbox item being raised and a run stalling, or something wider?
4. Is the `stead.ing` repository the website's source? And in "I actually found it extremely useful on its own without any form of UI", is "it" the backlog?

**ivy:** I think describing later projects will illustrate *standlone components vs. features that span several*. I'd say an oversimplification is that the Unix philosophy is a big part of how I think about it. I want a composable system where each component can be freely tossed, rewritten, replaced. The inbox is primarily for skift (the agent-first project management system).

Every service offers its own CLI. It's both a client and server. The `cli` repo is another unfortunately named repo. I wanted a quick way to get what I'd built on Linux to run in a macOS VM. I was at a loss for what to call this one-off solution. I'd also archive it. There's still a gap though... I think we need these services to properly support macOS if that's truly convenient.

For push notifications, I don't really like the idea of a bunch of services sending them. I'd have to consider it but mainly, I think everything flows into skift. The inbox and notifications are the way of bringing eventful, actionable things to humans' attention where it's needed and nothing else.

Yeah, the stead.ing repo holds the website's source.

> I actually found it extremely useful on its own without any form of UI

I found skift, mill, almanac, and other services useful without any form of UI.

## steading-ai/skift

### Round 1

**Interviewer:** You've called it "the agent-first project management system," and said the inbox and notifications both flow into it. In your own words: what is it, what does it do, and why did you make it? What makes it agent-first: who writes to it and who reads from it, and what does it do that a conventional tracker like GitHub Issues or Linear doesn't? And where does the DAG shape of the backlog come in?

**ivy:** I dubbed myself the Trello queen back in the day. I was a huge fan of Pivotal Tracker in its heyday. I used an app called Sprintly for a while, then I started using GitHub Projects and loving it even with its shortcomings. I dislike Jira, 'nuff said. I've also used Linear, Asana, and a whole bunch of others including plain old todo list apps like Todoist, Things, OmniPlan, etc.

Where those project management systems fell short is that they're all made for humans. Agents work best when they can work on a single task and nothing else. It's all about preserving focus and creating a separation of concerns. An agent-first project management system consists of (and these are my opinions) amazing fucking tools that let the agent work effectively and build its context. A complex web UI with a bunch of custom fields and unstructured data with decades of technical debt isn't a great starting point. Retrofitting it into an MCP and calling it agent-ready is even worse. Sure, they can *use* those but can they use them *well*?

I came to the same conclusion around the same time as the rest of the industry: directed acyclic graphs are a great way to reason about a backlog. That's (and this is an over-simplification) a Gantt chart in essence, but humans don't have the time or attention span for such things and they quickly fall out of date. I'm not here to be a fucking project manager, I'm here to direct the priorities and overall vision of the crazy things I'm trying to build!

So if you think about a package manager (we'll take a more complex example: Gentoo--a source-based distribution), it understands your goal, "install ffmpeg". Well... there's a whole mess of prerequisites to meet that goal. You need the right toolchain, the various codecs, backends, parsers, and other libraries and tools to build and *use* that tool. A software project is very similar except there's ambiguity. Anyone who says a plan should never change when using an agent is leaving *everything* on the table.

Here's a goal: I want to reimagine Japanese language learning through full immersion. I read Ultralearning a long time ago and had some takeaways but I'd really like to see how it fares when your tutor is an LLM with a corpus of data at its disposable where it can find truth in addition to leveraging its training data. I don't want to learn textbook, sterilized Japanese and cram for the JLPT. I want to use it to meet *my* specific goals of immersing myself in an awesome culture, making friends, and expanding beyond my English-speaking predominantly American echo chamber to take in new ideas, try new things, make cool stuff with people across the world. And I have a theory -- the best way to achieve that is to take the language app *out* of a single app and embed it in *other* apps. I want a central place to practice and for it to be multi-sensory. I want to demonstrate and grow my knowledge through writing, reading, speaking, hearing and I want every sense exercised as much as possible.

So... did you run away screaming yet? How does one achieve this? You break it down and agents are extremely effective at identifying the prerequisites and terrible at keeping it all in their context window. To ask a human or AI to understand every single dependency and validate every ambiguity and update its findings and plan... it goes off the rails quickly without an agent-to-human feedback loop.

What happens when you use skift:
- Your agent decomposes it into what an agent can readily act on and what needs a human.
- It breaks it down into *right-sized* tasks and identifies their dependencies.
- Upstream nodes can refer back to their downstream nodes when being acted on.
- Example: do some research on SRS and other study methods -> review findings and make a decision -> research libraries and algorithms (build vs. buy) -> ask a human for input if there isn't a clear winner -> build it!
  - This is just one branch of a massive tree of work.

The important thing is that you leave room for ambiguity and resteering. You might decide to put a feature on hold, prioritize one over another, or respec based on learnings.

With skift, agents attach the artifacts of their deliverables as a URI or blob. That could be a commit on a repo, a document, an asset, or something else entirely.

The other core principle of skift: Agents tell skift what kinds of work they can take. They don't ask for specific tasks. They get a lease and send a heartbeat. This allows skift to be a work queue. Also, it handles accounting for cost. You record estimates and actuals and for now, I'm just recording actuals. It opens the door to evals because you can look at a new model and run it through the exact checkouts and compare against previous delivered artifacts with a ledger tracking the tangible costs. You get immediate feedback about cost vs. effectiveness for your specific work rather than taking this week's trust-me-bro benchmark as gospel.

### Round 2

**Interviewer:**

1. When decomposition marks a node as needing a human, does it become an inbox item? Are people effectively another kind of worker taking human-kind work, or is the inbox a separate path?
2. In the SRS example, is "upstream" the prerequisite and "downstream" what depends on it, so that the agent doing the research can see what the research is for?
3. Who decomposes a goal: a thicket agent, or any agent with skift's tools? And when findings change the plan, who updates the graph?
4. What defines a "kind of work": capabilities like model, harness, skills, or repository access, or a type on the node? And is mill the main consumer of skift's leases, or one of several?

**ivy:** People *can* be workers but more likely, they're talking to an agent acting on their behalf. For example, while mill is working on nodes from skift, I have my own Claude Code open talking to skift. I use the new app or review the things that interest me. I see something that doesn't look right or I have a new directive. I talk to my agent, give it screenshots, hash out ideas. When I'm ready, I ask it to make changes in skift. If something needs keys or approval that an agent just doesn't have the ability to do on its own, then I have some work to do. But generally, I'm just talking to an agent with a separate role.

The benefit is that I'm not shotgunning feedback and random ideas at an agent that's actively working on building what I previously said. That reduces distractions. Imagine trying to write a complex algorithm and keep track of a massive codebase refactoring while a manager is peppering you with every random thought in their mind. Agents have a greater capacity for it but there's a visibile decline in their work with what I'd call *instruction density*. The less it needs to worry about, the better.

Yeah, upstream/downstream is ambiguous. Think about it like this: The node to implement SRS can reference all of the previous delivered artifacts. The course corrections, research, findings, assessment, and decisions. It isn't limited to a task body. It has access to the entire graph that feeds into that node its working on.

Skift nodes try to be as simple as possible. The data model is concerned with specification of work, delivered artifacts, priority, tagging, groups, and accounting. It doesn't specify a model. Mill's job is to make those decisions and figure out what workflow it should run, what models it picks, what harness it runs, etc.

## steading-ai/mill

### Round 1

**Interviewer:** There's also an `ivy/mill`: is it the same project that moved to Steading, a fork, a predecessor, or unrelated? In your own words: what is mill, what does it do, and why did you make it? How does it decide the workflow, model, and harness for a node: configuration you write, rules, or something that learns from the cost ledger? What is a mill run: one node, one agent session, or a multi-step workflow? And where does `mill-evals` fit?

**ivy:** The simplest way to describe mill: It is a system for tracking agent definitions, plugins, and skills. A router (a skill in itself) decides which agents, plugins, skills, etc. are run and in what order.

It's more linear, not a DAG, but sometimes I think... it's DAGs all the way down, but for my and everyone else's sanity, we're not doing that yet. This is the system that many would think of as an "orchestrator" or conflate to be a software factory. Its role is to supervise the claiming of skift nodes, execition, and delivery of artifacts. It opens pull requests and gets them merged. A core principle: It's the repository's responsibility to gate what merges without human review. If there's no status checks or code ownership, that's that. It merges.

So mill picks the shape of the workflow. Administrators manage the prompts and models and what to use, when.

A problem I see in the current implementation: Mill also is responsible for the provisioning of sandboxes, dev environment setup, and a whole bunch of things to operating the agent harnesses. I think it should move and that's the seed of the idea behind croft.

### Round 2

**Interviewer:**

1. Is `ivy/mill` the same project as `steading-ai/mill`, a fork, a predecessor, or unrelated?
2. Do administrators define the menu (agent definitions, prompts, models, rules for when) and the router choose from it per node? Does the router use skift's cost ledger?
3. Is a run one skift node through one linear workflow? Does every node end in a pull request, and how does mill deliver work that isn't code?
4. What is `mill-evals`?
5. Does mill declare to skift what kinds of work it can take, and what defines a kind?

**ivy:** Ah, sorry. ivy/mill needs to be deleted. It moved to steading-ai/mill. I grew everything out of my personal GitHub and realized it would be more fulfilling to build in the open and collaborate with others.

Yeah, mill reports back into skift for the actual costs. Also, it currently handles credential brokering for inference and there's no guardrails around egress from agents. That's where the gateway service comes in as a longer term vision. I have some sketches like llmrouter where I wanted to balance and route across OpenRouter and multiple subscriptions. That's a real need but also being able to limit what an agent can do based on its role.

The golden example workflows separate implementer from reviewer. An implementer gets a skift node, it formualate a plan. The reviewer approves or requests changes. Once they're in alignment, the implementer builds it and the reviewer validates that the acceptance criteria of the skift node are met. It doesn't do code review yet. That's a later consideration. At least at this point in the lifecycle, I care more that agents aren't cutting corners or deviating from the vision. The big part of this is allowing those workflow steps to have persistence between "rounds". These agents aren't amnesiac between planning, reviewing, implementing, etc. They have memory of prior rounds to prevent moving goal posts or waffling on decisions.

A mill deployment has the ability to specify which projects and/or tags it can take. This allows deploying mill for different kinds of development tasks like iOS/macOS development or other ecosystem or resource constrained work. It's extremely flexible.

mill-evals is just a scratch repo for doing research on different models. Doesn't matter here, but mill/steading should definitely have a solution for evals down the line oriented around making informed decisions about the configuration and prompting of agents in the steading ecosystem.

## steading-ai/croft

### Round 1

**Interviewer:** You said croft is the seed of what should move out of mill: sandbox provisioning, dev environment setup, and operating the agent harnesses. In your own words: what is croft, what does it do, and why does it need to be its own component? Where's the line between mill and croft: does mill decide what runs and in what order, and croft where and how? Is croft also where the gateway lives, or is the gateway its own component with `llmrouter` as a seed? And does croft exist yet, or is it a concept with a repository?

**ivy:** Croft is the fence that contains your agents. It compartmentalizes them and sets them up to work wherever you intend to run them. That could be an ephemeral VM on Daytona, an EC2 or Firecracker instance, a Docker container, a self-hosted VM/container, or bare metal. As I built mill I realized there are many other ways you'd want to run computing tasks, not just agents and solutions like Docker, Incus, KVM, EC2, etc. aren't as abstract as I'd like... which is a horrible thing to write. I started small on a single server and quickly exhausted my available resources. It's a Linux box. But I want to build Mac and iOS apps. When I wanted to build an Unreal Engine game, I started crashing my server. When I want to work on projects in parallel that all need a graphics card for running tests or local models, I run into resource contention. So the idea is: Can I define a budget and set up backends and let Croft figure out where it needs to run and with what degree of isolation?

The gateway is its own entirely separate project. I think they work together.

Croft is nothing more than a concept at this point. It's living in Mill right now and I keep implementing pieces of it and the actual bindings for the agent harnesses in every odd project.

### Round 2

**Interviewer:**

1. Is croft a general placement and isolation layer for any workload, with agents as the main consumer, or is it agent-specific?
2. Is a budget money, resources like cores, memory, and GPUs, or both? Does croft choose the degree of isolation, or do you set a floor per workload that croft honors?
3. Do the harness bindings, the code that drives Claude Code, Codex, and the rest, belong in croft, mill, or a component of their own?
4. Which repositories hold croft or harness-binding pieces that the wiki should steer people away from?

**ivy:** I'd say croft is the backend-agnostic sandbox runtime. For simplicity, let's say we ship the agent bindings and they speak ACP or another protocol for communicating with agents. With that in mind, the gateway could probably live here for some time. There's some isolation considerations down the line, but anything is better than nothing at all.

Budget... that's a question I haven't fully answered yet. I'd say there's a combination of a task declaring the resources it needs. Croft might set budgets for how much money spend so we aren't bankrupting ourselves. I think it ties back into skift. Skift's ledger allows to attribute the actual cost of a project and with that data, we can estimate future work down the line. This lets you share your grandiose idea and realize, "Oh, that's going to cost $10,000 to build. Okay, maybe we should scale it down a bit..."

I care more about accounting that budget at this point, to be clear. Cloud providers spend limits and cgroup limits are sufficient for now.

Yes, you got it right that "harness bindings" are the communication layer around Claude Code, OpenCode, Pi, Prime, etc.

### Round 3

**Interviewer:** Proofreading the croft page drafted from this interview. It said: "croft exists because Docker, Incus, KVM, and EC2 aren't abstract enough on their own, and because one Linux server runs out fast. It can't build Mac or iOS apps, it crashes building an Unreal Engine game, and parallel projects that each need a graphics card for tests or local models contend for it." And, under budgets: "A task declares the resources it needs, and croft might cap how much money is spent. Spending ties back to skift's ledger."

**ivy:** Croft doesn't exist *yet*. Mill is currently holding a lot of its responsibilities and those gaps you documented are in *mill* at the moment. Also, "aren't abstract enough on their own" is ambiguous and doesn't lead with the problem statement.

Docker, Podmain, Incus, KVM, Kubernetes, etc. are all solving the problem of shipping and running images in a particular way. The goal of Croft isn't to remake a container or VM runtime. It's to offer a single interface for executing tasks anywhere. It's motivated by skyrocketing hardware costs that make self-hosting difficult to sustain when work bursts past the resources you have available currently.

On budgets: I think this might be unrelated. I recall leaning more towards allowing the underlying providers decide the budget. They all have limits, budgets, and handling for resource exhaustion for the most part. I'd say it's out-of-scope at this point.

## steading-ai/thicket

### Round 1

**Interviewer:** A guess, built only from what you've said so far: thicket hosts conversational agents, the "agent with a separate role, acting on your behalf" from skift, run as a service instead of a Claude Code session on your laptop. A thicket agent has tools to read and change the backlog, investigate, and gather details; it's who you talk to from the inbox, and the steering side panel beside a live mill run. Mill's workers build; thicket agents are who you talk to. What's wrong or missing? How do `steading-ai/thicket` and `ivy/thicket` relate? And where do thicket agents run, and does thicket or mill define their tools, prompts, and roles?

**ivy:** I made thicket because at the time, the policy with regards to use of Claude subscriptions with third-party harnesses was ambiguous. They've since loosened up a bit but it's still ambiguous. There's already other solutions like OpenClaw and Hermes that are far more mature. I *think* it's probably better to just use those and come up with an assistant-agnostic way to embed them into the product suite. That said, I see thicket as a good reference implementation to just make it work. Also, I was exploring some other avenues for interacting with agents like having a live assistant that you can talk to over the phone.

I've kind of run out of steam on it, but the pain point is persistent. I want an easy way to talk to an agent that's plugged into the rest of the Steading stack and for those agents to have the ability to file new skift nodes, check on system status, etc.

Another use-case: I want thicket agents to ping me when my attention is needed on a blocked skift node or something needing my attention. I'm not sure *how* yet because there's crossover between skift nodes being blocked and systems like almanac needing my feedback or awareness. For example, I had this idea in mind of almanac watching RSS feeds, mailing lists, and CVE feeds for vulnerability reports. I had in mind that it could file a node in skift to investigate and a *thicket* agent could check my systems for impact and decide what to do. That could mean filing more nodes, flagging my attention, etc.

I imagine a software stack that's incredibly malleable. It doesn't have to be a software factory the way most people are thinking about it. If I wanted an algorithmic trading system, I could just as easily make an almanac workflow that connects to Alpaca's websocket and webhooks to receive account events, watch tickers, etc., watch EDGAR for filings, watch news and inform me on market changes. I could send trading hypotheses to thicket or another chat agent and it could file nodes to investigate further, run backtests, kick off research, report back with findings. I could basically have the whole system work autonomously and steer it whenever I feel like it. This would free me up to focus on the unique pieces like integrations, data sourcing, building the algorithmic trading system, etc. while automating the build, maintenance, and accelerating testing and rollout of new trading strategies.

That's the gist of it, at least.

### Round 2

**Interviewer:**

1. Is the principle that the suite's CLIs, MCP servers, and APIs are the seam any assistant plugs into, whether thicket, OpenClaw, Hermes, or Claude Code, with no assistant being the product?
2. Should the wiki say that skift holds what needs a human, including blocked nodes and nodes filed by almanac or others, and that your assistant is how it reaches you? Or is that still open?
3. Is the vision that skift, mill, almanac, and an assistant are a general system for decomposable work, with software as the first use?
4. How does `ivy/thicket` relate to `steading-ai/thicket`?

**ivy:** steading-ai/thicket is the canonical repo.

Yeah for assistant-agnosticism, I was thinking A2A, AG-UI, and/or ACP. I haven't decided or done research into OpenClaw and Hermes as of late.

Skift is focused on tracking the work. Have we talked about it yet? If the naming is ambiguous:
- Steading is like homesteading. It's yours to build and customize to suit your "land". Each product offers a unique piece of the puzzle and gives you the flexibility to start with what you need. You can swap them out. I'm standardizing on Go, Connect-RPC, PostgreSQL, sqlc, Ginkgo/Gomega, golangci-lint. The idea being that everything is consistent, you can rapidly iterate in an easy-to-read language (minimal footguns, low syntax complexity). golangci-lint with depguard lets you protect things like the DB being accessed directly by agents -- they frequently violate layer boundaries or drop to Go's testing package. Ginkgo keeps the business requirements expressed rather than being obfuscated in `TestThing...` functions.
- Skift is a mix of Middle English and Scandinavian -- shift, change, switch + plot, allot, divide, etc.
- Almanac is like a farmer's almanac. I think we covered that.
- Mill works on the stream of tasks.
- Croft provisions the sandboxes and enforces the boundaries.
- etc. -- kind of a cottagecore naming convention.

## steading-ai/almanac

### Round 1

**Interviewer:** A guess from what you've said in passing: almanac watches the outside world on your behalf. It runs workflows that subscribe to feeds, streams, and webhooks, like RSS, mailing lists, CVE reports, market data, filings, and news, and turns what matters into action: filing skift nodes, or telling you something changed. Skift tracks work; almanac is where new work and awareness come from. Is that close, and what's wrong or missing? What is an almanac workflow: configuration, code, or an agent with a prompt? And is `ivy/almanac` a duplicate like the others?

**ivy:** Yeah, I frequently start things under ivy/ and move them to steading-ai/ to share. Almanac is another example. I'll probably start working in the open in steading-ai though. It's a new GitHub org.

I was writing this to another agent when I realized we should just continue the interview:

> Let's start with what Almanac is meant to be. Borrowing the concept of a farmer's almanac, its role in the Steading suite is to receive events (webhooks, email, whatever), run routines on a schedule (basically cron), and get stuff into skift so it can be worked. Examples of workflows:
> - Fixing Renovate PRs when status checks fail so they can be automatically merged by Renovate
> - Watching Alpaca's websocket ticker and acting on activity
> - Watching RSS feeds for security vulnerabilities and investigating blast radius and making preemptive fixes
> - Basically any case where we want an agent in the loop when something eventful happens
>
> The idea is that Almanac is deterministic and event-driven. Other pieces of the Steading stack handle things like running agents and opening PRs, merging them, etc.

The first iteration was quick and dirty to make it work and see if we could make something that works. I've since had other ideas about this. Here's some thoughts I wrote on the Discord:

> technicalpickles — 11/21/26, 12:42
> What is the naming theme? i know what a mill is, what an almanac is, but  not the others or how they relate? I can only think of like home steading
> Ivy — 11/21/26, 20:08
> Yeah! Cottagecore
> Skift is a mix of Middle English and Scandinavian. Loose meaning is to share/portion/divide a lot.
> Mill works a “stream” of portioned tasks (skift).
> No idea about the name “almanac” though. I just needed something that loosely fit with the theme. The alternative was atlas but I kinda shrugged on that one
> technicalpickles — Yesterday at 06:20
> I like it, i think of farmer’s almanac with all the forecasts, moon cycles, etc. atlas is overdone
> Ivy — Yesterday at 08:55
> Ah! Right. That was the meaning
> This is the most philosoraptor moment I’ve had. I’m also drowsy and reading with my eyes half open but…
>
> Multi-agent systems were just the actor model the entire time? 🤯
> Eh okay, not that novel. I think I said that to @martinemde the other day when I was saying I want to crib Terraform’s architecture but make it multilingual.
> Ivy — Yesterday at 09:05
> Terraform’s provider model is quite beautiful. They’re just binaries that use RPC over stdio.
>
> I’d just pick connect-rpc. It didn’t exist back when terraform first came out IIRC. They used a fork of the Go RPC which uses gob (builtin binary serialization, bespoke to go). The tradeoff is that they get to just serialize raw Go types but version and runtime differences can lead to incompatibility.
>
> So if you lift that into protocol buffers with a versioned schema and use a language agnostic protocol, you get something language agnostic (whoops, said that twice). The core can handle the essentials and primitives while your alamanac workflows are free to crash and restart under process supervision.
> The other nice part is that a core that has a sufficiently small surface area gets you zero downtime restarts down the line: https://systemd.io/FILE_DESCRIPTOR_STORE/
>
> You can hand off open file descriptors and serialize process state, restart the process, and the new one can pick up without dropping any connections!
>
> When each workflow is socket activated, it holds no memory when it’s idle and starts lazily: https://0pointer.de/blog/projects/socket-activation.html

Basically, I want to take a step back and think about almanac's architecture. It's currently a bespoke duplication of Steading's product suite rather than integrating with them and keeping a single responsibility of watching the world.

This is a theme across the entire lineup. Each should have a clear unique responsibility and an API contract. Steading's components are meant to work together and be interchangeable. The idea is that you can freely adopt the parts that you haven't already built or just get it up and running if you're starting from nothing. It's a bit risky to approach this given the complexity of making it a distributed architecture from day one, but it's worth the thought and has resulted in a well-understood architecture that I can diagram and point to when discussing software factories.

### Round 2

**Interviewer:**

1. Deterministic, but "an agent in the loop." Almanac never runs an agent itself: it files a skift node, and mill or an assistant does the agentic work. Is that strict, including for Alpaca, and does almanac only ever file nodes, or can it notify you directly?
2. How settled the provider model is. Is the plugin-binary design the decided direction, with the core handling intake, scheduling, supervision, and handoff to skift, and each workflow deciding what an event means?
3. systemd and macOS. Is almanac Linux-first, or should it also run under something like launchd?
4. Publishing the Discord thread. Are you fine publishing technicalpickles' words, and what's their GitHub login?

**ivy:** Here's a link to that Discord conversation too: https://discord.com/channels/1570339313957212222/1570339315228352635/1573800712033534043

> 1. Deterministic, but "an agent in the loop."

Yes, you got that right. The idea is that the core supplies an SDK (Go, to start) and an API contract for workflows. The workflows themselves aren't necessarily deterministic. They're the glue defining the flow. One thing that almanac may offer is a decision model / classifier. Something like Jev (https://docs.typesafe.ai/llms.txt) that allows almanac to filter a firehose of events down to what's actually work acting on.

> 2. How settled the provider model is.

I'm fairly certain this is the way to go but I haven't started any work on it yet.

> 3. systemd and macOS.

We'll have to figure that out. Steading needs to be platform-agnostic. I built on Linux + systemd with everything for my own ease of rapid development. I want to collaborate with others though and they all have a mix of Linux/macOS.

> 4. Publishing the Discord thread.

The Discord server is public. No worries. technicalpickles is technicalpickles on GitHub.

### Round 3

**Interviewer:** Proofreading the almanac page drafted from this interview, which ivy edited directly.

**ivy**, in the page:

> almanac watches the world. Named after a farmer's almanac, its role is to:
>
> - Receive events such as webhooks, RSS feeds, and email
> - Run routines on a schedule
> - Get things into skift so they can be worked
> - Filter and de-duplicate events so only noteworthy/actionable work is started

> The current implementation is a quick first iteration (in Rust) that duplicates parts of the suite instead of integrating with them. It will be rewritten in Go against the rest of steading's components.

## steading-ai/agent-plugins

### Round 1

**Interviewer:** A guess from the name and from mill being "a system for tracking agent definitions, plugins, and skills": agent-plugins holds the agent definitions, plugins, and skills mill's router chooses from, like the golden implementer and reviewer workflow and skills that teach agents to use skift and the rest of the stack well. Is that close? And are they only for mill's workers, or also for your own assistant?

**ivy:** That repo's empty. Just an idea that crossed my mind and I don't recall what it was about.

## steading-ai/scratch-forge

### Round 1

**Interviewer:** Is it a scratch repository for experiments, like `mill-evals`, or something about a forge in the software-factory sense?

**ivy:** Just a test repo that mill was using for end-to-end tests.

## steading-ai/wiki

### Round 1

**Interviewer:** Why start the wiki now: was the trigger outdated and misleading docs, or something else, like opening Steading to collaborators? And is the wiki only Steading's own knowledge base, or is a wiki kept by people and agents together something Steading users would run themselves, like the other components?

**ivy:** Yeah, the docs in all the projects have rotted away. They encode a ton of bad assumptions and agent reasoning that over-applies things I've said into unchanging gospel. The goal of this interview is to get the current truth and uncertainty. The goal of the wiki is to create a shared space for those kinds of docs. It's based on Andrej Karpathy's LLM Wiki concept: https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f

The goal is to keep all of the organizational knowledge in one place and use it when making changes to the documentation in the projects, steering the work, and whatnot.

### Round 2

**Interviewer:** Proofreading the pages drafted from this interview, which each ended with a section of open questions and described ivy in the third person.

**ivy:** Yeah, reading these sections, I'd toss *Open Questions* entirely. Those are just *your* questions. I don't want that living in the wiki. Every reader (human or agent) will read and interpret and come up with their own questions.

Also, I edited out my name in the almanac page and I'd do the same across the whole wiki. "ivy found skift useful without any UI," is a pointless statement to make. (Sorry, being direct.) This is a shared project with collaborators, not my personal blog.

To clarify, I'm addressing the reader as *you*, *your*, etc. and framing things from their perspective as a user or contributor. It's a project wiki!

### Round 3

**Interviewer:** Proofreading the pages drafted from this interview.

**ivy:** Okay, let's think about the taxonomy. I had in mind that pages would have a `type:` field that we enforce with wikilint. We could capture key details across types of pages so they can be rendered and checked consistently. For example, I don't want to jump onto the wiki and be overwhelmed with hundreds of pages. I want sections. Take projects as an example: They all have different statuses.

### Round 4

**Interviewer:** A proposal. Every page has a required `type`: `project` for the component pages, `topic` for everything else, and the type is the section a page is listed in. A project adds a `status` from a fixed list (`concept`, `active`, `dormant`) and an optional `repository`; every page adds a one-sentence `summary` for listings. Since frontmatter can't hold a footnote, a project page states its status in a `## Status` section that cites a source, and wikilint requires it. Are those the right status names and sections, what should render them, and is that enough to cite a status?

**ivy:** I'll let you use your best judgment and let you know if I disagree with anything. 👍

## ivy/kb

### Round 1

**Interviewer:** For each `ivy/` repository, is it in scope as a Steading concept, out of scope, a duplicate of a `steading-ai/` repository, or one of the scattered croft or harness-binding pieces?

**ivy:** kb is an earlier iteration before Steading was a concept. The goal was to have a knowledge base with builtin agentic retrieval and to have a feedback loop. The system responds to prompts from agents asking for details or information. It returns answers with citations. Inaccuracies, stale docs, and conflicting facts result in corrections. It tries to do it all on its own when under the steading model, changes to incorrect docs are just skift nodes. kb probably didn't need its own agent. It probably could have just been a subagent. It probably didn't *need* an MCP server. I have something lighter weight in mind for the future.

### Round 2

**Interviewer:** Is the "something lighter weight" this wiki, built from dated sources, or something separate still to come?

**ivy:** I'm still letting the concept simmer. kb ingests knowledge from a variety of sources, starting with indexing Git repositories and keeping them up-to-date. I think that's still an effective model since it tracks provenance and gives you a single way to query it all. That said, there's a question of whether it makes sense to index the code and not just the Markdown docs. Still thinking about it. I'll come back to it. I think that's all for my personal repos.

## Tech stack

### Round 1

**Interviewer:** Proofreading the tech stack page drafted from this interview.

**ivy:** Okay, let me explain the tech stack more thoroughly:
- Go: It's fast to compile, produces a single static binary, the language is easy to learn, the quality of tooling is unmatched, it's type checked, memory safe, and garbage collected, it's steeped in (distributed) systems and orchestration software (Docker, Terraform, Kubernetes, etc.), it's (mostly) free from bikeshedding (go fmt and move on with your life).
- Connect-RPC: It's language-agnostic, type checked, works in the browser or on mobile, and is easy to understand. It's better than gRPC in that regard. The bindings are far nicer. I used to have a fondness for REST but it's another bike shed.
- PostgreSQL, pgx, sqlc, goose: https://youjustneedpostgres.com -- Database agnosticism is the one thing I don't care for here. PostgreSQL offers everything you could need (great types, message queueing, full-text search, vector search, unstructured documents, time series, caching, etc.) and you get the most value from actually using the database and all of its features rather than trying to support every database out there. SQLite is probably enough for everyone who's running a monolith on their home server, but here, it provides a ton of value. It's fine if every piece of Steading is a shared tenant on one single PostgreSQL server. The point is that you get all of the benefits, especially multi-process concurrent read/write and transactional guarantees. sqlc is great in the way that Connect-RPC/Protocol Buffers are great: let the agents write raw SQL queries then generate a package for interfacing with the store. Same goes for goose for database migrations. We don't need a query builder!
- Ginkgo/Gomega: This is where I deviate from using Go's `testing` package. I want my skift nodes to specify behavior and requirements. I want business logic to be documented in English. The result is that it's easier to get clear on the specification of the software.
- golangci-lint: It supports custom rules like depguard. Like I mentioned earlier: I just want some light guardrails for the agents.
- cobra/pflag/viper: They're the best libraries in Go for CLIs. YMMV, however.
- `log/slog`: 'nuff said
- mise/hk: Tool version management and Git hooks are great with these tools. I wouldn't choose anything else at this point.

Consistency across this core stack makes refactoring easier, avoids remaking the same tough decisions, and leads to software defined by specification of the desired behavior rather than reinventing everything. Consumers are free to build with whatever tech stack they want. Anyone can build their own by implementing the Connect-RPC interfaces.

## Governance

### Round 1

**Interviewer:** Thinking like someone coming into Steading: who is the wiki for, and where should a newcomer go first to talk to people? Where does the work live, how does someone pick something up with no DRIs, and what happens to an outside pull request? Are pull requests written by a contributor's agents welcome? What licenses, sign-off, or code of conduct apply? Who makes decisions, who can commit to the wiki, and what happens as people join?

**ivy:** Hm, yeah *governance* has this weird weight to it. Maybe it still belongs in the wiki and there are other pages relating to the meta of the project.

Readers could be human or AI. Most newcomers are going to start with stead.ing and the wiki is what I'm thinking of as the source of truth while stead.ing and project-specific sites are the landing pages / front doors. I use a convention: <project>.stead.ing, by the way.

The Discord server is the best way to join discussions: https://discord.gg/BjCu2vHgX4

Any maintainer in the GitHub org can merge. It's lightweight. "Steading is a solo project that is trying to grow a community," is sort of factual (sort of not) and undersells it. Steading *started* as a solo project but we're a group of five at this point. It's early-stage and I'm embracing a bit of anarchy.

### Round 2

**Interviewer:** Pages can only state what a source records. Can these statements stand for the governance stance?

1. Steading started as a solo project. It's now an early-stage group of five, and deliberately a little anarchic.
2. ivy is BDFL to some degree, with strong opinions on architecture, and won't compromise the stack's integrity for the sake of consensus.
3. Steading prefers soft power: it shows why an architectural decision is right rather than imposing it.
4. Modularity and interoperability are governance as much as architecture. You don't need permission to build your idea or to replace a piece of the stack, and you don't need to share anyone's opinions to benefit from the work.
5. Earn adoption, not allegiance.
6. The community is neighbors, not a workforce. There's no shared product roadmap imposed on anyone, and no one decides what counts as a worthwhile pursuit. Sometimes interests intersect and people build together; sometimes they don't.
7. Why Steading exists: this kind of automation should be a commodity that anyone can get, with no sales call needed, and the world is better with an open-source option. Nobody should have to build a software factory from scratch.
8. The architecture comes from using the tools, not from guessing what people might need.

The meta pages would split into governance (how decisions get made), contributing (how to get involved), and community (who's here), with a "Why" on the Steading page. For contributing: where is the backlog and how does someone pick up work? Are agent-written pull requests welcome? What licenses, sign-off, or code of conduct apply? Should the pages name the maintainers?

**ivy:** Sounds good. Yeah, I've gotta figure out the backlog... I'd like it to be skift but for now, GitHub issues and pull requests are fine. I'd recommend starting with an issue or discussion on Discord before spending tokens on a contribution. Agent-written PRs are preferred. ISC license. We aren't big enough for anything else. Security disclosures can go to the respective repositories, referring to their SECURITY.md and/or submitting privately through GitHub's disclosure feature. I'd leave names out. It's fine that it's ambiguous at this stage.

## Principles

### Round 1

**Interviewer:** You asked to come back to overall principles. Here is everything you've said that reads like one. Strike or correct what's wrong, and add what's missing.

1. Unix philosophy: a composable system where each component can be freely tossed, rewritten, or replaced.
2. Each component has a clear, unique responsibility and an API contract, and components are interchangeable.
3. Any single component can be deployed alone and give immediate value.
4. A single static binary, with its web UI bundled.
5. Every client uses the same APIs: the web UI, the CLI, and MCP. Every service offers its own CLI, both client and server. Services listen on Unix sockets or the network; MCP runs over stdio or HTTP.
6. Platform-agnostic: Linux and macOS.
7. Assistant-agnostic: assistants plug in through protocols like A2A, AG-UI, or ACP.
8. Agent-first tools: built so agents use them well, not human tools with an MCP retrofitted.
9. One task per agent, separate roles: builders aren't interrupted, and humans steer through their own agent, which keeps instruction density low.
10. Plans change: leave room for ambiguity and resteering.
11. The repository gates what merges without human review.
12. People are pulled in only for eventful, actionable things.
13. Record costs before enforcing budgets, and judge cost against effectiveness on your own work, not benchmarks.
14. The standard stack: Go, Connect-RPC, PostgreSQL, sqlc, Ginkgo/Gomega, and golangci-lint with depguard.
15. Cottagecore naming.
16. Built in the open, in `steading-ai`.

**ivy:**

> 12. People are pulled in only for eventful, actionable things.

I'd clarify. The goal is for the system to run autonomously, not act as a project manager and bother humans constantly with the tiny details. Under this model, every user is an executive, not a supervisor or IC. They might get involved in the small details, they might not. That's for them to decide. The main objective though is to only flag when there's something that absolutely needs a human. Like credentials, billing, and things that aren't sufficiently automated yet.

> 13. Record costs before enforcing budgets.

Maybe not a core principles, but more of a way of reasoning about how I built the system. It's a huge value-add to have my own corpus of past skift nodes to benchmark against. It makes evaluating new models, alternative harnesses, and planning easier. I don't really have a need for budgeting at this point but as I improve the system and it becomes more efficient at building to the point of consistently exhausting my tokens, it's something I'd like to get more of a handle on. That said, I find holds and priorities within skift to be sufficient at this stage and for my personal usage patterns.
