# First Unique Character in a String (387)

- **Date**: 2026-03-19
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 文字頻度カウントによる最初のユニーク文字探索

## Memorize

- 2パス: 1パス目で頻度カウント、2パス目で最初の count==1 を返す
- `[26]int` と `map[rune]int` どちらでも解ける

## Understand

### `[26]int` vs `map`

- `[26]int`: 連続メモリ、インデックスアクセス（直接アドレス計算）で定数倍が小さい
- `map`: ハッシュ計算、バケット探索、衝突処理のオーバーヘッド
- 入力が lowercase English letters に限定 → 固定長配列が最適
- どちらも Time O(N), Space O(1)（26文字が上限）

## Connect

- 49 (Group Anagrams) でも `[26]int` を使った — lowercase 限定なら配列が速いパターン
- 文字頻度カウントは文字列問題の基本操作
