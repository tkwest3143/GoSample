<h1>GoSumple</h1>
Go言語の練習用プロジェクト

<h2>プロジェクトについて</h2>

 - このプロジェクトを実行するには下記の配下に配置してください。 
   - {Goインストールフォルダ}/src/github/

 -Goのインストールは、Go公式のWebサイトまたはこちらを参考にしてください
   - https://nbaboston.net/1405.html

<h2>【gin_sumple】</h2>

Go言語のWebフレームワークであるGinフレームワークのためのフォルダです。<br>
チャットアプリを作っています。

<h3> -------環境--------</h3>

- DB
  - Postgresql

- エディタ
  - vsCode

- フレームワーク
  - gin
  - gorm

<h3>----フォルダ説明----</h3>

- github/GoSumple/gin_sumple/go/common
   - 共通的なパッケージ
     
     application.jsonからプロパティを取得する処理、ログファイルを出力する処理等を実装

- github/GoSumple/gin_sumple/go/db
  - データベース関連のパッケージ
  
  　主に単一テーブルの操作を実装

- github/GoSumple/gin_sumple/go/web
  - Webアプリでの各画面のGET、POST処理を実装するパッケージ
	
- github/GoSumple/gin_sumple/go/model
  - 複数テーブルからデータ取得等の操作を実装するパッケージ

- github/GoSumple/gin_sumple/content
  - Webの画面処理をまとめ
	
<h2>【参照】</h2>

- goについて書いた記事
  - https://nbaboston.net/category/go
