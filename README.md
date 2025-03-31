# 目的

GOを利用してクリーンアーキテクチャの学習を目的としてリポジトリとなります。

# TODO

- デプロイ
- フロントエンドの作成と接続

# SET UP

```
# create module
go mod init echo-rest-api-todo
# start db
docker compose up
# start app
GO_ENV=dev go run .
# run migrate
GO_ENV=dev go run migrate/migrate.go
```

# Architecture

![image](https://github.com/user-attachments/assets/f85779ec-3a7b-4619-872c-06507bb767d7)
