# Split BST (776)

- **Date**: 2026-04-25
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: BST の性質を活かして O(h) で 2 つに分割する

## Memorize

```go
func splitBST(root *TreeNode, target int) (*TreeNode, *TreeNode) {
    if root == nil {
        return nil, nil
    }
    if root.Val <= target {
        smaller, larger := splitBST(root.Right, target)
        root.Right = smaller
        return root, larger
    }
    smaller, larger := splitBST(root.Left, target)
    root.Left = larger
    return smaller, root
}
```

## Understand

### 計算量

- Time: **O(h)** where h は木の高さ
  - 平衡 BST: O(log N)
  - 歪んだ BST: O(N)
- Space: **O(h)** (再帰スタック)

### BST の性質を活かす理由

各再帰で `root.Val` と `target` を比較すると、片側の部分木の所属が**確定**する:
- `root.Val <= target` → root と root.Left は確実に smaller 側、root.Right が混在
- `root.Val > target` → root と root.Right は確実に larger 側、root.Left が混在

→ **片側の部分木だけ再帰**すれば良い (両方降りない) → O(h)

普通の二分木では順序がないので、各ノードを個別にチェックする必要がある → O(N)。

### アルゴリズムの構造

```
1. base case: root == nil → return (nil, nil)
2. root の所属を target との比較で決定
3. 混在する側の部分木を再帰で分割
4. root の該当子リンクを再帰結果で繋ぎ直す
5. (smaller_root, larger_root) を返す
```

### リンク繋ぎ直しの具体例

`root.val=4, target=2` のとき:
- 4 は target より大 → 4 は larger 側
- 4.Left = [2,1,3] を再帰分割 → smaller=[2,1], larger=[3]
- 4 の左に larger=[3] を繋ぐ:
  ```
        4
       / \
      3   6
         |\
         5 7
  ```
- 結果: smaller=[2,1], larger=[4,3,6,null,null,5,7]

### よくある間違い

- root を結果から消してしまう（`return root.Left, splitBST(root.Right, target)` のように）
  - root も結果に含める必要がある
- 多値返却を return 文に直接埋め込む（Go ではコンパイルエラー）
  - 一度変数で受けてから処理する

### 3 分割への拡張

「target 未満」「target と等しい」「target より大」の 3 グループに分ける:

```go
func splitBST3(root *TreeNode, target int) (*TreeNode, *TreeNode, *TreeNode) {
    if root == nil {
        return nil, nil, nil
    }
    if root.Val < target {
        less, equal, greater := splitBST3(root.Right, target)
        root.Right = less
        return root, equal, greater
    }
    if root.Val > target {
        less, equal, greater := splitBST3(root.Left, target)
        root.Left = greater
        return less, equal, root
    }
    // root.Val == target
    less := root.Left
    greater := root.Right
    root.Left = nil
    root.Right = nil
    return less, root, greater
}
```

ポイント: `root.Val == target` のとき、**子リンクを nil にクリア**しないと、equal 側のノードが less / greater のノードを子として持ってしまう。

## Connect

### 実務での応用

- **Range クエリ**: 2 回 split で範囲 [L, R] のサブツリー抽出 → 範囲集計が O(log N)
- **Treap**: Split / Merge を基本操作とする確率的バランス BST、列の挿入・削除・反転に強い
- **永続化データ構造**: 過去の状態を変更せず新バージョンを作る (MVCC、関数型 FP)
- **Sharded DB**: データを範囲で分割して別マシンに配置する論理的操作
- **時系列データベース**: 時刻ベースの BST 分割で履歴管理（InfluxDB など）

### 関連問題

- 700 (Search in a BST): BST 探索の基本
- 701 (Insert into a BST): BST 挿入
- 450 (Delete Node in a BST): 削除
- 1382 (Balance a BST): 不平衡 BST を平衡にする
- 938 (Range Sum of BST): 範囲合計、Split のアイデアを使うと O(log N) で解ける

### BST 系問題の共通テクニック

- **再帰**: 各ノードで「この部分木に対する答え」を計算、子から戻ってきた値を組み合わせる
- **順序の活用**: in-order traversal でソート済み列、片側だけ降りる、範囲チェックなど
- **境界値の伝播**: validate BST のように親から子へ min/max 制約を渡す
