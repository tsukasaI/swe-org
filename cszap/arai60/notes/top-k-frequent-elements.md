# Top K Frequent Elements (347)

- **Date**: 2026-03-15
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: Top K 問題を heap で最適化

## Memorize

- HashMap で頻度カウント → min heap サイズ k で上位 k 個を保持
- Push 後にサイズが k 超なら Pop — 703 と同じパターン
- 構造体 `Freq{v, freq}` で値と頻度をまとめて heap に入れる

## Understand

### 計算量の比較

| アプローチ | Time | Space |
|---|---|---|
| Sort | O(N log N) | O(N) |
| Max heap 全要素 | O(N log N) | O(N) |
| Min heap サイズ k | O(N log k) | O(N) |

### なぜ全要素 heap でも O(N log N) か

- N 個を push すると構築だけで O(N log N)
- サイズ k に制限すれば各 push/pop が O(log k) に改善

### よくあるミス

- Top K を求めるのに max heap を使ってしまう — サイズ k で保つなら min heap が正しい
- 型名と Less の向きの不一致（MaxHeap なのに `<`）— 面接で指摘される

## Connect

- 703 (Kth Largest in Stream) と同じ「min heap サイズ k」パターン
- Bucket sort で O(N) も可能（頻度を index とした配列）— さらなる Follow-up として知っておくとよい
