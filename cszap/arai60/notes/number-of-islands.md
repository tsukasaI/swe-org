# Number of Islands (200)

- **Date**: 2026-03-20
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: Grid 探索の DFS/BFS パターン

## Memorize

- grid 全体を走査 → `'1'` を見つけたら count++ → 繋がる `'1'` を全て `'0'` に書き換え
- 訪問済みを `'0'` に上書きすれば visited 配列不要
- 境界チェック: `i >= len(grid)` ではなく `>=`（0始まり）
- 4方向: 上下左右 `(i-1,j), (i+1,j), (i,j-1), (i,j+1)`

## Understand

### DFS（再帰）

- 再帰で4方向に展開、境界外 or `'0'` なら return
- Time O(N*M), Space O(N*M)（コールスタック、最悪ケースで全セルが `'1'`）

### BFS（キュー）

- キューに座標を入れて FIFO で処理
- Time O(N*M), Space O(N*M)（最悪ケース）
- 実際はキューに入るのは探索中の境界だけなので、コールスタックほど深くならないことが多い
- stack overflow のリスクがない — 大きい grid では BFS が安全

### よくあるミス

- 隣接セルの追加時に外側ループの `i, j` を使ってしまう → キューから取り出した座標を使うこと

## Connect

- Grid DFS/BFS パターン: 695 (Max Area of Island), 733 (Flood Fill), 994 (Rotting Oranges) など
- DFS は再帰 or 明示的スタック、BFS はキュー — 問題の性質で使い分け
