# logging

Go言語の標準ライブラリ`log/slog`を拡張したコンテキストベースのロギングライブラリです。

## 概要

このライブラリは、Go の contextを活用してログにコンテキスト情報を自動的に注入する機能を提供します。`sync.Map`を使用したスレッドセーフな実装により、並行処理環境でも安全に使用できます。

## 特徴

- **コンテキストベースの属性注入**: Go's contextを通じてログレコードに自動的に属性を追加
- **スレッドセーフ**: `sync.Map`を使用した並行処理対応
- **不変性**: 既存のcontextを変更せず、新しいcontextを作成
- **最小依存**: Go標準ライブラリのみを使用
- **slog互換**: 標準の`log/slog`と完全互換

## インストール

```bash
go get github.com/HMasataka/logging
```

## 基本的な使用方法

### 1. ログハンドラーの設定

```go
package main

import (
    "context"
    "log/slog"
    "os"

    "github.com/HMasataka/logging"
)

func main() {
    // カスタムハンドラーでロガーを作成
    logger := slog.New(logging.NewHandler(slog.NewJSONHandler(os.Stdout, nil)))
    slog.SetDefault(logger)
}
```

### 2. コンテキストに値を追加

```go
ctx := logging.WithValue(context.Background(), "request_id", "abc123")
ctx = logging.WithValue(ctx, "user_id", 42)
ctx = logging.WithValue(ctx, "metadata", map[string]string{
    "service": "api",
    "version": "v1.0.0",
})
```

### 3. コンテキスト情報付きでログ出力

```go
slog.InfoContext(ctx, "Processing request")
slog.ErrorContext(ctx, "Operation failed", "error", err)
```

出力例：

```json
{
  "time": "2025-11-23T15:49:08.187098+09:00",
  "level": "INFO",
  "msg": "Processing request",
  "request_id": "abc123",
  "user_id": 42,
  "metadata": { "service": "api", "version": "v1.0.0" }
}
```

## API リファレンス

### Core Functions

#### `WithValue(parent context.Context, key string, val any) context.Context`

コンテキストに新しいキーと値のペアを追加します。

```go
ctx := logging.WithValue(context.Background(), "session_id", "sess_123")
```

#### `HasLoggingContext(ctx context.Context) bool`

コンテキストにlogging用のSyncMapが含まれているかどうかを確認します。

```go
if logging.HasLoggingContext(ctx) {
    slog.InfoContext(ctx, "Logging context is available")
}
```

#### `HasValue(ctx context.Context, key string) bool`

コンテキストに指定されたキーの値が存在するかどうかを確認します。

```go
if logging.HasValue(ctx, "user_id") {
    slog.InfoContext(ctx, "User ID is present in context")
}
```

#### `NewHandler(handler slog.Handler) slog.Handler`

標準の`slog.Handler`をラップして、コンテキストベースの属性注入機能を追加します。

```go
handler := logging.NewHandler(slog.NewJSONHandler(os.Stdout, nil))
logger := slog.New(handler)
```

## 使用例

### 基本的な使用例

```go
package main

import (
    "context"
    "log/slog"
    "os"

    "github.com/HMasataka/logging"
)

func main() {
    // ロガー設定
    logger := slog.New(logging.NewHandler(slog.NewJSONHandler(os.Stdout, nil)))
    slog.SetDefault(logger)

    // コンテキストに値を追加
    ctx := logging.WithValue(context.Background(), "trace_id", "trace_123")
    ctx = logging.WithValue(ctx, "span_id", "span_456")

    processRequest(ctx)
}

func processRequest(ctx context.Context) {
    slog.InfoContext(ctx, "Request processing started")

    if err := someOperation(); err != nil {
        slog.ErrorContext(ctx, "Operation failed", "error", err)
        return
    }

    slog.InfoContext(ctx, "Request processing completed")
}
```

## 開発

### 必要要件

- Go 1.25.1以上
- [Task](https://taskfile.dev/)
