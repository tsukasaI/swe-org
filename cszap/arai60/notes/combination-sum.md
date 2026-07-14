# Combination Sum (39)

- **Date**: 2026-04-29
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: start index 型バックトラッキングで「同じ要素を無限回使える組合せ」を列挙

## Memorize

### 基本版（LC 39: distinct candidates、無限再利用）

```go
func combinationSum(candidates []int, target int) [][]int {
    var result [][]int
    var current []int
    currentSum := 0
    var backtrack func(start int)
    backtrack = func(start int) {
        if currentSum > target {
            return
        }
        if currentSum == target {
            result = append(result, append([]int{}, current...))
            return
        }
        for i := start; i < len(candidates); i++ {
            current = append(current, candidates[i])
            currentSum += candidates[i]
            backtrack(i)              // i+1 ではなく i → 同じ要素を再利用
            currentSum -= candidates[i]
            current = current[:len(current)-1]
        }
    }
    backtrack(0)
    return result
}
```

### 枝刈り強化版（ソート + break）

```go
func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    var result [][]int
    var current []int
    currentSum := 0
    var backtrack func(start int)
    backtrack = func(start int) {
        if currentSum == target {
            result = append(result, append([]int{}, current...))
            return
        }
        for i := start; i < len(candidates); i++ {
            if currentSum+candidates[i] > target {
                break   // ソート済みなので以降すべて NG
            }
            current = append(current, candidates[i])
            currentSum += candidates[i]
            backtrack(i)
            currentSum -= candidates[i]
            current = current[:len(current)-1]
        }
    }
    backtrack(0)
    return result
}
```

### LC 40 (Combination Sum II): 各要素 1 回 + 重複あり

```go
func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    var result [][]int
    var current []int
    currentSum := 0
    var backtrack func(start int)
    backtrack = func(start int) {
        if currentSum == target {
            result = append(result, append([]int{}, current...))
            return
        }
        for i := start; i < len(candidates); i++ {
            if currentSum+candidates[i] > target {
                break
            }
            if i > start && candidates[i] == candidates[i-1] {
                continue   // 同深さで同値はスキップ
            }
            current = append(current, candidates[i])
            currentSum += candidates[i]
            backtrack(i + 1)             // 各要素 1 回 → i+1
            currentSum -= candidates[i]
            current = current[:len(current)-1]
        }
    }
    backtrack(0)
    return result
}
```

## Understand

### 計算量

N = `len(candidates)`、T = target、M = `min(candidates)`

- Time: **O(N^(T/M))**（枝刈りなしの上界）
  - 再帰木の深さ T/M、各ノードで N 分岐
  - 厳密には各 leaf でコピー O(T/M) → O(N^(T/M) × T/M)
- Space (output 除く): **O(T/M)**（再帰スタック + current）

実際は枝刈りで大きく削れるが、最悪ケース（小さい M、大きい T）では指数が支配。

### LC 78 (Subsets) との違い

| | LC 78 (Subsets) | LC 39 (Combination Sum) |
|---|---|---|
| 目標 | 全部分集合 | sum == target の組合せ |
| 終了条件 | 全ノードで追加 | sum == target で追加、sum > target で打切 |
| 再帰呼び出し | `backtrack(i+1)` | `backtrack(i)` ← 再利用 |
| 制約 | なし | currentSum を追跡 |
| 計算量 | O(N × 2^N) | O(N^(T/M)) |

### `backtrack(i)` vs `backtrack(i+1)`

- **`i+1`**: 「使った要素は二度と使わない」→ 部分集合・順列・LC 40
- **`i`**: 「同じ要素を何度でも使ってよい」→ LC 39
- どちらも start を**進める方向**にしか動かさない → 順序を固定して重複組合せを防ぐ

### 重複組合せの回避（順序固定）

`candidates = [2,3,6,7], target = 7` で `[2,2,3]` を探すとき:
- index 順に非減少で選ぶ → `(0, 0, 1)` のみ
- もし `backtrack(0)` で全候補から選ぶと `[3,2,2]`, `[2,3,2]` も生成される

**start を進めることで「index の非減少順」を強制** → 同じ multiset の組合せは 1 度だけ生成。

### 枝刈り戦略

1. **`currentSum > target` で打切**（基本）
   - 全要素正の前提で、足し続けると単調増加 → これ以上深く行く意味なし

2. **ソート + `break`**（強化）
   - candidates をソートしておくと、`candidates[i] > target - currentSum` を見つけた時点で **以降すべて NG**
   - `continue` ではなく `break` できるので大幅に速い

3. **`if currentSum > target { return }` を消せる**
   - ソート + break 版では、そもそも超過する候補に入らないので不要
   - ただし `currentSum == target` の判定は必須

### つまずきポイント

- `backtrack(i+1)` と書いてしまう → 「同じ要素 1 回のみ」になり LC 39 の意図と違う
- ソートせずに `break` を使う → 早すぎる打切で正解を取りこぼす
- `current` のコピーを忘れて参照を append → 全要素が同じ最終状態
- 枝刈り `currentSum + candidates[i] > target` の前に append してしまう（無駄な操作）

## Connect

### LC 40 の拡張ポイント

LC 39 → LC 40 で変わるのは **2 箇所**:

1. **`backtrack(i)` → `backtrack(i+1)`**: 各要素 1 回のみ
2. **`i > start && candidates[i] == candidates[i-1] { continue }`**: 同深さで同値スキップ

### `i > start` vs `i > 0` の違い

| 条件 | 意味 | 結果 |
|---|---|---|
| `i > 0` | グローバルに重複スキップ | `[1,1,6]` も作れない |
| `i > start` | **同深さでのみ**重複スキップ | `[1,1,6]` OK、`[1,7]` の重複生成だけ防ぐ |

**直観**:
- `start` = この再帰呼び出しが見ている候補の左端
- `i == start` = この深さでの最初の選択 → 常に許可
- `i > start` で同値 = この深さで同値を 2 回目以降 → skip
- 異なる深さは独立した再帰呼び出しなので、`[1,1,...]` のように複数の同値を別の深さで使う組合せは生成可能

### バックトラッキング系の整理

| 問題 | start 進め方 | 重複処理 | 制約 |
|---|---|---|---|
| 78 (Subsets) | `i+1` | なし | なし |
| 90 (Subsets II) | `i+1` | `i > start` skip | なし |
| 39 (Combination Sum) | `i` (再利用) | なし (distinct) | sum == target |
| 40 (Combination Sum II) | `i+1` | `i > start` skip | sum == target |
| 46 (Permutations) | used[] 型 | なし | 全要素使用 |
| 47 (Permutations II) | used[] 型 | `!used[i-1]` skip | 全要素使用 |

### 設計の 4 軸

バックトラッキング問題は以下 4 軸で分類できる:

1. **要素の使用回数**: 1 回だけ / 無限回 / N 回など
2. **順序の意味**: あり (順列) / なし (組合せ・部分集合)
3. **入力の重複**: distinct / 重複あり
4. **終了条件**: 全要素使用 / sum 一致 / 制約満足など

LC 39 = (無限回, なし, distinct, sum 一致)
LC 40 = (1 回, なし, 重複あり, sum 一致)

### 関連問題

- 216 (Combination Sum III): 1〜9 から k 個選んで sum == target
- 377 (Combination Sum IV): **順列の数** を数える DP（バックトラッキングだと TLE）
- 322 (Coin Change): sum 系の最小コイン数 DP
- 518 (Coin Change II): sum 系の組合せ数 DP

LC 39 は **「列挙」**、LC 322/518 は **「数える / 最小値を求める」** → DP に転換。
列挙は本質的に指数なので、列挙以外の問いには DP の方が速い。
