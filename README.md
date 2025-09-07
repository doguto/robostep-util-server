# robostep-util-server

Notion のタスク更新を Discord に通知するためのユーティリティサーバーです。

## 概要

このプロジェクトは、Notion のWebhookを受け取り、タスクの変更内容を整形してDiscord のチャンネルに通知するGo製のWebサーバーです。タスク管理の効率化とチーム間の情報共有を目的としています。

## 機能

- Notion Webhook の受信と処理
- タスク情報の抽出と整形
- Discord への自動通知
- RESTful API エンドポイントの提供

## 必要条件

- Go 1.24 以上
- Discord Webhook URL
- Notion データベースとWebhook設定

## インストール

### 1. リポジトリのクローン

```bash
git clone https://github.com/doguto/robostep-util-server.git
cd robostep-util-server
```

### 2. 依存関係のインストール

```bash
go mod download
```

### 3. 環境変数の設定

プロジェクトルートに `.env` ファイルを作成し、以下の環境変数を設定してください：

```env
DISCORD_WEBHOOK_URL=your_discord_webhook_url_here
```

#### Discord Webhook URL の取得方法

1. Discord で通知を送信したいチャンネルの設定を開く
2. 「統合」→「Webhook」→「新しいWebhook」を作成
3. 生成されたWebhook URLをコピーして環境変数に設定

## ビルドと実行

### 開発環境での実行

```bash
go run cmd/main.go
```

### プロダクション用ビルド

```bash
go build -o robostep-util-server cmd/main.go
./robostep-util-server
```

サーバーはデフォルトでポート8080で起動します。

## API エンドポイント

### ヘルスチェック

```
GET /
```

サーバーの動作確認用エンドポイントです。

**レスポンス例:**
```
ok
```

### Notion タスク通知

```
POST /api/notion/notify_nhk_task
```

Notion からのWebhookを受信し、タスク更新をDiscordに通知します。

**リクエストボディ:**
Notion Webhook のペイロード形式

**レスポンス例:**
```json
{
  "message": "Task notification sent to Discord",
  "payload": { ... }
}
```

## 通知内容

Discord に送信される通知には以下の情報が含まれます：

- タスク名
- 担当者
- ステータス
- 期日
- Notion ページへのURL

**通知例:**
```
### タスクリストが更新されました！
　タスク名：**サンプルタスク**
　　担当者：**山田太郎**
ステータス：`進行中`
　　　期日：**2024-12-31**
　　　URL：https://notion.so/...
```

## プロジェクト構造

```
robostep-util-server/
├── cmd/
│   └── main.go              # エントリーポイント
├── internal/
│   ├── controllers/
│   │   └── notion_controller.go    # Notion関連のコントローラー
│   └── payloads/
│       └── notion_payload.go       # Notionペイロードの型定義
├── go.mod                   # Go モジュール定義
├── go.sum                   # 依存関係のハッシュ
├── .env                     # 環境変数（作成が必要）
├── .gitignore
└── README.md
```

## 開発

### 依存関係

- [Gin](https://github.com/gin-gonic/gin) - HTTP Webフレームワーク
- [godotenv](https://github.com/joho/godotenv) - 環境変数管理

