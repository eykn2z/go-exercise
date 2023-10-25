https://research.swtch.com/coro を読む

# 概要
coroutine packageがなぜgoに必要か
そもそもcoroutineとは

## function calls(subroutines)
> F calls G, which stops F and runs G.

関数Fが関数Gを呼び、Fが止まり、Gが実行を開始する

> G does its work, potentially calling and waiting for other functions, and eventually returns.

Gが、他の関数を呼び、待つ可能性もありながら、仕事を終え、最終的に実行を完了して返る。

> When G returns, G is gone and F continues running.

Gが返るとGは終了し、Fが続きを走らせる。

> In this pattern, only one function is running at a time, while its callers wait, all the way up the call stack.

このパターンでは一つのfunctionのみで一度、コールスタックの現在の位置に至るまで？呼び出し元を待ちながら実行している。

> In contrast to subroutines, coroutines run concurrently on different stacks, but it's still true that only one is running at a time, while its caller waits.

subroutineと対照的に、coroutineは違うスタック上で並行処理で実行する。しかし、呼び出し元が待ちながら一度に一つのみが実行するということは一致している。

> F starts G, but G does not run immediately. Instead, F must explicitly resume G, which then starts running.

FがGを開始するが、Gは即座に実行されない。代わりにGは明示的にGを実行させるために再開されなければいけない。

> At any point, G may turn around and yield back to F. 
> That pauses G and continues F from its resume operation. 
> Eventually F calls resume again, which pauses F and continues G from its yield.

いつでもGは現在の実行を一時中断して（向きを変える）Fに制御を戻す(yield back)するかもしれない。
それはGを停止し、それを再開させるオペレーションによってFを続ける。
最終的にFは、再び Fを停止し自分のyieldによってGを続ける再開を呼ぶ。

> On and on they go, back and forth, 
> until G returns, which cleans up G and continues F from its most recent resume, 
> with some signal to F that G is done 
> and that F should no longer try to resume G. 
> In this pattern, only one coroutine is running at a time, 
> while its caller waits on a different stack. 
> They take turns in a well-defined, coordinated manner.

これらは絶えず行ったり来たりする、
Gがクリーンアップし、一番直近の再開からFを続ける Gが返るまで
Gが終了するFへのいくつかのシグナルをもって
そしてこのFはGをこれ以上再開しようとしないはずだ。
このパターンの中では、たった一つのcoroutineが一度実行される、
呼び出し元が違うスタック上で待ちながら。
これらは明確に定義された、協力的な方法で交互に実行する

# Coroutines in Lua
# Generators in Python (Iterators in CLU)
# Coroutines, Threads, and Generators
##  "Concurrency"と"Parallelism"は違う
### 並行性（Concurrency）
タスクがアプリケーションの進行と並行して進行・管理される能力を指します。
並行処理はタスクが互いに独立して進行しているように見えるものであり、
実際のところ、これらのタスクが同時に（つまり物理的に同時に）実行されているかどうかは問題ではありません。
並行性は、単一のプロセッシングユニットでも発生し得ます。
例えば、コンテキストスイッチングを通じて。
### 並列性（Parallelism）
複数のタスクが同時にリテラルに実行される能力を指します。
並列処理は、複数のプロセッシングユニット（例えば、マルチコアプロセッサ）上で複数のタスクを同時に実行することを指します。

コルーチン、スレッド、ジェネレータの文脈での「並行性」は、タスクの管理と実行が重複している（つまり同時に進行しているように見える）が、必ずしも「同時」には実行されていないという状況を指しています。
一方で、スレッドが「並列性」を提供する場合、複数のタスクがリテラルに同時に実行されることを指します。

並行性はタスクの管理、並列性はタスクの同時実行に焦点を当てている。
## Coroutines
when one coroutine is running, the one that resumed it or yielded to it is not.
一つのコルーチンが実行されているとき、それを再開したものやそれにyieldしたものは実行されていない。
### 用途イメージ
1. Webサーバーとクライアントアプリケーション
   非同期I/O処理: コルーチンを使用すると、ウェブサーバーは非同期I/O操作をより効率的に処理でき、これによって高スループットを達成します。
   非同期APIリクエスト: クライアントアプリケーションでは、APIリクエストを並行して非同期に実行するためにコルーチンを利用します。
2. ゲーム開発
   同時アニメーション: 複数のアニメーションやロジックを同時に動かす際、コルーチンを使用して異なるタスクを同時に進行させます。
   ゲームロジックの制御: ゲーム中のキャラクターの動きやイベントを管理し、同時に進行させるためにもコルーチンが利用されます。
3. GUIアプリケーション
   UIの更新: コルーチンを使用して、バックグラウンドでデータを取得しつつ、UIをスムーズに更新する。
   ユーザーインタラクションの応答: コルーチンは、時間のかかる操作をバックグラウンドで実行しながら、UIをレスポンシブに保つのに役立ちます。
4. データサイエンスと機械学習
   データの同時処理: 大量のデータを分析する際、コルーチンを使用してデータのダウンロード、処理、分析を同時に実行します。
   非同期予測: モデルによる予測を非同期で行い、リソースを効率的に利用する。
5. ネットワークプログラム
   同時接続の管理: コルーチンを使用して、複数のネットワーク接続を同時に管理し、データを非同期で送受信します。
   非同期通信: クライアントとサーバーが非同期でデータをやり取りする際にコルーチンを利用する。
   コルーチンは、その協調的な特性から、どこでタスクが一時停止し、再開されるかをプログラマが制御できるため、タスクがどのようにスケジュールされるかについて高いコントロールを持ちます。これにより、複雑なフローを持つプロジェクトでも、シンプルで理解しやすいコードを書くことができます。

コルーチンは、その協調的な特性から、どこでタスクが一時停止し、再開されるかをプログラマが制御できるため、タスクがどのようにスケジュールされるかについて高いコントロールを持ちます。これにより、複雑なフローを持つプロジェクトでも、シンプルで理解しやすいコードを書くことができます。

## Threads
スレッドはコルーチン+並列性　強い
スケジュールのオーバーヘッドのコストがかかる
preemption:
コンピュータ上で実行中のプログラム（タスク）を強制的に一時中断し、他をプログラムの実行に切り替えること
割り込み処理
taxonomy:分類学

goのgoroutineは簡単に実装できるcoroutine
1つのgoルーチンから別のgoルーチンへのコンテキストスイッチは数百ナノ秒オーダー
ランタイム自体が一部を肩代わりしてるから、速く行える

javaの新しいスレッドも基本的に同じgoroutine
## Generators
コルーチンより機能が限られている(less power than coroutine)
yieldを使って関数を一時停止し、次に呼び出された時に再開する
コルーチンは任意の場所で止めて任意の場所で再開が可能

# なぜGoでCoroutine?
ほぼ同等(close enough)

goroutineによる並列性はレースコンディションを引き起こす可能性がある。
コルーチンは指定された同期点でのみコンテキストスイッチが行われるため、レースコンディションを避けられる可能性が高い。
この特定の例では、コルーチンがもたらす恩恵（例：データ共有における安全性、効率的なコンテキストスイッチなど）が強調されています。

```go
func (t *Tree[V]) All(yield func(v V)) {
    if t != nil {
        t.left.All(yield)
        yield(t.value)
        t.right.All(yield)
    }
}
//---
t.All(func(v V) {
fmt.Println(v)
})
//---
for v := range t.All {
fmt.Println(v)
}
```

be interaced(織り交ぜられる)

coroutines would provide an answer, letting us turn a function like (*Tree).All (a “push” iterator) into a function that returns a stream of values, one per call (a “pull” iterator).
### push iterator
### pull iterator

be indistinguishable from(~と区別がつかない)

CLU-likeとは？？






