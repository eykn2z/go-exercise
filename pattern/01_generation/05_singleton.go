package main

import (
	"fmt"
	"sync"
)

/*
===
singleton
===
クラスが１つのインスタンスのみを持つことを保証
インスタンスへの大域アクセスポイントを提供

DBやファイル等共有しげんへのアクセスを制御
グローバル変数と同じように、どこからでもオブジェクトのアクセスを許す
インスタンスが他のコードによって変更されるのを防止

singleton・・・別義：問題点の１つだけを解決する何らかの技法


・デフォルトコンストラクターをprivate
・コンストラクターとして機能する静的作成メソッドを作成
	・内部で非公開コンストラクターを呼び出してオブジェクトを生成し、静的なフィールドに保存
・次回以降のこのメソッドへの呼出は全てキャッシュされたオブジェクトを返す


・グローバル変数をより厳密に管理するとき
・DBインスタンス等

デメリット
・マルチスレッド環境において、複数のスレッドがシングルトン・オブジェクトを複数回生成しないよう工夫が必要
・ユニットテストが困難　非公開のため、静的メソッドを上書きすることが不可能
　singletonのモックを行うには、巧妙な方法を考える必要あり、orテスト放棄orSingletonの実装を諦める

他パターンとの関係
・FacadeはしばしばSingletonに変換可能
・Flywightも似ている
・Abstract Factories, Builders,PrototypesはSingletonsで実装可能
*/

// goroutineの例

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("single instance already created.")
		}
	}
	return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}
	fmt.Scanln()
}
