# Generate Parentheses (22)

- **Date**: 2026-04-29
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 制約付きバックトラッキング（ランニング制約で枝刈り）

## Memorize

### 基本版（[]string、可読性重視）

```go
func generateParenthesis(n int) []string {
    var result []string
    var current []string
    var backtrack func(open, close int)
    backtrack = func(open, close int) {
        if open == n && close == n {
            result = append(result, strings.Join(current, ""))
            return
        }
        if open < n {
            current = append(current, "(")
            backtrack(open+1, close)
            current = current[:len(current)-1]
        }
        if close < open {
            current = append(current, ")")
            backtrack(open, close+1)
            current = current[:len(current)-1]
        }
    }
    backtrack(0, 0)
    return result
}
```

### 最適化版（[]byte、定数倍高速）

```go
func generateParenthesis(n int) []string {
    var result []string
    current := make([]byte, 0, 2*n)
    var backtrack func(open, close int)
    backtrack = func(open, close int) {
        if open == n && close == n {
            result = append(result, string(current))
            return
        }
        if open < n {
            current = append(current, '(')
            backtrack(open+1, close)
            current = current[:len(current)-1]
        }
        if close < open {
            current = append(current, ')')
            backtrack(open, close+1)
            current = current[:len(current)-1]
        }
    }
    backtrack(0, 0)
    return result
}
```

## Understand

### 計算量

- Time: **O(C_n × n) = O(4^n / √n)**
  - C_n = n 番目のカタラン数 ≈ 4^n / (n^(3/2) × √π)
  - 各答えの構築/コピー O(n)
- Space (output 除く): **O(n)**
  - 再帰の深さ最大 2n
  - `current` の長さも最大 2n

### バックトラッキングの 3 要素

| 要素 | 内容 |
|---|---|
| 選択肢 | `(` または `)` |
| 状態 | `open` カウント、`close` カウント、`current` |
| 終了条件 | `open == n && close == n` (= `len(current) == 2n`) |

### ランニング制約

整形式 (well-formed) の条件を**途中**で満たすために:

| 制約 | 意味 |
|---|---|
| `open < n` で `(` を置く | 開き括弧が n を超えない |
| `close < open` で `)` を置く | 閉じ括弧が開きを超えない（`())(` のような不正形を防ぐ） |

両方の制約をクリアした選択肢のみ探索 → **無効な部分木は最初から作らない**。

### カタラン数 C_n

n 対の括弧で作れる整形式の数 = n 番目のカタラン数:

- C_0 = 1, C_1 = 1, C_2 = 2, C_3 = 5, C_4 = 14, C_5 = 42, ...
- 公式: C_n = (2n)! / ((n+1)! × n!) = `binom(2n, n) / (n+1)`
- 漸近: C_n ≈ 4^n / (n^(3/2) × √π)

カタラン数が現れる他の文脈:
- 二分木の構造数
- 凸多角形の三角形分割
- スタック操作の正当な順序

### 枝刈りの効果

| 戦略 | Time | コメント |
|---|---|---|
| 制約なし全列挙 + 後判定 | O(n × 4^n) | 2^(2n) = 4^n 全パターン × O(n) 検証 |
| ランニング制約あり | O(n × C_n) = O(4^n / √n) | √n ファクタ削減 |

n=20 で約 4.5 倍速い。指数の中では小さい差だが、**無駄な枝を 1 本も作らない**ことが本質。

### つまずきポイント

- `for open < n && close < n { ... }` のように **ループにしてしまう** → 無限ループ
  - 2 択の `if` を並べるだけで OK（または `for` で choice を回すなら open/close を実引数で渡す）
- `if close < n` と書いてしまう → `()(` のような状態で `)` を許してしまうので不整形
  - **正しくは `close < open`**
- `open`, `close` を closure 外の変数で管理しつつ更新を忘れる → 終了条件に到達せず無限再帰
  - 引数で渡すのが安全（状態の進退が呼び出しと同期する）
- `[]string` + `strings.Join` で書く → 動くが定数倍遅い
  - `[]byte` + `string(current)` がイディオマティック

### `(` を置く前提条件 vs 「制約」

設計上、`open < n` は「これ以上開けない」、`close < open` は「閉じすぎない」という**異なる種類**:

- `open < n`: **キャパ制約**（n という上限がある）
- `close < open`: **整形式制約**（順序を保つ）

似た構造の問題（N-Queens、数独、迷路など）でも、制約をこの 2 種類で分解すると整理しやすい。

## Connect

### バックトラッキングの分類（再掲・拡張）

| パターン | 例 | 状態管理 |
|---|---|---|
| start 進める型 | LC 78, 39, 40, 90 | start index |
| used[] 型 | LC 46, 47, 51 (N-Queens) | bool 配列 |
| include/exclude 型 | LC 78 alt | idx |
| **制約付き組み立て型** | **LC 22, 51, 37 (Sudoku)** | **カウンタ / 盤面** |

LC 22 は「列挙の制約 = ランニング制約で枝刈り」のテンプレ。

### 関連問題

- 20 (Valid Parentheses): スタックで判定 → 制約の理解に直結
- 32 (Longest Valid Parentheses): DP / スタック、別アプローチ
- 241 (Different Ways to Add Parentheses): 分割統治
- 301 (Remove Invalid Parentheses): BFS / バックトラッキング、より難しい
- 51 (N-Queens): 制約付きバックトラッキングの王道、列・対角の使用済み管理

### バックトラッキング汎用テンプレ（再掲）

```
backtrack(state):
    if 完成条件:
        record(state)
        return
    for choice in 選択肢:
        if 制約違反: continue        ← 枝刈り
        state を更新 (choose)
        backtrack(state)
        state を戻す (un-choose)
```

LC 22 では選択肢が `(` と `)` の 2 つだけ、それぞれに前置条件 (`open < n`, `close < open`) を付けることで「choice の前段で枝刈り」している。これは **`if 制約違反: continue` の 2 つの `if` 化**と等価。

### 「制約をいつチェックするか」の選択肢

1. **前段 (= ランニング制約)**: 選択する前にチェック → 無効な分岐を作らない（**LC 22 の方式**）
2. **後段 (= 完成後チェック)**: 完成した状態を検証して捨てる → シンプルだが遅い

可能なら前段が圧倒的に有利。`close < open` のように「途中で計算可能な制約」かどうかで判断する。
