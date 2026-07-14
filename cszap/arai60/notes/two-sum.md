# Two Sum (1)

- **Date**: 2026-03-17
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: HashMap による O(N) ペア探索

## Memorize

- HashMap に `target - v` をキー、index を値として格納
- 各要素を見るとき、自身が HashMap に存在すれば相方が既に見つかっている
- Time O(N), Space O(N)

## Understand

### 代替アプローチ: ソート + バイナリサーチ

- `(value, index)` のペア配列を作ってソート → 各要素に対してバイナリサーチ
- Time O(N log N), Space O(N)
- ソートするとインデックスが変わるので元のインデックスを保持する必要がある
- HashMap 解法の方が Time が良いので、こちらを選ぶ理由を説明できればよい

## Connect

- HashMap でペア探索: Three Sum (15), Four Sum (18) の基礎
- 「補数を HashMap に入れる」パターンは頻出
