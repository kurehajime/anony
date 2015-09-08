# anony

文中の人名をイニシャルに変換するコマンド

![anony](https://cloud.githubusercontent.com/assets/4569916/9704202/531bdade-54d8-11e5-8687-bdcdcd8a638d.gif)


人名を含んだ文章(例:test.txt)を・・・

>筒井康隆は、日本の小説家・劇作家・俳優である。ホリプロ所属。身長166cm。小松左京、星新一と並んで「SF御三家」とも称される。パロディやスラップスティックな笑いを得意とし、初期にはナンセンスなSF作品を多数発表。1970年代よりメタフィクションの手法を用いた前衛的な作品が増え、エンターテインメントや純文学といった境界を越える実験作を多数発表している。

anonyに渡すと・・・

```
$ anony test.txt
```

人名をイニシャルに変換します。

>T・Yは、日本の小説家・劇作家・俳優である。ホリプロ所属。身長166cm。K・S、H・Sと並んで「SF御三家」とも称される。パロディやスラップスティックな笑いを得意とし、初期にはナンセンスなSF作品を多数発表。1970年代よりメタフィクションの手法を用いた前衛的な作品が増え、エンターテインメントや純文学といった境界を越える実験作を多数発表している。



## インストール方法

[ここ](https://github.com/kurehajime/anony/releases)から実行ファイルをダウンロードできます。


Go言語の開発環境がある場合はgo getでもインストールできます。

```
go get -u github.com/kurehajime/anony
```

## オプション

  -e string  
    	文字コード (初期値 "utf-8")  
  -s	single initial  
    	イニシャルを一文字で表示
