このレポジトリは動きません

goでwebアプリケーションを開発するときにBaalというドキュメントフォーマットとSQLからどんなファイルを生成してどうやって開発しているかと言うのを示すためのサンプルです

## ジェネレート

### baal/*

[Baal](http://techblog.kayac.com/unity_advent_calendar_2016_20) はKAYACで使っているAPIドキュメントを兼ねたクライアントのコードをジェネレートするためのIDLです。

静的型付け言語のGoで開発するに当たってGoでのgenerateでも活用するようになりました。

`baal/*` 以下の`.faced` ふぁいるから `app/api/*.auto.go` `app/data/**/*.auto.go` を生成します

ジェネレータは[go-baal](https://github.com/shogo82148/go-baal) を使ってIDL解析し、text/template などで生成し x/tools/imports でフォーマットしています

### sql/

DSLとしてSQLでいいじゃん派でORMごとにオレオレDSLを覚えたくないので、SQLからコードを生成しています。

sqlをもとに `app/database/table/*.auto.go` と `app/database/row/*.auto.go` を生成しています。

ジェネレータは[schemalex](https://github.com/schemalex/schemalex)を使って解析し、text/template などで生成し x/tools/imports でフォーマットしています

このサンプルではありませんが、tableがmasterdataである場合はユニークインデックスをキーにしたキャッシュコードも自動生成しています。

## 開発手順

1. baal定義を書いてクライアントサイドエンジニアと同意を取る
2. sqlを書いてスキーマを決め、サーバーサイドエンジニアでレビューをする
3. baalからハンドラーをジェネレートする
4. `app/database/*` いかのコードをジェネレートする
5. app/database/row/以下に手でコードを書いてrowからdataに変換する関数を無心で書く
6. modelにAPIと対称となるコードを書く(ここに時間を使う)
7. apiとmodelを結合しHandlerを登録する
