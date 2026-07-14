# Minimum Size Subarray Sum (209)

- **Date**: 2026-04-25
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: sliding window で「最小長サブ配列で target 以上」を求める

## Memorize

### O(N) sliding window

```go
func minSubArrayLen(target int, nums []int) int {
    l, sum := 0, 0
    minLen := len(nums) + 1
    for r := 0; r < len(nums); r++ {
        sum += nums[r]
        for sum >= target {
            if r-l+1 < minLen {
                minLen = r - l + 1
            }
            sum -= nums[l]
            l++
        }
    }
    if minLen == len(nums)+1 {
        return 0
    }
    return minLen
}
```

### O(N log N) prefix sum + binary search (follow-up)

```go
func minSubArrayLen(target int, nums []int) int {
    n := len(nums)
    prefix := make([]int, n+1)
    for i := 0; i < n; i++ {
        prefix[i+1] = prefix[i] + nums[i]
    }
    minLen := n + 1
    for i := 0; i < n; i++ {
        wantedSum := prefix[i] + target
        // [i+1, n+1) で prefix[j] >= wantedSum を満たす最小 j (lower_bound)
        l, r := i+1, n+1
        for l < r {
            mid := (l + r) / 2
            if prefix[mid] >= wantedSum {
                r = mid
            } else {
                l = mid + 1
            }
        }
        if l <= n && l-i < minLen {
            minLen = l - i
        }
    }
    if minLen == n+1 {
        return 0
    }
    return minLen
}
```

## Understand

### 計算量

- Sliding window: Time O(N), Space O(1)
- Prefix sum + binary search: Time O(N log N), Space O(N)

### sliding window の枠組み

LC 3 のテンプレート B (最小有効ウィンドウ型):
- **状態**: `sum` (ウィンドウの合計)
- **無効 (違反)**: `sum < target` → 拡張すべき
- **有効**: `sum >= target` → 縮められるだけ縮めて min 更新
- **答え**: 最小ウィンドウサイズ

```
1. r を右に拡張、sum += nums[r]
2. sum >= target の間: minLen 更新、sum -= nums[l]、l++
3. r 進めて繰り返し
```

### なぜ全要素 > 0 が必要か

sliding window が成立するには **sum の単調性** が必要:
- r++ → sum 増加 (nums[r] > 0)
- l++ → sum 減少 (nums[l] > 0)

負の数があると:
- r++ で sum が**減る**可能性
- l++ で sum が**増える**可能性
- → 単調性崩壊、sliding window 不成立

LC 862 (負の数あり版) は **deque + prefix sum** が必要。

### prefix sum + binary search の原理

サブ配列 `nums[i..j-1]` の合計 = `prefix[j] - prefix[i]`

求めたい: `prefix[j] - prefix[i] >= target` を満たす最小 `j - i`

各 i で `prefix[j] >= prefix[i] + target` を満たす最小 j を二分探索 (lower_bound)。
- prefix は単調増加 (nums > 0) → 二分探索可能
- prefix の長さを `n+1` にして `prefix[0] = 0` を入れる (i=0 から始まるサブ配列対応)

### つまずきポイント

- 「違反」と「有効」の方向を逆に考える: `sum < target` が無効、`sum >= target` が有効
- センチネル値 `len(nums) + 1` を使うと「未更新判定」が自然
- 二分探索版で prefix の長さを `n` にすると i=0 始まりが扱えない → `n+1` 必須
- 二分探索のターゲットを `target` ではなく `prefix[i] + target` にする
- 二分探索範囲を `[0, n+1)` ではなく `[i+1, n+1)` にする (空サブ配列を除外)

## Connect

### 変種問題

#### 「ちょうど target」の最小サブ配列
- sum > target のときだけ縮める、sum == target で更新
- 全要素正なら sliding window OK
- 一般 (負あり) は **prefix sum + hashmap** (LC 560 の応用)

#### 「target 以上」の**最長**サブ配列
- 全要素正なら**自明**: 全体の合計が target 以上なら答えは n、そうでなければ 0
- 「最長」と「最短」で問題の性質が大きく異なる

### sliding window の 2 大型 (LC 3 のテンプレート再掲)

| 型 | 縮めるトリガー | 答え更新タイミング | 例 |
|---|---|---|---|
| A. 最大有効ウィンドウ | 無効 (違反) のとき | 有効状態で max 更新 | LC 3, 340, 424 |
| B. 最小有効ウィンドウ | 有効のとき (縮めて最小狙う) | 各 valid 状態で min 更新 | LC 76, 209 |

### 関連問題

- 76 (Minimum Window Substring): 同じ Type B の代表
- 560 (Subarray Sum Equals K): ちょうど target、prefix sum + hashmap
- 862 (Shortest Subarray with Sum at Least K): 負の数あり版、deque 必要
- 904 (Fruit Into Baskets): Type A の例
- 1004 (Max Consecutive Ones III): Type A の例
