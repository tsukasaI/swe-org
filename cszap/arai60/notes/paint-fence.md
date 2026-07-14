# Paint Fence (276)

- **Date**: 2026-03-23
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: DP の状態分割パターン

## Memorize

- 状態を「前と同じ色 (same)」と「前と違う色 (diff)」に分ける
- 漸化式:
  - `same[i] = diff[i-1]`（same の後に same → 3連続 NG なので、diff からしか来れない）
  - `diff[i] = (same[i-1] + diff[i-1]) * (k-1)`
- 初期値: `same[0] = 0`, `diff[0] = k`（n=1: 前がないので same は 0）
- 答え: `same[n-1] + diff[n-1]`

## Understand

### 計算量

- Time: O(N), Space: O(N)（配列版）
- Space O(1) に改善可: 各ステップで前の same, diff だけ必要なので変数2つで十分

### DP の考え方

- 「3本連続同色 NG」という制約を「前と同じ/違う」の2状態に分解
- same にできるのは前が diff のときだけ → 3連続を自然に防げる
- 早期リターン（n=1, n=2）は不要 — ループが 0〜1 回で正しく動く

## Connect

- DP の状態分割パターン: 制約を状態に変換して漸化式を立てる
- House Robber (198) も「前を選んだ/選ばなかった」の2状態 DP で同じ構造
- Climbing Stairs (70) も同様の線形 DP
