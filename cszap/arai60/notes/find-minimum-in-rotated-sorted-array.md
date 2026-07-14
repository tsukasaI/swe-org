# Find Minimum in Rotated Sorted Array (153)

- **Date**: 2026-04-25
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: ローテートされたソート済み配列の最小値を二分探索で求める

## Memorize

- `l, r := 0, len(nums)-1` で開始
- 比較対象は **`nums[r]`**（`nums[l]` ではない）
- `nums[m] < nums[r]` → 最小値は mid 以下 → `r = m`
- `nums[m] > nums[r]` → 最小値は mid より右 → `l = m + 1`
- ループ条件は `l < r`、終了時に `l == r` がそのまま答えのインデックス
- 返すのは `nums[l]`（インデックスじゃなく値）

## Understand

### 計算量

- Time: O(log N)
- Space: O(1)

### なぜ `nums[r]` と比較するのか

`nums[l]` との比較は曖昧になる。例 `[4,5,6,7,0,1,2]` で `m=3, nums[m]=7`:
- `nums[l]=4`, `nums[m]=7` → `nums[m] > nums[l]` だが、最小値は右側にある
- 一方 `[7,0,1,2,4,5,6]` で `m=3, nums[m]=2`:
- `nums[l]=7`, `nums[m]=2` → `nums[m] < nums[l]` だが、最小値は左側にある

→ `nums[l]` だと判定できない。

`nums[r]` との比較は常に決定的:
- `nums[m] > nums[r]` → 必ず mid と right の間に「下がる境目」がある → 最小値は右側
- `nums[m] < nums[r]` → mid から right までソート済み → 最小値は mid 以下

### `l < r` vs `l <= r`

- これは「境界探索」のパターンなので `l < r`
- `l <= r` だと `l == r` で入ったあと `l = m+1 = len(nums)` になり、`return nums[l]` が配列外アクセス
- `l <= r` は「特定の値を探す」標準パターン用

### つまずきポイント

- `return l` と書いてしまう → インデックスを返してしまう。問題は値を要求しているので `return nums[l]`

## Connect

### 154 (Find Minimum in Rotated Sorted Array II): 重複あり

- `nums[m] == nums[r]` のときどちらに最小値があるか判定不能
- `r--` で安全に範囲を縮める
- 最悪 O(N)、平均 O(log N)

```go
if nums[m] < nums[r] {
    r = m
} else if nums[m] > nums[r] {
    l = m + 1
} else {
    r--
}
```

### 33 (Search in Rotated Sorted Array): target 探索

- 性質: mid で分割すると**必ず片方の半分はソート済み**（境目は1箇所しかないため）
- `nums[l] <= nums[m]` → 左半分がソート済み、そうでなければ右半分
- ソート済みの半分で target の範囲チェック → 範囲内なら左/右、範囲外なら反対側
- 「最小値を見つけてから二分探索」も「一発で二分探索」も両方 O(log N)（big-O 同じ）

### 関連問題

- 81 (Search in Rotated Sorted Array II): 33 + 重複ありバージョン
- 二分探索の境界探索パターン: 278 (First Bad Version), 35 (Search Insert Position)
