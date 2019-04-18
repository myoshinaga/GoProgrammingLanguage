#  5.1 関数宣言

関数宣言は以下の要素で構成される。
- 名前
- パラメーターのリスト
- 省略可能な結果のリスト
- 本体

```
func name(parameter-list) (result-list) {
    body
}
```

パラメータ列や結果列は同一の型をまとめて宣言できる。
```
func f(i, j, k int, s, t, string) {/* ... */}
func f(i int, j int, k int, s string, t string) {/* ... */}
```

「パラメータの型の列」と「結果の型の列」が同じであれば、
シグニチャ（関数の型）は同じである。
※C#やJavaのシグニチャは「関数名」「パラメータの型の列」。
パラメーターの型、または個数を変えて関数をオーバーロードできる。

- デフォルトパラメーター機能なし
- 引数は値渡し（ただしポインタ、スライス、マップ、関数、チャネルなどの何らかの種類の参照が含まれていれば別）

# 5.2 再帰
関数は再帰呼び出しできる。

多くのプログラミング言語の実装は、
- 固定長の関数呼び出しスタックを使っている
- スタックの大きさは64KBから2MBまでが普通
- 固定長スタックは再帰呼び出しの深さに制限があるため、大きなデータ構造を再帰的に走査する場合はスタックオーバーフローを避けるために注意する必要がある

それに対してGoは
- 最初は小さく、そして必要に応じてギガバイトの大きさへと拡張される可変長スタックを使っている
- オーバーフローを心配することなく再帰を安全に扱える

# 5.3 複数戻り値
- 多値呼び出しは、複数のパラメータを持つ関数呼び出しの引数として書けるが、製品のコードではめったに使われない
- 戻り値にわかりやすい名前をつけるとドキュメント化には役立つ
- 慣習では最後のboolは成功、errorはエラーを示す
- 空リターンは名前付き結果の変数のそれぞれを正しい順序で返す省略記法。
→あまり使わない方がよい。

# 5.4 エラー
- 原因が1つしかありえない失敗は「型：boolean」「変数名：ok」を返す
- さまざまな原因があり得る場合は「型：error」「変数名：err」を返す
 - errorのログ出力
 ```
 fmt.Println(err)
 fmt.Printf("%v", err)
 ```
  - errorがnilでない場合は他の結果は無視する
- エラーを振る舞いの1つとして扱い、例外（Exceptions）は使わない
- パニックはバグであることを示す本当に予期されていないエラーを報告するためだけに使う

## 5.4.1 エラー処理戦略

### 伝搬
エラーを検査して適切に対応するのは呼び出し元の責任
- 呼び出し元にエラーを返す
```
resp, err := http.Get(url)
if err != nil {
    return nil, err
}
```

- 必要な情報を付け加えてエラーを返す
```
doc, err := http.Parse(resp.Body)
resp.Body.Close()
if err != nil {
    return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
}
```

### 再試行
一時的あるいは予想できないエラーの場合は失敗した操作を再び試みる
- 再試行前に遅延
- 試行回数の制限
- 試行に費やす時間の制限

### プログラム停止
処理を進めるのが不可能な場合は、呼び出し元でエラーを表示してプログラムを停止する
- この処理はmainパッケージに限定するべき
- ライブラリ関数ではバグでない限り呼び出し元にエラーを伝搬すべき
```
// (main関数内で)
if err := WaitForServer(url); err != nil {
    fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
    os.Exit(1)
}
```

- ```log.Fatalf``` はデフォルトで日付と時刻がエラーメッセージの先頭につけられる
```
if err := WaitForServer(url); err != nil {
    log.Fatalf("Site is down: %v\n", err)
}

// 接頭辞の設定
log.SetPrefix("wait: )
// 日付と時刻の表示を抑制
log.SetFlags(0)
```

### エラーを記録して制限された機能で処理継続
```
// 接頭辞つき
if err := Ping(); err != nil {
    log.Printf("ping failed: %v; networking disabled", err)
}

// 標準エラーに直接表示
if err := Ping(); err != nil {
    fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled\n", err)
}
```

### エラー無視
エラーを安全に無視できるまれな場合
```
dir, err := ioutil.TempDir("", "scratch")
if err != nil {
    return fmt.Errorf("failed to create temp dir: %v", err)
}

// ...tempディレクトリ使用...

os.RemoveAll(dir) // エラーを無視 $TEMODIR は定期的に削除される
```

## 5.4.2 ファイルの終わり（EOF:End of File）
ファイルの終わり（End of file）の状態に対しては他のエラーと異なった対応をする
```
in := bufio.NewReader(os.Stdin)
for {
    r, _, err := in.ReadRune()
    if err = io.EOF {
        break // 読み込みを終了
    }
    if err != nil {
        return fmt.Errorf("read failed: %v", err)
    }
    // ..rを使う..
}
```

# 5.5 関数値
関数はファーストクラス値であり、関数値は他の値と同様に型を持ち、
変数に代入したり、関数へ渡したり、関数から返したりできる
```
func square(n int) int { return n * n }
func negative(n int) int { return -n }
func product(m, n int) int { return m * n }

f := square
fmt.Println(f(3))   // "9"

f := negative
fmt.Println(f(3))   // "-3"
fmt.Printf("%T\n", f)   // "func(int) int"

f = product // コンパイルエラー:func(int, int) int を func(int) int へ代入できない
```

関数型のゼロ値はnilであり、nilの関数を呼び出すとパニックになる
```
var f func(int) int
f(3)    // パニック:nilの関数の呼び出し
```

関数値はnilと比較できる
```
var f func(int) int
if f != nil {
    f(3)
}
```

関数値は比較可能ではないので、間数値同士は比較してはいけない
マップのキーとしても使えない

関数のデータだけではなく振る舞いについてもパラメーターかすることができる
```
// func Map(mapping func(rune) rune, s string) string
// mapping関数によって操作された全ての文字のコピーを返す

func add1(r rune) rune { return r + 1 }

fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
fmt.Println(strings.Map(add1, "VMS")) // "WNT"
fmt.Println(strings.Map(add1, "Admix")) // "Benjy"
```
