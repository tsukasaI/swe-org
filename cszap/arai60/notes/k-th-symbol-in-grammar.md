# K-th Symbol in Grammar (779)

- **Date**: 2026-04-25
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 再帰の親子関係で実体を作らずに k 番目の値を求める

## Memorize

### 再帰版

```go
func kthGrammar(n int, k int) int {
    if n == 1 && k == 1 {
        return 0
    }
    if k%2 == 1 {
        return kthGrammar(n-1, (k+1)/2)
    }
    return 1 - kthGrammar(n-1, (k+1)/2)
}
```

### ループ版（O(1) space）

```go
func kthGrammar(n int, k int) int {
    count := 0
    for n > 1 {
        if k%2 == 0 {
            count++
        }
        k = (k + 1) / 2
        n--
    }
    return count % 2
}
```

### popcount 解法（最もエレガント）

```go
import "math/bits"

func kthGrammar(n int, k int) int {
    return bits.OnesCount(uint(k-1)) % 2
}
```

## Understand

### 計算量

| 解法 | Time | Space |
|---|---|---|
| 素朴（実体構築） | O(2^n) | O(2^n) |
| 再帰 | O(n) | O(n) |
| ループ | O(n) | O(1) |
| popcount | O(log k) = O(n) | O(1) |

注: 制約 `k ≤ 2^(n-1)` から `log k ≤ n-1`、つまり `O(n) ≈ O(log k)`。

### 親子関係の発見

n 行目の k 番目は、n-1 行目の **`(k+1)/2`** 番目から派生:
- n 行目の 1, 2 番目 → n-1 行目の 1 番目
- n 行目の 3, 4 番目 → n-1 行目の 2 番目
- n 行目の k 番目 → n-1 行目の `(k+1)/2` 番目

`k/2` だと off-by-one（k=1 で 0 になる）→ `(k+1)/2` が正しい

### 反転の規則

展開ルール:
```
0 → 01   (左=0, 右=1)
1 → 10   (左=1, 右=0)
```

これより:
- **k が奇数** = 親の左の子 → 親と**同じ値**（反転なし）
- **k が偶数** = 親の右の子 → 親と**逆の値**（反転1回）

ポイント: 「右の子なら反転」は親の値に関係なく普遍的。なので各ステップを独立にカウントできる。

### popcount 解法の原理

ループ版の「k が偶数のとき count++」を観察:
- 「k が偶数」 ↔ 「k の最下位ビットが 0」 ↔ 「k-1 の最下位ビットが 1」
- 各ステップで k = (k+1)/2 する操作は k-1 の右シフトに相当
- 右シフトで「1」が落ちる回数 = k-1 のビット中の 1 の総数 = popcount(k-1)

→ 答え = `popcount(k-1) % 2`

### 完全二分木の解釈

```
           0           ← Row 1 (root)
         /   \
        0     1        ← Row 2
       / \   / \
      0   1 1   0      ← Row 3
```

- リーフ k 番目 = root から k-1 を 2 進数で表現したパス（0=左、1=右）
- パスで「右」を通った回数の偶奇 = リーフのラベル
- これが popcount の正体

### つまずきポイント

- 親の位置を `k/2` にすると off-by-one（k=1 で 0 になる）
- 反転条件を「k が奇数のとき」と勘違いする（実際は偶数）
- 反転条件を「親の値と現在の k の関係」で考えすぎる（実は現在の k の偶奇だけで決まる）

## Connect

### 関連の数学的構造

- **Gray code**: 連続する数の 2 進表現が 1 ビットずつしか違わない符号、ハードウェアで使う
- **Thue-Morse 数列**: `t_n = popcount(n) % 2`、まさにこの問題と同じ構造
- **再帰的フラクタル**: シェルピンスキーの三角形、Cantor 集合などの自己相似構造

### 関連問題

- 50 (Pow(x, n)): 同じ「実体を作らず O(log) で求める」パターン
- 372 (Super Pow): modular exponentiation、ビット表現で計算
- 89 (Gray Code): Gray code を生成する
- 191 (Number of 1 Bits): popcount そのもの
- 338 (Counting Bits): 0 から n までの popcount を全部求める

### 実務での応用

- **暗号系**: ハッシュ関数、暗号学的乱数生成での bit 操作
- **データ圧縮**: 自己相似性を利用した圧縮アルゴリズム
- **エラー訂正**: Gray code を使った差分検出
- **競技プログラミング**: 「k 番目の何か」系で実体を作らずビット計算で済ませるテクニック
