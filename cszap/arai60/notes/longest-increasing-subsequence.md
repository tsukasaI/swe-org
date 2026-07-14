# Longest Increasing Subsequence (300)

- **Date**: 2026-03-24
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: DP と二分探索の2つのアプローチ

## Memorize

### DP 解法 (O(N^2))

- `dp[i]` = `nums[i]` を末尾とする最長増加部分列の長さ
- 初期値: 全て 1（自分単体）
- `dp[i] = max(dp[i], dp[j]+1)` for all `j < i` where `nums[j] < nums[i]`
- 答えは `dp` の最大値

### tails + 二分探索 (O(N log N))

- `tails[i]` = 長さ `i+1` の増加部分列の末尾として取りうる最小値
- 新しい値が tails 末尾より大きい → 追加（部分列の長さが伸びた）
- そうでなければ → 二分探索で置き換え位置を探して置き換え（末尾を小さく保つ）
- `sort.SearchInts(tails, nums[i])` で「nums[i] 以上の最初の位置」を取得
- 答えは `len(tails)`

## Understand

### 計算量

| アプローチ | Time | Space |
|---|---|---|
| DP | O(N^2) | O(N) |
| tails + 二分探索 | O(N log N) | O(N) |

### なぜ tails を置き換えるか

- 末尾を小さく保つほど、後から続く要素を追加しやすくなる
- tails は常にソート済みなので二分探索が使える
- tails は実際の部分列そのものではない — 長さだけが正確

### strictly increasing → non-decreasing に変わったら

- DP: `nums[j] < nums[i]` を `nums[j] <= nums[i]` に変更

## Connect

- DP の「i より前の全 j を探す」パターン: Edit Distance (72) なども同じ構造
- 二分探索 + 貪欲: Patience Sorting アルゴリズムと同じ原理
