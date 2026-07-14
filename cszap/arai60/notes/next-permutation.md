# Next Permutation (31)

- **Date**: 2026-05-02
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 辞書順で次の permutation を in-place で構築する 4 ステップアルゴリズム

## Memorize

### O(N) in-place

```go
func nextPermutation(nums []int) {
    // Step 1: 右から見て最初に nums[i] < nums[i+1] となる pivot を探す
    i := len(nums) - 2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }

    // Step 2 & 3: pivot が見つかった場合、右側で「pivot より大きい最小値」を探して swap
    if i >= 0 {
        j := len(nums) - 1
        for j > i && nums[j] <= nums[i] {
            j--
        }
        nums[i], nums[j] = nums[j], nums[i]
    }

    // Step 4: pivot より右を reverse (元々非増加 → 反転で非減少 = 辞書順最小)
    for l, r := i+1, len(nums)-1; l < r; l, r = l+1, r-1 {
        nums[l], nums[r] = nums[r], nums[l]
    }
}
```

## Understand

### 計算量

- Time: **O(N)** — 各ステップが線形（pivot 探し + j 探し + reverse）
- Space: **O(1)** — in-place、追加メモリなし

### 4 ステップアルゴリズム

```
ステップ 1: pivot 位置 i を探す
  右から見て最初に「降順が崩れる位置」 = nums[i] < nums[i+1] となる最大の i
  見つからなければ全体が非増加 → i = -1

ステップ 2: 右側で「pivot より大きい最小値」を探す
  サフィックス nums[i+1..] は非増加なので、右から線形に
  nums[j] > nums[i] を満たす最大の j

ステップ 3: nums[i] と nums[j] を swap

ステップ 4: nums[i+1..end] を reverse
  サフィックスは swap 後も非増加 → reverse で非減少 = 辞書順最小
```

i = -1 の場合（全体が非増加 = 最後の permutation）:
- ステップ 2, 3 はスキップ
- ステップ 4 で全体を reverse → 最初の permutation (昇順)

### なぜこの手順で「次」が得られるか

辞書順で「次に大きい」とは:
1. **どこかをより大きく**しないといけない
2. その変更は **できるだけ右側**で起こすべき (左を変えると差が大きくなる)
3. 変更は **できるだけ小さい増分**にしたい
4. 変更後の右側は **できるだけ小さい順**にしたい (= 昇順)

これに対応:
- ステップ 1: 「右側で大きくできる最右の位置」を探す = 右から非増加が崩れる場所
- ステップ 2-3: 「最小の増分」 = 右側で pivot を超える最小値と swap
- ステップ 4: 「右側を昇順に」 = サフィックスを reverse

### サフィックスが swap 後も非増加である証明

j は「`nums[j] > nums[i]`」を満たす **最大** の index。よって:
- 旧 `nums[j] > nums[i]`
- 旧 `nums[j+1] <= nums[i]` (j+1 が存在するなら、j+1 は条件を満たさない)
- 旧 `nums[j-1] >= nums[j]` (元々サフィックスは非増加)

swap 後、新 `nums[j] = 旧 nums[i]`:

| 比較対象 | 評価 | 判定 |
|---|---|---|
| 新 `nums[j-1] vs nums[j]` | 旧 `nums[j-1] >= nums[j] > 旧 nums[i]` = 新 `nums[j]` | strict descending ✓ |
| 新 `nums[j] vs nums[j+1]` | 旧 `nums[i] vs 旧 nums[j+1] <= 旧 nums[i]` | non-increasing ✓ |

→ サフィックスは非増加を保つ → reverse で非減少。

### つまずきポイント

- pivot 探しの条件を `nums[i] > nums[i+1]` と書いてしまう → 降順を維持する条件で逆方向
- j 探しの条件を `nums[j] < nums[i]` と書いてしまう → 反対の値を選んでしまう
- j 探しを `>=` で書く → 重複時に同値を選んで swap が無効になる
- ステップ 4 で sort を呼ぶ (O(N log N)) → 不要、reverse で O(N)
- `sort.Slice(nums[i+1:], func(a, b int) ...)` で外側の i を shadow → 元配列にアクセスする closure が間違う
- 全降順 (i=-1) のエッジケース処理を忘れる → サフィックス reverse で対応できる

### Strict `>` が重要な理由 (LC 例 `[1,1,5]`)

- pivot: i=1 (nums[1]=1 < nums[2]=5)
- j 探し:
  - `>` (正しい): j=2 (nums[2]=5 > nums[1]=1) → swap → `[1,5,1]` ✓
  - `>=` (誤り): j=1 (nums[1]=1 >= nums[1]=1 で止まる) → 自分と swap → 変化なし

「異なる値」に swap することを strict `>` が保証する。

### エッジケース

| 入力 | i 値 | j 値 | 結果 |
|---|---|---|---|
| `[1,2,3]` | 1 | 2 | `[1,3,2]` |
| `[3,2,1]` | -1 | (skip) | `[1,2,3]` (全 reverse) |
| `[1,1,5]` | 1 | 2 | `[1,5,1]` |
| `[1]` | -1 | (skip) | `[1]` (no-op) |
| `[1,3,2]` | 0 | 2 | swap → `[2,3,1]` → reverse → `[2,1,3]` |

## Connect

### permutation 系の問題群

| 問題 | パターン | 解法 |
|---|---|---|
| 31 (Next Permutation) | 1 個次へ | 4 ステップ、in-place O(N) |
| 46 (Permutations) | 全列挙 (distinct) | バックトラッキング (used[]) |
| 47 (Permutations II) | 全列挙 (重複あり) | バックトラッキング + skip |
| 60 (Permutation Sequence) | k 番目の permutation | factorial number system |
| 556 (Next Greater Element III) | 整数版の next permutation | 31 を string で適用 |
| 1053 (Previous Permutation) | 1 個前へ | 31 の逆方向 |
| 484 (Find Permutation) | DI sequence から復元 | greedy + reverse |

### LC 60 (k-th permutation) との関係

- LC 31: 1 個次へ進む O(N)
- LC 60: k 個進むのに 31 を k 回適用すると O(kN) (k が大きいと遅い)
- LC 60 の最適: factorial number system で直接構築 O(N^2) (slice 削除コスト)

LC 31 → 31 を k 回 → too slow when k 大。LC 60 は k を直接 base-factorial 表現に変換して各桁を決める。

### Previous Permutation (LC 1053)

LC 31 のミラー:
- ステップ 1: `nums[i] > nums[i+1]` で pivot (右から見て非減少が崩れる位置)
- ステップ 2: 右で「pivot より小さい最大値」と swap
- ステップ 4: サフィックスを reverse (非減少 → 非増加 = 辞書順最大)

すべて不等号を反転するだけ。

### 関連概念

- **Lexicographic order**: 辞書順比較は文字列・配列・数列で同じルール
- **In-place algorithm**: 追加メモリ O(1) の制約 → swap と reverse の活用
- **Factorial number system**: k 番目の permutation を O(N^2) で構築可能

### キーポイント

- **辞書順「次」の構築は 4 ステップで決定的に作れる** (DP や全探索不要)
- 「**右側でできるだけ小さく変える**」原則
- swap 後のサフィックスが非増加を保つ性質 → reverse で線形に昇順化
- strict `>` で重複時の自己 swap を回避
