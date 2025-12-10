# Go Echo Sample

Echo と OpenAPI (oapi-codegen) で組んだシンプルなサーバーのサンプルです。

## 必要環境

- Go 1.25.x 以上

## 開発手順

- OpenAPI (`docs/openapi.yaml`) から生成:

  ```bash
  make generate
  ```

- サーバー起動:

  ```bash
  go run ./cmd/api
  ```

  ヘルスチェック: `GET http://localhost:8080/health`

## メモ

- API 定義を変えたら `make generate` を再実行。
- 依存を追加・更新したら `go mod tidy` を忘れずに。
