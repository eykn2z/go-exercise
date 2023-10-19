https://go.dev/talks/2011/lex.slide を読む
# Coroutine
coroutineを使うとmessyなコードから解放される,
lexerを例にとって示す
# templateパッケージsystem
inflexible: 非柔軟性
inexpressive: 非表現力
code too fragile: コードが壊れやすい

arbitrary: 独断的な

lexerの中でやることの中でこのslideの中でfocusすること
the stuff outside actions
action delimiters
identifiers
numeric constants
string constants
and others

# どうやってtokenizeするか？
use a tool such as lex or ragel: lex,ragelを使う
use regular expressions: 一般的な表現を使う
use states, actions, and a switch statement: states,actions,switchを使う
# toolの利用
悪くないが・・・
hard to get good errors (can be very important): いいエラーを得るのが難しい(超重要になりうる)
tend to require learning another language: 他の言語を必須としがち？
result can be large, even slow: 結果が大きくなるもしくは遅い
often a poor fit: あまり適していない？
but lexing is easy to do yourself! : だけどlexingを自分で簡単に実行できる
# Reguler expressions
別のブログに詳しく書いてある
overkill
slow
can explore the state space too much: ステートのスペースを過剰に探索しうる
misuse of a dynamic engine to ask static questions: 静的な質問を問うための動的なエンジンの誤用？
# 自分で書いてみよう!
> Plus, most programming languages lex pretty much the same tokens, 
> so once we learn how it's trivial to adapt the lexer for the next purpose.

多くのプログラミング言語はかなり同じトークンをlex(話す?)している
なので一度どれだけ次の目的のためにlexerを適応することが些細なことか学ぼう

> an argument both for and against tools

ツールに賛成・反対する理由の両方を
# state machine
```go
// One iteration:
switch state {
case state1: 
    state = action1()
case state2:
    state = action2()
case state3: 
    state = action3()
}
```
state machineは忘れっぽい

> Boring and repetitive and error-prone, but anyway:
Why switch?
After each action, you know where you want to be;
the new state is the result of the action.
But we throw the info away and recompute it from the state.
(A consequence of returning to the caller.)
A tool can compile that out, but so can we.

退屈で繰り返しで誤りがち
なぜswitchか
それぞれのアクション後、自分が何になりたいか知っている、つまり新しいステートはアクションの結果である
しかし我々はその情報を投げ出し、ステートからそれを再計算する。(callerに返却されるのが結果)
ツールは出力をコンパイル出来るが、自分達でもできるよ。

# stateとaction?
state...我々がどこにいて何を期待しているのか
action...何をするのか
actionは新しいstateの結果

# 再帰的定義
```go
// stateFn represents the state of the scanner
// as a function that returns the next state.
type stateFn func(*lexer) stateFn

// run lexes the input by executing state functions
// until the state is nil.
func run() {
    for state := startState; state != nil; {
        state = state(lexer)
    }
}
```
# clientにtokenをどのように利用可能にするか？
同時実行すべし
> Tokens can emerge at times that are inconvenient to stop to return to the caller.

トークンは呼び出し元に戻るのに不便なタイミングで出現することがある

> Run the state machine as a goroutine, emit values on a channel

ステートマシンをgoroutineで走らせ、チャンネル上に値を送出する

```go
// lexer holds the state of the scanner.
type lexer struct {
    name  string    // used only for error reports.
    input string    // the string being scanned.
    start int       // start position of this item.
    pos   int       // current position in the input.
    width int       // width of last rune read from input.
    items chan item // channel of scanned items.
}
```
lexer
```go
func lex(name, input string) (*lexer, chan item) {
    l := &lexer{
        name:  name,
        input: input,
        items: make(chan item),
    }
    go l.run()  // Concurrently run state machine.
    return l, l.items
}
```
ここまで [link](https://go.dev/talks/2011/lex.slide#23)

