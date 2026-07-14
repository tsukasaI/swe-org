# Best Time to Buy and Sell Stock II (122)

- **Date**: 2026-03-26
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 複数回売買可能な最大利益

## Memorize

- 上がった日の差分を全て加算するだけ
- `prices[i] - prices[i-1] > 0` なら `maxProfit += prices[i] - prices[i-1]`
- Time O(N), Space O(1)

## Understand

### なぜ全ての上昇を取れるか

- 何回でも売買可能 + 同日に売って買い直せる
- 連続上昇 `[1,2,3]` は `(2-1) + (3-2) = 3-1` と等価 — 分割しても一括でも同じ利益
- 下がる日はスキップすれば損失を回避

### 121 との違い

- 121: 1回だけ → minPrice を追跡して最大差を求める
- 122: 複数回 → 全ての正の差分を貪欲に加算

## Connect

- 121 (Best Time I): 1回だけの売買
- 309 (Best Time with Cooldown): 売った翌日は買えない → 状態 DP が必要
