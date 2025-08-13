## 1.ビルド(dbは重いのでビルドしない)
``` sh
docker compose build server app 
```

## 2-a.実行(ソース変更だけならこれで十分)
# → 最速で反映確認できる、開発用
``` sh
docker compose up
```

## 2-b.バックグラウンド実行(本番運用)
``` sh
docker compose up -d
```

## main: server/appだけビルドして実行(パッケージの変更あるときなど)
``` sh
docker compose build server app && docker compose up
```

### main: server/appだけビルドしてバックグラウンド実行(パッケージの変更ある本番運用)
``` sh
docker compose build server app && docker compose up -d
```

