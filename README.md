# nandokuka

---
## なにこれ
難読化シェル芸を手で撃つのは大変だ！=> ツールを作ろう！

難読化を手助けするツールです。難読化の参考程度に考えてもらえればOKです。
## Install
```sh
go get github.com/xztaityozx/nandokuka
```

## Require
- xxd

## Usage
```sh
nandokuka [GLOBAL OPTION] [COMMAND] [OPTION] [FILE]
```
### GLOBAL OPTION
- -d|--decode : デコードします

### COMMAND
#### ascii
ASCIIコードに置き換えるASCII難読化を行います。
ex)
	`date => $'\x64\x61\x74\x65'`
	
ASCII難読化は以下みたいにワンライナーをまとめて変換します

```sh
seq 30 | awk 'NR%2==0{print}'
↓
$'\x73\x65\x71\x20\x33\x30\x7c\x61\x77\x6b\x20\x27\x4e\x52\x25\x32\x3d\x3d\x30\x7b\x70\x72\x69\x6e\x74\x7d\x27'
```

変換された出力を eval すれば実行できます

---
#### base64
base64変換するbase64難読化を行います。
ex)
	`date =>  ZGF0ZQ==`

このまま使うとただのbase64エンコーダーです。

##### オプション
- `-e|--excutable`
  - コピペすれば実行できる形で出力します
- `-j,--jp`
  - 日本語base64難読化をします

---
#### echo
文字をechoやprintfで文字列1文字を出力し、コマンド列を作り上げる難読化を行います
ex)
	`date => $(echo d)$(echo a)$(echo t)$(echo e)`

##### オプション
- `-r|--random`
  - echoだけじゃなくprintfも使います
- `-v|--verbose`
  - 変換過程を出力します

---
#### gzip
gzipとgunzipを使ったgzip難読化をおこないます
ex)
	`date => 1f8b080078d81b5b00034b492c4905007a379eaa04000000`

gzipとxxdに通した結果を出力します

##### オプション
- `-e|--excutable`
  - コピペすれば実行できるような形で出力します

---
#### help
ヘルプを出力します。

---
#### symbol
1,2,A,zだけで難読化する記号オンリー難読化を行います

ex)
	`date => ${@:2$((1+2)):1}${@:2$((2+2*2)):1}${@:$((2*2*2-1)):1}${@:22:1}`

パラメータ展開を利用した難読化です

##### オプション
- `-p|--prefix`
  - 記号オンリー難読化に必要な材料も一緒に出力します
- `-s|--super`
  - A,zを使わない超・記号オンリー難読化を行います

---
#### version
バージョン情報を出力します

## その他
どのサブコマンドにも`--help`があるのでそちらもどうぞ
