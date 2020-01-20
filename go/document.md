# session管理について
domainに定義されたsessionにて管理を行います。
アプリケーションが落ちたらセッションがきえます。つら。

トランザクション
https://stackoverflow.com/questions/16184238/database-sql-tx-detecting-commit-or-rollback

# ディレクトリのこと
```
go
├── cmd // mainしか入れません。
│   └── main.go
├── conf // config関連はここ。サーバとかデータベースの設定とか
│   └── config.go
├── go.mod
├── go.sum
└── pkg
    ├── domain // 使用するモデルたち。dbでテーブル増やしたりしたら作る
    │   ├── error.go
    │   ├── logger
    │   │   └── logger.go
    │   ├── session.go
    │   └── user.go
    ├── infrastructure // 外のことを唯一知ってる
    │   ├── authentication
    │   │   └── authenticate.go
    │   ├── datastore
    │   │   └── sql_handler.go
    │   ├── dcontext
    │   │   └── dcontext.go
    │   ├── handler
    │   │   ├── account_handler.go
    │   │   └── app_handler.go
    │   └── server
    │       ├── middleware
    │       │   └── authenticate.go
    │       ├── response
    │       │   └── response.go
    │       ├── router
    │       │   └── router.go
    │       ├── server.go
    │       └── session
    │           └── session.go
    ├── interface
    │   ├── controller
    │   │   └── account_controller.go
    │   └── repository
    │       ├── account_repository.go
    │       ├── sql_handler.go
    │       └── transact.go
    └── usecase
        └── interactor
            ├── account_interactor.go
            ├── auth_handler.go
            └── repository.go
```
