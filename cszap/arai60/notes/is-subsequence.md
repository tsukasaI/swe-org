# Is Subsequence (392)

- **Date**: 2026-05-01
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 2 ポインタによる subsequence 判定、大量クエリ対応の前計算

## Memorize

### 基本版（2 ポインタ）

```go
func isSubsequence(s string, t string) bool {
    if len(s) > len(t) {
        return false
    }
    i := 0
    for j := 0; j < len(t) && i < len(s); j++ {
        if t[j] == s[i] {
            i++
        }
    }
    return i == len(s)
}
```

### Follow-up: 同じ t に対する大量クエリ (二分探索版)

```go
type matcher struct {
    indexOf map[byte][]int
}

func newMatcher(t string) *matcher {
    m := &matcher{indexOf: make(map[byte][]int)}
    for j := 0; j < len(t); j++ {
        m.indexOf[t[j]] = append(m.indexOf[t[j]], j)
    }
    return m
}

func (m *matcher) isSubsequence(s string) bool {
    prev := -1
    for i := 0; i < len(s); i++ {
        list, ok := m.indexOf[s[i]]
        if !ok {
            return false
        }
        idx := sort.SearchInts(list, prev+1)   // lower_bound for prev+1
        if idx == len(list) {
            return false
        }
        prev = list[idx]
    }
    return true
}
```

### Follow-up: DP テーブル版 (より高速、メモリ多い)

```go
// next[j][c] = t の position j 以降で最初に文字 c が現れる位置 (-1 if not found)
func buildNext(t string) [][26]int {
    n := len(t)
    next := make([][26]int, n+1)
    for c := 0; c < 26; c++ {
        next[n][c] = -1
    }
    for j := n - 1; j >= 0; j-- {
        for c := 0; c < 26; c++ {
            if t[j] == byte('a'+c) {
                next[j][c] = j
            } else {
                next[j][c] = next[j+1][c]
            }
        }
    }
    return next
}

func isSubsequenceWithNext(s string, next [][26]int) bool {
    j := 0
    for i := 0; i < len(s); i++ {
        c := int(s[i] - 'a')
        if j >= len(next) || next[j][c] == -1 {
            return false
        }
        j = next[j][c] + 1
    }
    return true
}
```

## Understand

### 計算量

| 解法 | 前計算 | クエリごと | k クエリ合計 |
|---|---|---|---|
| 2 ポインタ | なし | O(\|s\| + \|t\|) | O(k(\|s\| + \|t\|)) |
| 二分探索 | O(\|t\|) | O(\|s\| log \|t\|) | O(\|t\| + k\|s\| log \|t\|) |
| DP テーブル | O(\|t\| × 26) | O(\|s\|) | O(\|t\| × 26 + k\|s\|) |

Space:
- 2 ポインタ: O(1)
- 二分探索: O(\|t\|) (index リスト)
- DP テーブル: O(\|t\| × 26)

|t| が大きく、|s| が小さく、k が巨大なときほど前計算版が有利。

### 2 ポインタの本質

- `i` (s 用): マッチした文字数 = subsequence の進捗
- `j` (t 用): 探索中の位置、常に右に進む
- マッチ → 両方進める / 不一致 → j のみ進める
- 後戻り不要 = 線形時間で済む

「s の各文字を、t の中で順番に左から探す」greedy で正当性が保証される（より早くマッチを取るほど後の自由度が高い）。

### Follow-up の発想

「同じ t に大量の s を判定」シナリオは LC 392 の有名な follow-up:
> If there are lots of incoming S (~10^9), how would you change your code?

毎回 t 全体を見るのは無駄。**t を「どこに何があるか」のインデックス**に変換しておけば、各 s で「次に欲しい文字を二分探索 or O(1) lookup」で見つけられる。

二分探索版: 各文字の出現位置リストを保持 → `lower_bound(prev+1)` で次の位置。
DP テーブル版: `next[j][c]` を作っておけば、各クエリで「現在位置 j から文字 c の次の位置」を O(1) で取れる。

### つまずきポイント

- `i` も増やしてしまう（マッチしなくても）→ 飛ばしロジックが壊れる
- 終了条件で `i == len(s)` を `i >= len(s)` と書く → 同じだが意図が曖昧
- 二分探索で `prev` ではなく `prev+1` を渡すのを忘れる → 同じ位置を再選択
- DP テーブルの `next[n][c] = -1` の初期化を忘れる → 範囲外参照
- 大文字小文字が混ざるケース → 問題制約を確認 (LC 392 は lowercase のみ)

## Connect

### Subsequence vs Substring

| | Subsequence | Substring |
|---|---|---|
| 制約 | 順序保持、飛ばし OK | 順序保持、**連続必須** |
| 例 ("ace" vs "abcde") | OK | NG |
| 標準アルゴリズム | greedy 2 ポインタ | KMP / Z / Rabin-Karp |
| Time (1 クエリ) | O(\|s\| + \|t\|) | O(\|s\| + \|t\|) (KMP) |
| 大量クエリ前計算 | indexOf or next[j][c] | suffix automaton / suffix array |

### Substring matching の主要アルゴリズム

| アルゴリズム | Time | Space | 特徴 |
|---|---|---|---|
| Naive | O(\|s\| × \|t\|) | O(1) | 全 start position で比較 |
| **KMP** | O(\|s\| + \|t\|) | O(\|s\|) | 失敗関数 (LPS 配列) |
| Z-algorithm | O(\|s\| + \|t\|) | O(\|s\| + \|t\|) | Z 配列 |
| Rabin-Karp | O(\|s\| + \|t\|) avg | O(1) | rolling hash |
| Boyer-Moore | sublinear avg | O(σ) | 後ろから比較 |

LeetCode の代表問題: **LC 28 (Find the Index of the First Occurrence)**.

### 関連問題

- 28 (Find the Index of the First Occurrence): substring matching、KMP の練習
- 524 (Longest Word in Dictionary through Deleting): 各 word が s の subsequence か判定 + 最長探索
- 727 (Minimum Window Subsequence): t の中で s が subsequence になる最小ウィンドウ
- 792 (Number of Matching Subsequences): 多数 word の subsequence カウント、follow-up の本格応用
- 1143 (Longest Common Subsequence): DP の典型、subsequence の長さ系
- 1216 (Valid Palindrome III): k 文字以内削除で palindrome subsequence

### subsequence 系の階層

1. **判定** (LC 392): true/false
2. **数える** (LC 115 Distinct Subsequences): 同じ subsequence のパターン数 → DP
3. **最長 / 最小** (LC 1143 LCS, LC 727): DP
4. **Lexicographically smallest / largest**: stack + 制約

LC 392 はもっとも単純な「判定」レベル。これが基礎で、より複雑なものは DP になる。

### キーポイント

- subsequence の判定は **greedy で OK**（後戻り不要）→ 2 ポインタ
- substring の判定は **失敗時の戻り処理**が必要 → KMP / Z
- 「同じ t、多数クエリ」シナリオでは t を **インデックス化**しておく
- 「次の出現位置を高速取得」が前計算の本質
