# Coin Change (322)

- **Date**: 2026-03-27
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: DP で最小コイン数を求める

## Memorize

- `dp[i]` = 金額 `i` を作るのに必要な最小コイン数
- `dp[0] = 0`、それ以外は `amount + 1` で初期化（到達不可能を表す）
- `dp[i] = min(dp[i], dp[i-coin] + 1)` for each coin where `i - coin >= 0`
- 最後に `dp[amount] > amount` なら `-1`

## Understand

### 計算量

- Time: O(A * C)（A = amount, C = コイン種類数）
- Space: O(A)

### 初期値の選び方

- `amount + 1`: 安全 — ありえない大きな値として機能し、+1 してもオーバーフローしない
- `math.MaxInt`: NG — `dp[i-coin] + 1` でオーバーフローして負の値になり min が壊れる
- `0` や `-1`: `dp[0] = 0` が有効な値と区別できず条件分岐が複雑になる

### つまずきポイント

- `dp[i-coin] > 0` のチェックを入れると `dp[0] = 0` を使えなくなる
- `dp[0] = 0` は「0枚で金額0を作れる」という有効な値

## Connect

- 518 (Coin Change II): 組み合わせ総数 — `min` → `+=` に変わる（`dp[i] += dp[i-coin]`）
- 同じ「1D DP で全コインを試す」パターン: 完全ナップサック問題の典型
