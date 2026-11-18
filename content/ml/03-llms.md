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
