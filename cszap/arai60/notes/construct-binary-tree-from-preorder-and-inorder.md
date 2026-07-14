# Construct Binary Tree from Preorder and Inorder Traversal (105)

- **Date**: 2026-03-22
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 走査結果から木を再構築するパターン

## Memorize

- preorder の先頭が root → inorder で root の位置を探す → 左右に分割して再帰
- 分割の鍵: inorder の root 左側の要素数 = preorder の左部分木の要素数
  - 左の preorder: `preorder[1 : 1+leftCount]`
  - 右の preorder: `preorder[1+leftCount:]`
- preorder の並び順は「root, 左全部, 右全部」と保証されているので、要素数でスライス分割可能

## Understand

### 計算量

| アプローチ | Time | Space |
|---|---|---|
| `slices.Index` 毎回 | O(N^2) | O(N) |
| 事前に HashMap (値→index) | O(N) | O(N) |

### 3種類の走査順

- **preorder**: root → 左 → 右 — root は配列の先頭
- **inorder**: 左 → root → 右 — root の位置で左右の部分木を分割
- **postorder**: 左 → 右 → root — root は配列の末尾

### よくあるミス

- preorder の分割で map を使って左右を振り分ける → root 自身が混入する可能性
- 要素数ベースのスライス分割がシンプルで正確

## Connect

- 106 (Construct from Inorder and Postorder): root が postorder の末尾に変わるだけ、考え方は同じ
- 木の走査順の理解は 98 (Validate BST) の in-order 別解にも繋がる
