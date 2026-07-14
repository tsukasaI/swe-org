# Reverse Linked List (206)

- **Date**: 2026-03-15
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: Linked list reversal in iterative and recursive approaches

## Memorize

- Iterative: `l` (reversed side), `r` (remaining side) の2ポインタで1パスで反転
- 毎ステップ: next を保存 → `r.Next = l` で向きを変える → `l`, `r` を進める
- `var l *ListNode` で nil 初期化 — Go のポインタ zero value として自然な書き方

## Understand

- Iterative: Time O(N), Space O(1)
- Recursive: Time O(N), Space O(N) — コールスタック分
- Recursive は iterative と同じ l, r パターンを再帰に置き換えただけ
- 実用上は iterative を選ぶ — O(1) space で stack overflow リスクなし

## Connect

- Reverse は linked list 問題の基本操作: Reverse Linked List II (92), Palindrome Linked List (234), Add Two Numbers の正順版などで部品として使う
