# github cliのユーザーを変えるやつ

readmeにとりあえず書いとく
gh cliのインストールが必要

cloneしてきてgo buildでbuildして実行/usr/local/binとかに配置して使う。

# 改善点
とりあえず一旦グ茶って書いたからきれいに整理したい。
あとログインしたユーザーが分からないgh auth statusしても結果が取れない。そこも合わせてできるようにしたい。


あと、分けたいトークンをとりあえずprivateとenterpriseでしか分けてないけどこれだとこの二つしか指定できないから
なんかもっと柔軟にできるようにしたい

ゆくゆくはファイルごとに自動で認知してくれてswitch-ghだけで切り替わるようにしたい

---

## もはやghで切り替えれる
パーソナルアクセストークン取得してalias作成してそれぞれ切り替えれば変更できる
なのであんまりこのリポジトリ意味ない
