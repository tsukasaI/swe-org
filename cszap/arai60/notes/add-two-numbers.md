# Add Two Numbers (2)

- **Date**: 2026-03-14
- **Context**: arai60 LeetCode practice — coding interview prep
- **Goal**: Linked list arithmetic with carry handling

## Memorize

- Dummy node で結果リストを構築、`dummy.Next` を返す
- ループ条件: `l1 != nil || l2 != nil || carry == 1` — carry の条件がないと最上位桁の繰り上がりを見逃す
- `sum % 10` で現在の桁、carry は 0 or 1 なので明示的に `if sum > 9` で十分（`sum / 10` は汎用的だが今回は過剰）

## Understand

- Reverse order で格納されているので、先頭から順に足せば自然に下の桁から処理できる
- リストの長さが異なる場合、短い方は 0 として扱う
- carry は最大 1（9+9+1=19）なので bool 的に扱える

### 正順 (most significant digit first) の場合

- 上の桁から順に足しても繰り上がりを処理できない
- リストの長さが異なると桁の位置がずれる
- 対処法: リストを reverse してから同じアルゴリズムを適用するのがシンプル

## Connect

- Dummy node パターン: 結果リストを構築する linked list 問題全般で使える
- Carry の扱いは multiply strings (43) などの数値演算系問題にも共通
