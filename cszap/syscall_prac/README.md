# syscall_prac

C言語とシステムコールを使った HTTP サーバーの練習リポジトリです。

## 概要

- `GET /calc?query=2+10` のようなリクエストを受け取り、計算結果を返す HTTP サーバー

## 実行方法

### 1. イメージをビルド

```bash
docker build -t syscall_prac .
```

### 2. C ファイルをコンパイル

```bash
docker run --rm -v $(pwd):/workspace syscall_prac gcc -o server server.c
```

### 3. サーバーを起動

```bash
docker run --rm -it -v $(pwd):/workspace -p 8080:8080 syscall_prac ./server
```

### 4. 動作確認

```bash
curl "http://localhost:8080/calc?query=2+10"
```

---

## gdb でデバッグする場合

```bash
docker run --rm -it -v $(pwd):/workspace --cap-add=SYS_PTRACE syscall_prac gdb ./server
```

> `--cap-add=SYS_PTRACE` は gdb がプロセスをアタッチするために必要です。
