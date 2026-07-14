# Search Insert Position (35)

- **Date**: 2026-04-25
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: lower_bound 型二分探索で挿入位置を求める

## Memorize

- ループ条件 `l < r`、初期値 `l, r = 0, len(nums)`（`len(nums)-1` ではない）
- `nums[c] >= target` なら `r = c`（c は候補なので外さない）
- `nums[c] < target` なら `l = c + 1`（c は答えではないと確定）
- ループ終了時 `l == r`、`return l`

## Understand

### 計算量

- Time: O(log N)
- Space: O(1)

### 不変条件 (loop invariant)

ループ中、常に以下が成り立つ:
```
[0, l) の範囲: nums[i] < target (答えにはなりえない)
[r, len(nums)) の範囲: nums[i] >= target (答えの候補、r が最も左)
[l, r) の範囲: 未確定
```

- `l`: ここまで全て target 未満と確定した境界の右端
- `r`: ここから先は全て target 以上と確定した境界の左端
- 答えは `[l, r]` の中にある
- 終了時 `l == r` がそのまま答え

### なぜ `r = len(nums)` か

- 答えが `len(nums)` になりうる（target が全要素より大きいケース）
- 例: `nums = [1,2,3,4], target = 100` → 4 を返す必要がある
- `r = len(nums) - 1` で始めると 4 に到達できない

### なぜ `r = c`（`r = c - 1` ではない）か

- `nums[c] >= target` のとき、c は答えの候補
- `r = c - 1` だと c を範囲から除外してしまう
- 不変条件 `[r, len(nums))` は「答えの候補」なので c を r に含める

### なぜ `l < r`（`l <= r` ではない）か

- 不変条件が `[l, r)` の半開区間
- ループ終了は「未確定範囲が空 = `l == r`」のタイミング
- `l <= r` だと終了後 `l > r` になり不変条件が崩れる

### なぜ `>=`（`>` ではない）か

- `>=` = lower_bound: 「target 以上が現れる**最初の位置**」
- `>` = upper_bound: 「target より大きい値が現れる**最初の位置**」
- target が配列に存在するとき:
  - `>=`: target の位置を返す（問題の要求通り）
  - `>`: target の1つ後ろを返す（問題の要求と違う）
- 例 `[1,2,3,4], target=3`: `>=` → 2、`>` → 3

### LC 33 との違い（二分探索パターン2種類）

| | 境界探索 (LC 35) | 値探索 (LC 33) |
|---|---|---|
| 不変条件 | 範囲ベース `[l, r)` | 単一要素ベース |
| ループ条件 | `l < r` | `l <= r` |
| 右更新 | `r = c` | `r = m - 1` |
| 初期 r | `len(nums)` | `len(nums) - 1` |
| 戻り値 | `l`（位置） | mid 一致時 `m`、なければ `-1` |

### つまずきポイント

- `r = len(nums) - 1` で始めると末尾挿入ケースが取れない
- `r = c - 1` で書くと候補を取り逃がす
- `l <= r` だと不変条件と矛盾、無限ループや off-by-one
- `>` を使うと target 一致時に位置がずれる

## Connect

### lower_bound / upper_bound の双子

- `>=` を `>` に変えるだけで `upper_bound` になる
- C++ STL の `std::lower_bound` / `std::upper_bound` がこの2つに対応
- 両方持っておくと汎用的: 「`[lower_bound(t), upper_bound(t))` が target の出現区間」

### 関連問題

- 33 (Search in Rotated Sorted Array): 値探索パターン、回転対応
- 153 (Find Minimum in Rotated Sorted Array): 境界探索パターン、回転点検出
- 278 (First Bad Version): 境界探索の典型
- 34 (Find First and Last Position): lower_bound + upper_bound の組み合わせ
- 69 (Sqrt(x)): 境界探索で `x*x <= target` の最大 x を見つける
