# HomePage
Semi Homepage Project

# 環境構築
 * `git remote add origin [リポジトリURL]`で、リモートにこのリポジトリを登録
 * プロジェクトのコードをを`git pull origin develop`で持ってくる
 * 'bundle install --without production --path vendor/bundle'
 * `rails s` + `localhost:3000` でページが表示できていれば成功

# バージョン
 * Ruby : 2.4.1p111
 * Rails : 5.1.3

# Userのroleについて
  userの役割（というか権限）を決める  
  * member ・・・ 一般ゼミ生、内部ページ閲覧の機能まで
  * admin ・・・ 管理者ゼミ生、新規ユーザーの登録や、各種ページの作成等をweb上で行える権限。基本web係になると思われる。
  * owner ・・・ 所有者ゼミ生、今のところできることはadminと変わらないかも。サーバー係  
