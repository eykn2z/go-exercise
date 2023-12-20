package main

import "fmt"

/*
===
adapter
===
非互換なインターフェースのオブジェクト同士の協働を可能とする

オブジェクトをラップして、裏で行われる変換の詳細を隠蔽
ラップされたオブジェクトはアダプター内で動作していることすら認識しない
メートルとキロメートルで動作するオブジェクトを全てのデータをフィートやマイルに変換するアダプターでラップすること可能

1. 既存オブジェクトの一つと互換なインターフェースを実装
2. このインターフェースを使用して、既存オブジェクトはアダプターのメソッドを安全に呼び出す
3. 呼び出されると、アダプターはリクエストを２つ目のオブジェクトに渡す。ただし、２つ目のオブジェクトが期待する形式で渡す

1. 最低２つの非互換なクラスがある
2. クライアント・インターフェースを宣言し、クライアントとサービスとの情報伝達の方法を記述
3. クライアントインターフェースに従うアダプタークラスを作成
4. アダプタークラスにサービスオブジェクトへの参照を格納するためのフィールドを追加

メリット
・開放閉鎖の原則
	・クライアントインターフェースを開始てアダプターを連携する限り既存のクライアント側コード新しい種類のアダプターをプログラムに追加可能
デメリット
・一連の新規のインターフェースとクラスを追加する必要があるため複雑性追加
*/

type Client struct{}

func (c *Client) Insert(com Computer) {
	com.Insert()
}

type Computer interface {
	Insert()
}

type Mac struct{}

func (m *Mac) Insert() {
	fmt.Println("called Insert")
}

type Windows struct{}

func (w *Windows) InsertA() {
	fmt.Println("called InsertA")
}

type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) Insert() {
	w.windowMachine.InsertA()
}

func main() {
	client := &Client{}

	mac := &Mac{}
	client.Insert(mac)

	windows := &Windows{}
	windowsAdapter := &WindowsAdapter{
		windowMachine: windows,
	}

	client.Insert(windowsAdapter)
}
