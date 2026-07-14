# Max Area of Island (695)

- **Date**: 2026-03-20
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: DFS で島の面積を計算する応用

## Memorize

- Number of Islands (200) と同じ DFS パターン、違いは dfs が面積（int）を返す点
- `return 1 + dfs(上) + dfs(下) + dfs(左) + dfs(右)` で繋がるセル数を再帰的に合計
- 結果は `max(dfs(i,j), area)` で最大値を取る（`+=` で合計しない）

## Understand

- Time O(N*M), Space O(N*M)（コールスタック）
- 200 では count++ だけでよかったが、今回は各島の面積が必要 → dfs の戻り値を活用

## Connect

- 200 (Number of Islands) の直接的な応用
- 同じパターンで解ける: 463 (Island Perimeter), 827 (Making A Large Island)
