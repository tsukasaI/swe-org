# Permutations (46)

- **Date**: 2026-04-25
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: バックトラッキングの典型「choose → recurse → un-choose」テンプレート

## Memorize

### used[] 版（推奨、汎用的）

```go
func permute(nums []int) [][]int {
    var result [][]int
    var current []int
    used := make([]bool, len(nums))

    var backtrack func()
    backtrack = func() {
        if len(current) == len(nums) {
            result = append(result, append([]int{}, current...))
            return
        }
        for i := 0; i < len(nums); i++ {
            if used[i] {
                continue
            }
            used[i] = true
            current = append(current, nums[i])
            backtrack()
            used[i] = false
            current = current[:len(current)-1]
        }
    }
    backtrack()
    return result
}
```

### swap 版（省メモリ、エレガント）

```go
func permute(nums []int) [][]int {
    var result [][]int
    var backtrack func(start int)
    backtrack = func(start int) {
        if start == len(nums) {
            cp := make([]int, len(nums))
            copy(cp, nums)
            result = append(result, cp)
            return
        }
        for i := start; i < len(nums); i++ {
            nums[start], nums[i] = nums[i], nums[start]
            backtrack(start + 1)
            nums[start], nums[i] = nums[i], nums[start]
        }
    }
    backtrack(0)
    return result
}
```

## Understand

### 計算量

- Time: **O(n × n!)**（n! 個の順列、各コピー O(n)）
- Space (output 除く): **O(n)**（再帰スタック + current 配列）
- Output Space: **O(n × n!)**

### バックトラッキングの 3 要素

| 要素 | この問題での内容 |
|---|---|
| 選択肢 (choices) | nums のうち未使用の要素 |
| 状態 (state) | current（組み立て中）+ used[]（使用済み） |
| 終了条件 (base case) | len(current) == n |

### 「choose → recurse → un-choose」テンプレート

```
for 各選択肢 in 選択肢リスト:
    state を更新 (choose)
    recurse
    state を元に戻す (un-choose)
```

これがバックトラッキングの普遍的な型。順列だけでなく、組合せ・部分集合・N-Queens など同じ構造。

### 「コピーを result に追加」が必須の理由

`current` は再帰中に変化する。参照のまま `result` に追加すると、後続の choose/un-choose で書き換えられてしまう。

```go
// NG: 参照を共有
result = append(result, current)

// OK: コピーを追加
result = append(result, append([]int{}, current...))
// 等価:
cp := make([]int, len(current))
copy(cp, current)
result = append(result, cp)
```

### つまずきポイント

- `current` のコピーを忘れて参照を append → 全要素が同じ最終状態になる
- `un-choose` を忘れる → 状態が汚染される
- swap 版で「start から len(nums) まで」のループ範囲を間違える
- 重複あり (LC 47) に拡張するとき、ソートを忘れる

## Connect

### LC 47 - 重複あり順列

ソート + `!used[i-1]` の条件で重複生成を防ぐ:

```go
sort.Ints(nums)
// ...
if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
    continue   // 同じ深さで同値を 2 度試さない
}
```

なぜ `!used[i-1]` か:
- `!used[i-1]` (false) = 直前同値は使ってない = 同深さで以前試した（戻ってきた）= 重複生成 → skip
- `used[i-1]` (true) = 直前同値は浅い深さで使用中 = 別深さの選択肢 = OK

### LC 60 - K 番目の順列

factorial number system で直接構成:
- 1 番目の数字を決めると残りは (n-1)! 通り
- k を (n-1)! で割って index 決定、剰余を更新、残った候補から削除
- O(n^2)（slice 削除のため）、Fenwick tree で O(n log n)

### バックトラッキング系の関連問題

| 問題 | 違い |
|---|---|
| 46 (Permutations) | 順列、used[] |
| 47 (Permutations II) | 重複あり、`!used[i-1]` で skip |
| 39 (Combination Sum) | 組合せ、無限回使える |
| 40 (Combination Sum II) | 組合せ、各要素 1 回 |
| 78 (Subsets) | 部分集合 |
| 90 (Subsets II) | 重複あり部分集合 |
| 51 (N-Queens) | 制約付き順列 |
| 22 (Generate Parentheses) | 文字列の組み立て |

### swap 版 vs used[] 版

| | swap 版 | used[] 版 |
|---|---|---|
| メモリ | O(1) 追加（再帰スタック除く） | O(n) |
| 可読性 | やや難 | 明確 |
| 重複あり拡張 | 難（set 必要） | 易（`!used[i-1]`） |
| 推奨用途 | LC 46 単体 | 一般的、面接向け |

### バックトラッキングの汎用テンプレート

```
backtrack(state):
    if 完成条件:
        record(state)
        return
    for 選択肢 in 候補:
        if 制約違反: continue
        state を更新
        backtrack(state)
        state を戻す
```

これが「組合せ列挙」「制約充足」「探索」系問題の万能テンプレート。
