# Binary Tree Level Order Traversal (102)

- **Date**: 2026-03-21
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: BFS でレベルごとにノードを収集

## Memorize

- BFS でレベル管理、各レベルの値を配列にまとめて result に追加
- レベル管理の方法2つ:
  - `nextQueue` で次レベルを別に構築 — 読みやすい
  - `size := len(queue)` で現在レベルの要素数を取る — メモリ割り当てが少ない

## Understand

### 計算量

- Time: O(N) — 全ノード訪問
- Space: O(N) — キューに最大 N/2 個（バランス木の最下層）が入る
  - O(log N) ではない — キューの幅は木の高さではなくレベルの幅に依存

### 逆順レベル走査 (107)

- 結果を逆順にして返すだけでよい

## Connect

- BFS レベル管理パターン: 127 (Word Ladder) と同じ
- 103 (Zigzag Level Order): 偶数レベルを逆順にする
- 107 (Level Order Traversal II): 結果を逆順に返す
