# kadai_a
## 課題A

### 実行環境
macOSX(goでビルドできて、python3が実行できればどこでも動きます)

### 確認方法
```
// 単体テストの実行
$ go test ./

// ビルド
$ go build a.go

// データセットの作成。以下の例では3つのデータセットを作成
$ python3 create_dataset.py 3 > dataset

// データセットを使って確認
$ ./a < dataset
```

### テスト実行結果例
```
$ python3 create_dataset.py 3 > dataset

$ cat dataset
3
4
18 4 7 2 17 2 5
11
20 11 19 13 1 16 7 20 5 18 14 10 17 15
12
12 2 3 19 14 19 10 12 11 4 2 15 1 16 16 19 16

$ ./a < dataset
true 55
false 171
true 191
```