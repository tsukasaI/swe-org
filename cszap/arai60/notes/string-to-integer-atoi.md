# String to Integer atoi (8)

- **Date**: 2026-05-02
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 文字列パース + オーバーフロー処理、エッジケース全網羅

## Memorize

### 順次処理版 (推奨、シンプル)

```go
import "math"

func myAtoi(s string) int {
    i := 0

    // 1. 先頭 whitespace スキップ
    for i < len(s) && s[i] == ' ' {
        i++
    }

    // 2. 符号
    sign := 1
    if i < len(s) && (s[i] == '+' || s[i] == '-') {
        if s[i] == '-' {
            sign = -1
        }
        i++
    }

    // 3. 数字累積 + オーバーフロー判定
    result := 0
    for i < len(s) && s[i] >= '0' && s[i] <= '9' {
        result = 10*result + int(s[i]-'0')
        if result > math.MaxInt32 {
            if sign == 1 {
                return math.MaxInt32
            }
            return math.MinInt32
        }
        i++
    }

    return sign * result
}
```

### FSM 版 (状態遷移、可読性 + 拡張性)

```go
type atoiState int

const (
    atoiStart atoiState = iota
    atoiSigned
    atoiInNumber
    atoiEnd
)

type atoiInputClass int

const (
    classSpace atoiInputClass = iota
    classSign
    classDigit
    classOther
)

func classify(c byte) atoiInputClass {
    switch {
    case c == ' ':
        return classSpace
    case c == '+' || c == '-':
        return classSign
    case c >= '0' && c <= '9':
        return classDigit
    default:
        return classOther
    }
}

var atoiTable = [4][4]atoiState{
    // SPACE       SIGN          DIGIT         OTHER
    {atoiStart, atoiSigned, atoiInNumber, atoiEnd},   // START
    {atoiEnd, atoiEnd, atoiInNumber, atoiEnd},        // SIGNED
    {atoiEnd, atoiEnd, atoiInNumber, atoiEnd},        // IN_NUMBER
    {atoiEnd, atoiEnd, atoiEnd, atoiEnd},             // END
}

func myAtoiFSM(s string) int {
    state := atoiStart
    sign := 1
    result := 0
    for i := 0; i < len(s) && state != atoiEnd; i++ {
        c := s[i]
        cls := classify(c)
        state = atoiTable[state][cls]
        switch state {
        case atoiSigned:
            if c == '-' {
                sign = -1
            }
        case atoiInNumber:
            result = 10*result + int(c-'0')
            if result > math.MaxInt32 {
                if sign == 1 {
                    return math.MaxInt32
                }
                return math.MinInt32
            }
        }
    }
    return sign * result
}
```

## Understand

### 計算量

- Time: **O(N)** — 1 パス
- Space: **O(1)** — 変数数個のみ

### 4 段階アルゴリズム

```
1. 先頭 whitespace スキップ
2. 符号 (+/-) 読み取り (オプショナル)
3. 数字累積 (非数字に当たったら停止)
4. [INT_MIN, INT_MAX] にクランプ
```

各段階で **read pointer i を 1 つ**進めていく形が定番。

### オーバーフロー判定の非対称性

32-bit signed:
- `INT_MAX = 2147483647`
- `INT_MIN = -2147483648`

→ **負の方が絶対値が 1 大きい**。

判定式 `result > math.MaxInt32` (result は正の累積値):
- abs 値 = 2147483648 → triggered → 正なら MaxInt32、負なら MinInt32
- abs 値 = 2147483647 → not triggered → 正なら MaxInt32、負なら -2147483647 (範囲内)
- abs 値 > 2147483648 → triggered → 同上にクランプ

**1 つの比較式で正負両方の境界を吸収できる**のは、`|MinInt32| = MaxInt32 + 1` だから。これに頼らない実装も可能だが、より複雑になる。

### 値が int64 で安全な理由

Go の `int` は 64-bit (現代のプラットフォーム)。result が `MaxInt32 + 1` で必ず判定 → 早期 return → int64 オーバーフローには到達しない。

入力長 200 文字でも、10 桁目で必ず判定されるため安全。

### 累積式 `result = 10*result + digit`

各桁を読むたびに 10 倍して足す。順序が大事:
- 先に digit 加算してから × 10 すると不正確
- ただし最後に × sign するのは OK (`return sign * result`)

### つまずきポイント

- digit 読みループで `i++` を忘れる → 無限ループ (TLE)
- `% MaxInt32` で剰余を取って overflow 対処 → wrap around、クランプにならない
- 「非数字に当たったら停止」を忘れて全部読もうとする → `"123abc"` で失敗
- whitespace を「途中でも skip」してしまう → `"+0 123"` で 123 を返してしまう (正解は 0)
- `+` と `-` 両方ある (`"+-12"`) を「符号 2 回」と扱う → 仕様上は最初の符号 + 次が非数字で 0
- 空文字列・whitespace のみのケース処理を忘れる
- 符号だけ (`"+"`, `"-"`) のケース → 0

### エッジケース表

| 入力 | 出力 | 理由 |
|---|---|---|
| `""` | 0 | 空 |
| `"   "` | 0 | whitespace のみ |
| `"+"` / `"-"` | 0 | 符号後に数字なし |
| `"+-12"` | 0 | 符号後 `-` は非数字 |
| `" +0 123"` | 0 | スペースで停止 |
| `"1337c0d3"` | 1337 | 非数字で停止 |
| `"0-1"` | 0 | `0` 読んで `-` で停止 |
| `"-2147483648"` | -2147483648 | INT_MIN ぴったり |
| `"-2147483649"` | -2147483648 | INT_MIN にクランプ |
| `"2147483648"` | 2147483647 | INT_MAX にクランプ |
| `"words and 987"` | 0 | 先頭が非数字 |

## Connect

### FSM (Finite State Machine) 実装スタイル

複雑なパース処理を「状態 × 入力クラス → 次状態」の表で記述。

**メリット**:
- 「どの状態でどの入力が来たらどうなる」が表で一覧 → 仕様変更に強い
- 網羅性が表で保証される
- compiler の lexer (字句解析器) で頻出

**デメリット**:
- 単純な問題には冗長
- 表のメンテナンスが必要

**実用例**:
- regex エンジン (regex 自体が FSM で実装)
- HTTP/JSON parser
- compiler lexer
- protocol state machine (TCP, etc.)

### 関連問題

- 7 (Reverse Integer): 整数の reverse + overflow 判定
- 65 (Valid Number): atoi の浮動小数点版、FSM が真価を発揮
- 66 (Plus One): 配列で 1 加算、桁上がり処理
- 67 (Add Binary): 2 進数文字列の加算
- 415 (Add Strings): 大きな数の加算
- 43 (Multiply Strings): 大きな数の乗算
- 273 (Integer to English Words): 数値→英語、地獄のエッジケース

### atoi 系問題のテーマ

| 問題 | 入力 | 難所 |
|---|---|---|
| 8 (atoi) | 文字列 | 符号、whitespace、overflow |
| 65 (Valid Number) | 文字列 | 小数点、指数表記、符号位置 |
| 7 (Reverse Integer) | 整数 | overflow |
| 273 (Int to English) | 整数 | 桁の組み合わせ、特殊数 (11-19) |

LC 8 は基礎、LC 65 は FSM の練習、LC 273 は地道な分岐の練習。

### 「クランプ vs wrap」の重要性

数値オーバーフロー処理には 3 種類:
1. **Clamp (saturation)**: 範囲外なら境界値 (LC 8 の方式)
2. **Wrap (modulo)**: 範囲外なら剰余で循環 (C/C++ の signed overflow は UB だが unsigned は wrap)
3. **Error (exception)**: 範囲外なら例外 (Python int は無限精度なので無関係、Java BigInteger も)

LC 8 の仕様は **clamp**。`% INT_MAX` で wrap してしまうと WA。

### キーポイント

- atoi は「**仕様の正確な実装**」を問う問題、アルゴリズムより**仕様読解**が肝
- エッジケースを **表で管理**するのが安全
- overflow 判定は「**追加する前に**チェック」がイディオム
- 拡張性が必要なら FSM、シンプルでよいなら順次処理
