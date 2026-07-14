# Capacity To Ship Packages Within D Days (1011)

- **Date**: 2026-04-25
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 「答えに対する二分探索」で最小容量を求める

## Memorize

- `lo = max(weights)`, `hi = sum(weights)` で範囲を決める
- 境界探索パターン (`l < r`, `r = c`, return `lo`)
- 判定関数 `canShip(cap)`: greedy で各日詰められるだけ詰める、必要日数 ≤ days か返す
- `daysUsed` を直接管理すると off-by-one を防ぎやすい

```go
func canShip(weights []int, days, cap int) bool {
    daysUsed := 1
    sum := 0
    for _, w := range weights {
        sum += w
        if sum > cap {
            daysUsed++
            if daysUsed > days {
                return false
            }
            sum = w
        }
    }
    return true
}
```

## Understand

### 計算量

- N = `len(weights)`, W = `sum(weights)` (二分探索の範囲)
- Time: **O(N * log W)** = O(N * log(sum(weights)))
- Space: O(1)
- 制約: N ≤ 5*10^4, W ≤ 2.5*10^7, log W ≈ 25 → 約 1.25*10^6 操作で十分高速

### 「答えに対する二分探索」の本質

- これまでの二分探索は「配列のインデックス」を探した
- 今回は「**整数値そのもの**」を探す
- log の中身が「配列長 N」ではなく「**値の範囲 W**」になる
- パターン: 「容量を増やすと判定がどこかで OK に変わる」=「単調な述語」を二分探索

### 探索範囲の根拠

- **lo = max(weights)**: 最重量パッケージが単独で1日必要 → 容量がそれ未満では絶対不可
- **hi = sum(weights)**: 全部1日で運べる容量 → これ以上は無駄
- 注意: lo を `min(weights)` にしてしまうのは典型ミス（最重量が運べない）

### 単調性の厳密な議論

主張: 容量 X で D 日に運べるなら、容量 Y > X でも D 日に運べる。

証明: X での分割 `[S1][S2]...[Sd]` (sum(Si) ≤ X) があるとき、Y >= X なので sum(Si) ≤ Y も成立 → 同じ分割が Y でも使える → Y でも d ≤ D 日。

つまり最小必要日数 `f(cap)` は **cap に対して非増加**。`f(cap) <= D` という述語は「ある境界より上で常に true」 → 二分探索可能。

### canShip が greedy で正しい理由

**Exchange argument**:
- 順序固定の制約下では「各日に詰められるだけ詰める」が最適
- 1日目に少なく詰める解は、後ろの日から前倒しすることで greedy に変形できる（容量に余裕があるので可能）
- 前倒しで日数は増えない → greedy ≤ OPT

DP でも解けるが O(N^2) per call → TLE。greedy O(N) が必須。

### off-by-one バグの防ぎ方

「split 回数 (count)」と「使った日数 (daysUsed)」は別物:
- `count = N` → `daysUsed = N + 1`
- `count == days` チェックは1つズレやすい

→ `daysUsed` を直接管理するほうが意味が明確で、+1/-1 の混乱を防げる。

### つまずきポイント

- `lo = min(weights)` にしてしまう → 最重量が運べないケースで破綻
- `daysUsed` の初期値を 0 にする → 1日目をカウントしない off-by-one
- `sum > cap` の境界（`>` vs `>=`）: ちょうど cap と等しいときは積める → `>` が正解
- canShip で DP を使ってしまう → O(N^2) で TLE

## Connect

### 「答えに対する二分探索」の双子問題

- **410 (Split Array Largest Sum)**: m 個の連続サブ配列に分割、各合計の最大を最小化 → LC 1011 と全く同じ構造（容量↔分割合計、日数↔分割数）
- **1552 (Magnetic Force Between Two Balls)**: k 個のボールを最大の最小距離で配置 → 距離を二分探索、greedy で配置可能か判定
- **875 (Koko Eating Bananas)**: 食べる速度 K を二分探索、H 時間で全部食べきれるか
- **774 (Minimize Max Distance to Gas Station)**: ステーション追加で最大距離を最小化（実数の二分探索）

### 二分探索の2大パターン整理

| パターン | 探索対象 | 例 |
|---|---|---|
| インデックス探索 | 配列の位置 | LC 33, 35, 153 |
| 値探索（答え探索） | 整数値そのもの | LC 1011, 410, 1552, 875 |

「答え探索」を見抜くキーワード:
- 「最大値を最小化 / 最小値を最大化」
- 「ちょうど D 個に分割 / D 日以内に / D 個配置」
- 単調な judge 関数が書けるかをまず確認

### 関連の判定ロジック

- canShip 系の greedy は exchange argument で正当性を示す
- 各日/各回に「詰められるだけ詰める」が最適になるのは、順序固定 or 同等のもののみで成立
