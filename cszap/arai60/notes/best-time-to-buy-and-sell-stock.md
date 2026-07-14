# Best Time to Buy and Sell Stock (121)

- **Date**: 2026-03-26
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 最小値の追跡で最大利益を求める

## Memorize

- `minPrice` を追跡しながら `prices[i] - minPrice` の最大値を取る
- dp 配列は不要 — 前の最大利益を変数で保持するだけで O(1) space

## Understand

### 計算量

- Time: O(N), Space: O(1)（変数2つ: minPrice, maxProfit）

### dp 配列版

- `dp[i] = max(dp[i-1], prices[i] - minPrice)` — 動くが Space O(N) で冗長
- `dp[i]` は `dp[i-1]` しか参照しないので変数1つに圧縮可能

## Connect

### 122 (Best Time to Buy and Sell Stock II): 複数回売買可能

- 「上がった日」の差分を全て加算するだけ
- `prices[i] - prices[i-1] > 0` なら利益に加算
- 何回でも売買できるので、全ての上昇を取れる

### 関連問題

- 123 (Best Time III): 最大2回 — 状態を増やした DP
- 188 (Best Time IV): 最大 k 回
- 309 (Best Time with Cooldown): 売った翌日は買えない制約
