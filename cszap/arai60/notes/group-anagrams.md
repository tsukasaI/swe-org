# Group Anagrams (49)

- **Date**: 2026-03-17
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: HashMap のキー設計によるグループ化

## Memorize

- 各文字列を `[26]int` の文字頻度配列に変換し、HashMap のキーにする
- Go では固定長配列 `[26]int` を map のキーに使える（slice は不可）
- `map[[26]int][]string` でアナグラム同士が同じキーにグループ化される

## Understand

### キーの選択肢

| キー | Time | Space |
|---|---|---|
| `[26]int` 文字頻度 | O(N * M) | O(N) |
| ソートした文字列 | O(N * M log M) | O(N) |

- `[26]int` の方が Time が良い
- ソートは実装がシンプルだが、M が大きいと不利

### なぜ固定長配列がキーになるか

- Go の map キーは comparable であればよい — 固定長配列は comparable、slice は不可

## Connect

- 「同じ特徴を持つものをグループ化」→ HashMap のキー設計パターン
- Two Sum (1) と同じく HashMap で O(N) を実現する発想
