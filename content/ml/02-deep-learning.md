# Deep Learning

Neural network fundamentals, architectures and training practice.

## Backpropagation

**Q: Explain backpropagation.**

Backprop is the **chain rule** applied efficiently on the computation graph: run a forward pass and cache activations, then propagate `∂L/∂(output)` backwards, multiplying by each layer's local Jacobian. Cost ≈ 2× the forward pass.

```viz
forward:   x --W1--> z1 --ReLU--> a1 --W2--> z2 --softmax/CE--> L
backward:  dz2 = p - y
           dW2 = a1ᵀ dz2          da1 = dz2 W2ᵀ
           dz1 = da1 * (z1 > 0)   dW1 = xᵀ dz1
```

Key facts: gradients of a linear layer `y = xW` are `dW = xᵀ·dy`, `dx = dy·Wᵀ`; softmax + cross-entropy has the clean gradient `p − y`; autodiff frameworks implement exactly this (reverse-mode AD).

## Activation Functions & Vanishing Gradients

**Q: Why ReLU? What are vanishing/exploding gradients?**

| Activation | Formula | Notes |
|---|---|---|
| Sigmoid | `1/(1+e^{−x})` | saturates → gradients ≤ 0.25 vanish; outputs not zero-centred; use for binary outputs/gates |
| Tanh | `(eˣ−e^{−x})/(eˣ+e^{−x})` | zero-centred, still saturates |
| ReLU | `max(0, x)` | cheap, no saturation for x>0; "dead ReLU" units |
| Leaky ReLU / ELU | small slope for x<0 | avoids dead units |
| GELU / SiLU (Swish) | `x·Φ(x)`, `x·σ(x)` | smooth; standard in transformers |

Deep chains multiply many Jacobians: norms < 1 → **vanishing**, > 1 → **exploding**. Fixes: ReLU-family activations, careful **initialization** (Xavier/Glorot for tanh, **He/Kaiming** for ReLU), **normalization layers**, **residual connections** (gradient highway), **gradient clipping** (exploding, esp. RNNs), LSTM/GRU gating.

## Batch Norm vs Layer Norm

**Q: What does normalization do, and why do transformers use LayerNorm?**

Normalize activations to zero mean / unit variance, then apply learned scale γ and shift β. Smoother loss landscape, allows higher learning rates, less sensitivity to init.
- **BatchNorm**: statistics over the **batch** (per channel). Needs running averages for inference; behaves poorly with small or variable batches and sequences. Standard in CNNs.
- **LayerNorm**: statistics over the **features of each example** → independent of batch size, same at train/inference; standard in transformers/RNNs. **RMSNorm** drops the mean subtraction (LLaMA).
- Pre-LN (normalize before the sublayer) trains more stably than Post-LN in deep transformers.

```viz
tensor [batch N, features C]
BatchNorm: normalize each column (over N)      LayerNorm: normalize each row (over C)
```

## Convolutional Neural Networks

**Q: Why do CNNs work well for images?**

- **Local connectivity** + **weight sharing** → far fewer parameters than dense layers and **translation equivariance**; pooling/striding gives some invariance and growing receptive fields.
- Output size: `(W − K + 2P) / S + 1`. Params of a conv layer: `K·K·C_in·C_out (+ C_out)`.
- Milestones: LeNet → AlexNet (ReLU, dropout, GPUs) → VGG (3×3 stacks) → **ResNet** (residual connections, 100+ layers) → EfficientNet (compound scaling) → ConvNeXt; **Vision Transformers** (patches as tokens) win at large data scale.
- 1×1 convs mix channels cheaply; depthwise-separable convs (MobileNet) cut compute.

```viz
input 32x32x3 --conv3x3x16--> 32x32x16 --pool2--> 16x16x16 --conv--> ... --global avg pool--> FC --> logits
```

## RNN, LSTM & GRU

**Q: How does an LSTM address vanishing gradients?**

A vanilla RNN `h_t = tanh(W h_{t−1} + U x_t)` multiplies by `W` at every step → gradients vanish/explode over long sequences.
**LSTM** adds a **cell state** with additive updates controlled by gates:
`f = σ(·)` forget, `i = σ(·)` input, `o = σ(·)` output, `c_t = f ⊙ c_{t−1} + i ⊙ tanh(·)`, `h_t = o ⊙ tanh(c_t)`.
The additive path lets gradients flow when `f ≈ 1`. **GRU** merges gates (update/reset) with fewer parameters.
Limitations vs transformers: sequential (no parallelism over time), limited long-range memory. Still useful for streaming/low-latency on-device models; modern **state-space models** (Mamba) revisit recurrence with parallel training.

## Attention & Transformers

**Q: Explain self-attention and the transformer block.**

Each token builds a **query**, **key** and **value**; it attends to all tokens with weights `softmax(QKᵀ / √d_k)` and mixes their values:
`Attention(Q, K, V) = softmax(QKᵀ/√d_k + mask) V`.
- `√d_k` keeps dot products from growing with dimension (which would saturate softmax).
- **Multi-head**: h parallel attentions in lower-dimensional subspaces, concatenated → different heads capture different relations.
- **Causal mask** (decoder) blocks attention to future tokens.
- Block: `x + MHA(LN(x))` then `x + FFN(LN(x))` (FFN ≈ 4× hidden, GELU/SwiGLU).
- Cost: O(n²·d) time and O(n²) memory in sequence length → FlashAttention (IO-aware exact attention), sparse/linear attention, sliding windows.
- Variants: **encoder-only** (BERT: bidirectional, classification/embedding), **decoder-only** (GPT: generation), **encoder-decoder** (T5, translation).

```viz
tokens -> embed + position -> [ LN -> multi-head attention -> + ] -> [ LN -> FFN -> + ] x N -> LN -> logits
                                 \_______ residual ________/          \___ residual __/
```

## Embeddings

**Q: What are embeddings and how are they learned?**

Dense low-dimensional vectors where geometric closeness reflects semantic similarity.
- **word2vec** (skip-gram / CBOW with negative sampling), GloVe (co-occurrence factorization), contextual embeddings from transformers (BERT, sentence-transformers).
- **Two-tower / contrastive** models (CLIP, dual encoders): pull matching pairs together, push negatives apart (InfoNCE loss); used for retrieval and recommendations.
- Categorical ids (users, items) → embedding tables learned end-to-end.
- Serve with approximate nearest neighbour indexes (HNSW, IVF-PQ); measure cosine similarity.

## Training Tricks & Debugging

**Q: Your model doesn't train. How do you debug it?**

1. **Overfit a tiny batch** first — if you can't reach ~0 loss, there's a bug (labels, loss, data pipeline, frozen params).
2. Check inputs: normalization, shuffled labels, augmentation bugs, data leakage.
3. Learning rate: run an LR range test; use warmup + decay; watch for loss spikes/NaNs.
4. Monitor gradient norms, activation statistics, dead units; use **gradient clipping**.
5. Compare against a simple baseline; check the initial loss (`ln(num_classes)` for CE).
6. Scale up gradually: more data/model/epochs only after the small setup works.

Efficiency: **mixed precision** (bf16/fp16 + loss scaling), gradient accumulation, gradient checkpointing (trade compute for memory), data/tensor/pipeline parallelism, ZeRO/FSDP sharding.

## Transfer Learning & Fine-Tuning

**Q: When and how do you fine-tune a pretrained model?**

- Start from a model pretrained on large data (ImageNet, web text) — features transfer, needing far less labeled data.
- Small data / similar domain: **freeze** the backbone, train a new head (linear probe). More data or different domain: **fine-tune** all or top layers with a **lower LR** (discriminative LRs), warmup, early stopping.
- Parameter-efficient fine-tuning (**PEFT**): adapters, **LoRA** (low-rank update `W + BA`), prompt/prefix tuning — train <1% of params, keep one base model for many tasks.
- Beware **catastrophic forgetting**: mix in original data, use smaller LRs, or regularize toward the original weights.

## Generative Models: VAE, GAN, Diffusion

**Q: Compare VAEs, GANs and diffusion models.**

| | Idea | Pros | Cons |
|---|---|---|---|
| **VAE** | encoder → latent distribution, decoder reconstructs; loss = reconstruction + KL to prior (ELBO); reparameterization trick | stable, smooth latent space | blurry samples |
| **GAN** | generator vs discriminator minimax game | sharp samples, fast sampling | unstable training, mode collapse |
| **Diffusion** | learn to reverse a gradual noising process (predict the noise at step t) | best quality & diversity, stable training | slow sampling (mitigated by DDIM, distillation, latent diffusion) |
| **Autoregressive** | model each token given all previous ones (chain rule of probability) | exact likelihood, scales great (LLMs) | sequential generation |

## Regularization & Data Augmentation in Deep Nets

**Q: How do you regularize a large neural network?**

Weight decay (AdamW), dropout (less common in big transformers), **data augmentation** (images: random crop/flip/color jitter, Mixup/CutMix; text: back-translation, token masking; audio: SpecAugment), label smoothing, stochastic depth, early stopping, and — most effective — **more data**. Large models also regularize implicitly through SGD noise and architecture choices.
