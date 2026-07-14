# Longest Substring Without Repeating Characters (3)

- **Date**: 2026-04-25
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: sliding window で重複なしの最長部分文字列を求める

## Memorize

### set 版（直感的）

```go
func lengthOfLongestSubstring(s string) int {
    set := make(map[byte]struct{})
    l, maxLen := 0, 0
    for r := 0; r < len(s); r++ {
        for {
            if _, ok := set[s[r]]; !ok {
                break
            }
            delete(set, s[l])
            l++
        }
        set[s[r]] = struct{}{}
        if r-l+1 > maxLen {
            maxLen = r - l + 1
        }
    }
    return maxLen
}
```

### ジャンプ版（map[byte]int で last index）

```go
func lengthOfLongestSubstring(s string) int {
    hm := make(map[byte]int)
    l, maxLen := 0, 0
    for r := 0; r < len(s); r++ {
        if idx, ok := hm[s[r]]; ok {
            l = max(l, idx+1)   // l は単調非減少を保つ
        }
        hm[s[r]] = r
        if r-l+1 > maxLen {
            maxLen = r - l + 1
        }
    }
    return maxLen
}
```

### 配列最適化版（ASCII のみ）

```go
func lengthOfLongestSubstring(s string) int {
    lastIdx := [128]int{}
    for i := range lastIdx {
        lastIdx[i] = -1
    }
    l, maxLen := 0, 0
    for r := 0; r < len(s); r++ {
        if lastIdx[s[r]] >= l {
            l = lastIdx[s[r]] + 1
        }
        lastIdx[s[r]] = r
        if r-l+1 > maxLen {
            maxLen = r - l + 1
        }
    }
    return maxLen
}
```

## Understand

### 計算量

- Time: **O(N)**
  - r は単調に N 通り、l も単調に N 通り → 全体 2N
  - amortized 解析: 各文字は set/map に入る ≤ 1 回、出る ≤ 1 回
- Space: **O(min(N, σ))** (σ = 文字種類数)
  - ASCII なら σ ≤ 128 → 実用的に O(1)

### sliding window の本質

- l, r 2 ポインタで「重複のないウィンドウ」を維持
- r を右に進めて拡張、必要に応じて l を右に進めて縮小
- l も r も戻らない単調性が amortized O(N) の鍵

### `max(l, idx+1)` の重要性

ジャンプ版で `l = idx + 1` だけだと **l が逆戻り**するケースがある。

例 `s = "tmmzuxt"`:
- r=2 で 'm' 重複検出 → l=2 にジャンプ (m を除外)
- r=6 で 't' 重複検出 → t は index=0 にあった
- もし `l = 0 + 1 = 1` だと l=2 から l=1 に逆戻り → "mmzuxt" を含み 'm' 重複
- 正しくは `l = max(2, 1) = 2` で維持

→ **l は単調非減少** を保つために `max` が必要。

### 最適化の段階

| 版 | データ構造 | Time | Space | 定数倍 |
|---|---|---|---|---|
| set + 1 個ずつ縮小 | map[byte]struct{} | O(N) | O(σ) | 中 |
| ジャンプ | map[byte]int | O(N) | O(σ) | 中 |
| 配列 ジャンプ | [128]int | O(N) | O(1) | **小** |

big-O は同じ、配列版は定数倍で 2〜5 倍速い (ハッシュコスト回避)。

### つまずきポイント

- `for r := range s` (単変数) は index を返す、文字を返さない (Go の挙動)
- ジャンプ版で `max(l, idx+1)` を忘れる → 過去の重複に引きずられる
- `len(set) == r - l + 1` の不変条件は維持されるが、`r - l + 1` で書く方が意図が明確
- 空文字列の特別扱いは不要 (for ループが回らず maxLen=0 のまま)

## Connect

### sliding window の汎用テンプレート

```
1. l = 0, r 動かす
2. r を右に拡張: ウィンドウに s[r] を追加、状態を更新
3. ウィンドウが「違反」の間: l を右に進めて状態更新
4. 「有効」なら答えを更新 (最大なら拡張後、最小なら縮小しながら)
5. r を進める
```

### 4 つの観点で関連問題を整理

| 問題 | ウィンドウの状態 | 違反条件 | 答え |
|---|---|---|---|
| LC 3 | 文字 set/map | 重複あり | **最大** サイズ |
| LC 76 | 各文字数 (need vs have) | t を含まない | **最小** サイズ |
| LC 159/340 | 種類数 | 種類 > K | **最大** サイズ |
| LC 424 | 各文字数, 最頻出数 | (size - maxCount) > K | **最大** サイズ |
| LC 567 | 各文字数 | 異なる count | **存在判定** |

### 2 つのバリエーション

#### A. 「最大の有効ウィンドウ」型 (LC 3, 159, 340, 424)
```
violate しなくなるまで縮める
答え = max(答え, window size)
```

#### B. 「最小の有効ウィンドウ」型 (LC 76)
```
valid な間に縮め続ける
各 valid 状態で 答え = min(答え, window size) を更新
```

### キーポイント

- 状態を**高速に更新できる**ことが必須 (count, distinct, maxCount など)
- l も r も**右にしか動かない**単調性 → amortized O(N)
- 各問題で「ウィンドウの状態 / 違反条件 / 答えの形」の 3 点を埋める

### 関連問題

- 76 (Minimum Window Substring): 最小ウィンドウ系の代表
- 159 (At Most Two Distinct): K=2 の特殊版
- 340 (At Most K Distinct): 一般化
- 424 (Character Replacement): max count 追跡
- 567 (Permutation in String): 固定長ウィンドウ
- 438 (Find All Anagrams): 567 の延長
- 30 (Substring with Concatenation of All Words): 単語単位の sliding window
