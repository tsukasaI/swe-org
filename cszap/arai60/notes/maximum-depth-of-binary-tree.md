# Maximum Depth of Binary Tree (104)

- **Date**: 2026-03-21
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 再帰で木の深さを求める基本パターン

## Memorize

- `nil` なら 0 を返し、左右の深さの max + 1 を返す
- 二分木の全ノード訪問 → Time は O(N)（二分探索 O(log N) と混同しない）

## Understand

### 計算量

- Time: O(N) — 全ノードを1回ずつ訪問
- Space: O(H) where H = 木の高さ
  - バランス木: O(log N)
  - 偏った木（全ノードが片側）: O(N)

### 二分木 vs 二分探索木

- 二分木: 全ノード訪問が必要 → O(N)
- 二分探索木（BST）: 探索は O(log N) だが、それは値を探すとき — 全ノード処理は同じく O(N)

## Connect

- 木の再帰パターンの基本: 110 (Balanced Binary Tree), 543 (Diameter of Binary Tree) などに応用
- BFS（レベル管理）でも解ける — 127 (Word Ladder) で使ったレベル管理パターン
