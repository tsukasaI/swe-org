# Binary Tree Zigzag Level Order Traversal (103)

- **Date**: 2026-03-21
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: BFS レベル走査の応用 — 交互に方向を変える

## Memorize

- 102 (Level Order) と同じ BFS、違いは偶数レベルで element を逆順にするだけ
- キューへの追加は**常に左→右** — 追加順を変えると次レベル以降が崩れる
- `slices.Reverse(element)` で逆順にしてから result に追加

## Understand

### 計算量

- Time: O(N) — reverse は各レベルで O(レベル幅)、全レベル合計で O(N)
- Space: O(N) — キュー

### 逆順の実装方法

| 方法 | メリット | デメリット |
|---|---|---|
| `slices.Reverse` | シンプル、既存の Level Order に3行追加するだけ | 追加のパス |
| index で逆から埋める | reverse 不要、1パス | ループ内の分岐が増える |
| 先頭に挿入 `append([]int{v}, e...)` | — | 毎回全要素コピーで O(N^2) |

### よくあるミス

- fromLeft でキューへの追加順（左右）を変えてしまう → 次レベル以降のノード順が壊れる
- 値の収集と次レベルの構築は分離して考える

## Connect

- 102 (Level Order) の直接的な応用
- 107 (Level Order II): 結果を逆順 — 同じく result の後処理だけで解ける
