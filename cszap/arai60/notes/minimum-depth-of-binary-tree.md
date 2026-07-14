# Minimum Depth of Binary Tree (111)

- **Date**: 2026-03-21
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: maxDepth との違いを理解する

## Memorize

- leaf node = 左右どちらも nil のノード
- 片方の子が nil の場合、nil 側は leaf ではないのでそちらのパスを無視して反対側だけ再帰する
- maxDepth では不要な処理 — `min(0, right+1)` だと nil を最短として誤カウントするため

## Understand

### maxDepth との違い

- maxDepth: `max(left, right) + 1` だけでよい（nil が 0 を返しても max で無視される）
- minDepth: 片方が nil のとき `min(0, right+1) = 1` となり、nil パスを最短と誤判定する
- そのため片方 nil のケースを分岐して、子がある側のみ再帰する必要がある

### 計算量

- Time: O(N) — 全ノード訪問
- Space: O(H) — コールスタック（最悪 O(N)、バランス木 O(log N)）

## Connect

- 104 (Maximum Depth) の対になる問題
- BFS で解くと最初の leaf に到達した時点で return でき、偏った木では DFS より効率的
