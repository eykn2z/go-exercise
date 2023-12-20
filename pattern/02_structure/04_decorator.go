package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/*
===
decorator
===
wrapper
A,B,C,D機能があり、componentにA,Bの組み合わせ,B,Cの組み合わせ、A,B,C,Dの組み合わせ　どれでも付与できるようにしたい
Adapterはオブジェクトに対して異なるインターフェースを定義
Decoratorはインターフェースを強化
classを渡せば汎用的なdecoratorになれる？
*/

func main() {
	//fileReader, err := NewFileReader("example.txt")
	//if err != nil {
	//	fmt.Println("Error creating FileReader:", err)
	//	return
	//}
	//// 暗号化デコレータを適用
	//encryptedReader := NewEncryptionDecorator(fileReader)

	// サンプルテキストを持つコンポーネント
	textReader := NewTextReader("This is a sample text.")

	// 暗号化デコレータを適用
	encryptedReader := NewEncryptionDecorator(textReader)

	// デコレートされたリーダーを使用してテキストを読み取る
	buffer := make([]byte, 100)
	_, err := encryptedReader.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Println("Decrypted content:", string(buffer))
}

type Reader interface {
	Read([]byte) (int, error)
}

type FileReader struct {
	file *os.File
}

type TextReader struct {
	textReader io.Reader
}

func NewFileReader(filename string) (*FileReader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return &FileReader{file: file}, nil
}

func (fr *FileReader) Read(p []byte) (int, error) {
	return fr.file.Read(p)
}

func NewTextReader(text string) *TextReader {
	return &TextReader{textReader: strings.NewReader(text)}
}

func (tr *TextReader) Read(p []byte) (int, error) {
	return tr.textReader.Read(p)
}

type EncryptionDecorator struct {
	reader io.Reader
}

func NewEncryptionDecorator(reader io.Reader) *EncryptionDecorator {
	return &EncryptionDecorator{reader: reader}
}

func (ed *EncryptionDecorator) Read(p []byte) (int, error) {
	n, err := ed.reader.Read(p)
	if err != nil {
		return n, err
	}

	// 暗号化ロジック
	for i := 0; i < n; i++ {
		p[i] = encryptByte(p[i])
	}

	return n, nil
}

func encryptByte(b byte) byte {
	return b + 1
}
