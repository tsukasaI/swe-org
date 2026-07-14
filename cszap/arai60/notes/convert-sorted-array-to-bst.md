# Convert Sorted Array to Binary Search Tree (108)

- **Date**: 2026-03-21
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: ソート済み配列から height-balanced BST を構築

## Memorize

- 中央要素を root にして、左右を再帰的に処理
- `center := len(nums) / 2` — 奇数偶数の場合分け不要（Go の整数除算で切り捨て）
- ベースケース: `len(nums) == 0` → nil

## Understand

### 計算量

- Time: O(N) — 全ノードを1回ずつ作成
- Space: O(log N) — コールスタック（常にバランスするので log N が保証）

### なぜ height-balanced が保証されるか

- 常に中央を root にするので左右の要素数の差は最大1
- 再帰的に同じ操作を繰り返すので全レベルでバランスする

### Go のスライスとメモリ

- `nums[:center]` はコピーではなく元の配列を共有する（スライスヘッダ: ポインタ、長さ、容量のみ新規作成）
- 読み取りだけなら共有のまま、`append` で容量を超えると新しい配列が確保される
- index を引数に渡す方法でも書けるが、スライスで十分（コピーは発生しない）

## Connect

- 109 (Convert Sorted List to BST): linked list 版 — ランダムアクセスできないので工夫が必要
- 二分探索と同じ「中央を選んで左右に分割」のパターン
