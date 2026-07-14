# House Robber II (213)

- **Date**: 2026-03-25
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 円形制約の House Robber

## Memorize

- `nums[0:n-1]` と `nums[1:n]` の2回 House Robber を実行して max を取る
- `len(nums) == 1` のガードが必要 — 分割すると両方空になる

## Understand

### 計算量

- Time: O(N), Space: O(N)

### なぜ2回に分けるだけで正しいか

- 円形なので最初と最後が隣接
- 最適解は「最初を盗む（最後は盗めない）」か「最初を盗まない（最後は盗める）」のどちらか
- 両方試して max を取れば全ケースをカバー

## Connect

- 198 (House Robber) の直接的な拡張 — robLinier をそのまま再利用
- 円形制約を線形に分解するパターン: 918 (Maximum Sum Circular Subarray) にも同じ発想
