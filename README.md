<h1>GoSample</h1>
Go言語の練習用プロジェクト

<h2>プロジェクトについて</h2>

- このプロジェクトを実行するには下記の配下に配置してください。

  - {GoPath}/src/

- Go のインストールは、Go 公式の Web サイトまたはこちらを参考にしてください
  - https://nbaboston.net/1405.html

<h2>【gin_sumple】</h2>

Go 言語の Web フレームワークである Gin フレームワークのためのフォルダです。<br>
チャットアプリを作っています。

<h3> -------環境--------</h3>

- エディタ

  - vsCode

- フレームワーク
  - gin
  - gorm

<h3>----フォルダ説明----</h3>

- strike/go/common

  - 共通的なパッケージ

    application.json からプロパティを取得する処理、ログファイルを出力する処理等を実装

- strike/go/db

  - データベース関連のパッケージ

  主に単一テーブルの操作を実装

  - ※/main/DBCreate.go は DB 作成用 Go ファイルです。
  - gin_sumple\bat\DBCreate.bat を実行することで DB が作成されます。

- strike/go/web
  - Web アプリでの各画面の GET、POST 処理を実装するパッケージ
- strike/go/model

  - 複数テーブルからデータ取得等の操作を実装するパッケージ

- ./content
  - Web の画面処理をまとめ

<h3>----実行方法----</h3>
- bat/AppStart.batを実行でアプリケーションが起動できます。

<h2>【sample】</h2>

go 言語の簡単なプログラム一覧です。<br>
何をしているかはそれぞれのソースをご覧ください

<h2>【参照】</h2>

- go について書いた記事
  - https://nbaboston.net/category/go
