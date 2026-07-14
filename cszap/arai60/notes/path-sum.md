# Path Sum (112)

- **Date**: 2026-03-21
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: root-to-leaf パスの合計判定

## Memorize

- targetSum から各ノードの Val を引いていき、leaf で残りが Val と一致するか判定
- leaf チェック必須: `root.Left == nil && root.Right == nil`
- `root == nil` で `targetSum == 0` を返す実装だと、leaf でないノードの先の nil で誤判定する

## Understand

### 計算量

- Time: O(N) — 全ノード訪問
- Space: O(H) — コールスタック（最悪 O(N)）

### なぜ leaf チェックが必要か

- `root = [1,2], targetSum = 1` の場合、root(1) の右は nil
- nil で `targetSum == 0` を返すと、leaf を通過していないのに true になる
- 「nil に到達」と「leaf を通過した」は別

## Connect

- 437 (Path Sum III): 任意ノード → 任意ノードのパス — prefix sum + HashMap（560 Subarray Sum Equals K と同じパターンを木に適用）
- 113 (Path Sum II): root-to-leaf パスを全列挙 — backtracking で経路を記録
