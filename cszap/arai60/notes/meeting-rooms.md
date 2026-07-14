# Meeting Rooms (252) / Meeting Rooms II (253)

- **Date**: 2026-04-30
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: 区間問題の基礎、ソート + 隣接比較 / min-heap / sweep line

## Memorize

### LC 252: 全会議に出席可能か

```go
func canAttendMeetings(intervals [][]int) bool {
    if len(intervals) == 0 {
        return true
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    for i := 1; i < len(intervals); i++ {
        if intervals[i-1][1] > intervals[i][0] {
            return false
        }
    }
    return true
}
```

### LC 253: 最低何部屋必要か (min-heap 版)

```go
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)         { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func minMeetingRooms(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    h := &IntHeap{}
    heap.Init(h)
    for _, iv := range intervals {
        if h.Len() > 0 && (*h)[0] <= iv[0] {
            heap.Pop(h)        // 終わった会議の部屋を再利用
        }
        heap.Push(h, iv[1])
    }
    return h.Len()
}
```

### LC 253: sweep line 版（heap 不要、より単純）

```go
func minMeetingRooms(intervals [][]int) int {
    n := len(intervals)
    starts := make([]int, n)
    ends := make([]int, n)
    for i, iv := range intervals {
        starts[i] = iv[0]
        ends[i] = iv[1]
    }
    sort.Ints(starts)
    sort.Ints(ends)
    rooms, maxRooms, j := 0, 0, 0
    for i := 0; i < n; i++ {
        if starts[i] < ends[j] {
            rooms++
            if rooms > maxRooms {
                maxRooms = rooms
            }
        } else {
            j++   // 1 つ終わったので部屋数据置き
        }
    }
    return maxRooms
}
```

## Understand

### 計算量

| 解法 | Time | Space |
|---|---|---|
| LC 252 (sort + 隣接比較) | O(N log N) | O(log N) (sort 内部) |
| LC 253 (min-heap) | O(N log N) | O(N) (heap) |
| LC 253 (sweep line) | O(N log N) | O(N) (starts, ends 分離) |

### 区間の重なり判定

2 区間 A, B が重なる条件は色々な書き方があるが、A が先に始まる前提なら:

- **重なる**: `A.end > B.start`
- **重ならない (touching OK)**: `A.end <= B.start`

「touching」(`A.end == B.start`) を重なりに含めるかは問題による:
- LC 252 / 253: touching は **OK**（同じ人が両会議に出席可能）
- 半開区間 `[start, end)` とみなすと自然

なので比較は **`>` (strictly greater)** を使う。`>=` だと touching を誤検出。

### LC 252: なぜ隣接比較で十分か

ソート後、もし `intervals[i]` と `intervals[j]` (i < j-1) が重なるなら、`intervals[i]` と `intervals[i+1]` も必ず重なる。

**証明**:
- start でソート済み: `intervals[i+1].start ≤ intervals[j].start`
- 重なり前提: `intervals[i].end > intervals[j].start ≥ intervals[i+1].start`
- → `intervals[i].end > intervals[i+1].start` → 重なる

つまり「飛ばしの重なり」が起きるなら「隣接の重なり」も必ず起きる → 隣接だけ調べれば十分。

### LC 253: min-heap のメンタルモデル

heap の中身 = **「現在進行中の会議の end 時刻リスト」**:

新しい会議 `iv` が始まるとき:
1. heap が空でない、かつ最早 end (`heap.Top()`) ≤ `iv.start` なら
   → その部屋は空いた → pop (再利用)
2. 新会議の end を push
3. ループ後の heap サイズ = 必要部屋数

**重要**: 「pop は最大 1 回」で十分。複数会議が終わっていても、1 つの新会議が同時に複数の部屋を使うわけではないので。

### LC 253: sweep line のメンタルモデル

時系列に **start/end イベント** を分けて並べる:
- start イベント → 部屋数 +1
- end イベント → 部屋数 -1
- 全イベント中の最大値 = 答え

実装は「starts と ends を別々にソート」して 2 ポインタで進める形が綺麗。

**touching の扱い**: `starts[i] < ends[j]` (`<`、`<=` ではない) を使うと、同時刻の end と start で end を先に処理 → 部屋を再利用。

### heap 版 vs sweep line 版の比較

| | heap 版 | sweep line 版 |
|---|---|---|
| 実装量 | 多い (heap interface 定義) | 少ない |
| 時間計算量 | O(N log N) | O(N log N) |
| 空間 | O(N) (heap) | O(N) (starts, ends) |
| ストリーミング (online) 対応 | 可（追加ごとに更新） | 不可（offline 限定） |
| 「**どの**会議がどの部屋に」追跡 | 可（heap top で識別） | 不可（数だけ） |
| 容量付き拡張 | 自然（部屋オブジェクトに置換） | やや不自然 |

**使い分け**:
- 「何部屋必要か」だけ → sweep line が簡潔
- 「部屋を割り当てる」or「容量制約あり」→ heap (= 部屋 ID 管理)

### `for` で複数 pop vs `if` で 1 回 pop

```go
// パターン A (推奨): 1 つ新会議 = 最大 1 つ古い会議が出ていく
if h.Len() > 0 && (*h)[0] <= iv[0] {
    heap.Pop(h)
}
heap.Push(h, iv[1])
return h.Len()                // 最終 size = 答え

// パターン B: 終わった会議をすべて取り出す
for h.Len() > 0 && (*h)[0] <= iv[0] {
    heap.Pop(h)
}
heap.Push(h, iv[1])
maxRooms = max(maxRooms, h.Len())   // 各時点の max を追跡
return maxRooms
```

両方正しいが意味論が異なる:
- A: heap は「これまで使った全部屋」（最終 size = 全使用部屋数 = 答え）
- B: heap は「現在進行中の部屋」（max を追跡する必要）

A の方が定数倍速いが、B の方が直感的（「heap = 進行中の会議」と一致）。

### つまずきポイント

- LC 252 で `>=` を使ってしまう → touching で誤判定
- LC 253 min-heap で「pop 後 push」でなく「常に push」してしまう → 部屋数が減らない
- sweep line で `<` と `<=` を逆にする → touching の扱いが間違う
- 「全 N 個の重なりペア」を全列挙しようとする → O(N^2) で遅い

## Connect

### 区間問題の 4 大パターン

| 問題 | 質問 | 解法 |
|---|---|---|
| LC 252 | 全部出席可能? | sort + 隣接比較 |
| LC 253 | 最低部屋数? | min-heap or sweep line |
| LC 56 (Merge Intervals) | 重なる区間をマージ | sort + 隣接合体 |
| LC 57 (Insert Interval) | 1 区間追加してマージ | linear scan |
| LC 435 (Non-overlapping) | 何個削れば重ならない? | sort by **end** + greedy |
| LC 1851 (Smallest Range) | 各点で最小被覆区間 | offline sort + heap |

### 「sort by start」 vs 「sort by end」

- LC 56, 252, 253: **start でソート**（時系列処理）
- LC 435: **end でソート**（greedy: 早く終わる方を残す）
- LC 1288 (Remove Covered): start ↑ + end ↓（包含判定）

ソートキーの選び方が解法を決める。「end でソートする greedy」は **interval scheduling** の典型。

### 動的な区間管理（実務）

毎回ソートし直すのが無駄なケース（カレンダーアプリ等）:

| データ構造 | 用途 | 操作の計算量 |
|---|---|---|
| **Interval Tree** (区間木) | 1 点が含まれる区間検索、区間と重なる区間検索 | O(log N + K) (K=結果数) |
| **Segment Tree** (セグメント木) + 座標圧縮 | 各時点の「同時進行数」、range update | O(log N) per op |
| **Balanced BST** (TreeMap) | 区間集合の挿入削除 + 隣接検索 | O(log N) |
| **Sweep line + sorted events** | offline 一括処理 | O(N log N) total |

LC 253 を「動的に追加/削除しながら最小部屋数」にすると、Segment Tree (range +1/-1, max query) が定番。

### sweep line の汎用テンプレ

```
1. start イベント (+1) と end イベント (-1) を全部用意
2. 時刻順にソート (同時刻は end を先に処理 = touching を許容)
3. 線形に走らせて累積値の最大を取る
```

応用例:
- LC 56, 253: 区間
- LC 218 (Skyline): 高さ付きイベント、heap で max-height
- LC 850 (Rectangle Area II): 2D sweep + segment tree
- LC 759 (Free Time): 全社員の空き時間

### キーポイント

- 区間は「start, end の 2 値」だが、それを**点イベント**に分解すると線形化できる
- ソートキー（start vs end）と比較演算（< vs ≤）で touching の扱いが決まる
- 「全ペア比較 O(N^2)」 → 「ソート + 線形 O(N log N)」 への変換が定番

### 容量制約付きへの拡張（実務観点）

**問題設定**: 各会議に `attendees` (人数)、各部屋に `capacity` (収容上限)。
- 制約: 会議は capacity ≥ attendees の部屋に割り当てる
- 目標: 必要部屋数を最小化、または与えられた部屋集合で feasibility 判定

**結論**: **NP-hard** (Bin Packing × Interval Scheduling on Multiple Machines)。多項式時間最適アルゴリズムは存在しない (P=NP 未解決)。

**ヒューリスティクス**:

| 方針 | 説明 | 落とし穴 |
|---|---|---|
| Best-Fit (start 順) | 各会議を「fit する最小の部屋」へ | 大きい部屋を無駄遣いするケース |
| First-Fit Decreasing | 大人数会議から順に詰める | start 順を崩すので時間制約と相性悪 |
| Multi-heap | 容量段階別 min-heap | 段階間の相互作用を扱えない |
| ILP / CP-SAT | 整数計画で最適解 | スケールしない (~数百会議まで) |

**Best-Fit が失敗する例**:
- 部屋: A (cap 100), B (cap 10)
- 会議:
  - M1: [0, 10], 5 人
  - M2: [5, 15], 5 人
  - M3: [12, 20], 50 人
- Best-Fit: M1→B, M2→A, M3 は B (cap 不足) も A (15 まで占有) も不可 → 3 部屋目必要
- 最適: M1→A, M2→B, M3→A (M1 終了後の 12-20) → 2 部屋で OK

→ 「smallest fitting」貪欲が裏目に出る。**未来の大人数会議のために大部屋を残す** という発想が必要だが、未来知識がないと最適化不可。

**実システム例**:
- Kubernetes scheduler: 同じ構造 (CPU/memory + node availability + 時間)
- Filter (制約) + Score (heuristic) のプラグイン構成で対応
- 完璧でなくても「実用的に妥当」を狙う

**面接での示唆**: 「容量制約あり」を見た瞬間、簡単な解は無いと察する。素朴な heap 拡張では最適にならない。NP-hard を認めた上で実務的妥協を提示するのがプロの解答。
