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

Deep chains multiply many Jacobians: norms below 1 → **vanishing**, above 1 → **exploding**. Fixes: ReLU-family activations, careful **initialization** (Xavier/Glorot for tanh, **He/Kaiming** for ReLU), **normalization layers**, **residual connections** (gradient highway), **gradient clipping** (exploding, esp. RNNs), LSTM/GRU gating.

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


## Self-Supervised & Contrastive Learning

**Q: How do models learn without labels?**

Create a supervised task from the data itself:
- **Masked prediction**: BERT (mask tokens), MAE (mask image patches) — reconstruct what's hidden.
- **Next-token prediction**: GPT-style language modelling.
- **Contrastive**: SimCLR/MoCo (two augmentations of the same image are positives, others negatives), CLIP (image–caption pairs). Loss = **InfoNCE**: `−log exp(sim(q,k⁺)/τ) / Σ exp(sim(q,k)/τ)`; needs many negatives (large batches or memory queues).
- **Non-contrastive**: BYOL, DINO — avoid collapse with stop-gradient / momentum teachers.
Pretrain on huge unlabeled data, then fine-tune or linear-probe with few labels.

## Mixture of Experts (MoE)

**Q: What is a mixture-of-experts layer and why use it?**

Replace a dense FFN with E expert FFNs and a **router** that sends each token to the top-k experts (k = 1–2). Parameters grow ~E×, compute per token stays ~constant → more capacity for the same FLOPs (Switch Transformer, Mixtral, many frontier LLMs).
Challenges: **load balancing** (auxiliary loss so experts are used evenly), expert capacity limits / token dropping, all-to-all communication across devices (expert parallelism), and higher memory to hold all experts.

## Distributed Training

**Q: How do you train a model that doesn't fit on one GPU?**

| Strategy | Splits | Communication |
|---|---|---|
| **Data parallel** (DDP) | batch across GPUs, full model copy each | all-reduce gradients every step |
| **ZeRO / FSDP** | shards optimizer states, gradients and parameters | gather params on demand |
| **Tensor parallel** | individual matrices across GPUs (within a node) | all-reduce per layer (needs NVLink) |
| **Pipeline parallel** | layers across GPUs/nodes; micro-batches fill the pipeline | activations between stages; "bubble" overhead |
| **Expert parallel** | MoE experts across GPUs | all-to-all token routing |

Large runs combine them ("3D parallelism"), plus mixed precision (bf16), activation checkpointing, and gradient accumulation to reach large effective batch sizes.

## Graph Neural Networks

**Q: How do GNNs work and where are they used?**

**Message passing**: each node updates its embedding by aggregating its neighbours' messages: `h_v' = σ(W·h_v + Σ_{u∈N(v)} W'·h_u / |N(v)|)` (GCN, GraphSAGE; GAT uses attention weights). k layers ⇒ k-hop receptive field.
Uses: fraud rings, recommendations (user–item graphs, PinSage), molecules/drug discovery, traffic/ETA, knowledge graphs.
Issues: **over-smoothing** with many layers, neighbourhood explosion → **neighbour sampling** (GraphSAGE) or cluster batching for large graphs.

## Object Detection & Segmentation

**Q: How do object detectors work? How are they evaluated?**

- **Two-stage** (Faster R-CNN): a region proposal network suggests boxes, then RoI Align crops features and a head classifies + refines each box. Accurate, slower.
- **One-stage** (YOLO, SSD, RetinaNet): predict class scores and box offsets densely over a grid of anchors in one pass; RetinaNet's **focal loss** fixes the extreme foreground/background imbalance. Fast, the default for real time.
- **Anchor-free** (FCOS, CenterNet) predict centers/distances directly; **DETR** treats detection as **set prediction** with a transformer and Hungarian matching between predictions and ground truth — no anchors, no NMS.
- Post-processing: score threshold + **non-maximum suppression** (see ML Coding).
- **Metrics**: IoU decides a true positive; **AP** = area under the precision–recall curve per class; **mAP** averages classes; COCO averages mAP over IoU thresholds 0.5:0.95.
- **Segmentation**: semantic (a class per pixel: U-Net encoder–decoder with skip connections, DeepLab with atrous convolutions), instance (Mask R-CNN adds a mask head), panoptic (both); **SAM** segments anything from point/box prompts.
- Practical: small objects (feature pyramids, higher resolution), class imbalance, label noise in boxes, deployment latency (quantization, smaller backbones).

## Model Compression: Distillation, Quantization, Pruning

**Q: Your model is too slow/big for production. What are your options?**

- **Knowledge distillation**: train a small student on the teacher's **soft targets** (probabilities at temperature T) plus the true labels: `L = α·CE(y, s) + (1 − α)·T²·KL(softmax(t/T) ‖ softmax(s/T))`. Soft targets carry "dark knowledge" (which wrong classes are close). LLMs distill from teacher-generated data or logits.
- **Quantization**: store/compute weights (and activations) in int8/int4/fp8 instead of fp32/bf16 → 2–8× smaller, faster on supporting hardware. **Post-training quantization** needs only calibration data; **quantization-aware training** simulates rounding during training for better accuracy. Per-channel scales, outlier handling (LLM.int8, SmoothQuant), GPTQ/AWQ for LLM weights.
- **Pruning**: remove weights (unstructured magnitude pruning: sparse but rarely faster on GPUs) or whole channels/heads/layers (**structured**: real speedups); NVIDIA 2:4 sparsity is a hardware-supported middle ground. Fine-tune after pruning.
- **Architecture**: efficient blocks (depthwise-separable convs, MobileNet/EfficientNet), smaller/shallower models, low-rank factorization, early exit, caching.
- **Serving tricks**: batching, compiling (TensorRT, ONNX Runtime, torch.compile), mixed precision, caching embeddings/results.
- Always measure the accuracy drop on the metric that matters and the real latency on the target hardware.

## Diffusion Models in Depth

**Q: How are diffusion models trained and sampled? What is classifier-free guidance?**

- **Forward process**: gradually add Gaussian noise; in closed form `x_t = √ᾱ_t · x₀ + √(1 − ᾱ_t) · ε`.
- **Training**: sample an image, a timestep t and noise ε; the network (U-Net or transformer/DiT) predicts ε from `(x_t, t, condition)`; loss = MSE between true and predicted noise. Simple and stable.
- **Sampling**: start from pure noise and iteratively denoise. DDPM needs hundreds to ~1000 steps; **DDIM** and ODE solvers take 20–50; distillation / consistency models get to 1–4 steps.
- **Conditioning**: text embeddings (CLIP/T5) injected via cross-attention. **Classifier-free guidance**: train with the condition randomly dropped, then sample with `ε = ε_uncond + w·(ε_cond − ε_uncond)`; larger w = closer prompt adherence, less diversity.
- **Latent diffusion** (Stable Diffusion): diffuse in a VAE's compressed latent space (e.g. 8× smaller per side) — far cheaper at high resolution.
- Newer: **flow matching / rectified flow** (straight-line paths, fewer steps), transformer backbones (DiT), video diffusion with temporal attention.
- Evaluation: FID (realism vs a reference set), CLIP score (prompt alignment), human preference ratings.

## Vision Transformers & CLIP

**Q: How do transformers handle images? How does CLIP enable zero-shot classification?**

- **ViT**: split the image into 16×16 patches, linearly embed each, add position embeddings and a [CLS] token, run a standard transformer encoder. Less inductive bias than CNNs → needs large data (or strong augmentation/distillation: DeiT); scales very well. Variants: Swin (windowed attention, hierarchical), MAE pretraining (mask 75% of patches, reconstruct).
- **CLIP**: an image encoder and a text encoder trained **contrastively** on hundreds of millions of image–caption pairs (InfoNCE over the batch: matching pairs on the diagonal).
- **Zero-shot classification**: embed prompts like "a photo of a {class}", embed the image, pick the most similar text — no task-specific training. Also powers image search, text-to-image conditioning and the vision side of many multimodal LLMs (SigLIP uses a sigmoid loss instead of softmax).
- Limitations: biases from web data, weak at counting/spatial relations, prompt sensitivity (prompt ensembles help).
