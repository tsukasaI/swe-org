# Pow(x, n) (50)

- **Date**: 2026-04-25
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 高速累乗 (binary exponentiation) で O(log n) で x^n を求める

## Memorize

### 再帰版

```go
func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    if n < 0 {
        n = -n
        x = 1 / x
    }
    if n%2 == 0 {
        return myPow(x*x, n/2)
    }
    return x * myPow(x*x, n/2)
}
```

### ループ版（O(1) space）

```go
func myPow(x float64, n int) float64 {
    if n < 0 {
        x = 1 / x
        n = -n
    }
    result := 1.0
    for n > 0 {
        if n&1 == 1 {
            result *= x
        }
        x *= x
        n >>= 1
    }
    return result
}
```

## Understand

### 計算量

- Time: O(log n)
- Space: O(log n)（再帰版）/ O(1)（ループ版）

### アルゴリズムの本質

`n` を**ビット表現**で見る:
- `n = 13 = 1101` なら `x^13 = x^8 * x^4 * x^1`
- ビットが 1 の位置に対応する `x^(2^k)` を結果に掛ける
- `x` を毎回 2乗することで `x^1, x^2, x^4, x^8, ...` と準備していく

漸化式:
```
pow(x, n) = pow(x*x, n/2)        if n is even
pow(x, n) = x * pow(x*x, n/2)    if n is odd
pow(x, 0) = 1                    base case
```

### トレース例: `x=2, n=10` (`1010`)

| n (bin) | n & 1 | result | x |
|---|---|---|---|
| 1010 | 0 | 1 | 2 |
| 101 | 1 (×2) | 4 | 4→16 |
| 10 | 0 | 4 | 16→256 |
| 1 | 1 (×256) | 1024 | 256→65536 |
| 0 | end | 1024 | - |

`x^10 = x^8 * x^2 = 256 * 4 = 1024` ✓

### 負の n の扱い

`x^(-n) = 1 / x^n` なので、`x = 1/x, n = -n` に変換。

**罠**: `n = INT_MIN` のとき `-n` はオーバーフロー。
- Go (int は 64bit): 問題なし
- Java/C++ (int は 32bit): `long` キャストが必要

### エッジケース

- `n = 0`: 1（base case）
- `x = 0, n > 0`: 自然に再帰で 0 になる（特別ハンドリング不要）
- `x = 0, n <= 0`: 問題の制約で発生しない（"Either x is not zero or n > 0"）

### 浮動小数点の精度

- float64 は仮数部 52 ビット → 約 15-16 桁
- O(log n) でも累積誤差が出る（n=10^9 で 30 回程度の掛け算）
- 対策:
  - `math/big` の `big.Float` で任意精度
  - `math.Exp(n * math.Log(x))` で 1 回の log+exp に置き換え（x>0 限定、極大 n で overflow リスク）
  - LeetCode は 1e-5 の許容誤差なので普通は気にしなくて良い

### つまずきポイント

- 素朴な `for i := 0; i < n; i++ { result *= x }` は O(n) で TLE
- 負の n を変換し忘れる
- `n & 1` ではなく `n % 2` でも動くが、ビット演算の方が速度・意図が明確
- 再帰版は stack overflow のリスクは低い（log n 段なので n=10^9 でも 30 段）

## Connect

### Fast exponentiation の汎用テンプレート

「結合則を満たす二項演算 + 単位元 (= モノイド)」があれば O(log n) で n 乗できる:

```go
type Monoid[T any] interface {
    Mul(a, b T) T
    Identity() T
}

func FastPow[T any](x T, n int, m Monoid[T]) T {
    result := m.Identity()
    for n > 0 {
        if n&1 == 1 {
            result = m.Mul(result, x)
        }
        x = m.Mul(x, x)
        n >>= 1
    }
    return result
}
```

### 応用例

- **数の累乗**: `LC 50` (この問題)
- **行列累乗 → フィボナッチ O(log n)**: `[[1,1],[1,0]]^n` で `F(n)` が得られる
- **modular exponentiation**: 暗号 (RSA など)、`(a*b) mod p` を Mul とする
- **文字列繰り返し**: `Mul = concat`、O(log n) で `s` を n 回繰り返した文字列の特定位置を求める
- **グラフの k 乗ステップ移動**: 隣接行列の k 乗で k ステップ後の到達可能性

### 関連問題

- 70 (Climbing Stairs): フィボナッチ系、行列累乗で O(log n) も可能
- 372 (Super Pow): `x^(huge n) mod 1337`、modular exponentiation
- 509 (Fibonacci Number): 通常 O(n) DP、行列累乗で O(log n)
