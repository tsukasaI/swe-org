# syscall_prac

C言語とシステムコールを使った HTTPサーバーとクライアント。

## 仕様

`GET /calc?query=<数値>+<数値>+...` で整数の合計を返します

| request | response |
| --- | --- |
| GET /calc?query=2+10 | 200 OK body=12 |
| GET /calc/query=1+2+3+4 | 200 OK body=10 |
| GET /calc | 400 Bad Request |
| GET /not_found | 404 Not Found |

## ビルド

```bash
cc -Wall -Wextra -o server server.c
cc -Wall -Wextra -o client client.c
```

## 実行

```bash
./server
# server_f = 3
# listening on port: 8080
```

別ターミナルで:

```bash
# 自作クライアント
./client
# result: 12

# curl
curl 'http://localhost:8080/calc?query=2+10'
# 12
```

サーバー停止: `Ctrl+C` (shutting downを出して終了)

## 構成

### server.c

- interative HTTP server
- socket / setsockopt / bind / listen / accept / read / write / close / sigaction を直接コールする

### client .c

- 固定のリクエストを送る HTTP client
- socket / connect / write / read / close を直接コールする

## 制約 / 既知の挙動

- リクエストヘッダは1回の `read` で全て届く前提。本格運用では `\r\n\r\n` までループする必要あり
- iterative server なので同時接続を同時処理しない(順次処理)
- macOSで動作確認をしたため`accept` の `EINTR` を使っています。 `signal()` ではなく `sigaction()` を使用(macOS の `signal` は`SA_RESTART` 暗黙有効でループから抜けられない)
- リクエストの `+` を演算子として扱う。URL エンコード `%2B` は未対応(`atoi`が先頭の数字だけパースして残りを無視するため、結果がズレる)
- `atoi` を使用しているためオーバーフローや非数値を厳密に検出しない
