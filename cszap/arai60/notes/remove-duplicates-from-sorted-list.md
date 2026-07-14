# Remove Duplicates from Sorted List (83 & 82)

- **Date**: 2026-03-14
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: Handle duplicate removal in sorted linked lists

## Memorize

- Sorted list → duplicates are always adjacent, so one pass with inner loop suffices
- Dummy node (sentinel node): create a fake head to simplify edge cases where the real head might be removed
- Return `dummy.Next`, not `dummy` — the dummy's Val is never used in comparisons

## Understand

### 83: Keep one copy of each duplicate

- Track `cur`, use inner loop to skip nodes with same Val, then link `cur.Next` to the first different node
- Time O(N), Space O(1)

### 82: Remove all duplicates entirely

- Dummy node + `cur` を「最後に確定した（残す）ノード」として扱い、`cur.Next` 以降を調べる
- `cur` を削除対象自身にすると、前のノードに戻れず繋ぎ替えできない — `cur` は常に確定済みノードに置く
- Inner loop で `next.Val == next.Next.Val` を比較し、`next` を最後の重複ノードまで進める
- 重複があれば `cur.Next = next.Next` で重複グループ全体をスキップ、なければ `cur` を進める
- Time O(N), Space O(1)

## Connect

- Dummy node pattern applies to: remove nth node, merge sorted lists, partition list, etc.
- Sorted constraint is key — without it, would need HashMap O(N) space for duplicate detection
