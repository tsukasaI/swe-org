# Subarray Sum Equals K (560)

- **Date**: 2026-03-20
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: Prefix sum + HashMap でサブ配列の和を効率的にカウント

## Memorize

- `prefixSum[r] - prefixSum[l] = k` → `prefixSum[r] - k = prefixSum[l]`
- HashMap に累積和の出現回数を記録、各ステップで `sum - k` を探す
- `hm[0] = 1` を初期値として入れる — 先頭から始まる部分配列（`sum == k`）を拾うため
- `result += hm[sum-k]`（`++` ではない）— 同じ prefixSum が複数回出現しうる

## Understand

### なぜ sliding window が使えないか

- `nums[i]` に負の数がありえる
- r を進めても和が大きくなる保証がなく、l を進めても和が小さくなる保証がない
- sliding window は単調性が前提

### prefix sum + HashMap のパターン

- Two Sum (1) の「補数を HashMap で探す」と同じ発想
- 違い: Two Sum はペアを1つ見つければよいが、この問題は全カウントなので出現回数を記録する

## Connect

- Prefix sum + HashMap: 連続部分配列の和に関する問題全般に適用（523, 974 など）
- Sliding window が使える条件: 要素が非負、または単調性が保証されている場合
