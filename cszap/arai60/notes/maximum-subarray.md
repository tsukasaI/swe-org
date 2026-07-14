# Maximum Subarray (53)

- **Date**: 2026-03-24
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: Kadane's algorithm で最大部分配列和を O(N) で求める

## Memorize

- 累積和が負になったらリセット（0 に戻す）— 負の累積和を引き継いでも最大にならない
- `maxSum` を毎ステップ更新、`sum` をリセット前に比較する
- 初期値は `math.MinInt` が安全（制約依存の値は避ける）

## Understand

### 計算量

- Time: O(N), Space: O(1)
- O(N^2) は `10^5` でタイムアウト — 目安として `10^8` が操作回数の上限

### 開始・終了 index も返す場合

- `l, r` ポインタを追加
- `sum < 0` でリセットするとき `l` を次の要素に更新
- `maxSum` を更新するとき `r` を現在の index に更新、`l` も記録

### Kadane's algorithm の本質

- 各要素で「今の部分配列を続けるか、ここから新たに始めるか」を判断
- `sum = max(nums[i], sum + nums[i])` とも書ける

## Connect

- 918 (Maximum Sum Circular Subarray): Kadane's の応用 — 最小部分配列和を引く方法
- 152 (Maximum Product Subarray): 積の場合は負×負=正があるので min/max 両方追跡
