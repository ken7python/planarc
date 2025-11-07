# PlanArc

学習の「計画・実行・振り返り」を1つにまとめた、ToDo管理×学習記録のWebアプリです。タスクの消化状況や科目別の学習時間を見える化し、気分や楽しみの記録、AIによるフィードバックまで日々の学びを優しく伴走します。

## 概要（何ができるか）
- ToDo管理: 科目・優先度（MUST/WANT）・日付で整理、完了チェック、未完了タスクの繰越/戻し
- 学習記録: 手動入力/ストップウォッチで開始・終了時刻を記録し、学習時間を自動集計
- 可視化/分析: 日別の時間割表示、期間指定の科目別円グラフ/積み上げ棒グラフ
- 日次ステータス: 気分（4段階）と「今日の楽しみ」を保存
- AI振り返り: 当日の実績とメモから、先生/親/友達の人格で200字前後のコメントをSSEで生成・保存
- 認証/セキュリティ: JWT認証、CORS制御、レートリミット（必要時）

## アピールポイント（高校生エンジニアとして）
- 実運用を意識した設計: API/フロント分離、JWT認証、環境変数による設定、Docker化で再現性ある開発体験
- データドリブンなUI: 科目別/日別の集計とグラフで「がんばり」を可視化し継続を支援
- プロダクト志向の体験設計: ToDoと学習ログ、気分・楽しみの記録、AIフィードバックを1画面で回せる日課フロー
- モダンな技術選定: Vue 3 + Vite、Go + Gin + Gorm、SSEによる逐次出力表示
- 成長のための工夫: 音声入力UI、PWA対応、ホットリロード環境で開発効率を最大化

## 技術スタック / 構成
- フロントエンド: Vue 3 + Vite, Chart.js, @event-calendar/core, Element Plus, PWA
- バックエンド: Go (Gin, Gorm, JWT, CORS, rate limiter), OpenAI API（SSEストリーミング）
- データベース: MySQL
- インフラ/ツール: Docker Compose, .env による設定切替

ディレクトリ:
- `app/` フロントエンド
- `server/` バックエンドAPI
- `mysql-init/` DB初期化（必要に応じて利用）
- `run-dev.sh` ローカル開発の起動ヘルパー
- `run-aws.sh` デプロイ用ヘルパー（例）

## クイックスタート（Docker 推奨）
前提: Docker と Docker Compose をインストール済み

1) 環境変数を準備
- ルートの`.env`（MySQLコンテナ用）例:
  - `MYSQL_ROOT_PASSWORD=...`
  - `MYSQL_DATABASE=PlanArc`
  - `MYSQL_USER=xxxx`
  - `MYSQL_PASSWORD=xxxxxxxx`
- `server/.env`（API用）を作成し、以下を設定してください（サンプル）:
  - `DB_USER=xxxx`
  - `DB_PASSWORD=xxxxxxxx`
  - `DB_HOST=db`（Docker Compose の MySQL サービス名）
  - `DB_PORT=3306`
  - `DB_NAME=PlanArc`
  - `SECRET_KEY=（十分に長いランダム文字列）`
  - `OPENAI_API_KEY=（あなたのOpenAI APIキー）`
  - `CORS_GO=ON`（ローカル開発でのCORS許可・本番環境はNginxでCORS管理しています）
- `app/.env.local`（フロント用）例:
  - `VITE_API_URL=http://localhost:8080/api`

2) ビルドと起動
- ビルド: `docker compose build server app`
- 起動: `docker compose up`

3) アクセス
- フロントエンド: `http://localhost:4173`
- API ヘルスチェック: `http://localhost:8080/`（`Welcome to the PlanArc API`）

4) 開発の小技
- ホットリロードで開発
- API CORSを切り替えたい場合は `server/.env` の `CORS_GO` を調整

## ローカル実行（Docker なし）
前提: Node.js 20+, Go 1.21+, MySQL 8.x

- MySQL を起動し、`DB_*` が接続可能な状態にする
- バックエンド（`server/`）
  - `server/.env` を作成（上記と同様）
  - 依存取得: `go mod download`
  - 実行: `go run .`
- フロントエンド（`app/`）
  - 依存取得: `npm install`
  - `.env.local` に `VITE_API_URL=http://localhost:8080/api`
  - 実行: `npm run dev`（デフォルトでポート 4173）

## 主要機能のAPI（抜粋）
- 認証: `POST /api/accounts/register`, `POST /api/accounts/login`, `GET /api/accounts/profile`
- 科目: `GET /api/subject/`, `POST /api/subject/add`, `POST /api/subject/edit`
- 勉強ログ: `GET /api/studylog/?date=YYYY-MM-DD`, `POST /api/studylog/add`, `POST /api/studylog/delete`
- ToDo: `GET /api/todo/?date=...`, `GET /api/todo/group?date=...`, `POST /api/todo/add|check|edit`
- 未完了: `GET /api/unfinished/`, `POST /api/unfinished/move|back|delete`
- ステータス: `GET /api/status/?date=...`, `POST /api/status/enjoyment|mood`
- コメント: `POST /api/comment/ask`（SSE）, `GET /api/comment/?date=...`

## 本番デプロイ
- 公開先: https://planarc.kencode.tech
- 環境変数の適切な管理（Secret Manager 等）と CORS 設定、DB バックアップの運用を推奨

## トラブルシュート
- APIに接続できない: `server/.env` の `DB_HOST/PORT/USER/PASSWORD/NAME` と DB の起動状態を確認
- CORSエラー: `CORS_GO` と `AllowOrigins` の設定、`VITE_API_URL` が一致しているか確認
- OpenAIエラー: `OPENAI_API_KEY` が未設定/失効していないか、SSEをブロックするプロキシが無いか確認
- フロントが起動しない: `app/node_modules` の権限/キャッシュ再作成（`rm -rf node_modules && npm install`）

---
PlanArc は、がんばりを見える化し、やさしく背中を押す学習パートナーです。改善案・コントリビューションも歓迎します！
