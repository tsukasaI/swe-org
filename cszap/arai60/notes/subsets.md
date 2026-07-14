# Subsets (78)

- **Date**: 2026-04-27
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: バックトラッキングで全部分集合を列挙する

## Memorize

### start index 型（汎用的、推奨）

```go
func subsets(nums []int) [][]int {
    var result [][]int
    var current []int
    var backtrack func(start int)
    backtrack = func(start int) {
        result = append(result, append([]int{}, current...))
        for i := start; i < len(nums); i++ {
            current = append(current, nums[i])
            backtrack(i + 1)
            current = current[:len(current)-1]
        }
    }
    backtrack(0)
    return result
}
```

### include/exclude 型

```go
func subsets(nums []int) [][]int {
    var result [][]int
    var current []int
    var backtrack func(idx int)
    backtrack = func(idx int) {
        if idx == len(nums) {
            result = append(result, append([]int{}, current...))
            return
        }
        // 含める
        current = append(current, nums[idx])
        backtrack(idx + 1)
        current = current[:len(current)-1]
        // 含めない
        backtrack(idx + 1)
    }
    backtrack(0)
    return result
}
```

### ビット演算（反復、再帰なし）

```go
func subsets(nums []int) [][]int {
    n := len(nums)
    var result [][]int
    for mask := 0; mask < (1 << n); mask++ {
        var subset []int
        for i := 0; i < n; i++ {
            if mask&(1<<i) != 0 {
                subset = append(subset, nums[i])
            }
        }
        result = append(result, subset)
    }
    return result
}
```

## Understand

### 計算量

3 つの解法すべて同じ:
- Time: **O(N × 2^N)**（部分集合 2^N 個、各コピー O(N)）
- Space (output 除く): O(N)（再帰スタック + current）
- Output Space: O(N × 2^N)

### 部分集合の総数

各要素に対して「含める / 含めない」の 2 択 → **2^N** 通り。

### LC 46 (順列) との違い

| | LC 46 (順列) | LC 78 (部分集合) |
|---|---|---|
| 順序 | 意味あり | 意味なし |
| 全要素使用 | yes | no |
| 総数 | n! | 2^n |
| 追加タイミング | リーフだけ | 全ノード（start index 型） |
| 状態管理 | used[] | start index |

### start index 型 vs include/exclude 型

| | start index 型 | include/exclude 型 |
|---|---|---|
| 木の形 | 不規則 (各ノードで for) | 完全二分木 |
| 追加タイミング | 全ノード | リーフのみ |
| 状態の進め方 | start を進める | idx を進める |
| 出力順 | 辞書順に近い | 異なる順 |
| 拡張先 | 組合せ系 (LC 39, 40) | 真偽分岐系 |

両方書けると本物。**面接では start index 型が汎用的**。

### start index 型の決定木 (`nums=[1,2,3]`)

```
                    backtrack(0), current=[]
                    📝 add []
                   /         |         \
               i=0         i=1         i=2
              [1]          [2]          [3]
        backtrack(1)  backtrack(2)  backtrack(3)
        📝 add [1]    📝 add [2]    📝 add [3]
         /     \           |
       i=1     i=2         i=2
      [1,2]   [1,3]       [2,3]
       ↓       ↓           ↓
   📝 add [1,2] 📝 add [1,3] 📝 add [2,3]
       |
       i=2
      [1,2,3]
       ↓
   📝 add [1,2,3]
```

各ノード = 1 つの部分集合。start を進めることで「すでに使った要素は二度と使わない」を保証。

### つまずきポイント

- `current` のコピー忘れて参照を append → 全要素が同じ最終状態になる
- start index 型で `start` を `i` ではなく `i+1` にし忘れる → 同じ要素を 2 度使う
- include/exclude 型で base case の `idx == len(nums)` を `idx >= len(nums) - 1` などに間違える
- ビット版で `mask & (1 << i)` の **`!= 0` 比較**を忘れる（`mask & (1 << i)` 自体は int）

## Connect

### LC 90 - Subsets II（重複あり）

ソート + `i > start` の条件で重複を排除:

```go
sort.Ints(nums)
for i := start; i < len(nums); i++ {
    if i > start && nums[i] == nums[i-1] {
        continue   // 同じ深さで同値は最初だけ
    }
    // ...
}
```

**`i > 0` ではなく `i > start`** がポイント:
- 各深さの「最初の反復」(i == start) は常に許可
- 同じ深さの 2 回目以降で同値ならスキップ
- 異なる深さでは独立 → `[1,2,2]` のような部分集合は生成可能

### バックトラッキング系の関連問題

| 問題 | パターン | 重複処理 |
|---|---|---|
| 78 (Subsets) | 部分集合 | なし |
| 90 (Subsets II) | 部分集合、重複あり | `i > start` で skip |
| 46 (Permutations) | 順列 | なし |
| 47 (Permutations II) | 順列、重複あり | `!used[i-1]` で skip |
| 39 (Combination Sum) | 組合せ、無限再利用 | start index 型 |
| 40 (Combination Sum II) | 組合せ、各 1 回 | `i > start` で skip |
| 22 (Generate Parentheses) | 制約付き組み立て | 制約で枝刈り |

### バックトラッキングの 3 種類のパターン

1. **start 進める型**: 部分集合・組合せ向き、順序を固定して重複防止
2. **used[] 型**: 順列向き、全要素から選ぶ
3. **include/exclude 型**: 各要素の 2 値選択

問題の特性に応じて使い分け。
