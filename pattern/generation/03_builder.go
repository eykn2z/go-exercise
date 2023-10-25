package main

/*
builder,directorからなる
builderは任意のパーツ制作の置き場
directorが同パーツを組み合わせて生成する: クライアントコードからプロダクト構築を隠蔽する
directorの中にパーツの組み込みを行うので段階的に構築が可能

※builderは

例)
ドアは木、壁は木で白、2階建の家にしたい
ドアは木、壁はコンクリート、平家にしたい
この場合家にドアの木を設置するcomponentがほしい

*/
