# Linked List Cycle (141 & 142)

- **Date**: 2026-03-14
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: Understand cycle detection in linked lists and Floyd's algorithm

## Memorize

- HashMap approach: use `*ListNode` (pointer) as key, not `Val` — same value can exist in different nodes
- Floyd's Tortoise and Hare: slow moves 1 step, fast moves 2 steps — relative speed is 1, so they always meet in a cycle
- Floyd's Phase 2 (find cycle start): after meeting, move one pointer to head, both advance 1 step — they meet at cycle start

## Understand

### Why Floyd's Phase 2 works

- Let A = distance from head to cycle start, B = distance from cycle start to meeting point, C = cycle length
- slow travels `A + B`, fast travels `2(A + B)`
- fast also travels `A + B + nC` (n full cycles), so `2(A + B) = A + B + nC` → `A + B = nC` → `A = nC - B`
- From meeting point, advancing A steps: `B + A = B + nC - B = nC` — back to cycle start
- From head, advancing A steps: arrives at cycle start
- Therefore both pointers meet at cycle start without knowing A

### Initial pointer positions matter

- `slow, fast := head, head` (same start) — matches the math proof, use `do-while` style (move then compare)
- `slow, fast := head, head.Next` (offset start) — works for detection only, but breaks Phase 2 assumptions

## Connect

- 141 (Linked List Cycle): detection only — HashMap O(N) space or Floyd's O(1) space
- 142 (Linked List Cycle II): find cycle start node — Floyd's Phase 1 + Phase 2
- Cycle length: after meeting, advance one pointer until they meet again, count steps
