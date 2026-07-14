# Unique Email Addresses (929)

- **Date**: 2026-03-17
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 文字列処理 + Set でユニークカウント

## Memorize

- `@` で分割 → local の `+` 以降を除去 → `.` を除去 → domain と結合して set に入れる
- `strings.Split`, `strings.ReplaceAll` で十分シンプルに書ける

## Understand

- Time O(N * M)（N: メール数、M: メール文字列長）, Space O(N)
- `fmt.Sprintf` vs `+` 結合: Sprintf はリフレクション+フォーマット解析で遅い、`+` の方が速い。ただし制約が小さければ可読性で選んでよい

## Connect

- Set でユニークカウント: 349 (Intersection of Two Arrays) と同じパターン
- 文字列の正規化 → HashMap/Set のキーにする: 49 (Group Anagrams) と同じ発想
