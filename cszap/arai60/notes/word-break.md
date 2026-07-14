# Word Break (139)

- **Date**: 2026-03-26
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 文字列分割の DP

## Memorize

- `dp[i]` = `s[:i]` が辞書の単語で分割可能か（bool）
- `dp[0] = true`（空文字列）、`dp` の長さは `len(s) + 1`
- `dp[i]` が true かつ `s[i:j]` が辞書にあれば `dp[j] = true`
- 内側ループは `j = i+1` から `j <= len(s)` まで（`<` だと末尾を見逃す）

## Understand

### 計算量

- Time: O(N^2)（ナイーブ）→ O(N * W) に最適化可能（W = 辞書の最大単語長）
- Space: O(N + M)（dp 配列 + 辞書 set）
- 辞書の最大単語長が 20 なので `j` を `min(i+21, len(s))` に制限 → 内側ループ最大 20 回

### off-by-one エラーの防ぎ方

- `dp` のサイズが `len(s) + 1` で `dp[0]` が空文字列 → インデックスが文字列と1つずれる
- 小さい例（`s="ab"`, `wordDict=["ab"]`）で dp と文字列スライスの対応を紙に書いて確認してからコードに入る

### つまずきポイント

- `dp` の長さを `len(s)` にすると `dp[0]` が空文字列ではなく `s[:1]` に対応してしまう
- 内側ループの `j < len(s)` だと最後の単語を見逃す → `j <= len(s)`
- `s[i:i]` は空文字列 → `j` は `i+1` から始める

## Connect

- 140 (Word Break II): 全ての分割パターンを列挙 — backtracking + memoization
- 辞書 set パターン: 127 (Word Ladder) と同じく set lookup で効率化
