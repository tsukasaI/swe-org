# Unique Paths (62)

- **Date**: 2026-03-24
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 2D グリッド DP と空間最適化

## Memorize

- `grid[i][j] = grid[i-1][j] + grid[i][j-1]`（上 + 左）
- 1行目と1列目は全て 1（一直線にしか行けない）

## Understand

### 計算量

| アプローチ | Time | Space |
|---|---|---|
| 2D grid | O(N*M) | O(N*M) |
| 1D 配列 | O(N*M) | O(M) |

### 1D 配列での空間最適化

- `dp[j] = dp[j] + dp[j-1]` を左から右に更新
- `dp[j]`（更新前）= 上の行の値、`dp[j-1]`（更新済み）= 現在の行の左の値
- 上の行と現在の行の値が1つの配列に自然に共存する

### 障害物がある場合 (63)

- 障害物セルを 0 として扱う — そのセルからの経路数は 0

## Connect

- 2D DP の空間最適化パターン: 前の行しか使わない場合は 1D に圧縮可能
- 64 (Minimum Path Sum), 120 (Triangle) にも同じ最適化が適用できる
