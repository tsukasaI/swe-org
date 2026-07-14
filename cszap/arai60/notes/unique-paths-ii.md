# Unique Paths II (63)

- **Date**: 2026-03-25
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 障害物ありの 2D グリッド DP

## Memorize

- 62 (Unique Paths) と同じ 1D DP、障害物セルは `dp[j] = 0` にする
- 1行目の初期化: 障害物があったら `break` でそれ以降は 0 のまま（到達不可能）
- `dp[j] = dp[j-1] + dp[j]`（障害物でなければ）

## Understand

### 計算量

- Time: O(M*N), Space: O(N)

### 障害物の処理

- 障害物セル: `dp[j] = 0` — そのセルを経由する経路は 0
- 1行目: 障害物より右は1行目だけでは到達不可能 → 0
- スタート地点が障害物: `dp[0] = 0` のまま全て 0 → 正しく 0 を返す

### 62 との違い

- 動き方（右・下のみ）は同じなので 1D DP の空間最適化がそのまま適用可能
- 障害物のあるなしだけが差分 — 漸化式に条件分岐を追加するだけ

## Connect

- 62 (Unique Paths) の直接的な拡張
- 64 (Minimum Path Sum): 同じグリッド DP で経路の最小コストを求める
