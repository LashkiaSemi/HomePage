# semi homepage

## topページの話
お知らせの部分は、`活動内容`として管理しています。
活動内容のうち、日付(処理用)が今日以降の項目を表示します。

## ログ排出ルールの話
基本的にerrorが起きたら、ハンドリングするときに注釈をつけてwrap。
個人的に`failed to ***`って感じで統一しました。

実際に`log.Printf()`するのは`handler`で行います。例外的に、リポジトリやミドルウェアで出力しているところがあるけど。

### prefixの話

- `[error]` = サーバ側のエラー。直す必要がある。
- `[warn]` = エラーじゃないけど、こんなことが起こったよ。っていう。
    - id指定してリソースを取ってきたけど、中身が空だった。
    - ログインしているか確認するためにcookieから値取ってきたら、空(未ログイン)だった。とか
- `[info]` = サーバのinfo。あんま見なくていいよ

## ディレクトリの話
- go
    - アプリケーション
    - go上で、`go run cmd/main.go`ってやると起動します。goのインストールが必要。。。
- design_proto
    - design_proto上でgulp sass-watchってやるとsassの自動コンパイルができるよ
    - cssとかのデザインするために置いてるやつ

### トップレベル

```
homepage
├── README.md
├── db/             // データベース関連
├── design_proto/   // デザインのプロトタイプ作る。CSSの編集とかしかしない。
├── go/             // アプリケーションを突っ込んでる
├── go.mod
└── go.sum
```

### db
```
db
└── init.d/
    ├── ddl.sql   // 全テーブル定義。当初あったものは同じ定義になっているはず。
    ├── dml.sql   // 当初なかったtableにデータを入れるためのdml
    └── dump.sql  // 開発時に使ったdbのダンプ
```

### design_proto
```
design_proto
├── gulpfile.js
├── package-lock.json
├── package.json
└── src/
    ├── index.html        // コンポーネント単位の表示テスト用
    ├── base.html         // html/内のファイルを作りやすくするためのコピペ元
    ├── css/              // いじらない！sassのコンパイル出力先
    │   ├── admin.css
    │   └── homepage.css
    ├── html/             // 各ページテスト用のhtml
    │   ├── activity.html
    │   ├── ...
    │   └── society.html
    └── sass/             // sass書くところ。cssを変更するときはこっちで変更してる
        ├── admin.scss
        └── homepage.scss
```

### go
厳密ではないかもしれないが、クリーンアーキテクチャっぽく書いてます。

```
go
├── Dockerfile
├── cmd/
│   └── main.go  // これ走らせれば動く
├── go.mod       // いじらない！依存関係のファイル
├── go.sum       // いじらない！依存関係のファイル
├── pkg
│   ├── configs
│   │   └── constant.go // 定数を設定している。コネクション周りの設定も入ってるよ。
│   ├── entity           // エンティティ部分。オブジェクトをゴニョゴニョする。
│   │   ├── activity.go
│   │   ├── ...
│   │   └── user.go
│   ├── infrastructure   // インフラ部分。
│   │   ├── auth
│   │   │   ├── auth.go  // パスワードの検証とかハッシュまわりのコード
│   │   │   ├── jwt.go   // jwt関連はここ
│   │   │   └── session.go // セッション周り管理。cookieとか。
│   │   ├── database
│   │   │   └── database.go // データベースアクセスの実装
│   │   ├── dcontext        // 使ってねえわ多分
│   │   │   └── dcontext.go
│   │   ├── handler         // 各種ハンドラ。リクエストを受けてレスポンスをする場所。
│   │   │   ├── app.go      // 全リソースのハンドラをまとめる子
│   │   │   ├── common.go   // ハンドラの操作で使ってる共通操作みたいなもの
│   │   │   ├── activity.go
│   │   │   ├── ....
│   │   │   └── user.go
│   │   └── server
│   │       ├── middleware
│   │       │   └── authentication.go // 認証周りのmiddleware
│   │       ├── response
│   │       │   └── response.go // テンプレートエンジンを良い感じにしてくれる。
│   │       └── server.go // サーバの起動とかルーティングとかの定義
│   ├── interface           // インターフェース部分
│   │   ├── controller     // エンティティをレスポンス向けの整形をしている
│   │   │   ├── common.go  // controllerで共通して使う機能
│   │   │   ├── activity.go
│   │   │   ├── ...
│   │   │   └── user.go
│   │   └── repository
│   │       ├── datastore.go // DIキメるためのインターフェースを定義
│   │       ├── activity.go
│   │       ├── ...
│   │       └── user.go
│   └── usecase                // ユースケース部分。
│       └── interactor  // 俗に言うビジネスロジック担当。実際の処理はせず、他の層にデータを投げ続けるイメージ。
│           ├── verify.go  // DI用のインターフェース。パスワードハッシュが必要なので。
│           ├── activity.go
│           ├── ...
│           └── user.go
├── public  // ファイルアップロード先
│   ├── lectures/     // レクチャーのファイルアップロード先
│   └── researches/   // 卒業研究のファイルアップロード先
├── static  // 静的ファイルの置き場所
│   ├── css
│   │   ├── admin.css
│   │   └── homepage.css
│   └── image
│       └── header.png
└── template  // テンプレートファイル。_から始まるファイルは部品
    ├── _footer.html // footer部分
    ├── _header.html // header部分
    ├── login.html
    ├── logout.html
    ├── error.html   // errorページの拡張用
    ├── index.html   // topページ
    ├── admin        // 管理者サイト用のテンプレ
    ├── activity     // 以下、各種リソース用のテンプレ
    ├── equipment
    ├── job
    ├── lecture
    ├── link
    ├── member
    ├── research
    └── society

```