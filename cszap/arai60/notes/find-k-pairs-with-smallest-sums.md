# Find K Pairs with Smallest Sums (373)

- **Date**: 2026-03-15
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: ソート済み配列のペア問題を heap で効率的に解く

## Memorize

- 初期化: `nums1[i] + nums2[0]` を `min(len(nums1), k)` 個だけ min heap に入れる
- k 回 pop し、pop した `(i, j)` に対して `j+1` が範囲内なら `(i, j+1)` を push
- Node に `i, j` を持たせて pop 後に次の候補を特定する
- min heap なので pop は常に最小 → 全ペアを列挙せず小さい順に k 個取れる

## Understand

### なぜ全ペア列挙が不要か

- 両配列がソート済みなので、`(i, j)` を pop したとき次の候補は `(i, j+1)` だけでよい
- `(i+1, j)` は別の `nums1[i+1]` の行として初期化時にすでに入っている

### 計算量

| アプローチ | Time | Space |
|---|---|---|
| 全ペア + sort | O(MN log MN) | O(MN) |
| 全ペア + max heap サイズ k | O(MN log k) | O(k) |
| Min heap + 逐次展開 | O(k log k) | O(k) |

### 初期化で全要素を入れない理由

- k 個しか結果が不要なので `nums1[k]` 以降は結果に入らない
- min heap で全要素を入れると pop が最小を捨ててしまうので、max heap のサイズ k 制限とは挙動が異なる

## Connect

- 「ソート済みデータから k 個」→ heap で逐次展開パターン: Merge K Sorted Lists (23) も同じ発想
- Top K 系でも min heap サイズ k（703, 347）と min heap 逐次展開（373）で使い分けが必要
