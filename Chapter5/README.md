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
- 空リターンは名前付き結果の変数のそれぞれをただし順序で返す省略記法。
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

# 5.6 無名関数

- 名前付き関数はパッケージレベルでだけ宣言できる
- 関数値を表す関数リテラルは全ての式内で使える
- 関数リテラルはfunc予約語の後に名前がなく、無名関数と呼ばれる

```
func add1(r rune) rune { return r + 1 }
fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"

// 無名関数を使用
strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000") // "IBM.:111"
```

- 上記のような方法で定義された関数はレキシカルなスコープ全体へアクセスできる

```
// squaresは呼び出されるごとに次の平方数を返す関数を返す
func squares() func() int{
    var x int
    return func() int {
        x++
        return x * x
    }
}

func main() {
    f := squares()
    fmt.Println(f()) // "1"
    fmt.Println(f()) // "4"
    fmt.Println(f()) // "9"
    fmt.Println(f()) // "16"
}
```

- 上記の例は関数値が単なるコードではなく、状態を持つことができることを示している
- 無名内部関数はそれを囲んでいる関数suquaresのローカル変数へアクセスしたり更新したりできる
- 関数を参照型として分類し、間数値が比較可能でない理由はこれらの隠蔽された変数への参照のため
- このような間数値はクロージャと呼ばれる技法を用いて実装されている

## 5.6.1 警告：ループ変数の補足

- レキシカルスコープ規則の落とし穴に注意

- 複数のディレクトリを作成して、作成したディレクトリを後で削除するプログラム

```
var rmdirs []func()
for _, d := range tempDirs() {
    dir := d                // 注意：必要！
    os.MkdirAll(dir, 0755)  // 親ディレクトリも作成する
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir)
    })
}
// ...何らかの処理...

for _, rmdir := range rmdirs {
    rmdir()     // 削除処理
}
```

- 下記の記載方法は正しくない

```
var rmdirs []func()
for _, dir := range tempDirs() {
    os.MkdirAll(dir, 0755)
    rmdirs = append(rmsirs, func() {
        os.RemoveAll(dir)   // 注意：ループ変数を参照、正しくない！
    })
}
```

- このループにより作成されるすべての間数値はアドレス化可能なメモリ位置を共有する
- dirの値はループで更新されるので、削除関数を呼び出した時には繰り返しの最後に代入された値を保持しており、RemoveAllの全ての呼び出しは同じディレクトリに対して行われる

- この問題をさけるために、内側の変数は外側の変数と全く同じ名前がつけられる

```
for _, dir := range tempDirs() {
    dir := dir  // 内側のdirを宣言し、外側のdirで初期化される
    // ...
}
```

- ループ変数の捕捉の問題はgo分やdeferを使っているときに最も多く発生する

# 5.7 可変個引数関数

- 可変個の引数で呼び出すことができる関数
- 可変個引数を宣言するには最後のパラメーターの型の前に省略記号"..."をつける


```
func sum(vals ...int) int {
    total := 0
    for _, val := range vals {
        total += val
    }
    return total
}

fmt.Println(sum())              // "0"
fmt.Println(sum(3))             // "3"
fmt.Println(sum(1, 2, 3, 4))    // "10"
```

- 呼び出し元は暗黙的に配列を割り当てて、引数をその配列へコピーして、関数にその配列全体のスライスを渡す
- 上記の最後の呼び出しは下記の呼び出しと同じように振る舞う
- 下記コードは引数がすでにスライスの中にある場合の可変個引数関数の呼び出し方であり、最後の引数の後に省略記号を書く

```
values := []int{1, 2, 3, 4}
fmt.Println(sum(valus...))  // "10"
```

- ...intパラメータは関数本体内ではスライスとして振る舞うが、可変個引数関数の型は普通のスライスパラメータを持つ関数の型とは明確に異なる

```
func f(...int) {}
func g([]int) {}

fmt.Printf("%T\n", f)   // "func(...int)"
fmt.Printf("%T\n", g)   // "func([]int)"
```

- 可変個引数関数はたいていは文字列をフォーマットするために使われる

```
// 先頭に行番号を持つようにフォーマットしてエラーメッセージを構築
func errorf(linenum int, format string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
    fmt.Fprintf(os.Stderr, format, args...)
    fmt.Println(os.Stderr)
}

linum, name := 12, "count"
errorf(linenum, "undefined: %s", name)  // "Line 12: undefined: count"
```

- interface{}型はすべての型を受け付けることができる

# 5.8 遅延関数呼び出し

- defer文は、普通の関数やメソッドの呼び出しの前に予約語deferを付けたもの
- 関数と引数の式はdefer文が実行されるときに評価されるが、実際の呼び出しはdefer文を含む関数が完了するまで遅延される
- returnを実行したり関数の最後に到達したりという正常な完了であっても、パニックによる異常な完了であっても、何個でも呼び出しを遅延することができる
- 呼び出しは遅延された順序の逆順に実行される
- defer文はオープンとクローズ、接続と切断、ロックとアンロックなどのように一対となる操作で使われることが多い

- 資源を解放するdefer文の正しい位置は、資源の獲得に成功した直後

```
func title(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // ...何らかの処理...

    return nil
}
```

- 開かれたファイルを閉じる場合

```
package ioutil

func ReadFile(filename string) ([]byte, error) {
    f, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    return ReadAll(f)
}
```

- ミューテックスをアンロックする場合

```
var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
    mu.Lock()
    defer mu.Unlock()
    return m[key]
}
```

- 複雑な関数をデバッグする場合に「入った(on entry)」処理と「出た(on exit)」処理を一対にするためにも使うことができる
- defer文での最後の丸かっこを忘れると、「入った」処理は出る際に行われ、「出た」処理は全く行われない

```
func bigSlowOperation() {
    defer trace("bigSlowOperation")()   // 最後の追加の丸かっこを忘れないこと
    // ...大量の処理...
    time.Sleep(10 * time.Second)    // スリープによって遅い操作を模倣する
}

func trace(mst string) func() {
    start := time.Now()
    log.Printf("enter %s", msg)
    return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
```

- bigSlowOperationが呼び出されるごとに、関数に入った時間・出た時間のログを出力する

```
$ go build gopl.io/ch5/trace
$ ./trace
2015/11/18 09:53:26 enter bigSlowOperation
2015/11/18 09:53:36 exit bigSlowOperation(10.000589217s)
```

- ループ内のdefer文は遅延実行されることで問題ないか確認が必要

```
for _, filename := range filenames {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()     // 注意：危険、ファイル奇術師が枯渇する可能性がある
    // ...fを処理する...
}
```

- １つの解決方法は、defer文を含むループ本体を、繰り返しごとに呼ばれる別の関数に移すこと

```
for _, filename := range filenames {
    if err := doFile(filename); err != nil {
        return err
    }
}

func doFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()
    // ...fを処理する...
}
```

# 5.9 パニック

- 境界外への配列アクセスやnilポインタによる参照などの誤りを検出した場合はパニック（panic）になる
- 典型的なパニックでは、通常の実行は停止し、そのゴルーチン内でのすべての遅延関数の呼び出しが行われ、プログラムはログメッセージを表示してクラッシュする
- ログメッセージにはパニック値（panic value）、スタックトレース（stack trace）が含まれる
→パニックしたプログラムに関するバグレポートにはこのログメッセージを常に含めるべき

- すべてのパニックがランタイムから発生するわけではない
- 組み込みのpanic関数を直接呼び出せる

```
switch s := suit(drawCard()); s{
    case "Spades":      // ...
    case "Hearts":      // ...
    case "Diamonds":    // ...
    case "Clubs":       // ...
    default:
        panic(fmt.Sprintf("incalid suit %q", s))    // ジョーカー？
}
```

Goのパニック機構は他言語の例外に似ているが、パニックが使われる状況はかなり異なる

- パニックはプログラムをクラッシュさせるので、一般的にプログラム内での論理的不整合などのように重大なエラーに対して使われる
- 「予期される」エラー、つまり誤った入力、誤った設定、失敗したI/Oなどから発生する類のエラーはerror値を使って処理されるべき

- パニックが発生した場合には、すべての遅延された関数は、スタックの最上位の関数の遅延された関数から始まってmain関数まで逆順に実行される

```
func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
```

- 標準出力には次のように表示される

```
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.f(0x0)
        C:/Users/mana/go/src/github.com/myoshinaga/GoProgrammingLanguage/Chapter5/10/main.go:10 +0x17b
main.f(0x1)
        C:/Users/mana/go/src/github.com/myoshinaga/GoProgrammingLanguage/Chapter5/10/main.go:12 +0x156
main.f(0x2)
        C:/Users/mana/go/src/github.com/myoshinaga/GoProgrammingLanguage/Chapter5/10/main.go:12 +0x156
main.f(0x3)
        C:/Users/mana/go/src/github.com/myoshinaga/GoProgrammingLanguage/Chapter5/10/main.go:12 +0x156
main.main()
        C:/Users/mana/go/src/github.com/myoshinaga/GoProgrammingLanguage/Chapter5/10/main.go:6 +0x31
exit status 2
```

- スタックダンプの出力もできる

```
func main() {
    defer printStack()
    f(3)
}

func printStack() {
    var buf [4096]byte
    n := runtime.Stack(buf[:], false)
    os.Stdout.Write(buf[:n])
}
```

# 5.10 リカバー

パニックになっても、プログラムを終了させないようにリカバー（回復）することができる

- 遅延された関数内で組み込みのrecover関数が呼び出され、かつdefer文を含む関数がパニックになっているとrecoverはパニックの現状を終了させてパニック値を返す
- パニックになっていた関数は止まった場所から続けることはないが、正常にリターンする
- パニックになっていないときにrecoverが呼び出された場合、何も影響はなくnilを返す

```
// パーサにバグがあった場合、クラッシュする代わりに、パニックを普通の構文解析エラーに変える
func Parse(input string) (s *Syntax, err error) {
    defer func() {
        if p := recover(); p != nil {
            err = fmtErrorf("internal error: %v", p)
        }
    }()
    // ...パーサ...
}
```

- パニックから無条件に回復することは好ましくない
→パニック後のパッケージの状態はほとんど定義されていなかったり文書化されていなかったりするため
- 同じパッケージ内でパニックから回復することは、複雑なエラーや予期しないエラーの処理を簡単にするのに役立つ
- しかし、一般的な規則として他のパッケージからのパニックからの回復を試みるべきではない
- 公開APIは失敗をerrorとして報告すべき
- 呼び出し元が提供するコールバックのような自分で管理していない関数から発生するパニックからも回復すべきではない