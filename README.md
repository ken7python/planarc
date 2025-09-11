# PlanArc

## 概要
PlanArcは、勉強計画と記録を一体的に管理するためのウェブアプリケーションです。ユーザー登録、科目ごとの勉強時間やTODO管理、未完了タスクの繰越、気分・楽しみの記録、AIによる振り返りコメントなどを提供します。

## 技術スタック
- **フロントエンド:** Vue 3 + Vite
- **バックエンド:** Go (Gin, Gorm)
- **データベース:** MySQL
- **その他:** Docker Compose, Google Generative AI (Gemini)

## ディレクトリ構成
- `app/` - フロントエンドコード
- `server/` - バックエンドAPI
- `mysql-init/` - データベース初期化スクリプト
- `run-dev.sh` - 開発用起動スクリプト
- `run-aws.sh` - AWS用起動スクリプト

## 主な機能
- アカウント登録・ログイン・プロフィール取得
- 科目の登録・編集
- 勉強ログの記録と取得
- ToDoリストの作成、編集、完了チェック
- 未完了タスクの移動・戻し
- 気分・楽しみ状況の記録
- Geminiによる学習振り返りコメント生成

## デプロイ
本番環境は [https://planarc.kencode.tech](https://planarc.kencode.tech) で公開しています。

=======
## 開発環境の起動
DockerとDocker Composeを利用して簡単に環境構築できます。

```sh
# ビルド
docker compose build server app
# 起動
docker compose up
```

ソースコードの変更のみを素早く反映したい場合は `./run-dev.sh` を利用してください。
