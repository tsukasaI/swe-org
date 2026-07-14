# Number of Connected Components in an Undirected Graph (323)

- **Date**: 2026-03-21
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: グラフの連結成分を DFS でカウント

## Memorize

- edges から隣接リスト `map[int][]int` を構築（無向なので双方向に追加）
- 0 から n-1 を走査、未訪問なら count++ して DFS で繋がるノードを全て訪問済みにする
- Number of Islands (200) と同じパターン: grid → ノードとエッジに変わっただけ

## Understand

### 隣接リストの構築

```go
adjacentList := make(map[int][]int)
for _, v := range edges {
    adjacentList[v[0]] = append(adjacentList[v[0]], v[1])
    adjacentList[v[1]] = append(adjacentList[v[1]], v[0])
}
```

### 訪問管理

- `visited` の set（`map[int]struct{}`）で管理
- adjacentList の delete ではなく visited を使う理由: 孤立ノード（edges がない）は adjacentList に存在しないため、delete 方式だとカウントできない
- DFS の先頭で visited チェック — 無限ループ防止

### 計算量

- Time: O(N + E) — 各ノード1回訪問、各辺1回たどる
- Space: O(N + E) — 隣接リスト O(N + E)、visited O(N)、コールスタック O(N)

### つまずきポイント

- 孤立ノードも1つの連結成分 — メインループは adjacentList ではなく全ノード（0..n-1）を走査する
- 訪問済みチェックがないと無向グラフで無限再帰になる

## Connect

- 200 (Number of Islands) の汎用グラフ版
- Union-Find でも解ける — 辺を順に union して最後に連結成分数を返す
