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