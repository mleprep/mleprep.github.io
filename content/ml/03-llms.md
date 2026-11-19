# LLMs & GenAI

Large language model questions now common in ML/AI engineering interviews.

## How LLMs Are Trained

**Q: Walk through the LLM training pipeline.**

```viz
web-scale text --> tokenizer --> PRETRAINING (next-token prediction, trillions of tokens)
                                      |
                                      v   base model: knows a lot, doesn't follow instructions
                     SUPERVISED FINE-TUNING (instruction -> response demonstrations)
                                      |
                                      v
         PREFERENCE TUNING (RLHF with a reward model + PPO, or DPO / variants)
                                      |
                                      v   assistant model: helpful, harmless, follows instructions
                     (+ RL on verifiable tasks: math, code, tool use, reasoning)
```

- **Pretraining** objective: minimize cross-entropy of the next token (causal LM). Compute ≈ `6 · params · tokens` FLOPs.
- **Scaling laws** (Kaplan, Chinchilla): loss falls as a power law in params, data and compute; compute-optimal training uses ~20 tokens per parameter (modern models often train far beyond that to make inference cheaper).
- Data quality (dedup, filtering, mixing code/math) matters as much as size.

## Tokenization

**Q: What is BPE and why does tokenization matter?**

**Byte-Pair Encoding** starts from bytes/characters and repeatedly merges the most frequent adjacent pair into a new token until the vocabulary size (~32k–200k) is reached. Variants: WordPiece, Unigram (SentencePiece).
- Trade-off: bigger vocab → shorter sequences but larger embedding/softmax matrices.
- Consequences: arithmetic and character-level tasks are hard (digits split oddly), non-English text costs more tokens, whitespace and casing produce different tokens.

## Positional Encodings

**Q: How do transformers know token order?**

Attention is permutation-invariant, so position must be injected:
- **Sinusoidal** (original transformer) or **learned absolute** embeddings added to inputs.
- **Relative** schemes: T5 bias, **ALiBi** (linear distance penalty in attention scores).
- **RoPE** (rotary): rotate query/key vectors by position-dependent angles so `q·k` depends on relative offset; standard in modern LLMs; context extension via position interpolation / NTK / YaRN scaling.

## Inference: KV Cache, Batching, Quantization

**Q: Why is LLM inference expensive and how is it optimized?**

- Generation is **autoregressive**: one forward pass per output token. The **prefill** phase (processing the prompt) is compute-bound; **decode** (one token at a time) is **memory-bandwidth-bound**.
- **KV cache**: store past keys/values so each new token only computes its own; memory ≈ `2 · layers · heads · d_head · seq_len · batch · bytes`. Reduced by **MQA/GQA** (share K/V across heads) and paged memory (**PagedAttention**/vLLM).
- **Continuous batching**: add/remove requests from the running batch each step → high GPU utilization.
- **Quantization**: int8/int4/fp8 weights (GPTQ, AWQ) → less memory and bandwidth.
- **Speculative decoding**: a small draft model proposes several tokens; the large model verifies them in one pass (same output distribution).
- Others: FlashAttention, tensor parallelism, distillation into smaller models, prompt caching.

## Fine-Tuning: SFT, LoRA & PEFT

**Q: How does LoRA work and when would you fine-tune vs prompt?**

- **LoRA**: freeze `W`, learn a low-rank update `ΔW = B·A` (`A: r×d`, `B: d×r`, r ≈ 8–64) on attention/MLP matrices. Trains ~0.1–1% of params; adapters are tiny and swappable; can merge into W for zero inference overhead. **QLoRA** = LoRA on a 4-bit quantized base.
- Choose: **prompting / few-shot** first (cheap, flexible) → **RAG** if the model lacks knowledge → **fine-tuning** for format/style/behaviour consistency, domain language, latency/cost reduction (smaller tuned model), or tasks prompting can't reach.
- Data quality beats quantity: a few thousand clean examples often suffice for SFT.

## RLHF and DPO

**Q: Explain RLHF. What does DPO change?**

1. Collect **preference pairs** (prompt, chosen, rejected) from humans or AI feedback.
2. Train a **reward model** `r(x, y)` with a Bradley–Terry loss: `−log σ(r(chosen) − r(rejected))`.
3. Optimize the policy with **PPO** to maximize reward minus a **KL penalty** to the SFT model (prevents reward hacking / drifting into gibberish).

**DPO** skips the explicit reward model and RL loop: a closed-form derivation turns the same objective into a simple classification-style loss on preference pairs, using the frozen reference model's log-probs. Simpler and more stable; widely used along with variants (IPO, KTO, ORPO). RL with **verifiable rewards** (unit tests, math answers) powers modern reasoning models.

## Retrieval-Augmented Generation (RAG)

**Q: Design a RAG system over company documents.**

```viz
INDEXING:  docs -> parse/clean -> chunk (300-800 tokens, overlap) -> embed -> vector DB (+ BM25 index, metadata)

QUERY:     question -> (rewrite / expand) -> hybrid retrieve top-50 (dense + BM25)
                    -> rerank (cross-encoder) -> top-5 chunks
                    -> prompt: system + context chunks + question -> LLM -> answer with citations
```

Design points:
- **Chunking** respects structure (headings, tables); store metadata for filtering and ACLs.
- **Hybrid search** (embeddings + keyword) is more robust than either alone; **rerankers** boost precision.
- Ground answers in the context, ask for citations, allow "I don't know".
- **Evaluate**: retrieval recall@k / MRR on a labeled set, answer faithfulness and relevance (LLM-as-judge + human spot checks), latency and cost.
- Failure modes: wrong chunk granularity, stale index, missing permissions filtering, context overflow, prompt injection via retrieved docs.

## Decoding Strategies

**Q: Greedy, beam search, temperature, top-k, top-p — when to use which?**

- **Greedy**: argmax each step; deterministic, can be repetitive.
- **Beam search**: keep the b best partial sequences; good for translation/summarization with a "correct" answer; bland for open-ended text.
- **Temperature** `T`: divide logits by T; T < 1 sharper (more deterministic), T > 1 more diverse.
- **Top-k**: sample from the k most likely tokens. **Top-p (nucleus)**: sample from the smallest set whose probability mass ≥ p (adapts to the distribution's shape).
- Repetition/frequency penalties, stop sequences, constrained/structured decoding (JSON grammars) for reliable output formats.

## Evaluating LLMs & Hallucinations

**Q: How do you evaluate an LLM application, and reduce hallucinations?**

- **Benchmarks** (MMLU, GSM8K, HumanEval, …) for general capability — beware contamination.
- **Task-specific evals**: a curated, versioned test set from real traffic; exact match / unit tests where possible; **LLM-as-judge** with rubrics (validated against humans); pairwise preference win rates.
- Track quality, latency, cost, and safety (toxicity, jailbreaks, PII leakage) per release; A/B test in production.
- **Hallucinations** (fluent but false): ground with RAG and citations, constrain to the provided context, lower temperature, self-consistency (sample several answers and compare), verification steps/tools (calculators, code execution, search), and teach the model to abstain.


## Agents & Tool Use

**Q: How do LLM agents work, and what makes them reliable?**

An agent runs a loop: the model receives the goal + available **tools** (function schemas), decides on a call, the runtime executes it, results go back into the context, repeat until done.

```viz
user goal -> LLM (plan / choose tool) -> tool call {name, args} -> runtime executes (search, code, DB, API)
                ^                                                         |
                +-------------------- observation / result ---------------+
                stop when the model returns a final answer (or a step/cost budget is hit)
```

Reliability levers: clear tool descriptions and typed schemas; small, composable tools; structured outputs; guardrails on side-effecting actions (confirmation, sandboxing, least privilege); step limits and timeouts; memory (summaries, retrieval) for long tasks; evaluating end-to-end task success, not single responses. Beware **prompt injection** from tool outputs (treat retrieved content as data, not instructions).

## Prompt Engineering

**Q: What prompting techniques actually matter?**

- Be specific about the task, audience, format and constraints; give the **why**.
- **Few-shot examples** that cover edge cases; consistent formatting.
- **Structured output** (JSON schema / tags) for downstream parsing.
- Let the model **reason before answering** (chain-of-thought / extended thinking) on multi-step problems; ask it to verify against criteria.
- Put long documents first and the question last; quote the relevant parts before answering.
- Decompose complex tasks into chained prompts; evaluate prompts against a test set like code.

## Long Context

**Q: How do models handle very long inputs, and what breaks?**

Techniques: RoPE scaling / position interpolation, training on long sequences, efficient attention (FlashAttention, sliding window + global tokens), KV-cache compression, ring/sequence parallelism.
Failure modes: "lost in the middle" (information in the middle is used less), cost growing with length, distractors reducing accuracy. Mitigations: RAG to select relevant chunks, put key instructions at the start/end, summarize or structure long inputs, evaluate with needle-in-a-haystack *and* realistic multi-hop tasks.

## Vector Search & ANN Indexes

**Q: How do vector databases find nearest neighbours among a billion embeddings?**

Exact search is O(N·d) per query — too slow at scale, so we trade a little recall for speed with **approximate nearest neighbour (ANN)** indexes:

| Index | Idea | Trade-offs |
|---|---|---|
| **HNSW** | multi-layer proximity graph; greedy search from a sparse top layer down to a dense bottom layer | excellent recall/latency; memory heavy (vectors + graph); slow-ish inserts; the common default |
| **IVF** | k-means "coarse" clusters; search only the `nprobe` closest clusters | fast, tunable; recall drops if neighbours sit in unprobed clusters |
| **PQ** (product quantization) | split each vector into m sub-vectors, replace each by one of 256 centroids → m bytes per vector | 10–100× compression; approximate distances; combine as **IVF-PQ** for billions of vectors |
| **DiskANN / ScaNN** | graph on SSD / anisotropic quantization | billion scale on one machine; high throughput |

- Knobs: `efSearch` (HNSW) and `nprobe` (IVF) trade recall for latency; measure **recall@k** against exact search.
- **Cosine similarity** = dot product on L2-normalized vectors — normalize at index and query time.
- **Filtering** (tenant, date, ACL): pre-filtering can break graph search; post-filtering can return too few results → filtered-HNSW variants, partitioning by tenant, or over-fetching.
- **Hybrid search**: combine BM25 and vector scores (reciprocal rank fusion) — better than either alone for RAG.
- Updates/deletes: HNSW tombstones and periodic rebuilds; keep the raw vectors elsewhere so indexes can be rebuilt with new parameters or embedding models (re-embedding everything when the model changes!).

## Reasoning Models & Test-Time Compute

**Q: What makes "reasoning" models different, and how do you spend more compute at inference?**

- **Chain-of-thought**: generating intermediate steps before the answer improves multi-step problems (math, code, planning).
- **Test-time compute scaling**: accuracy keeps rising with more inference compute — longer thinking, **self-consistency** (sample many solutions, majority vote), **best-of-N** with a verifier/reward model, tree search over partial solutions.
- **Training reasoners**: reinforcement learning with **verifiable rewards** (unit tests pass, math answer correct) teaches models to produce long, self-correcting reasoning (e.g. PPO or GRPO — group-relative advantages over several samples per prompt, no separate value model). **Outcome** reward models score final answers; **process** reward models score each step.
- **Distillation**: reasoning traces from a large model fine-tune smaller ones cheaply.
- Trade-offs: latency and cost grow with thinking tokens → route easy queries to fast models, cap reasoning budgets; reasoning can still be wrong — verify with tools (code execution, search) where possible.

## LLM Safety & Guardrails

**Q: How do you make an LLM application safe and robust?**

- **Threats**: harmful content requests and **jailbreaks** (role-play, obfuscation, many-shot), **prompt injection** — direct, or *indirect* via retrieved documents, emails, web pages and tool outputs that contain instructions — data exfiltration through tools/links, PII leakage, hallucinated facts.
- **Model-level**: safety fine-tuning (RLHF, constitutional AI / RLAIF), refusal training balanced against over-refusal.
- **System-level** (defense in depth):
  - treat all retrieved/tool content as **data, not instructions**; keep a clear instruction hierarchy (system > developer > user > tool output);
  - input and output **classifiers** (toxicity, PII, policy violations); structured outputs validated against schemas;
  - **least privilege** for tools, sandboxed code execution, allow-lists for URLs/actions, **human confirmation** for side-effecting actions (payments, emails, deletes);
  - rate limiting, abuse monitoring, audit logs.
- **Evaluation**: red-teaming (human + automated), jailbreak and injection benchmarks, regression suites run on every model/prompt change; track both harmful-output rate and false refusals.

## Multimodal LLMs

**Q: How do LLMs understand images (and audio)?**

- **Late-fusion recipe** (LLaVA-style): a pretrained vision encoder (CLIP/SigLIP ViT) turns the image into patch embeddings → a small **projector** (MLP, or a Q-Former/perceiver that compresses to fewer tokens) maps them into the LLM's embedding space → the LLM reads them like text tokens.
- **Training**: (1) align — train the projector on image–caption pairs with the encoder and LLM frozen; (2) **visual instruction tuning** — fine-tune on image + instruction → answer data (often LLM-generated), unfreezing more of the model.
- **Early fusion / native multimodal** models tokenize images (and audio) and train jointly from the start — better cross-modal reasoning, can also *generate* images as tokens.
- High-resolution images: tile into crops plus a global thumbnail; token count (cost) grows with resolution.
- Audio: speech encoders (Whisper-style) feeding the LLM, or end-to-end speech models for low-latency voice.
- Evaluation: VQA, document/chart understanding (OCR-heavy), grounding, hallucination of objects not in the image.
