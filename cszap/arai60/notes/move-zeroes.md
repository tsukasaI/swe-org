# Move Zeroes (283)

- **Date**: 2026-04-30
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 2 ポインタ (read/write 型) で in-place 移動、操作数最小化

## Memorize

### read/write 型 2 ポインタ（基本）

```go
func moveZeroes(nums []int) {
    write := 0
    for read := 0; read < len(nums); read++ {
        if nums[read] != 0 {
            nums[write], nums[read] = nums[read], nums[write]
            write++
        }
    }
}
```

### 操作数最小化版（先頭の非ゼロをスキップ）

```go
func moveZeroes(nums []int) {
    write := 0
    // 最初の 0 を探す（ここまでは swap 不要）
    for write < len(nums) && nums[write] != 0 {
        write++
    }
    for read := write; read < len(nums); read++ {
        if nums[read] != 0 {
            nums[write], nums[read] = nums[read], nums[write]
            write++
        }
    }
}
```

### 順序保持なし版（参考、LC 27 系）

```go
// 注意: LC 283 では順序保持必須なので使えない
func moveZeroesUnordered(nums []int) {
    l, r := 0, len(nums)-1
    for l <= r {
        if nums[l] == 0 {
            nums[l], nums[r] = nums[r], nums[l]
            r--
        } else {
            l++
        }
    }
}
```

## Understand

### 計算量

- Time: **O(N)** — 各要素を高々 2 回 (read で 1 回、write で swap)
- Space: **O(1)** — in-place

### 2 ポインタの役割

| ポインタ | 役割 |
|---|---|
| `read` | 配列を左から右にスキャン |
| `write` | 「次に非ゼロを置く位置」 = 既に処理済み非ゼロ列の右端 |

不変条件:
- `nums[0..write)` = 元の非ゼロ要素を順序保持して詰めたもの
- `nums[write..read)` = すべて 0 (まだ後ろに送っていない)
- `nums[read..)` = 未処理

### swap が「順序を保持」できる理由

`read` が非ゼロを見つけて `write` と swap するとき:
- `nums[write]` は必ず 0（不変条件より）
- 0 を read 側に送っても、read はもう前に戻らないので問題ない
- 非ゼロは write 位置に移るので順序は保たれる

「0 と非ゼロの swap」しか起きない → 非ゼロ同士の順序は崩れない。

### 操作数最小化の意味

LC 283 follow-up: "Could you minimize the total number of operations done?"

基本版で `[1,2,3,4,0]` を処理すると:
- read=0,1,2,3 で各 swap (`nums[w], nums[r] = nums[r], nums[w]`、ただし w==r で自分と自分の swap)
- 不要な書き込みが 4 回発生

最適化版 (先頭の非ゼロをスキップ):
- 最初の 0 (index 4) まで write を進めるだけ
- 以降 swap も 0 回
- **無意味な write を完全に省ける**

代替実装: swap せず「非ゼロを書く + 後で 0 で埋める」も等価:

```go
func moveZeroes(nums []int) {
    write := 0
    for _, v := range nums {
        if v != 0 {
            nums[write] = v
            write++
        }
    }
    for i := write; i < len(nums); i++ {
        nums[i] = 0
    }
}
```

ただし全要素非ゼロでも 2 回ループするので、操作数最小化観点では「先頭スキップ + swap」の方が良い。

### つまずきポイント

- `read` と `write` の進み方を混同 → write は「非ゼロ確定時のみ」進める
- `write++` を swap 前にしてしまう → 位置がズレる
- 順序保持なし版を LC 283 に使ってしまう → WA（順序が崩れる）
- 「全要素 0」「全要素非ゼロ」「`[0]`」のエッジケースを忘れる

## Connect

### LC 27 (Remove Element) との比較

| | LC 283 (Move Zeroes) | LC 27 (Remove Element) |
|---|---|---|
| 対象 | 0 を後ろへ | 指定値を削除 |
| 順序保持 | 必須 | 任意 |
| 配列長 | 不変 (in-place 移動) | 残った要素数を返す |
| 最適化 | read/write 型 swap | 両端ポインタで write 数最小化 |

LC 27 は順序自由なので「両端から走らせる」最適化が可能 → write 回数 = 残す要素数のみ。
LC 283 は順序保持なので read/write 型に固定。

### 2 ポインタのパターン分類

| パターン | 例 | 特徴 |
|---|---|---|
| read/write 型 (slow/fast) | LC 26, 27, 283 | 同じ方向、write は read を追従 |
| 両端から (left/right) | LC 11, 15, 167 | 反対方向、収束 |
| sliding window | LC 3, 76, 209 | 同じ方向、両方右に進むだけ |

read/write 型の本質: **「圧縮 / フィルタリング in-place」**。

### 関連問題

- 26 (Remove Duplicates from Sorted Array): write は「異なる値が出るたびに進める」
- 27 (Remove Element): 順序自由なら両端ポインタが最適
- 80 (Remove Duplicates from Sorted Array II): 各値最大 2 回まで OK
- 75 (Sort Colors): 3 色問題、write 2 つ + read 1 つ (Dutch National Flag)
- 88 (Merge Sorted Array): 後ろから write すると O(1) 追加スペース

### read/write 型の汎用テンプレ

```go
write := 0
for read := 0; read < len(nums); read++ {
    if 残すべき条件(nums[read]) {
        nums[write] = nums[read]   // または swap
        write++
    }
}
// write は新しい配列長になる
```

「残すべき条件」が問題ごとに違う:
- LC 27: `nums[read] != val`
- LC 283: `nums[read] != 0`
- LC 26: `read == 0 || nums[read] != nums[read-1]`
