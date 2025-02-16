# OpenChef

OpenChef is a **Go-based** AI agent orchestrator that embeds **NATS** for messaging and exposes an **OpenAI-compatible** HTTP API. It unifies your custom AI logic (across any language or framework) into a single, **easy-to-run** system—whether you need a simple “direct model” or a sophisticated multi-agent swarm.

---

## Why OpenChef?

**1. Single Binary, Instant Setup**  
- No complicated cloud or multi-service overhead—just download/run one binary.  
- Embedded **NATS** means you get robust, high-speed messaging out-of-the-box.

**2. Future-Proof Your AI**  
- Start small with a direct agent (like a single custom LLM) and scale up to orchestrated or swarm-based multi-agent workflows.  
- Plug in advanced AI logic or external libraries in Python, Node, or Go—OpenChef unifies them via NATS.

**3. OpenAI-Like API**  
- Your existing tools (LangChain, openai-python) can talk to OpenChef by changing `api_base`.  
- Let your custom “model” or agent appear as if it's a standard OpenAI endpoint, speeding up integration.

**4. Collaborative Agent Ecosystem**  
- Easily register new agents, from a Summarizer to a Database Query bot, and let them cooperate or self-organize.  
- Use “Orchestrator” mode for step-by-step control, or “Swarm” mode for fully decentralized, claim-based problem-solving.

**5. Simple Local UI** *(optional)*  
- A minimal dashboard to see “Registered Agents,” recent requests, or partial logs—helping you debug and track usage.

---

## Key Features

1. **Embedded NATS** – High-performance messaging, zero extra infra needed.  
2. **OpenAI-Endpoints** – `POST /v1/chat/completions`, returning JSON or partial SSE chunks.  
3. **Modes**:
   - **Direct**: route to a single agent subject (`myAgent`)  
   - **Orchestrator**: built-in “supervisor agent” that calls multiple sub-agents in sequence/parallel  
   - **Swarm**: broadcast tasks, advanced “claim” logic for truly distributed solutions  
4. **Agent Registration** – Agents can connect from Python/Node/Go, sign up with a name and subject, handle requests.  
5. **Optional File Endpoint** – If you want a local `/v1/files` for storing metadata or small assets.

---

## Quick Example

```bash
# 1) Run OpenChef
./openchef run --port=8080

# 2) A Python agent using NATS
pip install nats-py
python agent.py  # Subscribes to "model.direct.myAgent.request"

# 3) Call OpenChef's endpoint
curl -X POST http://localhost:8080/v1/chat/completions \
     -H "Content-Type: application/json" \
     -d '{
       "model": "myAgent",
       "messages": [{"role": "user", "content": "Hello?"}]
     }'
# => Passes request to your agent, returns AI response in OpenAI JSON


## Quickstart

1. **Download / Build**

```bash
# Example if you have source
git clone https://github.com/yourorg/openchef.git
cd openchef
go build -o openchef cmd/openchef/main.go


## Development

1. Clone the repository
```bash
git clone https://github.com/epuerta9/openchef.git
cd openchef
```

2. Create your environment file
```bash
cp .env.template .env
```

3. Install dependencies and tools
```bash
make setup
```

4. Start the development server
```bash
make dev
```

## Documentation
See the [docs](./docs) directory for detailed documentation.



