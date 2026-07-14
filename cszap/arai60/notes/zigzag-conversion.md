# Zigzag Conversion (6)

- **Date**: 2026-05-02
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: zigzag パターンの周期性を利用した文字列再構成

## Memorize

### シミュレーション版 (推奨、面接向け)

```go
func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    buf := make([][]byte, numRows)
    moveDown := false
    p := 0
    for i := 0; i < len(s); i++ {
        buf[p] = append(buf[p], s[i])
        if p == 0 || p == numRows-1 {
            moveDown = !moveDown
        }
        if moveDown {
            p++
        } else {
            p--
        }
    }
    result := make([]byte, 0, len(s))
    for _, b := range buf {
        result = append(result, b...)
    }
    return string(result)
}
```

### インデックス計算版 (mod 演算)

```go
func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    cycle := (numRows - 1) * 2
    buf := make([][]byte, numRows)
    for i := 0; i < len(s); i++ {
        m := i % cycle
        var row int
        if m < numRows {
            row = m              // 下りの位置
        } else {
            row = cycle - m      // 上りで戻る位置
        }
        buf[row] = append(buf[row], s[i])
    }
    result := make([]byte, 0, len(s))
    for _, b := range buf {
        result = append(result, b...)
    }
    return string(result)
}
```

## Understand

### 計算量

- Time: **O(N)** — 各文字を 1 回処理 + 結合
- Space: **O(N)** — buf 配列に N 文字格納

### 周期 = 2n - 2

zigzag の 1 周期 = 「下に n 文字 + 上に (n-2) 文字」（両端は折り返しで 1 回ずつ）
- n=3: 周期 = 4
- n=4: 周期 = 6
- n=5: 周期 = 8

```
n=4 の 1 周期:
  P(0)            ← row 0
   A(1)           ← row 1
    Y(2)          ← row 2
     P(3)         ← row 3 で折り返し
    A(4)          ← row 2 (上り)
   L(5)           ← row 1 (上り)
   ↓ 次の周期 (i=6 で row 0 へ)
```

n=1 のとき 2n-2 = 0 → cycle=0 は除算エラー → **早期 return** で対応。

### 各行の周期内位置

`numRows = n`, 周期 = `cycle = 2n - 2` のとき、`i % cycle` の値で行が決まる:

| 行 | 該当する `i % cycle` |
|---|---|
| 0 | 0 |
| n-1 (最下) | n-1 |
| 中間行 r (0 < r < n-1) | r または cycle - r |

中間行が 2 つの位置を持つのは「下りで通る」+「上りで通る」両方含むから。
最上 (0) と最下 (n-1) は 1 周期に 1 回しか通らない。

### シミュレーション版の動き

ポインタ `p` と方向フラグ `moveDown` を持って 1 文字ずつ振り分け:
- 端 (`p == 0` or `p == n-1`) で方向反転
- 各文字を `buf[p]` に追加して p を進める

**初期化のコツ**: `moveDown = false` で開始。p=0 で最初の flip により `moveDown=true` (下り) に → 自然に下から始まる。

### つまずきポイント

- `numRows == 1` で cycle = 0 になる → 0 除算 / 無限ループ → 早期 return 必須
- インデックス計算版で「中間行の対称性 (`row = cycle - m`)」を見落とす
- シミュレーション版で方向反転を「flip 後に進める」ではなく「進めてから flip」 → 範囲外
- byte と rune の混同（LC 6 は ASCII のみだが、Unicode 入力なら注意）

### エッジケース

| 入力 | 出力 |
|---|---|
| `s="A", numRows=1` | "A" (早期 return) |
| `s="AB", numRows=1` | "AB" |
| `s="AB", numRows=2` | "AB" (cycle=2: A→row0, B→row1) |
| `s="ABC", numRows=5` | "ABC" (numRows > len(s)、下りのみで終わる) |

## Connect

### シミュレーション版 vs インデックス計算版

| | シミュレーション | インデックス計算 |
|---|---|---|
| 実装難度 | 易 | 中 |
| バグの入りやすさ | 低 | 高 |
| 面接適性 | ◎ | △ |
| 数学的美しさ | ○ | ◎ |
| 速度 | 同じ O(N) | 同じ O(N) |

**面接では迷わずシミュレーション版**。インデックス計算版は「気付いたらこっちもある」程度の知識として持つ。

### 周期パターン問題

zigzag のような「**周期で繰り返すパターン**」を扱う問題:

| 問題 | 周期の定義 |
|---|---|
| 6 (Zigzag) | `2n - 2` 文字 |
| 1041 (Robot Bounded In Circle) | 4 回繰り返して元の位置に戻るか |
| 1759 (Count Number of Homogenous Substrings) | 連続同一文字のラン |

「i % period」で位置を決定するパターンは zigzag 系の典型。

### 文字列構築問題の系譜

| 問題 | 操作 |
|---|---|
| 6 (Zigzag) | row 単位で再構成 |
| 38 (Count and Say) | 前項を読み上げて次項生成 |
| 67 (Add Binary) | 桁上がり処理しながら結合 |
| 415 (Add Strings) | 同上、10 進数版 |
| 273 (Integer to English Words) | 数値 → 英語 |
| 12 (Integer to Roman) | 数値 → ローマ数字 |

zigzag は「**配置パターンを理解して再構成**」型の単純例。アルゴリズムというより観察力が問われる。

### キーポイント

- zigzag 系は「**周期 2n-2**」を見抜けるかが第一関門
- 中間行の **2 箇所性** (下りと上りで通る) を理解する
- 実装は**シミュレーション**が安全、インデックス計算は対称性の理解が必要
- `numRows == 1` の特殊ケース処理を忘れない
