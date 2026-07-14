# Merge Two Binary Trees (617)

- **Date**: 2026-03-21
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 2つの木を同時に再帰で走査するパターン

## Memorize

- 2つのノードを同時に再帰で降りる
- ベースケース3パターン:
  - 両方 nil → nil を返す
  - 片方 nil → nil でない方をそのまま返す
  - 両方存在 → Val を足して新ノードを作り、Left/Right を再帰

## Understand

- Time: O(min(N, M)) — 重なっている部分だけ訪問、片方 nil ならそのまま返す
- Space: O(min(N, M)) — 再帰の深さは重なり部分の高さに依存
- 新しいノードを作るので元の木を破壊しない

## Connect

- 2つの木を同時に走査するパターン: 100 (Same Tree), 101 (Symmetric Tree) にも応用
- 104 (Max Depth), 111 (Min Depth) と同じ再帰の基本構造
