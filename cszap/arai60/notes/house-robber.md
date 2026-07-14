# House Robber (198)

- **Date**: 2026-03-25
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 隣接制約のある線形 DP

## Memorize

- `dp[i] = max(dp[i-2] + nums[i], dp[i-1])`（盗むか盗まないか）
- `dp[0] = nums[0]`, `dp[1] = max(nums[0], nums[1])`
- `dp[1]` を `nums[1]` にすると `[2, 1]` のようなケースで失敗する

## Understand

### 計算量

- Time: O(N), Space: O(N)
- Space O(1) に改善可: 前の2つの値だけ変数で保持すればよい

### よくあるミス

- `dp[1] = nums[1]` にしてしまう → `nums[0] > nums[1]` のケースで最大値を見逃す

## Connect

- 213 (House Robber II): 円形 — `nums[0:n-1]` と `nums[1:n]` の2回実行して max を取る
- 276 (Paint Fence) と同じ「前の状態に依存する線形 DP」パターン
- 337 (House Robber III): 木構造版 — DFS + DP
