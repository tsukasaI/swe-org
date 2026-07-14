# Valid Parentheses (20)

- **Date**: 2026-03-15
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: Stack-based bracket matching

## Memorize

- 閉じ括弧 → 開き括弧のマップを用意し、閉じ括弧が来たら stack の top と比較
- Stack が空の状態で閉じ括弧が来たら即 `return false` — 空チェックを忘れると panic
- 最後に `len(stack) == 0` で未対応の開き括弧がないか確認

## Understand

- Stack は LIFO なので「最も内側の括弧から閉じる」という制約に自然に対応する
- Time O(N), Space O(N)
- Edge cases: 閉じ括弧のみ (`")"`)、開き括弧のみ (`"("`)、空文字列

## Connect

- Stack パターンは括弧系問題全般に共通: Generate Parentheses (22), Longest Valid Parentheses (32), Min Remove to Make Valid (1249) など
