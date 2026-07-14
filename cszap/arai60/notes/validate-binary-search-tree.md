# Validate Binary Search Tree (98)

- **Date**: 2026-03-22
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: BST の検証 — 範囲の伝播パターン

## Memorize

- ヘルパー `dfs(node, lower, upper)` で各ノードの取りうる値の範囲を持ち回る
- `node.Val <= lower || node.Val >= upper` なら false
- 左の子: upper を node.Val に、右の子: lower を node.Val に狭める
- 直接の親子チェックは不要 — 範囲チェックが全てカバーする

## Understand

### 計算量

- Time: O(N) — 全ノード訪問
- Space: O(N) — コールスタック（最悪ケース）

### よくあるミス

- 直接の親子だけチェックする → 祖先との関係を見逃す（例: 右部分木に root より小さい値）
- 初期値の min/max に `Node.val` の範囲と同じ値を使うと、境界値で壊れる

### 別解: in-order traversal

- BST を 左→自分→右 の順で訪問すると結果は必ず昇順
- 前のノードの値より現在の値が大きいかを確認するだけで検証できる
- min/max の初期値問題が発生しないメリットがある

## Connect

- 範囲の伝播パターン: 再帰で制約を狭めていく手法は BST 系問題に共通
- in-order traversal: 98, 230 (Kth Smallest Element in a BST), 二分探索木の基本操作
