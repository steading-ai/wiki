---
type: project
status: concept
summary: Brokers inference and limits what agents can do.
---
# The gateway

The gateway stands between agents and what they reach. It brokers credentials for inference, balances and routes inference across providers such as OpenRouter and multiple subscriptions, and limits what an agent can do based on its role.[^mill]

The gateway is its own project, separate from [croft](croft.md), and the two work together.[^croft] It may live in croft for a while; isolation comes later, and anything is better than nothing at all.[^croft]

## Status

The gateway is a longer-term vision. Credential brokering for inference happens in [mill](mill.md) for now, with no guardrails on what agents can reach, and [llmrouter](https://github.com/steading-ai/llmrouter) is a sketch of the routing.[^mill]

[^mill]: [Projects interview: mill](../sources/202611150814-projects-in-ivys-words.md#steading-aimill)
[^croft]: [Projects interview: croft](../sources/202611150814-projects-in-ivys-words.md#steading-aicroft)
