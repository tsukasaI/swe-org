# Kth Largest Element in a Stream (703)

- **Date**: 2026-03-15
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: Heap を使ったストリーム処理の効率化

## Memorize

- Min heap をサイズ k で保つ → top が常に k 番目に大きい要素
- `Add`: push → サイズが k 超なら pop → top を返す
- Go の `container/heap`: `heap.Interface` を実装（Len, Less, Swap, Push, Pop）
- Min heap: `Less` で `h[i] < h[j]`、Max heap: `h[i] > h[j]`
- Heap の Push/Pop は O(log N) — ソート全体 O(N log N) ではない

## Understand

### なぜ min heap でサイズ k か

- 大きい方から k 個だけ保持 → 最小値（top）が k 番目に大きい要素
- Max heap だと内部配列はソート順不定、index アクセスで k 番目は取れない

### 計算量の比較

| アプローチ | Add の Time | Space |
|---|---|---|
| Sort 毎回 | O(N log N) | O(N) |
| Min heap サイズ k | O(log k) | O(k) |

### 面接での進め方

- まず動く解法（sort）を出してから heap に最適化 — 面接で好印象

## Connect

- Heap パターン: Top K 系問題に共通 (215. Kth Largest Element in an Array, 347. Top K Frequent Elements)
- Go の heap は interface ベースなので型名と Less の向きを一致させる（MinHeap なら `<`）
