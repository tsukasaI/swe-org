# Word Ladder (127)

- **Date**: 2026-03-21
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: BFS で最短変換列を求める

## Memorize

- 最短経路 → BFS（DFS ではない）
- 隣接ノードの探索: 各文字を a-z に置き換えて set で存在確認（全ペア比較 O(N^2) より効率的）
- BFS のレベル管理: `size := len(queue)` で現在レベルの要素数を取り、`for range size` で処理後に count++
- endWord が wordList にない場合は即 `return 0`
- 訪問済み管理: set から delete する（別途 visited 不要）

## Understand

### Time / Space

- Time: O(N * L^2) — N 単語 × L 文字 × 26 置換 × L（文字列生成・比較）
- Space: O(N * L) — set とキュー

### BFS レベル管理パターン

```
count := 1
for len(queue) > 0 {
    size := len(queue)
    for range size {
        // 現在レベルの全要素を処理
    }
    count++  // レベルを進める
}
```

- count++ の位置を間違えると距離がずれる
- endWord を見つけたときは `count + 1`（次のレベルの単語なので）

### つまずきポイント

- sliding window 的にキューを1つずつ処理するのではなく、レベル単位で処理する発想が重要
- endWord が wordList にない場合のガード忘れ
- `break` で内側ループを抜けただけでは return できない

## Connect

- BFS 最短経路パターン: 200 (Number of Islands) の BFS 版と同じキュー操作
- set + delete で訪問管理: 349 (Intersection of Two Arrays) と同じ発想
- 隣接ノード生成: 全ペア比較 vs 文字置換 — 制約に応じて効率的な方を選ぶ
