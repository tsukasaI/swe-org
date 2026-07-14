# Intersection of Two Arrays (349)

- **Date**: 2026-03-17
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: Set を使った配列の共通要素抽出

## Memorize

- Set (map[int]struct{}) で片方の配列をユニーク化 → もう片方を走査して存在チェック
- 重複排除の方法2つ:
  - 2つ目の set (`intersectionMap`) で結果の重複を防ぐ — 意図が明確
  - `delete(num1Set, v)` で一度見つけたら削除 — set 1つで済む、Space 削減

## Understand

- Time O(N+M), Space O(N) (delete 版) or O(N+M) (2 set 版)
- delete 版は `num1Set` を破壊するので、後で再利用できない点に注意

## Connect

- Set パターン: Two Sum (1) の HashMap と同じく「存在チェックを O(1) にする」発想
- 350 (Intersection of Two Arrays II) では重複も含めて返す — set ではなく map[int]int で頻度カウントが必要
